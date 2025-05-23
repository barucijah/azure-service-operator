// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
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

func Test_SqlDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabase, SqlDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabase runs a test to see if a specific instance of SqlDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabase(subject SqlDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabase
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

// Generator of SqlDatabase instances for property testing - lazily instantiated by SqlDatabaseGenerator()
var sqlDatabaseGenerator gopter.Gen

// SqlDatabaseGenerator returns a generator of SqlDatabase instances for property testing.
func SqlDatabaseGenerator() gopter.Gen {
	if sqlDatabaseGenerator != nil {
		return sqlDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabase(generators)
	sqlDatabaseGenerator = gen.Struct(reflect.TypeOf(SqlDatabase{}), generators)

	return sqlDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = SqlDatabase_SpecGenerator()
	gens["Status"] = SqlDatabase_STATUSGenerator()
}

func Test_SqlDatabaseGetProperties_Resource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseGetProperties_Resource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseGetProperties_Resource_STATUS, SqlDatabaseGetProperties_Resource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseGetProperties_Resource_STATUS runs a test to see if a specific instance of SqlDatabaseGetProperties_Resource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseGetProperties_Resource_STATUS(subject SqlDatabaseGetProperties_Resource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseGetProperties_Resource_STATUS
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

// Generator of SqlDatabaseGetProperties_Resource_STATUS instances for property testing - lazily instantiated by
// SqlDatabaseGetProperties_Resource_STATUSGenerator()
var sqlDatabaseGetProperties_Resource_STATUSGenerator gopter.Gen

// SqlDatabaseGetProperties_Resource_STATUSGenerator returns a generator of SqlDatabaseGetProperties_Resource_STATUS instances for property testing.
// We first initialize sqlDatabaseGetProperties_Resource_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseGetProperties_Resource_STATUSGenerator() gopter.Gen {
	if sqlDatabaseGetProperties_Resource_STATUSGenerator != nil {
		return sqlDatabaseGetProperties_Resource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS(generators)
	sqlDatabaseGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetProperties_Resource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS(generators)
	sqlDatabaseGetProperties_Resource_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetProperties_Resource_STATUS{}), generators)

	return sqlDatabaseGetProperties_Resource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["Colls"] = gen.PtrOf(gen.AlphaString())
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
	gens["Users"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseGetProperties_Resource_STATUS(gens map[string]gopter.Gen) {
	gens["RestoreParameters"] = gen.PtrOf(RestoreParametersBase_STATUSGenerator())
}

func Test_SqlDatabaseOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseOperatorSpec, SqlDatabaseOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseOperatorSpec runs a test to see if a specific instance of SqlDatabaseOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseOperatorSpec(subject SqlDatabaseOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseOperatorSpec
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

// Generator of SqlDatabaseOperatorSpec instances for property testing - lazily instantiated by
// SqlDatabaseOperatorSpecGenerator()
var sqlDatabaseOperatorSpecGenerator gopter.Gen

// SqlDatabaseOperatorSpecGenerator returns a generator of SqlDatabaseOperatorSpec instances for property testing.
func SqlDatabaseOperatorSpecGenerator() gopter.Gen {
	if sqlDatabaseOperatorSpecGenerator != nil {
		return sqlDatabaseOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	sqlDatabaseOperatorSpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseOperatorSpec{}), generators)

	return sqlDatabaseOperatorSpecGenerator
}

func Test_SqlDatabaseResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseResource, SqlDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseResource runs a test to see if a specific instance of SqlDatabaseResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseResource(subject SqlDatabaseResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseResource
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

// Generator of SqlDatabaseResource instances for property testing - lazily instantiated by
// SqlDatabaseResourceGenerator()
var sqlDatabaseResourceGenerator gopter.Gen

// SqlDatabaseResourceGenerator returns a generator of SqlDatabaseResource instances for property testing.
// We first initialize sqlDatabaseResourceGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseResourceGenerator() gopter.Gen {
	if sqlDatabaseResourceGenerator != nil {
		return sqlDatabaseResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseResource(generators)
	sqlDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseResource{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseResource(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseResource(generators)
	sqlDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseResource{}), generators)

	return sqlDatabaseResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseResource(gens map[string]gopter.Gen) {
	gens["CreateMode"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseResource is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseResource(gens map[string]gopter.Gen) {
	gens["RestoreParameters"] = gen.PtrOf(RestoreParametersBaseGenerator())
}

func Test_SqlDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabase_STATUS, SqlDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabase_STATUS runs a test to see if a specific instance of SqlDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabase_STATUS(subject SqlDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabase_STATUS
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

// Generator of SqlDatabase_STATUS instances for property testing - lazily instantiated by SqlDatabase_STATUSGenerator()
var sqlDatabase_STATUSGenerator gopter.Gen

// SqlDatabase_STATUSGenerator returns a generator of SqlDatabase_STATUS instances for property testing.
// We first initialize sqlDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabase_STATUSGenerator() gopter.Gen {
	if sqlDatabase_STATUSGenerator != nil {
		return sqlDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabase_STATUS(generators)
	sqlDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForSqlDatabase_STATUS(generators)
	sqlDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabase_STATUS{}), generators)

	return sqlDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResource_STATUSGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseGetProperties_Resource_STATUSGenerator())
}

func Test_SqlDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabase_Spec, SqlDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabase_Spec runs a test to see if a specific instance of SqlDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabase_Spec(subject SqlDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabase_Spec
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

// Generator of SqlDatabase_Spec instances for property testing - lazily instantiated by SqlDatabase_SpecGenerator()
var sqlDatabase_SpecGenerator gopter.Gen

// SqlDatabase_SpecGenerator returns a generator of SqlDatabase_Spec instances for property testing.
// We first initialize sqlDatabase_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabase_SpecGenerator() gopter.Gen {
	if sqlDatabase_SpecGenerator != nil {
		return sqlDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabase_Spec(generators)
	sqlDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabase_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForSqlDatabase_Spec(generators)
	sqlDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabase_Spec{}), generators)

	return sqlDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabase_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabase_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(SqlDatabaseOperatorSpecGenerator())
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseResourceGenerator())
}
