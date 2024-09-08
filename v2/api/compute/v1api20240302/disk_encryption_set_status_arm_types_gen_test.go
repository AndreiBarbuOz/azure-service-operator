// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

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

func Test_ApiErrorBase_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiErrorBase_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiErrorBase_STATUS_ARM, ApiErrorBase_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiErrorBase_STATUS_ARM runs a test to see if a specific instance of ApiErrorBase_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApiErrorBase_STATUS_ARM(subject ApiErrorBase_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiErrorBase_STATUS_ARM
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

// Generator of ApiErrorBase_STATUS_ARM instances for property testing - lazily instantiated by
// ApiErrorBase_STATUS_ARMGenerator()
var apiErrorBase_STATUS_ARMGenerator gopter.Gen

// ApiErrorBase_STATUS_ARMGenerator returns a generator of ApiErrorBase_STATUS_ARM instances for property testing.
func ApiErrorBase_STATUS_ARMGenerator() gopter.Gen {
	if apiErrorBase_STATUS_ARMGenerator != nil {
		return apiErrorBase_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiErrorBase_STATUS_ARM(generators)
	apiErrorBase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ApiErrorBase_STATUS_ARM{}), generators)

	return apiErrorBase_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApiErrorBase_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiErrorBase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApiError_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApiError_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApiError_STATUS_ARM, ApiError_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApiError_STATUS_ARM runs a test to see if a specific instance of ApiError_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApiError_STATUS_ARM(subject ApiError_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApiError_STATUS_ARM
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

// Generator of ApiError_STATUS_ARM instances for property testing - lazily instantiated by
// ApiError_STATUS_ARMGenerator()
var apiError_STATUS_ARMGenerator gopter.Gen

// ApiError_STATUS_ARMGenerator returns a generator of ApiError_STATUS_ARM instances for property testing.
// We first initialize apiError_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApiError_STATUS_ARMGenerator() gopter.Gen {
	if apiError_STATUS_ARMGenerator != nil {
		return apiError_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiError_STATUS_ARM(generators)
	apiError_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ApiError_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApiError_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForApiError_STATUS_ARM(generators)
	apiError_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ApiError_STATUS_ARM{}), generators)

	return apiError_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApiError_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApiError_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.AlphaString())
	gens["Message"] = gen.PtrOf(gen.AlphaString())
	gens["Target"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApiError_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApiError_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Details"] = gen.SliceOf(ApiErrorBase_STATUS_ARMGenerator())
	gens["Innererror"] = gen.PtrOf(InnerError_STATUS_ARMGenerator())
}

func Test_DiskEncryptionSet_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiskEncryptionSet_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiskEncryptionSet_STATUS_ARM, DiskEncryptionSet_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiskEncryptionSet_STATUS_ARM runs a test to see if a specific instance of DiskEncryptionSet_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDiskEncryptionSet_STATUS_ARM(subject DiskEncryptionSet_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiskEncryptionSet_STATUS_ARM
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

// Generator of DiskEncryptionSet_STATUS_ARM instances for property testing - lazily instantiated by
// DiskEncryptionSet_STATUS_ARMGenerator()
var diskEncryptionSet_STATUS_ARMGenerator gopter.Gen

// DiskEncryptionSet_STATUS_ARMGenerator returns a generator of DiskEncryptionSet_STATUS_ARM instances for property testing.
// We first initialize diskEncryptionSet_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiskEncryptionSet_STATUS_ARMGenerator() gopter.Gen {
	if diskEncryptionSet_STATUS_ARMGenerator != nil {
		return diskEncryptionSet_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSet_STATUS_ARM(generators)
	diskEncryptionSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiskEncryptionSet_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDiskEncryptionSet_STATUS_ARM(generators)
	diskEncryptionSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DiskEncryptionSet_STATUS_ARM{}), generators)

	return diskEncryptionSet_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDiskEncryptionSet_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiskEncryptionSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiskEncryptionSet_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiskEncryptionSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(EncryptionSetIdentity_STATUS_ARMGenerator())
	gens["Properties"] = gen.PtrOf(EncryptionSetProperties_STATUS_ARMGenerator())
}

func Test_EncryptionSetIdentity_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSetIdentity_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSetIdentity_STATUS_ARM, EncryptionSetIdentity_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSetIdentity_STATUS_ARM runs a test to see if a specific instance of EncryptionSetIdentity_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSetIdentity_STATUS_ARM(subject EncryptionSetIdentity_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSetIdentity_STATUS_ARM
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

// Generator of EncryptionSetIdentity_STATUS_ARM instances for property testing - lazily instantiated by
// EncryptionSetIdentity_STATUS_ARMGenerator()
var encryptionSetIdentity_STATUS_ARMGenerator gopter.Gen

// EncryptionSetIdentity_STATUS_ARMGenerator returns a generator of EncryptionSetIdentity_STATUS_ARM instances for property testing.
// We first initialize encryptionSetIdentity_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionSetIdentity_STATUS_ARMGenerator() gopter.Gen {
	if encryptionSetIdentity_STATUS_ARMGenerator != nil {
		return encryptionSetIdentity_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity_STATUS_ARM(generators)
	encryptionSetIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionSetIdentity_STATUS_ARM(generators)
	encryptionSetIdentity_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity_STATUS_ARM{}), generators)

	return encryptionSetIdentity_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSetIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSetIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(
		EncryptionSetIdentity_Type_STATUS_ARM_None,
		EncryptionSetIdentity_Type_STATUS_ARM_SystemAssigned,
		EncryptionSetIdentity_Type_STATUS_ARM_SystemAssignedUserAssigned,
		EncryptionSetIdentity_Type_STATUS_ARM_UserAssigned))
}

// AddRelatedPropertyGeneratorsForEncryptionSetIdentity_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSetIdentity_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator())
}

func Test_EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM, EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM runs a test to see if a specific instance of EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM(subject EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM
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

// Generator of EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM instances for property testing - lazily
// instantiated by EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator()
var encryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator gopter.Gen

// EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator returns a generator of EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM instances for property testing.
func EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator() gopter.Gen {
	if encryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator != nil {
		return encryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM(generators)
	encryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM{}), generators)

	return encryptionSetIdentity_UserAssignedIdentities_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSetIdentity_UserAssignedIdentities_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
}

func Test_EncryptionSetProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of EncryptionSetProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEncryptionSetProperties_STATUS_ARM, EncryptionSetProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEncryptionSetProperties_STATUS_ARM runs a test to see if a specific instance of EncryptionSetProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEncryptionSetProperties_STATUS_ARM(subject EncryptionSetProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual EncryptionSetProperties_STATUS_ARM
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

// Generator of EncryptionSetProperties_STATUS_ARM instances for property testing - lazily instantiated by
// EncryptionSetProperties_STATUS_ARMGenerator()
var encryptionSetProperties_STATUS_ARMGenerator gopter.Gen

// EncryptionSetProperties_STATUS_ARMGenerator returns a generator of EncryptionSetProperties_STATUS_ARM instances for property testing.
// We first initialize encryptionSetProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EncryptionSetProperties_STATUS_ARMGenerator() gopter.Gen {
	if encryptionSetProperties_STATUS_ARMGenerator != nil {
		return encryptionSetProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetProperties_STATUS_ARM(generators)
	encryptionSetProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEncryptionSetProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForEncryptionSetProperties_STATUS_ARM(generators)
	encryptionSetProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(EncryptionSetProperties_STATUS_ARM{}), generators)

	return encryptionSetProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForEncryptionSetProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEncryptionSetProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["EncryptionType"] = gen.PtrOf(gen.OneConstOf(DiskEncryptionSetType_STATUS_ARM_ConfidentialVmEncryptedWithCustomerKey, DiskEncryptionSetType_STATUS_ARM_EncryptionAtRestWithCustomerKey, DiskEncryptionSetType_STATUS_ARM_EncryptionAtRestWithPlatformAndCustomerKeys))
	gens["FederatedClientId"] = gen.PtrOf(gen.AlphaString())
	gens["LastKeyRotationTimestamp"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["RotationToLatestKeyVersionEnabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForEncryptionSetProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEncryptionSetProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ActiveKey"] = gen.PtrOf(KeyForDiskEncryptionSet_STATUS_ARMGenerator())
	gens["AutoKeyRotationError"] = gen.PtrOf(ApiError_STATUS_ARMGenerator())
	gens["PreviousKeys"] = gen.SliceOf(KeyForDiskEncryptionSet_STATUS_ARMGenerator())
}

func Test_InnerError_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InnerError_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInnerError_STATUS_ARM, InnerError_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInnerError_STATUS_ARM runs a test to see if a specific instance of InnerError_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInnerError_STATUS_ARM(subject InnerError_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InnerError_STATUS_ARM
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

// Generator of InnerError_STATUS_ARM instances for property testing - lazily instantiated by
// InnerError_STATUS_ARMGenerator()
var innerError_STATUS_ARMGenerator gopter.Gen

// InnerError_STATUS_ARMGenerator returns a generator of InnerError_STATUS_ARM instances for property testing.
func InnerError_STATUS_ARMGenerator() gopter.Gen {
	if innerError_STATUS_ARMGenerator != nil {
		return innerError_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInnerError_STATUS_ARM(generators)
	innerError_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(InnerError_STATUS_ARM{}), generators)

	return innerError_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForInnerError_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInnerError_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Errordetail"] = gen.PtrOf(gen.AlphaString())
	gens["Exceptiontype"] = gen.PtrOf(gen.AlphaString())
}

func Test_KeyForDiskEncryptionSet_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KeyForDiskEncryptionSet_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKeyForDiskEncryptionSet_STATUS_ARM, KeyForDiskEncryptionSet_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKeyForDiskEncryptionSet_STATUS_ARM runs a test to see if a specific instance of KeyForDiskEncryptionSet_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKeyForDiskEncryptionSet_STATUS_ARM(subject KeyForDiskEncryptionSet_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KeyForDiskEncryptionSet_STATUS_ARM
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

// Generator of KeyForDiskEncryptionSet_STATUS_ARM instances for property testing - lazily instantiated by
// KeyForDiskEncryptionSet_STATUS_ARMGenerator()
var keyForDiskEncryptionSet_STATUS_ARMGenerator gopter.Gen

// KeyForDiskEncryptionSet_STATUS_ARMGenerator returns a generator of KeyForDiskEncryptionSet_STATUS_ARM instances for property testing.
// We first initialize keyForDiskEncryptionSet_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func KeyForDiskEncryptionSet_STATUS_ARMGenerator() gopter.Gen {
	if keyForDiskEncryptionSet_STATUS_ARMGenerator != nil {
		return keyForDiskEncryptionSet_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS_ARM(generators)
	keyForDiskEncryptionSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(KeyForDiskEncryptionSet_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS_ARM(generators)
	keyForDiskEncryptionSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(KeyForDiskEncryptionSet_STATUS_ARM{}), generators)

	return keyForDiskEncryptionSet_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["KeyUrl"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForKeyForDiskEncryptionSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["SourceVault"] = gen.PtrOf(SourceVault_STATUS_ARMGenerator())
}

func Test_SourceVault_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SourceVault_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSourceVault_STATUS_ARM, SourceVault_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSourceVault_STATUS_ARM runs a test to see if a specific instance of SourceVault_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSourceVault_STATUS_ARM(subject SourceVault_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SourceVault_STATUS_ARM
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

// Generator of SourceVault_STATUS_ARM instances for property testing - lazily instantiated by
// SourceVault_STATUS_ARMGenerator()
var sourceVault_STATUS_ARMGenerator gopter.Gen

// SourceVault_STATUS_ARMGenerator returns a generator of SourceVault_STATUS_ARM instances for property testing.
func SourceVault_STATUS_ARMGenerator() gopter.Gen {
	if sourceVault_STATUS_ARMGenerator != nil {
		return sourceVault_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSourceVault_STATUS_ARM(generators)
	sourceVault_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(SourceVault_STATUS_ARM{}), generators)

	return sourceVault_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSourceVault_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSourceVault_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}