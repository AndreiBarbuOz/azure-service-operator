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

func Test_BastionHostIPConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHostIPConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHostIPConfiguration, BastionHostIPConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHostIPConfiguration runs a test to see if a specific instance of BastionHostIPConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHostIPConfiguration(subject BastionHostIPConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHostIPConfiguration
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

// Generator of BastionHostIPConfiguration instances for property testing - lazily instantiated by
// BastionHostIPConfigurationGenerator()
var bastionHostIPConfigurationGenerator gopter.Gen

// BastionHostIPConfigurationGenerator returns a generator of BastionHostIPConfiguration instances for property testing.
// We first initialize bastionHostIPConfigurationGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BastionHostIPConfigurationGenerator() gopter.Gen {
	if bastionHostIPConfigurationGenerator != nil {
		return bastionHostIPConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostIPConfiguration(generators)
	bastionHostIPConfigurationGenerator = gen.Struct(reflect.TypeOf(BastionHostIPConfiguration{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostIPConfiguration(generators)
	AddRelatedPropertyGeneratorsForBastionHostIPConfiguration(generators)
	bastionHostIPConfigurationGenerator = gen.Struct(reflect.TypeOf(BastionHostIPConfiguration{}), generators)

	return bastionHostIPConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForBastionHostIPConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHostIPConfiguration(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBastionHostIPConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBastionHostIPConfiguration(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BastionHostIPConfigurationPropertiesFormatGenerator())
}

func Test_BastionHostIPConfigurationPropertiesFormat_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHostIPConfigurationPropertiesFormat via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHostIPConfigurationPropertiesFormat, BastionHostIPConfigurationPropertiesFormatGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHostIPConfigurationPropertiesFormat runs a test to see if a specific instance of BastionHostIPConfigurationPropertiesFormat round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHostIPConfigurationPropertiesFormat(subject BastionHostIPConfigurationPropertiesFormat) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHostIPConfigurationPropertiesFormat
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

// Generator of BastionHostIPConfigurationPropertiesFormat instances for property testing - lazily instantiated by
// BastionHostIPConfigurationPropertiesFormatGenerator()
var bastionHostIPConfigurationPropertiesFormatGenerator gopter.Gen

// BastionHostIPConfigurationPropertiesFormatGenerator returns a generator of BastionHostIPConfigurationPropertiesFormat instances for property testing.
// We first initialize bastionHostIPConfigurationPropertiesFormatGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BastionHostIPConfigurationPropertiesFormatGenerator() gopter.Gen {
	if bastionHostIPConfigurationPropertiesFormatGenerator != nil {
		return bastionHostIPConfigurationPropertiesFormatGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostIPConfigurationPropertiesFormat(generators)
	bastionHostIPConfigurationPropertiesFormatGenerator = gen.Struct(reflect.TypeOf(BastionHostIPConfigurationPropertiesFormat{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostIPConfigurationPropertiesFormat(generators)
	AddRelatedPropertyGeneratorsForBastionHostIPConfigurationPropertiesFormat(generators)
	bastionHostIPConfigurationPropertiesFormatGenerator = gen.Struct(reflect.TypeOf(BastionHostIPConfigurationPropertiesFormat{}), generators)

	return bastionHostIPConfigurationPropertiesFormatGenerator
}

// AddIndependentPropertyGeneratorsForBastionHostIPConfigurationPropertiesFormat is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHostIPConfigurationPropertiesFormat(gens map[string]gopter.Gen) {
	gens["PrivateIPAllocationMethod"] = gen.PtrOf(gen.OneConstOf(IPAllocationMethod_Dynamic, IPAllocationMethod_Static))
}

// AddRelatedPropertyGeneratorsForBastionHostIPConfigurationPropertiesFormat is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBastionHostIPConfigurationPropertiesFormat(gens map[string]gopter.Gen) {
	gens["PublicIPAddress"] = gen.PtrOf(BastionHostSubResourceGenerator())
	gens["Subnet"] = gen.PtrOf(BastionHostSubResourceGenerator())
}

func Test_BastionHostPropertiesFormat_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHostPropertiesFormat via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHostPropertiesFormat, BastionHostPropertiesFormatGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHostPropertiesFormat runs a test to see if a specific instance of BastionHostPropertiesFormat round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHostPropertiesFormat(subject BastionHostPropertiesFormat) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHostPropertiesFormat
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

// Generator of BastionHostPropertiesFormat instances for property testing - lazily instantiated by
// BastionHostPropertiesFormatGenerator()
var bastionHostPropertiesFormatGenerator gopter.Gen

// BastionHostPropertiesFormatGenerator returns a generator of BastionHostPropertiesFormat instances for property testing.
// We first initialize bastionHostPropertiesFormatGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BastionHostPropertiesFormatGenerator() gopter.Gen {
	if bastionHostPropertiesFormatGenerator != nil {
		return bastionHostPropertiesFormatGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostPropertiesFormat(generators)
	bastionHostPropertiesFormatGenerator = gen.Struct(reflect.TypeOf(BastionHostPropertiesFormat{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostPropertiesFormat(generators)
	AddRelatedPropertyGeneratorsForBastionHostPropertiesFormat(generators)
	bastionHostPropertiesFormatGenerator = gen.Struct(reflect.TypeOf(BastionHostPropertiesFormat{}), generators)

	return bastionHostPropertiesFormatGenerator
}

// AddIndependentPropertyGeneratorsForBastionHostPropertiesFormat is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHostPropertiesFormat(gens map[string]gopter.Gen) {
	gens["DisableCopyPaste"] = gen.PtrOf(gen.Bool())
	gens["DnsName"] = gen.PtrOf(gen.AlphaString())
	gens["EnableFileCopy"] = gen.PtrOf(gen.Bool())
	gens["EnableIpConnect"] = gen.PtrOf(gen.Bool())
	gens["EnableShareableLink"] = gen.PtrOf(gen.Bool())
	gens["EnableTunneling"] = gen.PtrOf(gen.Bool())
	gens["ScaleUnits"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForBastionHostPropertiesFormat is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBastionHostPropertiesFormat(gens map[string]gopter.Gen) {
	gens["IpConfigurations"] = gen.SliceOf(BastionHostIPConfigurationGenerator())
}

func Test_BastionHostSubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHostSubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHostSubResource, BastionHostSubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHostSubResource runs a test to see if a specific instance of BastionHostSubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHostSubResource(subject BastionHostSubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHostSubResource
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

// Generator of BastionHostSubResource instances for property testing - lazily instantiated by
// BastionHostSubResourceGenerator()
var bastionHostSubResourceGenerator gopter.Gen

// BastionHostSubResourceGenerator returns a generator of BastionHostSubResource instances for property testing.
func BastionHostSubResourceGenerator() gopter.Gen {
	if bastionHostSubResourceGenerator != nil {
		return bastionHostSubResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHostSubResource(generators)
	bastionHostSubResourceGenerator = gen.Struct(reflect.TypeOf(BastionHostSubResource{}), generators)

	return bastionHostSubResourceGenerator
}

// AddIndependentPropertyGeneratorsForBastionHostSubResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHostSubResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_BastionHost_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BastionHost_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBastionHost_Spec, BastionHost_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBastionHost_Spec runs a test to see if a specific instance of BastionHost_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForBastionHost_Spec(subject BastionHost_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BastionHost_Spec
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

// Generator of BastionHost_Spec instances for property testing - lazily instantiated by BastionHost_SpecGenerator()
var bastionHost_SpecGenerator gopter.Gen

// BastionHost_SpecGenerator returns a generator of BastionHost_Spec instances for property testing.
// We first initialize bastionHost_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BastionHost_SpecGenerator() gopter.Gen {
	if bastionHost_SpecGenerator != nil {
		return bastionHost_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHost_Spec(generators)
	bastionHost_SpecGenerator = gen.Struct(reflect.TypeOf(BastionHost_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBastionHost_Spec(generators)
	AddRelatedPropertyGeneratorsForBastionHost_Spec(generators)
	bastionHost_SpecGenerator = gen.Struct(reflect.TypeOf(BastionHost_Spec{}), generators)

	return bastionHost_SpecGenerator
}

// AddIndependentPropertyGeneratorsForBastionHost_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBastionHost_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBastionHost_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBastionHost_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BastionHostPropertiesFormatGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(Sku_Name_Basic, Sku_Name_Standard))
}