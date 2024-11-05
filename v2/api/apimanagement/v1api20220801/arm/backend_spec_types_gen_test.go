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

func Test_BackendAuthorizationHeaderCredentials_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendAuthorizationHeaderCredentials via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendAuthorizationHeaderCredentials, BackendAuthorizationHeaderCredentialsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendAuthorizationHeaderCredentials runs a test to see if a specific instance of BackendAuthorizationHeaderCredentials round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendAuthorizationHeaderCredentials(subject BackendAuthorizationHeaderCredentials) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendAuthorizationHeaderCredentials
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

// Generator of BackendAuthorizationHeaderCredentials instances for property testing - lazily instantiated by
// BackendAuthorizationHeaderCredentialsGenerator()
var backendAuthorizationHeaderCredentialsGenerator gopter.Gen

// BackendAuthorizationHeaderCredentialsGenerator returns a generator of BackendAuthorizationHeaderCredentials instances for property testing.
func BackendAuthorizationHeaderCredentialsGenerator() gopter.Gen {
	if backendAuthorizationHeaderCredentialsGenerator != nil {
		return backendAuthorizationHeaderCredentialsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendAuthorizationHeaderCredentials(generators)
	backendAuthorizationHeaderCredentialsGenerator = gen.Struct(reflect.TypeOf(BackendAuthorizationHeaderCredentials{}), generators)

	return backendAuthorizationHeaderCredentialsGenerator
}

// AddIndependentPropertyGeneratorsForBackendAuthorizationHeaderCredentials is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendAuthorizationHeaderCredentials(gens map[string]gopter.Gen) {
	gens["Parameter"] = gen.PtrOf(gen.AlphaString())
	gens["Scheme"] = gen.PtrOf(gen.AlphaString())
}

func Test_BackendContractProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendContractProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendContractProperties, BackendContractPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendContractProperties runs a test to see if a specific instance of BackendContractProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendContractProperties(subject BackendContractProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendContractProperties
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

// Generator of BackendContractProperties instances for property testing - lazily instantiated by
// BackendContractPropertiesGenerator()
var backendContractPropertiesGenerator gopter.Gen

// BackendContractPropertiesGenerator returns a generator of BackendContractProperties instances for property testing.
// We first initialize backendContractPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackendContractPropertiesGenerator() gopter.Gen {
	if backendContractPropertiesGenerator != nil {
		return backendContractPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendContractProperties(generators)
	backendContractPropertiesGenerator = gen.Struct(reflect.TypeOf(BackendContractProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendContractProperties(generators)
	AddRelatedPropertyGeneratorsForBackendContractProperties(generators)
	backendContractPropertiesGenerator = gen.Struct(reflect.TypeOf(BackendContractProperties{}), generators)

	return backendContractPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForBackendContractProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendContractProperties(gens map[string]gopter.Gen) {
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(BackendContractProperties_Protocol_Http, BackendContractProperties_Protocol_Soap))
	gens["ResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Title"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBackendContractProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackendContractProperties(gens map[string]gopter.Gen) {
	gens["Credentials"] = gen.PtrOf(BackendCredentialsContractGenerator())
	gens["Properties"] = gen.PtrOf(BackendPropertiesGenerator())
	gens["Proxy"] = gen.PtrOf(BackendProxyContractGenerator())
	gens["Tls"] = gen.PtrOf(BackendTlsPropertiesGenerator())
}

func Test_BackendCredentialsContract_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendCredentialsContract via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendCredentialsContract, BackendCredentialsContractGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendCredentialsContract runs a test to see if a specific instance of BackendCredentialsContract round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendCredentialsContract(subject BackendCredentialsContract) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendCredentialsContract
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

// Generator of BackendCredentialsContract instances for property testing - lazily instantiated by
// BackendCredentialsContractGenerator()
var backendCredentialsContractGenerator gopter.Gen

// BackendCredentialsContractGenerator returns a generator of BackendCredentialsContract instances for property testing.
// We first initialize backendCredentialsContractGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackendCredentialsContractGenerator() gopter.Gen {
	if backendCredentialsContractGenerator != nil {
		return backendCredentialsContractGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendCredentialsContract(generators)
	backendCredentialsContractGenerator = gen.Struct(reflect.TypeOf(BackendCredentialsContract{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendCredentialsContract(generators)
	AddRelatedPropertyGeneratorsForBackendCredentialsContract(generators)
	backendCredentialsContractGenerator = gen.Struct(reflect.TypeOf(BackendCredentialsContract{}), generators)

	return backendCredentialsContractGenerator
}

// AddIndependentPropertyGeneratorsForBackendCredentialsContract is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendCredentialsContract(gens map[string]gopter.Gen) {
	gens["Certificate"] = gen.SliceOf(gen.AlphaString())
	gens["CertificateIds"] = gen.SliceOf(gen.AlphaString())
	gens["Header"] = gen.MapOf(
		gen.AlphaString(),
		gen.SliceOf(gen.AlphaString()))
	gens["Query"] = gen.MapOf(
		gen.AlphaString(),
		gen.SliceOf(gen.AlphaString()))
}

// AddRelatedPropertyGeneratorsForBackendCredentialsContract is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackendCredentialsContract(gens map[string]gopter.Gen) {
	gens["Authorization"] = gen.PtrOf(BackendAuthorizationHeaderCredentialsGenerator())
}

func Test_BackendProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendProperties, BackendPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendProperties runs a test to see if a specific instance of BackendProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendProperties(subject BackendProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendProperties
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

// Generator of BackendProperties instances for property testing - lazily instantiated by BackendPropertiesGenerator()
var backendPropertiesGenerator gopter.Gen

// BackendPropertiesGenerator returns a generator of BackendProperties instances for property testing.
func BackendPropertiesGenerator() gopter.Gen {
	if backendPropertiesGenerator != nil {
		return backendPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForBackendProperties(generators)
	backendPropertiesGenerator = gen.Struct(reflect.TypeOf(BackendProperties{}), generators)

	return backendPropertiesGenerator
}

// AddRelatedPropertyGeneratorsForBackendProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackendProperties(gens map[string]gopter.Gen) {
	gens["ServiceFabricCluster"] = gen.PtrOf(BackendServiceFabricClusterPropertiesGenerator())
}

func Test_BackendProxyContract_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendProxyContract via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendProxyContract, BackendProxyContractGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendProxyContract runs a test to see if a specific instance of BackendProxyContract round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendProxyContract(subject BackendProxyContract) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendProxyContract
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

// Generator of BackendProxyContract instances for property testing - lazily instantiated by
// BackendProxyContractGenerator()
var backendProxyContractGenerator gopter.Gen

// BackendProxyContractGenerator returns a generator of BackendProxyContract instances for property testing.
func BackendProxyContractGenerator() gopter.Gen {
	if backendProxyContractGenerator != nil {
		return backendProxyContractGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendProxyContract(generators)
	backendProxyContractGenerator = gen.Struct(reflect.TypeOf(BackendProxyContract{}), generators)

	return backendProxyContractGenerator
}

// AddIndependentPropertyGeneratorsForBackendProxyContract is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendProxyContract(gens map[string]gopter.Gen) {
	gens["Password"] = gen.PtrOf(gen.AlphaString())
	gens["Url"] = gen.PtrOf(gen.AlphaString())
	gens["Username"] = gen.PtrOf(gen.AlphaString())
}

func Test_BackendServiceFabricClusterProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendServiceFabricClusterProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendServiceFabricClusterProperties, BackendServiceFabricClusterPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendServiceFabricClusterProperties runs a test to see if a specific instance of BackendServiceFabricClusterProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendServiceFabricClusterProperties(subject BackendServiceFabricClusterProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendServiceFabricClusterProperties
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

// Generator of BackendServiceFabricClusterProperties instances for property testing - lazily instantiated by
// BackendServiceFabricClusterPropertiesGenerator()
var backendServiceFabricClusterPropertiesGenerator gopter.Gen

// BackendServiceFabricClusterPropertiesGenerator returns a generator of BackendServiceFabricClusterProperties instances for property testing.
// We first initialize backendServiceFabricClusterPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackendServiceFabricClusterPropertiesGenerator() gopter.Gen {
	if backendServiceFabricClusterPropertiesGenerator != nil {
		return backendServiceFabricClusterPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendServiceFabricClusterProperties(generators)
	backendServiceFabricClusterPropertiesGenerator = gen.Struct(reflect.TypeOf(BackendServiceFabricClusterProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendServiceFabricClusterProperties(generators)
	AddRelatedPropertyGeneratorsForBackendServiceFabricClusterProperties(generators)
	backendServiceFabricClusterPropertiesGenerator = gen.Struct(reflect.TypeOf(BackendServiceFabricClusterProperties{}), generators)

	return backendServiceFabricClusterPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForBackendServiceFabricClusterProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendServiceFabricClusterProperties(gens map[string]gopter.Gen) {
	gens["ClientCertificateId"] = gen.PtrOf(gen.AlphaString())
	gens["ClientCertificatethumbprint"] = gen.PtrOf(gen.AlphaString())
	gens["ManagementEndpoints"] = gen.SliceOf(gen.AlphaString())
	gens["MaxPartitionResolutionRetries"] = gen.PtrOf(gen.Int())
	gens["ServerCertificateThumbprints"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBackendServiceFabricClusterProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackendServiceFabricClusterProperties(gens map[string]gopter.Gen) {
	gens["ServerX509Names"] = gen.SliceOf(X509CertificateNameGenerator())
}

func Test_BackendTlsProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackendTlsProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackendTlsProperties, BackendTlsPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackendTlsProperties runs a test to see if a specific instance of BackendTlsProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForBackendTlsProperties(subject BackendTlsProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackendTlsProperties
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

// Generator of BackendTlsProperties instances for property testing - lazily instantiated by
// BackendTlsPropertiesGenerator()
var backendTlsPropertiesGenerator gopter.Gen

// BackendTlsPropertiesGenerator returns a generator of BackendTlsProperties instances for property testing.
func BackendTlsPropertiesGenerator() gopter.Gen {
	if backendTlsPropertiesGenerator != nil {
		return backendTlsPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackendTlsProperties(generators)
	backendTlsPropertiesGenerator = gen.Struct(reflect.TypeOf(BackendTlsProperties{}), generators)

	return backendTlsPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForBackendTlsProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackendTlsProperties(gens map[string]gopter.Gen) {
	gens["ValidateCertificateChain"] = gen.PtrOf(gen.Bool())
	gens["ValidateCertificateName"] = gen.PtrOf(gen.Bool())
}

func Test_Backend_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Backend_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackend_Spec, Backend_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackend_Spec runs a test to see if a specific instance of Backend_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForBackend_Spec(subject Backend_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Backend_Spec
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

// Generator of Backend_Spec instances for property testing - lazily instantiated by Backend_SpecGenerator()
var backend_SpecGenerator gopter.Gen

// Backend_SpecGenerator returns a generator of Backend_Spec instances for property testing.
// We first initialize backend_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Backend_SpecGenerator() gopter.Gen {
	if backend_SpecGenerator != nil {
		return backend_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackend_Spec(generators)
	backend_SpecGenerator = gen.Struct(reflect.TypeOf(Backend_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackend_Spec(generators)
	AddRelatedPropertyGeneratorsForBackend_Spec(generators)
	backend_SpecGenerator = gen.Struct(reflect.TypeOf(Backend_Spec{}), generators)

	return backend_SpecGenerator
}

// AddIndependentPropertyGeneratorsForBackend_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackend_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForBackend_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackend_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(BackendContractPropertiesGenerator())
}

func Test_X509CertificateName_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of X509CertificateName via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForX509CertificateName, X509CertificateNameGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForX509CertificateName runs a test to see if a specific instance of X509CertificateName round trips to JSON and back losslessly
func RunJSONSerializationTestForX509CertificateName(subject X509CertificateName) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual X509CertificateName
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

// Generator of X509CertificateName instances for property testing - lazily instantiated by
// X509CertificateNameGenerator()
var x509CertificateNameGenerator gopter.Gen

// X509CertificateNameGenerator returns a generator of X509CertificateName instances for property testing.
func X509CertificateNameGenerator() gopter.Gen {
	if x509CertificateNameGenerator != nil {
		return x509CertificateNameGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForX509CertificateName(generators)
	x509CertificateNameGenerator = gen.Struct(reflect.TypeOf(X509CertificateName{}), generators)

	return x509CertificateNameGenerator
}

// AddIndependentPropertyGeneratorsForX509CertificateName is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForX509CertificateName(gens map[string]gopter.Gen) {
	gens["IssuerCertificateThumbprint"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}