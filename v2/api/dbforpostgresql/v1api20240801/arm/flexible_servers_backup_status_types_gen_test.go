// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_FlexibleServersBackup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersBackup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersBackup_STATUS, FlexibleServersBackup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersBackup_STATUS runs a test to see if a specific instance of FlexibleServersBackup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersBackup_STATUS(subject FlexibleServersBackup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersBackup_STATUS
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

// Generator of FlexibleServersBackup_STATUS instances for property testing - lazily instantiated by
// FlexibleServersBackup_STATUSGenerator()
var flexibleServersBackup_STATUSGenerator gopter.Gen

// FlexibleServersBackup_STATUSGenerator returns a generator of FlexibleServersBackup_STATUS instances for property testing.
// We first initialize flexibleServersBackup_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersBackup_STATUSGenerator() gopter.Gen {
	if flexibleServersBackup_STATUSGenerator != nil {
		return flexibleServersBackup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersBackup_STATUS(generators)
	flexibleServersBackup_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersBackup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersBackup_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersBackup_STATUS(generators)
	flexibleServersBackup_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersBackup_STATUS{}), generators)

	return flexibleServersBackup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersBackup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersBackup_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersBackup_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersBackup_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ServerBackupProperties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_ServerBackupProperties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ServerBackupProperties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForServerBackupProperties_STATUS, ServerBackupProperties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForServerBackupProperties_STATUS runs a test to see if a specific instance of ServerBackupProperties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForServerBackupProperties_STATUS(subject ServerBackupProperties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ServerBackupProperties_STATUS
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

// Generator of ServerBackupProperties_STATUS instances for property testing - lazily instantiated by
// ServerBackupProperties_STATUSGenerator()
var serverBackupProperties_STATUSGenerator gopter.Gen

// ServerBackupProperties_STATUSGenerator returns a generator of ServerBackupProperties_STATUS instances for property testing.
func ServerBackupProperties_STATUSGenerator() gopter.Gen {
	if serverBackupProperties_STATUSGenerator != nil {
		return serverBackupProperties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForServerBackupProperties_STATUS(generators)
	serverBackupProperties_STATUSGenerator = gen.Struct(reflect.TypeOf(ServerBackupProperties_STATUS{}), generators)

	return serverBackupProperties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForServerBackupProperties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForServerBackupProperties_STATUS(gens map[string]gopter.Gen) {
	gens["BackupType"] = gen.PtrOf(gen.OneConstOf(ServerBackupProperties_BackupType_STATUS_CustomerOnDemand, ServerBackupProperties_BackupType_STATUS_Full))
	gens["CompletedTime"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
}
