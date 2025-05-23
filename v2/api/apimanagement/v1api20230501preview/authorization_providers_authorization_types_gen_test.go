// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501preview

import (
	"encoding/json"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	v20230501ps "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20230501preview/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_AuthorizationError_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationError_STATUS to AuthorizationError_STATUS via AssignProperties_To_AuthorizationError_STATUS & AssignProperties_From_AuthorizationError_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForAuthorizationError_STATUS, AuthorizationError_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAuthorizationError_STATUS tests if a specific instance of AuthorizationError_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAuthorizationError_STATUS(subject AuthorizationError_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.AuthorizationError_STATUS
	err := copied.AssignProperties_To_AuthorizationError_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AuthorizationError_STATUS
	err = actual.AssignProperties_From_AuthorizationError_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AuthorizationError_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationError_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationError_STATUS, AuthorizationError_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationError_STATUS runs a test to see if a specific instance of AuthorizationError_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationError_STATUS(subject AuthorizationError_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationError_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AuthorizationError_STATUS instances for property testing - lazily instantiated by
// AuthorizationError_STATUSGenerator()
var authorizationError_STATUSGenerator gopter.Gen

// AuthorizationError_STATUSGenerator returns a generator of AuthorizationError_STATUS instances for property testing.
func AuthorizationError_STATUSGenerator() gopter.Gen {
	if authorizationError_STATUSGenerator != nil {
		return authorizationError_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationError_STATUS(generators)
	authorizationError_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthorizationError_STATUS{}), generators)

	return authorizationError_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationError_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationError_STATUS(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
}

func Test_AuthorizationProvidersAuthorization_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationProvidersAuthorization to hub returns original",
		prop.ForAll(RunResourceConversionTestForAuthorizationProvidersAuthorization, AuthorizationProvidersAuthorizationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForAuthorizationProvidersAuthorization tests if a specific instance of AuthorizationProvidersAuthorization round trips to the hub storage version and back losslessly
func RunResourceConversionTestForAuthorizationProvidersAuthorization(subject AuthorizationProvidersAuthorization) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20220801s.AuthorizationProvidersAuthorization
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual AuthorizationProvidersAuthorization
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AuthorizationProvidersAuthorization_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationProvidersAuthorization to AuthorizationProvidersAuthorization via AssignProperties_To_AuthorizationProvidersAuthorization & AssignProperties_From_AuthorizationProvidersAuthorization returns original",
		prop.ForAll(RunPropertyAssignmentTestForAuthorizationProvidersAuthorization, AuthorizationProvidersAuthorizationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAuthorizationProvidersAuthorization tests if a specific instance of AuthorizationProvidersAuthorization can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAuthorizationProvidersAuthorization(subject AuthorizationProvidersAuthorization) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.AuthorizationProvidersAuthorization
	err := copied.AssignProperties_To_AuthorizationProvidersAuthorization(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AuthorizationProvidersAuthorization
	err = actual.AssignProperties_From_AuthorizationProvidersAuthorization(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AuthorizationProvidersAuthorization_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvidersAuthorization via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvidersAuthorization, AuthorizationProvidersAuthorizationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvidersAuthorization runs a test to see if a specific instance of AuthorizationProvidersAuthorization round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvidersAuthorization(subject AuthorizationProvidersAuthorization) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvidersAuthorization
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AuthorizationProvidersAuthorization instances for property testing - lazily instantiated by
// AuthorizationProvidersAuthorizationGenerator()
var authorizationProvidersAuthorizationGenerator gopter.Gen

// AuthorizationProvidersAuthorizationGenerator returns a generator of AuthorizationProvidersAuthorization instances for property testing.
func AuthorizationProvidersAuthorizationGenerator() gopter.Gen {
	if authorizationProvidersAuthorizationGenerator != nil {
		return authorizationProvidersAuthorizationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization(generators)
	authorizationProvidersAuthorizationGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorization{}), generators)

	return authorizationProvidersAuthorizationGenerator
}

// AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization(gens map[string]gopter.Gen) {
	gens["Spec"] = AuthorizationProvidersAuthorization_SpecGenerator()
	gens["Status"] = AuthorizationProvidersAuthorization_STATUSGenerator()
}

func Test_AuthorizationProvidersAuthorizationOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationProvidersAuthorizationOperatorSpec to AuthorizationProvidersAuthorizationOperatorSpec via AssignProperties_To_AuthorizationProvidersAuthorizationOperatorSpec & AssignProperties_From_AuthorizationProvidersAuthorizationOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForAuthorizationProvidersAuthorizationOperatorSpec, AuthorizationProvidersAuthorizationOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAuthorizationProvidersAuthorizationOperatorSpec tests if a specific instance of AuthorizationProvidersAuthorizationOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAuthorizationProvidersAuthorizationOperatorSpec(subject AuthorizationProvidersAuthorizationOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.AuthorizationProvidersAuthorizationOperatorSpec
	err := copied.AssignProperties_To_AuthorizationProvidersAuthorizationOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AuthorizationProvidersAuthorizationOperatorSpec
	err = actual.AssignProperties_From_AuthorizationProvidersAuthorizationOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AuthorizationProvidersAuthorizationOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvidersAuthorizationOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvidersAuthorizationOperatorSpec, AuthorizationProvidersAuthorizationOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvidersAuthorizationOperatorSpec runs a test to see if a specific instance of AuthorizationProvidersAuthorizationOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvidersAuthorizationOperatorSpec(subject AuthorizationProvidersAuthorizationOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvidersAuthorizationOperatorSpec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AuthorizationProvidersAuthorizationOperatorSpec instances for property testing - lazily instantiated by
// AuthorizationProvidersAuthorizationOperatorSpecGenerator()
var authorizationProvidersAuthorizationOperatorSpecGenerator gopter.Gen

// AuthorizationProvidersAuthorizationOperatorSpecGenerator returns a generator of AuthorizationProvidersAuthorizationOperatorSpec instances for property testing.
func AuthorizationProvidersAuthorizationOperatorSpecGenerator() gopter.Gen {
	if authorizationProvidersAuthorizationOperatorSpecGenerator != nil {
		return authorizationProvidersAuthorizationOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	authorizationProvidersAuthorizationOperatorSpecGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorizationOperatorSpec{}), generators)

	return authorizationProvidersAuthorizationOperatorSpecGenerator
}

func Test_AuthorizationProvidersAuthorization_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationProvidersAuthorization_STATUS to AuthorizationProvidersAuthorization_STATUS via AssignProperties_To_AuthorizationProvidersAuthorization_STATUS & AssignProperties_From_AuthorizationProvidersAuthorization_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForAuthorizationProvidersAuthorization_STATUS, AuthorizationProvidersAuthorization_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAuthorizationProvidersAuthorization_STATUS tests if a specific instance of AuthorizationProvidersAuthorization_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAuthorizationProvidersAuthorization_STATUS(subject AuthorizationProvidersAuthorization_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.AuthorizationProvidersAuthorization_STATUS
	err := copied.AssignProperties_To_AuthorizationProvidersAuthorization_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AuthorizationProvidersAuthorization_STATUS
	err = actual.AssignProperties_From_AuthorizationProvidersAuthorization_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AuthorizationProvidersAuthorization_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvidersAuthorization_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvidersAuthorization_STATUS, AuthorizationProvidersAuthorization_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvidersAuthorization_STATUS runs a test to see if a specific instance of AuthorizationProvidersAuthorization_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvidersAuthorization_STATUS(subject AuthorizationProvidersAuthorization_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvidersAuthorization_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AuthorizationProvidersAuthorization_STATUS instances for property testing - lazily instantiated by
// AuthorizationProvidersAuthorization_STATUSGenerator()
var authorizationProvidersAuthorization_STATUSGenerator gopter.Gen

// AuthorizationProvidersAuthorization_STATUSGenerator returns a generator of AuthorizationProvidersAuthorization_STATUS instances for property testing.
// We first initialize authorizationProvidersAuthorization_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationProvidersAuthorization_STATUSGenerator() gopter.Gen {
	if authorizationProvidersAuthorization_STATUSGenerator != nil {
		return authorizationProvidersAuthorization_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorization_STATUS(generators)
	authorizationProvidersAuthorization_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorization_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorization_STATUS(generators)
	AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization_STATUS(generators)
	authorizationProvidersAuthorization_STATUSGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorization_STATUS{}), generators)

	return authorizationProvidersAuthorization_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorization_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorization_STATUS(gens map[string]gopter.Gen) {
	gens["AuthorizationType"] = gen.PtrOf(gen.OneConstOf(AuthorizationContractProperties_AuthorizationType_STATUS_OAuth2))
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Oauth2GrantType"] = gen.PtrOf(gen.OneConstOf(AuthorizationContractProperties_Oauth2GrantType_STATUS_AuthorizationCode, AuthorizationContractProperties_Oauth2GrantType_STATUS_ClientCredentials))
	gens["Parameters"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization_STATUS(gens map[string]gopter.Gen) {
	gens["Error"] = gen.PtrOf(AuthorizationError_STATUSGenerator())
}

func Test_AuthorizationProvidersAuthorization_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AuthorizationProvidersAuthorization_Spec to AuthorizationProvidersAuthorization_Spec via AssignProperties_To_AuthorizationProvidersAuthorization_Spec & AssignProperties_From_AuthorizationProvidersAuthorization_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForAuthorizationProvidersAuthorization_Spec, AuthorizationProvidersAuthorization_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAuthorizationProvidersAuthorization_Spec tests if a specific instance of AuthorizationProvidersAuthorization_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForAuthorizationProvidersAuthorization_Spec(subject AuthorizationProvidersAuthorization_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230501ps.AuthorizationProvidersAuthorization_Spec
	err := copied.AssignProperties_To_AuthorizationProvidersAuthorization_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AuthorizationProvidersAuthorization_Spec
	err = actual.AssignProperties_From_AuthorizationProvidersAuthorization_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AuthorizationProvidersAuthorization_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AuthorizationProvidersAuthorization_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAuthorizationProvidersAuthorization_Spec, AuthorizationProvidersAuthorization_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAuthorizationProvidersAuthorization_Spec runs a test to see if a specific instance of AuthorizationProvidersAuthorization_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAuthorizationProvidersAuthorization_Spec(subject AuthorizationProvidersAuthorization_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AuthorizationProvidersAuthorization_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of AuthorizationProvidersAuthorization_Spec instances for property testing - lazily instantiated by
// AuthorizationProvidersAuthorization_SpecGenerator()
var authorizationProvidersAuthorization_SpecGenerator gopter.Gen

// AuthorizationProvidersAuthorization_SpecGenerator returns a generator of AuthorizationProvidersAuthorization_Spec instances for property testing.
// We first initialize authorizationProvidersAuthorization_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AuthorizationProvidersAuthorization_SpecGenerator() gopter.Gen {
	if authorizationProvidersAuthorization_SpecGenerator != nil {
		return authorizationProvidersAuthorization_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorization_Spec(generators)
	authorizationProvidersAuthorization_SpecGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorization_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorization_Spec(generators)
	AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization_Spec(generators)
	authorizationProvidersAuthorization_SpecGenerator = gen.Struct(reflect.TypeOf(AuthorizationProvidersAuthorization_Spec{}), generators)

	return authorizationProvidersAuthorization_SpecGenerator
}

// AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorization_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAuthorizationProvidersAuthorization_Spec(gens map[string]gopter.Gen) {
	gens["AuthorizationType"] = gen.PtrOf(gen.OneConstOf(AuthorizationContractProperties_AuthorizationType_OAuth2))
	gens["AzureName"] = gen.AlphaString()
	gens["Oauth2GrantType"] = gen.PtrOf(gen.OneConstOf(AuthorizationContractProperties_Oauth2GrantType_AuthorizationCode, AuthorizationContractProperties_Oauth2GrantType_ClientCredentials))
}

// AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAuthorizationProvidersAuthorization_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(AuthorizationProvidersAuthorizationOperatorSpecGenerator())
}
