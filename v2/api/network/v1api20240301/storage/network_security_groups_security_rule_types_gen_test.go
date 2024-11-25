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

func Test_ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded, ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()
var applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator != nil {
		return applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded{}), generators)

	return applicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator
}

func Test_ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded, ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded runs a test to see if a specific instance of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(subject ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
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

// Generator of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instances for
// property testing - lazily instantiated by
// ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator()
var applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator gopter.Gen

// ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator returns a generator of ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instances for property testing.
func ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator() gopter.Gen {
	if applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator != nil {
		return applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(generators)
	applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded{}), generators)

	return applicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkSecurityGroupsSecurityRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule, NetworkSecurityGroupsSecurityRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule(subject NetworkSecurityGroupsSecurityRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule
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

// Generator of NetworkSecurityGroupsSecurityRule instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRuleGenerator()
var networkSecurityGroupsSecurityRuleGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRuleGenerator returns a generator of NetworkSecurityGroupsSecurityRule instances for property testing.
func NetworkSecurityGroupsSecurityRuleGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRuleGenerator != nil {
		return networkSecurityGroupsSecurityRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(generators)
	networkSecurityGroupsSecurityRuleGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule{}), generators)

	return networkSecurityGroupsSecurityRuleGenerator
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NetworkSecurityGroupsSecurityRule_SpecGenerator()
	gens["Status"] = NetworkSecurityGroupsSecurityRule_STATUSGenerator()
}

func Test_NetworkSecurityGroupsSecurityRuleOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRuleOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRuleOperatorSpec, NetworkSecurityGroupsSecurityRuleOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRuleOperatorSpec runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRuleOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRuleOperatorSpec(subject NetworkSecurityGroupsSecurityRuleOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRuleOperatorSpec
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

// Generator of NetworkSecurityGroupsSecurityRuleOperatorSpec instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRuleOperatorSpecGenerator()
var networkSecurityGroupsSecurityRuleOperatorSpecGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRuleOperatorSpecGenerator returns a generator of NetworkSecurityGroupsSecurityRuleOperatorSpec instances for property testing.
func NetworkSecurityGroupsSecurityRuleOperatorSpecGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRuleOperatorSpecGenerator != nil {
		return networkSecurityGroupsSecurityRuleOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	networkSecurityGroupsSecurityRuleOperatorSpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRuleOperatorSpec{}), generators)

	return networkSecurityGroupsSecurityRuleOperatorSpecGenerator
}

func Test_NetworkSecurityGroupsSecurityRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_STATUS, NetworkSecurityGroupsSecurityRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_STATUS runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_STATUS(subject NetworkSecurityGroupsSecurityRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule_STATUS
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

// Generator of NetworkSecurityGroupsSecurityRule_STATUS instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRule_STATUSGenerator()
var networkSecurityGroupsSecurityRule_STATUSGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRule_STATUSGenerator returns a generator of NetworkSecurityGroupsSecurityRule_STATUS instances for property testing.
// We first initialize networkSecurityGroupsSecurityRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSecurityRule_STATUSGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRule_STATUSGenerator != nil {
		return networkSecurityGroupsSecurityRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(generators)
	networkSecurityGroupsSecurityRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(generators)
	networkSecurityGroupsSecurityRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_STATUS{}), generators)

	return networkSecurityGroupsSecurityRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_STATUS(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator())
}

func Test_NetworkSecurityGroupsSecurityRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupsSecurityRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec, NetworkSecurityGroupsSecurityRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec runs a test to see if a specific instance of NetworkSecurityGroupsSecurityRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupsSecurityRule_Spec(subject NetworkSecurityGroupsSecurityRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupsSecurityRule_Spec
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

// Generator of NetworkSecurityGroupsSecurityRule_Spec instances for property testing - lazily instantiated by
// NetworkSecurityGroupsSecurityRule_SpecGenerator()
var networkSecurityGroupsSecurityRule_SpecGenerator gopter.Gen

// NetworkSecurityGroupsSecurityRule_SpecGenerator returns a generator of NetworkSecurityGroupsSecurityRule_Spec instances for property testing.
// We first initialize networkSecurityGroupsSecurityRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupsSecurityRule_SpecGenerator() gopter.Gen {
	if networkSecurityGroupsSecurityRule_SpecGenerator != nil {
		return networkSecurityGroupsSecurityRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	networkSecurityGroupsSecurityRule_SpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(generators)
	networkSecurityGroupsSecurityRule_SpecGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupsSecurityRule_Spec{}), generators)

	return networkSecurityGroupsSecurityRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(gens map[string]gopter.Gen) {
	gens["Access"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["DestinationPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["DestinationPortRanges"] = gen.SliceOf(gen.AlphaString())
	gens["Direction"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Priority"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["SourceAddressPrefixes"] = gen.SliceOf(gen.AlphaString())
	gens["SourcePortRange"] = gen.PtrOf(gen.AlphaString())
	gens["SourcePortRanges"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupsSecurityRule_Spec(gens map[string]gopter.Gen) {
	gens["DestinationApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator())
	gens["OperatorSpec"] = gen.PtrOf(NetworkSecurityGroupsSecurityRuleOperatorSpecGenerator())
	gens["SourceApplicationSecurityGroups"] = gen.SliceOf(ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbeddedGenerator())
}