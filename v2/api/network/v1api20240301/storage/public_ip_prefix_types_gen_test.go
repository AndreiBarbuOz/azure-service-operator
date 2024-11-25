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

func Test_NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded, NatGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded runs a test to see if a specific instance of NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewaySpec_PublicIPPrefix_SubResourceEmbedded(subject NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded
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

// Generator of NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded instances for property testing - lazily instantiated
// by NatGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator()
var natGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator gopter.Gen

// NatGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator returns a generator of NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded instances for property testing.
func NatGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator() gopter.Gen {
	if natGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator != nil {
		return natGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	natGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NatGatewaySpec_PublicIPPrefix_SubResourceEmbedded{}), generators)

	return natGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator
}

func Test_NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded, NatGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded runs a test to see if a specific instance of NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded(subject NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded
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

// Generator of NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded instances for property testing - lazily
// instantiated by NatGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator()
var natGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator gopter.Gen

// NatGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator returns a generator of NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded instances for property testing.
func NatGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator() gopter.Gen {
	if natGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator != nil {
		return natGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded(generators)
	natGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(NatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded{}), generators)

	return natGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForNatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGateway_STATUS_PublicIPPrefix_SubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPPrefix_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefix via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefix, PublicIPPrefixGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefix runs a test to see if a specific instance of PublicIPPrefix round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefix(subject PublicIPPrefix) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefix
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

// Generator of PublicIPPrefix instances for property testing - lazily instantiated by PublicIPPrefixGenerator()
var publicIPPrefixGenerator gopter.Gen

// PublicIPPrefixGenerator returns a generator of PublicIPPrefix instances for property testing.
func PublicIPPrefixGenerator() gopter.Gen {
	if publicIPPrefixGenerator != nil {
		return publicIPPrefixGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPublicIPPrefix(generators)
	publicIPPrefixGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefix{}), generators)

	return publicIPPrefixGenerator
}

// AddRelatedPropertyGeneratorsForPublicIPPrefix is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPPrefix(gens map[string]gopter.Gen) {
	gens["Spec"] = PublicIPPrefix_SpecGenerator()
	gens["Status"] = PublicIPPrefix_STATUSGenerator()
}

func Test_PublicIPPrefixOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefixOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefixOperatorSpec, PublicIPPrefixOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefixOperatorSpec runs a test to see if a specific instance of PublicIPPrefixOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefixOperatorSpec(subject PublicIPPrefixOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefixOperatorSpec
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

// Generator of PublicIPPrefixOperatorSpec instances for property testing - lazily instantiated by
// PublicIPPrefixOperatorSpecGenerator()
var publicIPPrefixOperatorSpecGenerator gopter.Gen

// PublicIPPrefixOperatorSpecGenerator returns a generator of PublicIPPrefixOperatorSpec instances for property testing.
func PublicIPPrefixOperatorSpecGenerator() gopter.Gen {
	if publicIPPrefixOperatorSpecGenerator != nil {
		return publicIPPrefixOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	publicIPPrefixOperatorSpecGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefixOperatorSpec{}), generators)

	return publicIPPrefixOperatorSpecGenerator
}

func Test_PublicIPPrefixSku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefixSku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefixSku, PublicIPPrefixSkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefixSku runs a test to see if a specific instance of PublicIPPrefixSku round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefixSku(subject PublicIPPrefixSku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefixSku
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

// Generator of PublicIPPrefixSku instances for property testing - lazily instantiated by PublicIPPrefixSkuGenerator()
var publicIPPrefixSkuGenerator gopter.Gen

// PublicIPPrefixSkuGenerator returns a generator of PublicIPPrefixSku instances for property testing.
func PublicIPPrefixSkuGenerator() gopter.Gen {
	if publicIPPrefixSkuGenerator != nil {
		return publicIPPrefixSkuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefixSku(generators)
	publicIPPrefixSkuGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefixSku{}), generators)

	return publicIPPrefixSkuGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPPrefixSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPPrefixSku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPPrefixSku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefixSku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefixSku_STATUS, PublicIPPrefixSku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefixSku_STATUS runs a test to see if a specific instance of PublicIPPrefixSku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefixSku_STATUS(subject PublicIPPrefixSku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefixSku_STATUS
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

// Generator of PublicIPPrefixSku_STATUS instances for property testing - lazily instantiated by
// PublicIPPrefixSku_STATUSGenerator()
var publicIPPrefixSku_STATUSGenerator gopter.Gen

// PublicIPPrefixSku_STATUSGenerator returns a generator of PublicIPPrefixSku_STATUS instances for property testing.
func PublicIPPrefixSku_STATUSGenerator() gopter.Gen {
	if publicIPPrefixSku_STATUSGenerator != nil {
		return publicIPPrefixSku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefixSku_STATUS(generators)
	publicIPPrefixSku_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefixSku_STATUS{}), generators)

	return publicIPPrefixSku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPPrefixSku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPPrefixSku_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tier"] = gen.PtrOf(gen.AlphaString())
}

func Test_PublicIPPrefix_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefix_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefix_STATUS, PublicIPPrefix_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefix_STATUS runs a test to see if a specific instance of PublicIPPrefix_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefix_STATUS(subject PublicIPPrefix_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefix_STATUS
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

// Generator of PublicIPPrefix_STATUS instances for property testing - lazily instantiated by
// PublicIPPrefix_STATUSGenerator()
var publicIPPrefix_STATUSGenerator gopter.Gen

// PublicIPPrefix_STATUSGenerator returns a generator of PublicIPPrefix_STATUS instances for property testing.
// We first initialize publicIPPrefix_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPPrefix_STATUSGenerator() gopter.Gen {
	if publicIPPrefix_STATUSGenerator != nil {
		return publicIPPrefix_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefix_STATUS(generators)
	publicIPPrefix_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefix_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefix_STATUS(generators)
	AddRelatedPropertyGeneratorsForPublicIPPrefix_STATUS(generators)
	publicIPPrefix_STATUSGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefix_STATUS{}), generators)

	return publicIPPrefix_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPPrefix_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPPrefix_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IpPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrefixLength"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicIPAddressVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPPrefix_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPPrefix_STATUS(gens map[string]gopter.Gen) {
	gens["CustomIPPrefix"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocation_STATUSGenerator())
	gens["IpTags"] = gen.SliceOf(IpTag_STATUSGenerator())
	gens["LoadBalancerFrontendIpConfiguration"] = gen.PtrOf(SubResource_STATUSGenerator())
	gens["NatGateway"] = gen.PtrOf(NatGateway_STATUS_PublicIPPrefix_SubResourceEmbeddedGenerator())
	gens["PublicIPAddresses"] = gen.SliceOf(ReferencedPublicIpAddress_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(PublicIPPrefixSku_STATUSGenerator())
}

func Test_PublicIPPrefix_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PublicIPPrefix_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPublicIPPrefix_Spec, PublicIPPrefix_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPublicIPPrefix_Spec runs a test to see if a specific instance of PublicIPPrefix_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPublicIPPrefix_Spec(subject PublicIPPrefix_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PublicIPPrefix_Spec
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

// Generator of PublicIPPrefix_Spec instances for property testing - lazily instantiated by
// PublicIPPrefix_SpecGenerator()
var publicIPPrefix_SpecGenerator gopter.Gen

// PublicIPPrefix_SpecGenerator returns a generator of PublicIPPrefix_Spec instances for property testing.
// We first initialize publicIPPrefix_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PublicIPPrefix_SpecGenerator() gopter.Gen {
	if publicIPPrefix_SpecGenerator != nil {
		return publicIPPrefix_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefix_Spec(generators)
	publicIPPrefix_SpecGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefix_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPublicIPPrefix_Spec(generators)
	AddRelatedPropertyGeneratorsForPublicIPPrefix_Spec(generators)
	publicIPPrefix_SpecGenerator = gen.Struct(reflect.TypeOf(PublicIPPrefix_Spec{}), generators)

	return publicIPPrefix_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPublicIPPrefix_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPublicIPPrefix_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PrefixLength"] = gen.PtrOf(gen.Int())
	gens["PublicIPAddressVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPublicIPPrefix_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPublicIPPrefix_Spec(gens map[string]gopter.Gen) {
	gens["CustomIPPrefix"] = gen.PtrOf(SubResourceGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["IpTags"] = gen.SliceOf(IpTagGenerator())
	gens["NatGateway"] = gen.PtrOf(NatGatewaySpec_PublicIPPrefix_SubResourceEmbeddedGenerator())
	gens["OperatorSpec"] = gen.PtrOf(PublicIPPrefixOperatorSpecGenerator())
	gens["Sku"] = gen.PtrOf(PublicIPPrefixSkuGenerator())
}

func Test_ReferencedPublicIpAddress_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ReferencedPublicIpAddress_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForReferencedPublicIpAddress_STATUS, ReferencedPublicIpAddress_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForReferencedPublicIpAddress_STATUS runs a test to see if a specific instance of ReferencedPublicIpAddress_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForReferencedPublicIpAddress_STATUS(subject ReferencedPublicIpAddress_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ReferencedPublicIpAddress_STATUS
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

// Generator of ReferencedPublicIpAddress_STATUS instances for property testing - lazily instantiated by
// ReferencedPublicIpAddress_STATUSGenerator()
var referencedPublicIpAddress_STATUSGenerator gopter.Gen

// ReferencedPublicIpAddress_STATUSGenerator returns a generator of ReferencedPublicIpAddress_STATUS instances for property testing.
func ReferencedPublicIpAddress_STATUSGenerator() gopter.Gen {
	if referencedPublicIpAddress_STATUSGenerator != nil {
		return referencedPublicIpAddress_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForReferencedPublicIpAddress_STATUS(generators)
	referencedPublicIpAddress_STATUSGenerator = gen.Struct(reflect.TypeOf(ReferencedPublicIpAddress_STATUS{}), generators)

	return referencedPublicIpAddress_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForReferencedPublicIpAddress_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForReferencedPublicIpAddress_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}