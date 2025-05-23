/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-openapi/spec"
	"github.com/rotisserie/eris"
	"golang.org/x/sync/errgroup"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/config"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/jsonast"
)

const LoadTypesStageID = "loadTypes"

/*
	LoadTypes creates a PipelineStage to load Swagger data.

This information is derived from the Azure Swagger specifications. We parse the Swagger specs and look for
any actions that appear to be ARM resources (have PUT methods with types we can use and appropriate names in the
action path). Then for each resource, we use the existing JSON AST parser to extract the status type
(the type-definition part of swagger is the same as JSON Schema).
*/
func LoadTypes(
	idFactory astmodel.IdentifierFactory,
	config *config.Configuration,
	log logr.Logger,
) *Stage {
	return NewStage(
		LoadTypesStageID,
		"Load all types from Swagger files",
		func(ctx context.Context, state *State) (*State, error) {
			log.V(1).Info(
				"Loading Swagger data",
				"source", config.SchemaRoot)

			swaggerTypes, err := loadSwaggerData(ctx, idFactory, config, log)
			if err != nil {
				return nil, eris.Wrapf(err, "unable to load Swagger data")
			}

			log.V(1).Info(
				"Loaded Swagger data",
				"resources", len(swaggerTypes.ResourceDefinitions),
				"otherDefinitions", len(swaggerTypes.OtherDefinitions))

			if len(swaggerTypes.ResourceDefinitions) == 0 || len(swaggerTypes.OtherDefinitions) == 0 {
				return nil, eris.Errorf("Failed to load swagger information")
			}

			resourceToStatus, statusTypes, err := generateStatusTypes(swaggerTypes)
			if err != nil {
				return nil, err
			}

			resourceToSpec, specTypes, err := generateSpecTypes(swaggerTypes)
			if err != nil {
				return nil, err
			}

			// put all definitions into a new set
			defs := make(astmodel.TypeDefinitionSet)
			defs.AddTypes(statusTypes)
			defs.AddTypes(specTypes)

			for resourceName, resourceInfo := range swaggerTypes.ResourceDefinitions {
				spec := resourceToSpec.MustGetDefinition(resourceName)
				status := resourceToStatus.MustGetDefinition(resourceName)

				specType := spec.Type()
				specType = addRequiredSpecFields(specType)
				statusType := status.Type()

				resourceType := astmodel.NewResourceType(specType, statusType)

				// add on ARM Type, URI, and supported operations
				resourceType = resourceType.WithARMType(resourceInfo.ARMType).WithARMURI(resourceInfo.ARMURI)
				resourceType = resourceType.WithScope(resourceInfo.Scope)
				resourceType = resourceType.WithSupportedOperations(resourceInfo.SupportedOperations)

				resourceDefinition := astmodel.MakeTypeDefinition(resourceName, resourceType)

				// document origin of resource
				sourceFile := strings.TrimPrefix(resourceInfo.SourceFile, config.SchemaRoot)
				resourceDefinition = resourceDefinition.
					WithDescription(
						"Generator information:",
						fmt.Sprintf(" - Generated from: %s", filepath.ToSlash(sourceFile)),
						fmt.Sprintf(" - ARM URI: %s", resourceInfo.ARMURI),
					)

				err = defs.AddAllowDuplicates(resourceDefinition)
				if err != nil {
					return nil, err
				}

				err = addObjectResourceLinkIfNeeded(defs, spec, resourceName)
				if err != nil {
					return nil, eris.Wrapf(err, "failed to add resource link to %s", spec.Name())
				}
				err = addObjectResourceLinkIfNeeded(defs, status, resourceName)
				if err != nil {
					return nil, eris.Wrapf(err, "failed to add resource link to %s", status.Name())
				}
			}

			return state.WithDefinitions(defs), nil
		})
}

var requiredSpecFields = astmodel.NewObjectType().WithProperties(
	astmodel.NewPropertyDefinition(astmodel.NameProperty, "name", astmodel.StringType))

func addRequiredSpecFields(t astmodel.Type) astmodel.Type {
	return astmodel.BuildAllOfType(t, requiredSpecFields)
}

// generateStatusTypes returns the statusTypes for the input Swagger types
// all types (apart from Resources) are renamed to have "_STATUS" as a
// suffix, to avoid name clashes.
func generateStatusTypes(
	swaggerTypes jsonast.SwaggerTypes,
) (astmodel.TypeDefinitionSet, astmodel.TypeDefinitionSet, error) {
	resourceLookup, otherTypes, err := renamed(swaggerTypes, true, astmodel.StatusSuffix)
	if err != nil {
		return nil, nil, err
	}

	newResources := make(astmodel.TypeDefinitionSet)
	// often the top-level type in Swagger has a bad name like "CreateParametersX"
	// we'll try to substitute that with a better name here
	for resourceName, resourceDef := range resourceLookup {
		statusTypeName := resourceDef.Type().(astmodel.TypeName) // always a TypeName, see 'renamed' comment
		desiredStatusName := resourceName.WithName(resourceName.Name() + astmodel.StatusSuffix)

		if statusTypeName == desiredStatusName {
			newResources.Add(resourceDef)
			continue // nothing to do
		}

		targetType, err := otherTypes.FullyResolve(statusTypeName)
		if err != nil {
			panic("didn't find type in set after renaming; shouldn't be possible")
		}

		if _, isOneOf := astmodel.AsOneOfType(targetType); isOneOf {
			// We must avoid copying the body of a OneOf type into the status of the
			// resource because they are currently partially formed - they won't be
			// standalone independent types until after conversion to objects.
			targetType = statusTypeName
		}

		err = otherTypes.AddAllowDuplicates(astmodel.MakeTypeDefinition(desiredStatusName, targetType))
		if err != nil {
			// unable to use desiredName
			newResources.Add(resourceDef)
		} else {
			newResources.Add(astmodel.MakeTypeDefinition(resourceName, desiredStatusName))
		}
	}

	return newResources, otherTypes, nil
}

// Note that the first result is for mapping resource names → types, so it is always TypeName→TypeName.
// The second contains all the renamed types.
func renamed(
	swaggerTypes jsonast.SwaggerTypes,
	status bool,
	suffix string,
) (astmodel.TypeDefinitionSet, astmodel.TypeDefinitionSet, error) {
	renamer := astmodel.NewRenamingVisitorFromLambda(
		func(typeName astmodel.InternalTypeName) astmodel.InternalTypeName {
			return typeName.WithName(typeName.Name() + suffix)
		})

	var errs []error
	otherTypes := make(astmodel.TypeDefinitionSet)
	for _, typeDef := range swaggerTypes.OtherDefinitions {
		renamedDef, err := renamer.RenameDefinition(typeDef)
		if err != nil {
			errs = append(errs, err)
		} else {
			otherTypes.Add(renamedDef)
		}
	}

	resources := make(astmodel.TypeDefinitionSet)
	for resourceName, resourceDef := range swaggerTypes.ResourceDefinitions {
		// resourceName is not renamed as this is a lookup for the resource type
		typeToRename := resourceDef.SpecType
		if status {
			typeToRename = resourceDef.StatusType
		}
		renamedType, err := renamer.Rename(typeToRename)
		if err != nil {
			errs = append(errs, err)
		} else {
			resources.Add(astmodel.MakeTypeDefinition(resourceName, renamedType))
		}
	}

	err := kerrors.NewAggregate(errs)
	if err != nil {
		return nil, nil, err
	}

	return resources, otherTypes, nil
}

func generateSpecTypes(
	swaggerTypes jsonast.SwaggerTypes,
) (astmodel.TypeDefinitionSet, astmodel.TypeDefinitionSet, error) {
	// TODO: I think this should be renamed to "_Spec" for consistency, but will do in a separate PR for cleanliness #Naming
	// the alternative is that we place them in their own package
	resources, otherTypes, err := renamed(swaggerTypes, false, "")
	if err != nil {
		return nil, nil, err
	}

	// Fix-up: rename top-level resource types so that a Resource
	// always points to a _Spec type.
	// TODO: remove once preceding TODO is resolved and everything is consistently named _Spec #Naming
	{
		renames := make(astmodel.TypeAssociation)
		for typeName := range otherTypes {
			if _, ok := resources[typeName]; ok {
				// would be a clash with resource name
				renames[typeName] = typeName.WithName(typeName.Name() + "Spec")
			}
		}

		renamer := astmodel.NewRenamingVisitor(renames)
		newOtherTypes, renameErr := renamer.RenameAll(otherTypes)
		if renameErr != nil {
			panic(renameErr)
		}

		otherTypes = newOtherTypes

		// for resources we must only rename the Type not the Name,
		// since this is used as a lookup:
		newResources := make(astmodel.TypeDefinitionSet)
		for rName, rType := range resources {
			newType, renameErr := renamer.Rename(rType.Type())
			if renameErr != nil {
				panic(renameErr)
			}
			newResources.Add(astmodel.MakeTypeDefinition(rName, newType))
		}

		rewriter := astmodel.TypeVisitorBuilder[any]{
			VisitObjectType: func(this *astmodel.TypeVisitor[any], it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
				// strip all readonly props
				var propsToRemove []astmodel.PropertyName
				it.Properties().ForEach(func(prop *astmodel.PropertyDefinition) {
					if prop.ReadOnly() {
						propsToRemove = append(propsToRemove, prop.PropertyName())
					}
				})

				it = it.WithoutSpecificProperties(propsToRemove...)

				return astmodel.IdentityVisitOfObjectType(this, it, ctx)
			},
		}.Build()

		otherTypes, err = rewriter.VisitDefinitions(otherTypes, nil)
		if err != nil {
			return nil, nil, err
		}

		resources = newResources
	}

	return resources, otherTypes, nil
}

func loadSwaggerData(
	ctx context.Context,
	idFactory astmodel.IdentifierFactory,
	config *config.Configuration,
	log logr.Logger,
) (jsonast.SwaggerTypes, error) {
	schemas, err := loadAllSchemas(
		ctx,
		idFactory,
		config,
		log)
	if err != nil {
		return jsonast.SwaggerTypes{}, err
	}

	loader := jsonast.NewCachingFileLoader(schemas)

	typesByGroup := make(map[astmodel.LocalPackageReference][]typesFromFile)
	countLoaded := 0
	for schemaPath, schema := range schemas {
		logInfoSparse(
			log,
			"Loading Swagger files",
			"loaded", countLoaded,
			"total", len(schemas))
		countLoaded++

		extractor := jsonast.NewSwaggerTypeExtractor(
			config,
			idFactory,
			schema.Swagger,
			schemaPath,
			*schema.Package, // always set during generation
			loader,
			log)

		types, err := extractor.ExtractTypes(ctx)
		if err != nil {
			return jsonast.SwaggerTypes{}, eris.Wrapf(err, "error processing %q", schemaPath)
		}

		typesByGroup[*schema.Package] = append(typesByGroup[*schema.Package], typesFromFile{types, schemaPath})
	}

	log.Info(
		"Loaded Swagger files",
		"loaded", countLoaded,
		"total", len(schemas))

	return mergeSwaggerTypesByGroup(idFactory, typesByGroup, config)
}

var (
	loadingLock sync.Mutex
	lastLogTime *time.Time
)

func logInfoSparse(log logr.Logger, message string, keysAndValues ...interface{}) {
	loadingLock.Lock()
	defer loadingLock.Unlock()

	shouldDisplay := lastLogTime == nil || time.Since(*lastLogTime) > 800*time.Millisecond
	if shouldDisplay {
		log.Info(message, keysAndValues...)
		now := time.Now()
		lastLogTime = &now
	}
}

func mergeSwaggerTypesByGroup(
	idFactory astmodel.IdentifierFactory,
	m map[astmodel.LocalPackageReference][]typesFromFile,
	cfg *config.Configuration,
) (jsonast.SwaggerTypes, error) {
	result := jsonast.SwaggerTypes{
		ResourceDefinitions: make(jsonast.ResourceDefinitionSet),
		OtherDefinitions:    make(astmodel.TypeDefinitionSet),
	}

	for pkg, group := range m {
		merged, err := mergeTypesForPackage(group, idFactory, cfg)
		if err != nil {
			return result, eris.Wrapf(err, "merging swagger types for %s", pkg)
		}

		for rn, rt := range merged.ResourceDefinitions {
			if _, ok := result.ResourceDefinitions[rn]; ok {
				panic("duplicate resource generated")
			}

			result.ResourceDefinitions[rn] = rt
		}

		err = result.OtherDefinitions.AddTypesAllowDuplicates(merged.OtherDefinitions)
		if err != nil {
			return result, eris.Wrapf(err, "when combining swagger types for %s", pkg)
		}
	}

	return result, nil
}

type typesFromFile struct {
	jsonast.SwaggerTypes
	filePath string
}

// mergeTypesForPackage merges the types for a single package from multiple files into our master set.
// typesFromFiles is a slice of typesFromFile, each of which contains the types for a single file.
// idFactory is used to generate new identifiers for types that collide.
func mergeTypesForPackage(
	typesFromFiles []typesFromFile,
	idFactory astmodel.IdentifierFactory,
	cfg *config.Configuration,
) (*jsonast.SwaggerTypes, error) {
	// Sort into order by filePath so we're deterministic
	slices.SortFunc(
		typesFromFiles,
		func(left typesFromFile, right typesFromFile) int {
			return strings.Compare(left.filePath, right.filePath)
		})

	typeNameCounts := make(map[astmodel.InternalTypeName]int)
	for _, typesFromFile := range typesFromFiles {
		for name := range typesFromFile.OtherDefinitions {
			// TODO: This is very hacky
			if strings.Contains(name.String(), "network/v1api20220701/SubResource") {
				def := typesFromFile.OtherDefinitions[name]
				ot, ok := def.Type().(*astmodel.ObjectType)
				if ok {
					prop, foundProp := ot.Property("Id")
					if foundProp {
						ot = ot.WithProperty(prop.MakeOptional())
						def = def.WithType(ot)
					}
				}
				typesFromFile.OtherDefinitions[name] = def
			}

			typeNameCounts[name] += 1
		}
	}

	// a set of renamings, one per file
	renames := make([]astmodel.TypeAssociation, len(typesFromFiles))
	for ix := range typesFromFiles {
		renames[ix] = make(astmodel.TypeAssociation)
	}

	for name, count := range typeNameCounts {
		colliding := findCollidingTypeNames(typesFromFiles, name, count)
		if colliding == nil {
			continue
		}

		names := make([]string, len(colliding))
		for ix, ttc := range colliding {
			newName := generateRenaming(idFactory, name, typesFromFiles[ttc.typesFromFileIx].filePath, typeNameCounts)
			names[ix] = newName.Name()
			renames[ttc.typesFromFileIx][name] = newName
		}
	}

	for ix := range typesFromFiles {
		renamesForFile := renames[ix]
		if len(renamesForFile) > 0 {
			typesFromFiles[ix] = applyRenames(renamesForFile, typesFromFiles[ix])
		}
	}

	mergedResult := &jsonast.SwaggerTypes{
		ResourceDefinitions: make(jsonast.ResourceDefinitionSet),
		OtherDefinitions:    make(astmodel.TypeDefinitionSet),
	}

	for _, typesFromFile := range typesFromFiles {
		for _, t := range typesFromFile.OtherDefinitions {
			// for consistent results we always sort typesFromFiles first (at top of this function)
			// so that we always pick the same one when there are multiple
			_ = mergedResult.OtherDefinitions.AddAllowDuplicates(t)
			// errors ignored since we already checked for structural equality
			// it’s possible for types to refer to different typenames in which case they are not TypeEquals Equal
			// but they might be structurally equal
		}

		defs := mergedResult.OtherDefinitions
		for rn, rt := range typesFromFile.ResourceDefinitions {
			if foundRT, ok := mergedResult.ResourceDefinitions[rn]; ok {
				// We have two resources with the same name, if they have exact same structure, they're the same
				scopesEqual := foundRT.Scope == rt.Scope
				specsEqual := structurallyIdentical(foundRT.SpecType, defs, rt.SpecType, defs)
				statusesEqual := structurallyIdentical(foundRT.StatusType, defs, rt.StatusType, defs)
				if scopesEqual && specsEqual && statusesEqual {
					// They're the same resource, we're good.
					continue
				}

				// Check for renames and apply them if available
				delete(mergedResult.ResourceDefinitions, rn)
				newNameA, renamedA := tryRename(rn, rt, cfg)
				newNameB, renamedB := tryRename(rn, foundRT, cfg)

				if _, collides := mergedResult.ResourceDefinitions[newNameA]; collides {
					err := eris.Errorf(
						"merging file %s: renaming %s to %s resulted in a collision with an existing type",
						typesFromFile.filePath,
						rn.Name(),
						newNameA.Name())

					return nil, err
				}

				if _, collides := mergedResult.ResourceDefinitions[newNameB]; collides {
					err := eris.Errorf(
						"merging file %s: renaming %s to %s resulted in a collision with an existing type",
						typesFromFile.filePath,
						rn.Name(),
						newNameB.Name())

					return nil, err
				}

				if newNameA != newNameB && (renamedA || renamedB) {
					// One or other or both was successfully renamed, we can keep going
					mergedResult.ResourceDefinitions[newNameA] = rt
					mergedResult.ResourceDefinitions[newNameB] = foundRT
					continue
				}

				// Names are still the same. We have a collision.
				err := eris.Errorf(
					"merging file %s: duplicate resource types generated with name %s",
					typesFromFile.filePath,
					rn.Name())

				return nil, err
			}

			mergedResult.ResourceDefinitions[rn] = rt
		}
	}

	return mergedResult, nil
}

func tryRename(
	name astmodel.InternalTypeName,
	rsrc jsonast.ResourceDefinition,
	cfg *config.Configuration,
) (astmodel.InternalTypeName, bool) {
	if cfg == nil {
		// No renames configured
		return name, false
	}

	for _, ren := range cfg.TypeLoaderRenames {
		if !ren.AppliesToType(name) {
			// Doesn't apply to us, not our name
			continue
		}

		if ren.Scope != nil &&
			!strings.EqualFold(*ren.Scope, string(rsrc.Scope)) {
			// Doesn't apply to us, not our scope
			continue
		}

		if ren.RenameTo == nil {
			// No rename configured
			return name, false
		}

		return name.WithName(*ren.RenameTo), true
	}

	return name, false
}

type typeAndSource struct {
	def             astmodel.TypeDefinition
	typesFromFileIx int
}

// findCollidingTypeNames finds any types with the given name that collide, and returns
// the definition as well as the index of the file it was found in
func findCollidingTypeNames(
	typesFromFiles []typesFromFile,
	name astmodel.InternalTypeName,
	duplicateCount int,
) []typeAndSource {
	if duplicateCount == 1 {
		// cannot collide
		return nil
	}

	typesToCheck := make([]typeAndSource, 0, duplicateCount)
	for typesIx, types := range typesFromFiles {
		if def, ok := types.OtherDefinitions[name]; ok {
			typesToCheck = append(typesToCheck, typeAndSource{def: def, typesFromFileIx: typesIx})
			// short-circuit
			if len(typesToCheck) == duplicateCount {
				break
			}
		}
	}

	first := typesToCheck[0]
	for _, other := range typesToCheck[1:] {
		if !structurallyIdentical(
			first.def.Type(),
			typesFromFiles[first.typesFromFileIx].OtherDefinitions,
			other.def.Type(),
			typesFromFiles[other.typesFromFileIx].OtherDefinitions,
		) {
			if name != first.def.Name() || name != other.def.Name() {
				panic("assert")
			}

			_, firstOk := first.def.Type().(*astmodel.ResourceType)
			_, otherOk := other.def.Type().(*astmodel.ResourceType)
			if firstOk || otherOk {
				panic("cannot rename resources")
			}

			return typesToCheck
		}
		// Else: they are structurally identical and it is okay to pick one.
		// Note that when types are structurally identical they aren’t necessarily type.Equals equal;
		// for this reason we must ignore errors when adding to the overall result set using AddAllowDuplicates.
	}

	return nil
}

// generateRenaming finds a new name for a type based upon the file it is in
func generateRenaming(
	idFactory astmodel.IdentifierFactory,
	original astmodel.InternalTypeName,
	filePath string,
	typeNames map[astmodel.InternalTypeName]int,
) astmodel.InternalTypeName {
	name := filepath.Base(filePath)
	name = strings.TrimSuffix(name, filepath.Ext(name))

	// Prefix the typename with the filename
	result := astmodel.MakeInternalTypeName(
		original.InternalPackageReference(),
		idFactory.CreateIdentifier(name+original.Name(), astmodel.Exported))

	// see if there are any collisions: add Xs until there are no collisions
	// TODO: this might result in non-determinism depending on iteration order
	// in the calling method
	for _, ok := typeNames[result]; ok; _, ok = typeNames[result] {
		result = astmodel.MakeInternalTypeName(
			result.InternalPackageReference(),
			result.Name()+"X",
		)
	}

	return result
}

func applyRenames(
	renames astmodel.TypeAssociation,
	typesFromFile typesFromFile,
) typesFromFile {
	visitor := astmodel.NewRenamingVisitor(renames)

	// visit all other types
	newOtherTypes, err := visitor.RenameAll(typesFromFile.OtherDefinitions)
	if err != nil {
		panic(err)
	}

	// visit all resource types
	newResourceTypes := make(jsonast.ResourceDefinitionSet)
	for rn, rt := range typesFromFile.ResourceDefinitions {
		newSpecType, err := visitor.Rename(rt.SpecType)
		if err != nil {
			panic(err)
		}

		newStatusType, err := visitor.Rename(rt.StatusType)
		if err != nil {
			panic(err)
		}

		newResourceTypes[rn] = jsonast.ResourceDefinition{
			SpecType:            newSpecType,
			StatusType:          newStatusType,
			SourceFile:          rt.SourceFile,
			ARMURI:              rt.ARMURI,
			ARMType:             rt.ARMType,
			SupportedOperations: rt.SupportedOperations,
			Scope:               rt.Scope,
		}
	}

	typesFromFile.OtherDefinitions = newOtherTypes
	typesFromFile.ResourceDefinitions = newResourceTypes
	return typesFromFile
}

// structurallyIdentical checks if the two provided types are structurally identical
// all the way to their leaf nodes (recursing into TypeNames)
func structurallyIdentical(
	leftType astmodel.Type,
	leftDefinitions astmodel.TypeDefinitionSet,
	rightType astmodel.Type,
	rightDefinitions astmodel.TypeDefinitionSet,
) bool {
	// we cannot simply recurse when we hit TypeNames as there can be cycles in types.
	// instead we store all TypeNames that need to be checked in here, and
	// check them one at a time until there is nothing left to be checked:
	type pair struct {
		left  astmodel.InternalTypeName
		right astmodel.InternalTypeName
	}
	toCheck := []pair{}            // queue of pairs to check
	checked := map[pair]struct{}{} // set of pairs that have been enqueued

	override := astmodel.EqualityOverrides{}
	override.InternalTypeName = func(
		left astmodel.InternalTypeName,
		right astmodel.InternalTypeName,
	) bool {
		// note that this relies on Equals implementations preserving the left/right order
		p := pair{left, right}
		if _, ok := checked[p]; !ok {
			checked[p] = struct{}{}
			toCheck = append(toCheck, p)
		}

		// conditionally true, as long as all pairs are equal
		return true
	}

	// check the provided types
	if !astmodel.TypeEquals(leftType, rightType, override) {
		return false
	}

	// check all TypeName pairs until there are none left to check
	for len(toCheck) > 0 {
		next := toCheck[0]
		toCheck = toCheck[1:]
		if !astmodel.TypeEquals(
			leftDefinitions[next.left].Type(),
			rightDefinitions[next.right].Type(),
			override) {
			return false
		}
	}

	// if we didn’t find anything that didn’t match then they are identical
	return true
}

// TODO: is there, perhaps, a way to detect these without hardcoding these paths?
var skipDirectories = []string{
	"/examples/",
	"/quickstart-templates/",
	"/control-plane/",
	"/data-plane/",
}

func shouldSkipDir(filePath string) bool {
	p := filepath.ToSlash(filePath)

	for _, skipDir := range skipDirectories {
		if strings.Contains(p, skipDir) {
			return true
		}
	}

	return false
}

// loadAllSchemas walks all .json files in the given rootPath in directories
// of the form "Microsoft.GroupName/…/2000-01-01/…" (excluding those matching
// shouldSkipDir), and returns those files in a map of path→swagger spec.
func loadAllSchemas(
	ctx context.Context,
	idFactory astmodel.IdentifierFactory,
	config *config.Configuration,
	log logr.Logger,
) (map[string]jsonast.PackageAndSwagger, error) {
	rootPath := config.SchemaRoot
	localPathPrefix := config.LocalPathPrefix()
	overrides := config.Status.Overrides

	var mutex sync.Mutex
	schemas := make(map[string]jsonast.PackageAndSwagger)

	log.Info("Loading schemas", "rootPath", rootPath)
	var eg errgroup.Group
	eg.SetLimit(10)

	countFound := 0
	err := filepath.Walk(rootPath, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		if shouldSkipDir(filePath) {
			return filepath.SkipDir // this is a magic error
		}

		if !fileInfo.IsDir() &&
			filepath.Ext(filePath) == ".json" {

			group := groupFromPath(filePath, rootPath, overrides)
			version := versionFromPath(filePath, rootPath)
			if group == "" || version == "" {
				return nil
			}

			pkg := astmodel.MakeLocalPackageReference(
				localPathPrefix,
				idFactory.CreateGroupName(group),
				astmodel.GeneratorVersion,
				version)

			// We need the file if the version is short (e.g. "v1") because those are often shared between
			// resource providers.
			// Alternatively, we need the file if we have configuration for the group
			fileNeeded := len(version) < 10 ||
				config.ObjectModelConfiguration.IsGroupConfigured(pkg)
			if !fileNeeded {
				return nil
			}

			// all files are loaded in parallel to speed this up
			logInfoSparse(
				log,
				"Scanning for schemas",
				"found", countFound)
			countFound++

			eg.Go(func() error {
				var swagger spec.Swagger

				log.V(1).Info(
					"Loading OpenAPI spec",
					"file", filePath)

				fileContent, err := os.ReadFile(filePath)
				if err != nil {
					return eris.Wrapf(err, "unable to read swagger file %q", filePath)
				}

				err = swagger.UnmarshalJSON(fileContent)
				if err != nil {
					return eris.Wrapf(err, "unable to parse swagger file %q", filePath)
				}

				mutex.Lock()
				schemas[filePath] = jsonast.PackageAndSwagger{Package: &pkg, Swagger: swagger}
				mutex.Unlock()

				return nil
			})
		}

		return nil
	})

	egErr := eg.Wait() // for files to finish loading
	log.Info(
		"Scanning for schemas",
		"found", countFound)

	if err != nil {
		return nil, err
	}

	if egErr != nil {
		return nil, egErr
	}

	return schemas, nil
}

func groupFromPath(filePath string, rootPath string, overrides []config.SchemaOverride) string {
	filePath = filepath.ToSlash(filePath)
	group := jsonast.SwaggerGroupRegex.FindString(filePath)

	// see if there is a config override for this file
	for _, schemaOverride := range overrides {
		configSchemaPath := filepath.ToSlash(filepath.Join(rootPath, schemaOverride.BasePath))
		if strings.HasPrefix(filePath, configSchemaPath) {
			// a forced namespace: use it
			if schemaOverride.Namespace != "" {
				return schemaOverride.Namespace
			}

			// found a suffix override: apply it
			if schemaOverride.Suffix != "" {
				group = group + "." + schemaOverride.Suffix
				return group
			}
		}
	}

	return group
}

// supports date-based versions or v1, v2 (as used by common types)
var swaggerVersionRegex = regexp.MustCompile(`/(\d{4}-\d{2}-\d{2}(-preview|-privatepreview)?)|(v\d+)|(\d+\.\d+)/`)

func versionFromPath(filePath string, rootPath string) string {
	// we want to ignore anything in the root path, since, e.g.
	// the specs can be nested inside a directory that matches the swaggerVersionRegex
	// (and indeed this is the case with the /v2/ package)
	filePath = strings.TrimPrefix(filePath, rootPath)
	// must trim leading & trailing '/' as golang does not support lookaround
	fp := filepath.ToSlash(filePath)
	return strings.Trim(swaggerVersionRegex.FindString(fp), "/")
}

func addResource(spec astmodel.TypeDefinition, resourceName astmodel.TypeName) (astmodel.TypeDefinition, error) {
	visitor := astmodel.TypeVisitorBuilder[any]{
		VisitObjectType: func(this *astmodel.TypeVisitor[any], it *astmodel.ObjectType, ctx interface{}) (astmodel.Type, error) {
			it = it.WithResource(resourceName).WithIsResource(true)
			return it, nil
		},
	}.Build()

	updatedSpec, err := visitor.VisitDefinition(spec, nil)
	if err != nil {
		return astmodel.TypeDefinition{}, err
	}
	return updatedSpec, nil
}

func compareObjectTypeIgnoreIsResource(left *astmodel.ObjectType, right *astmodel.ObjectType) bool {
	left = left.ClearResources().WithIsResource(false)
	right = right.ClearResources().WithIsResource(false)

	return astmodel.TypeEquals(left, right)
}

func addObjectResourceLinkIfNeeded(defs astmodel.TypeDefinitionSet, def astmodel.TypeDefinition, resourceName astmodel.TypeName) error {
	resolvedDef, err := resolveDefAlias(defs, def)
	if err != nil {
		return err
	}

	// If the resolved name is the same as def, we don't have an alias and there is nothing to do
	if astmodel.TypeEquals(resolvedDef.Name(), def.Name()) {
		return nil
	}

	updatedDef, err := addResource(resolvedDef, resourceName)
	if err != nil {
		return err
	}

	existing, ok := defs[updatedDef.Name()]
	// Ensure that the updated def is equal to the existing def, except for new resource details. We don't want to accidentally change
	// a resources structure.
	updatedDefMostlyEqual := astmodel.TypeEquals(existing.Type(), updatedDef.Type(), astmodel.EqualityOverrides{ObjectType: compareObjectTypeIgnoreIsResource})
	if !ok || updatedDefMostlyEqual {
		defs[updatedDef.Name()] = updatedDef
	}

	return nil
}

// resolveDefAlias resolves the given definition and returns a fully resolved type (meaning it is not a TypeName pointing to another TypeName)
// This function caters to two scenarios:
//  1. A direct alias: TypeName -> TypeName, this is pretty self-explanatory.
//  2. An alias indirected through an AllOf with 2 Types. One must be a TypeName and the other must be an ObjectType with a single Name property.
//     Connascence alert: the type structure expected here is a direct result of including the name property as part of an AllOf
//     in swagger_type_extractor.go ExtractResourceTypes.
func resolveDefAlias(defs astmodel.TypeDefinitionSet, def astmodel.TypeDefinition) (astmodel.TypeDefinition, error) {
	resolvedDef, err := defs.FullyResolveDefinition(def)
	if err != nil {
		return astmodel.TypeDefinition{}, err
	}

	allOf, ok := resolvedDef.Type().(*astmodel.AllOfType)
	if !ok {
		return resolvedDef, nil
	}

	if allOf.Types().Len() != 2 {
		return resolvedDef, nil
	}

	var result astmodel.TypeDefinition
	var found bool
	var foundName bool
	allOf.Types().ForEach(func(t astmodel.Type, ix int) {
		if ot, ok := t.(*astmodel.ObjectType); ok {
			if _, ok := ot.Property(astmodel.NameProperty); ok && ot.Properties().Len() == 1 {
				foundName = true
			}
		}

		if name, ok := t.(astmodel.TypeName); ok {
			found = true
			result, err = resolveDefAlias(defs, astmodel.MakeTypeDefinition(astmodel.InternalTypeName{}, name))
		}
	})

	if err != nil {
		return astmodel.TypeDefinition{}, err
	}
	if !found || !foundName {
		return resolvedDef, nil
	}

	return result, nil
}
