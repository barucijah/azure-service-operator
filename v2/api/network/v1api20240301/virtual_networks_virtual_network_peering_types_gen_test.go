// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240301

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301/storage"
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

func Test_VirtualNetworksVirtualNetworkPeering_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering to hub returns original",
		prop.ForAll(RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering tests if a specific instance of VirtualNetworksVirtualNetworkPeering round trips to the hub storage version and back losslessly
func RunResourceConversionTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.VirtualNetworksVirtualNetworkPeering
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual VirtualNetworksVirtualNetworkPeering
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

func Test_VirtualNetworksVirtualNetworkPeering_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering to VirtualNetworksVirtualNetworkPeering via AssignProperties_To_VirtualNetworksVirtualNetworkPeering & AssignProperties_From_VirtualNetworksVirtualNetworkPeering returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering tests if a specific instance of VirtualNetworksVirtualNetworkPeering can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.VirtualNetworksVirtualNetworkPeering
	err := copied.AssignProperties_To_VirtualNetworksVirtualNetworkPeering(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworksVirtualNetworkPeering
	err = actual.AssignProperties_From_VirtualNetworksVirtualNetworkPeering(&other)
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

func Test_VirtualNetworksVirtualNetworkPeering_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeering via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering, VirtualNetworksVirtualNetworkPeeringGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeering round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering(subject VirtualNetworksVirtualNetworkPeering) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeering
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

// Generator of VirtualNetworksVirtualNetworkPeering instances for property testing - lazily instantiated by
// VirtualNetworksVirtualNetworkPeeringGenerator()
var virtualNetworksVirtualNetworkPeeringGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeeringGenerator returns a generator of VirtualNetworksVirtualNetworkPeering instances for property testing.
func VirtualNetworksVirtualNetworkPeeringGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeeringGenerator != nil {
		return virtualNetworksVirtualNetworkPeeringGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering(generators)
	virtualNetworksVirtualNetworkPeeringGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering{}), generators)

	return virtualNetworksVirtualNetworkPeeringGenerator
}

// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering(gens map[string]gopter.Gen) {
	gens["Spec"] = VirtualNetworksVirtualNetworkPeering_SpecGenerator()
	gens["Status"] = VirtualNetworksVirtualNetworkPeering_STATUSGenerator()
}

func Test_VirtualNetworksVirtualNetworkPeeringOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeeringOperatorSpec to VirtualNetworksVirtualNetworkPeeringOperatorSpec via AssignProperties_To_VirtualNetworksVirtualNetworkPeeringOperatorSpec & AssignProperties_From_VirtualNetworksVirtualNetworkPeeringOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeeringOperatorSpec, VirtualNetworksVirtualNetworkPeeringOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeeringOperatorSpec tests if a specific instance of VirtualNetworksVirtualNetworkPeeringOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeeringOperatorSpec(subject VirtualNetworksVirtualNetworkPeeringOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.VirtualNetworksVirtualNetworkPeeringOperatorSpec
	err := copied.AssignProperties_To_VirtualNetworksVirtualNetworkPeeringOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworksVirtualNetworkPeeringOperatorSpec
	err = actual.AssignProperties_From_VirtualNetworksVirtualNetworkPeeringOperatorSpec(&other)
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

func Test_VirtualNetworksVirtualNetworkPeeringOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeeringOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringOperatorSpec, VirtualNetworksVirtualNetworkPeeringOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringOperatorSpec runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeeringOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeeringOperatorSpec(subject VirtualNetworksVirtualNetworkPeeringOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeeringOperatorSpec
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

// Generator of VirtualNetworksVirtualNetworkPeeringOperatorSpec instances for property testing - lazily instantiated by
// VirtualNetworksVirtualNetworkPeeringOperatorSpecGenerator()
var virtualNetworksVirtualNetworkPeeringOperatorSpecGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeeringOperatorSpecGenerator returns a generator of VirtualNetworksVirtualNetworkPeeringOperatorSpec instances for property testing.
func VirtualNetworksVirtualNetworkPeeringOperatorSpecGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeeringOperatorSpecGenerator != nil {
		return virtualNetworksVirtualNetworkPeeringOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	virtualNetworksVirtualNetworkPeeringOperatorSpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeeringOperatorSpec{}), generators)

	return virtualNetworksVirtualNetworkPeeringOperatorSpecGenerator
}

func Test_VirtualNetworksVirtualNetworkPeering_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering_STATUS to VirtualNetworksVirtualNetworkPeering_STATUS via AssignProperties_To_VirtualNetworksVirtualNetworkPeering_STATUS & AssignProperties_From_VirtualNetworksVirtualNetworkPeering_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_STATUS, VirtualNetworksVirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_STATUS tests if a specific instance of VirtualNetworksVirtualNetworkPeering_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_STATUS(subject VirtualNetworksVirtualNetworkPeering_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.VirtualNetworksVirtualNetworkPeering_STATUS
	err := copied.AssignProperties_To_VirtualNetworksVirtualNetworkPeering_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworksVirtualNetworkPeering_STATUS
	err = actual.AssignProperties_From_VirtualNetworksVirtualNetworkPeering_STATUS(&other)
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

func Test_VirtualNetworksVirtualNetworkPeering_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeering_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_STATUS, VirtualNetworksVirtualNetworkPeering_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_STATUS runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeering_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_STATUS(subject VirtualNetworksVirtualNetworkPeering_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeering_STATUS
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

// Generator of VirtualNetworksVirtualNetworkPeering_STATUS instances for property testing - lazily instantiated by
// VirtualNetworksVirtualNetworkPeering_STATUSGenerator()
var virtualNetworksVirtualNetworkPeering_STATUSGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeering_STATUSGenerator returns a generator of VirtualNetworksVirtualNetworkPeering_STATUS instances for property testing.
// We first initialize virtualNetworksVirtualNetworkPeering_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksVirtualNetworkPeering_STATUSGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeering_STATUSGenerator != nil {
		return virtualNetworksVirtualNetworkPeering_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(generators)
	virtualNetworksVirtualNetworkPeering_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(generators)
	virtualNetworksVirtualNetworkPeering_STATUSGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_STATUS{}), generators)

	return virtualNetworksVirtualNetworkPeering_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["EnableOnlyIPv6Peering"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LocalSubnetNames"] = gen.SliceOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PeerCompleteVnets"] = gen.PtrOf(gen.Bool())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Connected, VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Disconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_STATUS_Initiated))
	gens["PeeringSyncLevel"] = gen.PtrOf(gen.OneConstOf(
		VirtualNetworkPeeringPropertiesFormat_PeeringSyncLevel_STATUS_FullyInSync,
		VirtualNetworkPeeringPropertiesFormat_PeeringSyncLevel_STATUS_LocalAndRemoteNotInSync,
		VirtualNetworkPeeringPropertiesFormat_PeeringSyncLevel_STATUS_LocalNotInSync,
		VirtualNetworkPeeringPropertiesFormat_PeeringSyncLevel_STATUS_RemoteNotInSync))
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["RemoteSubnetNames"] = gen.SliceOf(gen.AlphaString())
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_STATUS(gens map[string]gopter.Gen) {
	gens["LocalAddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["LocalVirtualNetworkAddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunities_STATUSGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["RemoteVirtualNetworkAddressSpace"] = gen.PtrOf(AddressSpace_STATUSGenerator())
	gens["RemoteVirtualNetworkEncryption"] = gen.PtrOf(VirtualNetworkEncryption_STATUSGenerator())
}

func Test_VirtualNetworksVirtualNetworkPeering_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworksVirtualNetworkPeering_Spec to VirtualNetworksVirtualNetworkPeering_Spec via AssignProperties_To_VirtualNetworksVirtualNetworkPeering_Spec & AssignProperties_From_VirtualNetworksVirtualNetworkPeering_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_Spec, VirtualNetworksVirtualNetworkPeering_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_Spec tests if a specific instance of VirtualNetworksVirtualNetworkPeering_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksVirtualNetworkPeering_Spec(subject VirtualNetworksVirtualNetworkPeering_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.VirtualNetworksVirtualNetworkPeering_Spec
	err := copied.AssignProperties_To_VirtualNetworksVirtualNetworkPeering_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworksVirtualNetworkPeering_Spec
	err = actual.AssignProperties_From_VirtualNetworksVirtualNetworkPeering_Spec(&other)
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

func Test_VirtualNetworksVirtualNetworkPeering_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworksVirtualNetworkPeering_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_Spec, VirtualNetworksVirtualNetworkPeering_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_Spec runs a test to see if a specific instance of VirtualNetworksVirtualNetworkPeering_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksVirtualNetworkPeering_Spec(subject VirtualNetworksVirtualNetworkPeering_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworksVirtualNetworkPeering_Spec
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

// Generator of VirtualNetworksVirtualNetworkPeering_Spec instances for property testing - lazily instantiated by
// VirtualNetworksVirtualNetworkPeering_SpecGenerator()
var virtualNetworksVirtualNetworkPeering_SpecGenerator gopter.Gen

// VirtualNetworksVirtualNetworkPeering_SpecGenerator returns a generator of VirtualNetworksVirtualNetworkPeering_Spec instances for property testing.
// We first initialize virtualNetworksVirtualNetworkPeering_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksVirtualNetworkPeering_SpecGenerator() gopter.Gen {
	if virtualNetworksVirtualNetworkPeering_SpecGenerator != nil {
		return virtualNetworksVirtualNetworkPeering_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(generators)
	virtualNetworksVirtualNetworkPeering_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(generators)
	virtualNetworksVirtualNetworkPeering_SpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworksVirtualNetworkPeering_Spec{}), generators)

	return virtualNetworksVirtualNetworkPeering_SpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(gens map[string]gopter.Gen) {
	gens["AllowForwardedTraffic"] = gen.PtrOf(gen.Bool())
	gens["AllowGatewayTransit"] = gen.PtrOf(gen.Bool())
	gens["AllowVirtualNetworkAccess"] = gen.PtrOf(gen.Bool())
	gens["AzureName"] = gen.AlphaString()
	gens["DoNotVerifyRemoteGateways"] = gen.PtrOf(gen.Bool())
	gens["EnableOnlyIPv6Peering"] = gen.PtrOf(gen.Bool())
	gens["LocalSubnetNames"] = gen.SliceOf(gen.AlphaString())
	gens["PeerCompleteVnets"] = gen.PtrOf(gen.Bool())
	gens["PeeringState"] = gen.PtrOf(gen.OneConstOf(VirtualNetworkPeeringPropertiesFormat_PeeringState_Connected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Disconnected, VirtualNetworkPeeringPropertiesFormat_PeeringState_Initiated))
	gens["PeeringSyncLevel"] = gen.PtrOf(gen.OneConstOf(
		VirtualNetworkPeeringPropertiesFormat_PeeringSyncLevel_FullyInSync,
		VirtualNetworkPeeringPropertiesFormat_PeeringSyncLevel_LocalAndRemoteNotInSync,
		VirtualNetworkPeeringPropertiesFormat_PeeringSyncLevel_LocalNotInSync,
		VirtualNetworkPeeringPropertiesFormat_PeeringSyncLevel_RemoteNotInSync))
	gens["RemoteSubnetNames"] = gen.SliceOf(gen.AlphaString())
	gens["UseRemoteGateways"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksVirtualNetworkPeering_Spec(gens map[string]gopter.Gen) {
	gens["LocalAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["LocalVirtualNetworkAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["OperatorSpec"] = gen.PtrOf(VirtualNetworksVirtualNetworkPeeringOperatorSpecGenerator())
	gens["RemoteAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["RemoteBgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesGenerator())
	gens["RemoteVirtualNetwork"] = gen.PtrOf(SubResourceGenerator())
	gens["RemoteVirtualNetworkAddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
}
