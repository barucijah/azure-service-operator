/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package pipeline

import (
	"context"
	"strings"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
)

// ImprovePropertyDescriptionsStageID is the unique identifier for this pipeline stage
const ImprovePropertyDescriptionsStageID = "improvePropertyDescriptions"

// ImprovePropertyDescriptions adds documentation to any properties missing it by copying the description from the
// corresponding type
func ImprovePropertyDescriptions() *Stage {
	stage := NewStage(
		ImprovePropertyDescriptionsStageID,
		"Improve property descriptions by copying from the corresponding type",
		func(ctx context.Context, state *State) (*State, error) {
			visitor := createPropertyImprovingVisitor(state.Definitions())
			result := make(astmodel.TypeDefinitionSet)
			for _, def := range state.Definitions() {
				newDef, err := visitor.VisitDefinition(def, nil)
				if err != nil {
					return nil, err
				}

				result.Add(newDef)
			}

			return state.WithDefinitions(result), nil
		})

	return stage
}

func createPropertyImprovingVisitor(defs astmodel.TypeDefinitionSet) astmodel.TypeVisitor[any] {
	visitor := astmodel.TypeVisitorBuilder[any]{
		VisitObjectType: func(
			this *astmodel.TypeVisitor[any],
			it *astmodel.ObjectType,
			ctx any,
		) (astmodel.Type, error) {
			result := it
			for _, prop := range it.Properties().AsSlice() {

				// If the property already has a description, use it
				if prop.Description() != "" {
					continue
				}

				// If the property type isn't a typename, we can't do anything
				tn, haveTypeName := astmodel.AsInternalTypeName(prop.PropertyType())
				if !haveTypeName {
					continue
				}

				// if the type name doesn't reference a type in our definitions, we can't do anything
				def, haveDef := defs[tn]
				if !haveDef {
					continue
				}

				// If there's no description on the type, there's nothing to do
				desc := def.Description()
				if len(desc) == 0 {
					continue
				}

				prop = prop.WithDescription(
					strings.Join(desc, " "))

				result = result.WithProperty(prop)
			}

			return result, nil
		},
	}

	return visitor.Build()
}
