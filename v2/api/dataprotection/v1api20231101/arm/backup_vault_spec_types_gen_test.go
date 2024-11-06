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

func Test_AzureMonitorAlertSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AzureMonitorAlertSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAzureMonitorAlertSettings, AzureMonitorAlertSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAzureMonitorAlertSettings runs a test to see if a specific instance of AzureMonitorAlertSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForAzureMonitorAlertSettings(subject AzureMonitorAlertSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AzureMonitorAlertSettings
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

// Generator of AzureMonitorAlertSettings instances for property testing - lazily instantiated by
// AzureMonitorAlertSettingsGenerator()
var azureMonitorAlertSettingsGenerator gopter.Gen

// AzureMonitorAlertSettingsGenerator returns a generator of AzureMonitorAlertSettings instances for property testing.
func AzureMonitorAlertSettingsGenerator() gopter.Gen {
	if azureMonitorAlertSettingsGenerator != nil {
		return azureMonitorAlertSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings(generators)
	azureMonitorAlertSettingsGenerator = gen.Struct(reflect.TypeOf(AzureMonitorAlertSettings{}), generators)

	return azureMonitorAlertSettingsGenerator
}

// AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAzureMonitorAlertSettings(gens map[string]gopter.Gen) {
	gens["AlertsForAllJobFailures"] = gen.PtrOf(gen.OneConstOf(AzureMonitorAlertSettings_AlertsForAllJobFailures_Disabled, AzureMonitorAlertSettings_AlertsForAllJobFailures_Enabled))
}

func Test_BackupVaultSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupVaultSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupVaultSpec, BackupVaultSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupVaultSpec runs a test to see if a specific instance of BackupVaultSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupVaultSpec(subject BackupVaultSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupVaultSpec
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

// Generator of BackupVaultSpec instances for property testing - lazily instantiated by BackupVaultSpecGenerator()
var backupVaultSpecGenerator gopter.Gen

// BackupVaultSpecGenerator returns a generator of BackupVaultSpec instances for property testing.
// We first initialize backupVaultSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackupVaultSpecGenerator() gopter.Gen {
	if backupVaultSpecGenerator != nil {
		return backupVaultSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVaultSpec(generators)
	backupVaultSpecGenerator = gen.Struct(reflect.TypeOf(BackupVaultSpec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVaultSpec(generators)
	AddRelatedPropertyGeneratorsForBackupVaultSpec(generators)
	backupVaultSpecGenerator = gen.Struct(reflect.TypeOf(BackupVaultSpec{}), generators)

	return backupVaultSpecGenerator
}

// AddIndependentPropertyGeneratorsForBackupVaultSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupVaultSpec(gens map[string]gopter.Gen) {
	gens["ReplicatedRegions"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBackupVaultSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackupVaultSpec(gens map[string]gopter.Gen) {
	gens["FeatureSettings"] = gen.PtrOf(FeatureSettingsGenerator())
	gens["MonitoringSettings"] = gen.PtrOf(MonitoringSettingsGenerator())
	gens["SecuritySettings"] = gen.PtrOf(SecuritySettingsGenerator())
	gens["StorageSettings"] = gen.SliceOf(StorageSettingGenerator())
}

func Test_BackupVault_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of BackupVault_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForBackupVault_Spec, BackupVault_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForBackupVault_Spec runs a test to see if a specific instance of BackupVault_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForBackupVault_Spec(subject BackupVault_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual BackupVault_Spec
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

// Generator of BackupVault_Spec instances for property testing - lazily instantiated by BackupVault_SpecGenerator()
var backupVault_SpecGenerator gopter.Gen

// BackupVault_SpecGenerator returns a generator of BackupVault_Spec instances for property testing.
// We first initialize backupVault_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func BackupVault_SpecGenerator() gopter.Gen {
	if backupVault_SpecGenerator != nil {
		return backupVault_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVault_Spec(generators)
	backupVault_SpecGenerator = gen.Struct(reflect.TypeOf(BackupVault_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForBackupVault_Spec(generators)
	AddRelatedPropertyGeneratorsForBackupVault_Spec(generators)
	backupVault_SpecGenerator = gen.Struct(reflect.TypeOf(BackupVault_Spec{}), generators)

	return backupVault_SpecGenerator
}

// AddIndependentPropertyGeneratorsForBackupVault_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForBackupVault_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForBackupVault_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForBackupVault_Spec(gens map[string]gopter.Gen) {
	gens["Identity"] = gen.PtrOf(DppIdentityDetailsGenerator())
	gens["Properties"] = gen.PtrOf(BackupVaultSpecGenerator())
}

func Test_CrossRegionRestoreSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CrossRegionRestoreSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCrossRegionRestoreSettings, CrossRegionRestoreSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCrossRegionRestoreSettings runs a test to see if a specific instance of CrossRegionRestoreSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForCrossRegionRestoreSettings(subject CrossRegionRestoreSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CrossRegionRestoreSettings
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

// Generator of CrossRegionRestoreSettings instances for property testing - lazily instantiated by
// CrossRegionRestoreSettingsGenerator()
var crossRegionRestoreSettingsGenerator gopter.Gen

// CrossRegionRestoreSettingsGenerator returns a generator of CrossRegionRestoreSettings instances for property testing.
func CrossRegionRestoreSettingsGenerator() gopter.Gen {
	if crossRegionRestoreSettingsGenerator != nil {
		return crossRegionRestoreSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCrossRegionRestoreSettings(generators)
	crossRegionRestoreSettingsGenerator = gen.Struct(reflect.TypeOf(CrossRegionRestoreSettings{}), generators)

	return crossRegionRestoreSettingsGenerator
}

// AddIndependentPropertyGeneratorsForCrossRegionRestoreSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCrossRegionRestoreSettings(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(CrossRegionRestoreSettings_State_Disabled, CrossRegionRestoreSettings_State_Enabled))
}

func Test_CrossSubscriptionRestoreSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CrossSubscriptionRestoreSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCrossSubscriptionRestoreSettings, CrossSubscriptionRestoreSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCrossSubscriptionRestoreSettings runs a test to see if a specific instance of CrossSubscriptionRestoreSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForCrossSubscriptionRestoreSettings(subject CrossSubscriptionRestoreSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CrossSubscriptionRestoreSettings
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

// Generator of CrossSubscriptionRestoreSettings instances for property testing - lazily instantiated by
// CrossSubscriptionRestoreSettingsGenerator()
var crossSubscriptionRestoreSettingsGenerator gopter.Gen

// CrossSubscriptionRestoreSettingsGenerator returns a generator of CrossSubscriptionRestoreSettings instances for property testing.
func CrossSubscriptionRestoreSettingsGenerator() gopter.Gen {
	if crossSubscriptionRestoreSettingsGenerator != nil {
		return crossSubscriptionRestoreSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings(generators)
	crossSubscriptionRestoreSettingsGenerator = gen.Struct(reflect.TypeOf(CrossSubscriptionRestoreSettings{}), generators)

	return crossSubscriptionRestoreSettingsGenerator
}

// AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCrossSubscriptionRestoreSettings(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(CrossSubscriptionRestoreSettings_State_Disabled, CrossSubscriptionRestoreSettings_State_Enabled, CrossSubscriptionRestoreSettings_State_PermanentlyDisabled))
}

func Test_DppIdentityDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DppIdentityDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDppIdentityDetails, DppIdentityDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDppIdentityDetails runs a test to see if a specific instance of DppIdentityDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForDppIdentityDetails(subject DppIdentityDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DppIdentityDetails
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

// Generator of DppIdentityDetails instances for property testing - lazily instantiated by DppIdentityDetailsGenerator()
var dppIdentityDetailsGenerator gopter.Gen

// DppIdentityDetailsGenerator returns a generator of DppIdentityDetails instances for property testing.
// We first initialize dppIdentityDetailsGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DppIdentityDetailsGenerator() gopter.Gen {
	if dppIdentityDetailsGenerator != nil {
		return dppIdentityDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDppIdentityDetails(generators)
	dppIdentityDetailsGenerator = gen.Struct(reflect.TypeOf(DppIdentityDetails{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDppIdentityDetails(generators)
	AddRelatedPropertyGeneratorsForDppIdentityDetails(generators)
	dppIdentityDetailsGenerator = gen.Struct(reflect.TypeOf(DppIdentityDetails{}), generators)

	return dppIdentityDetailsGenerator
}

// AddIndependentPropertyGeneratorsForDppIdentityDetails is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDppIdentityDetails(gens map[string]gopter.Gen) {
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDppIdentityDetails is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDppIdentityDetails(gens map[string]gopter.Gen) {
	gens["UserAssignedIdentities"] = gen.MapOf(
		gen.AlphaString(),
		UserAssignedIdentityDetailsGenerator())
}

func Test_FeatureSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FeatureSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFeatureSettings, FeatureSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFeatureSettings runs a test to see if a specific instance of FeatureSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForFeatureSettings(subject FeatureSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FeatureSettings
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

// Generator of FeatureSettings instances for property testing - lazily instantiated by FeatureSettingsGenerator()
var featureSettingsGenerator gopter.Gen

// FeatureSettingsGenerator returns a generator of FeatureSettings instances for property testing.
func FeatureSettingsGenerator() gopter.Gen {
	if featureSettingsGenerator != nil {
		return featureSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFeatureSettings(generators)
	featureSettingsGenerator = gen.Struct(reflect.TypeOf(FeatureSettings{}), generators)

	return featureSettingsGenerator
}

// AddRelatedPropertyGeneratorsForFeatureSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFeatureSettings(gens map[string]gopter.Gen) {
	gens["CrossRegionRestoreSettings"] = gen.PtrOf(CrossRegionRestoreSettingsGenerator())
	gens["CrossSubscriptionRestoreSettings"] = gen.PtrOf(CrossSubscriptionRestoreSettingsGenerator())
}

func Test_ImmutabilitySettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ImmutabilitySettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForImmutabilitySettings, ImmutabilitySettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForImmutabilitySettings runs a test to see if a specific instance of ImmutabilitySettings round trips to JSON and back losslessly
func RunJSONSerializationTestForImmutabilitySettings(subject ImmutabilitySettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ImmutabilitySettings
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

// Generator of ImmutabilitySettings instances for property testing - lazily instantiated by
// ImmutabilitySettingsGenerator()
var immutabilitySettingsGenerator gopter.Gen

// ImmutabilitySettingsGenerator returns a generator of ImmutabilitySettings instances for property testing.
func ImmutabilitySettingsGenerator() gopter.Gen {
	if immutabilitySettingsGenerator != nil {
		return immutabilitySettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForImmutabilitySettings(generators)
	immutabilitySettingsGenerator = gen.Struct(reflect.TypeOf(ImmutabilitySettings{}), generators)

	return immutabilitySettingsGenerator
}

// AddIndependentPropertyGeneratorsForImmutabilitySettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForImmutabilitySettings(gens map[string]gopter.Gen) {
	gens["State"] = gen.PtrOf(gen.OneConstOf(ImmutabilitySettings_State_Disabled, ImmutabilitySettings_State_Locked, ImmutabilitySettings_State_Unlocked))
}

func Test_MonitoringSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MonitoringSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMonitoringSettings, MonitoringSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMonitoringSettings runs a test to see if a specific instance of MonitoringSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForMonitoringSettings(subject MonitoringSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MonitoringSettings
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

// Generator of MonitoringSettings instances for property testing - lazily instantiated by MonitoringSettingsGenerator()
var monitoringSettingsGenerator gopter.Gen

// MonitoringSettingsGenerator returns a generator of MonitoringSettings instances for property testing.
func MonitoringSettingsGenerator() gopter.Gen {
	if monitoringSettingsGenerator != nil {
		return monitoringSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForMonitoringSettings(generators)
	monitoringSettingsGenerator = gen.Struct(reflect.TypeOf(MonitoringSettings{}), generators)

	return monitoringSettingsGenerator
}

// AddRelatedPropertyGeneratorsForMonitoringSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMonitoringSettings(gens map[string]gopter.Gen) {
	gens["AzureMonitorAlertSettings"] = gen.PtrOf(AzureMonitorAlertSettingsGenerator())
}

func Test_SecuritySettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecuritySettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecuritySettings, SecuritySettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecuritySettings runs a test to see if a specific instance of SecuritySettings round trips to JSON and back losslessly
func RunJSONSerializationTestForSecuritySettings(subject SecuritySettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecuritySettings
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

// Generator of SecuritySettings instances for property testing - lazily instantiated by SecuritySettingsGenerator()
var securitySettingsGenerator gopter.Gen

// SecuritySettingsGenerator returns a generator of SecuritySettings instances for property testing.
func SecuritySettingsGenerator() gopter.Gen {
	if securitySettingsGenerator != nil {
		return securitySettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSecuritySettings(generators)
	securitySettingsGenerator = gen.Struct(reflect.TypeOf(SecuritySettings{}), generators)

	return securitySettingsGenerator
}

// AddRelatedPropertyGeneratorsForSecuritySettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSecuritySettings(gens map[string]gopter.Gen) {
	gens["ImmutabilitySettings"] = gen.PtrOf(ImmutabilitySettingsGenerator())
	gens["SoftDeleteSettings"] = gen.PtrOf(SoftDeleteSettingsGenerator())
}

func Test_SoftDeleteSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SoftDeleteSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSoftDeleteSettings, SoftDeleteSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSoftDeleteSettings runs a test to see if a specific instance of SoftDeleteSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForSoftDeleteSettings(subject SoftDeleteSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SoftDeleteSettings
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

// Generator of SoftDeleteSettings instances for property testing - lazily instantiated by SoftDeleteSettingsGenerator()
var softDeleteSettingsGenerator gopter.Gen

// SoftDeleteSettingsGenerator returns a generator of SoftDeleteSettings instances for property testing.
func SoftDeleteSettingsGenerator() gopter.Gen {
	if softDeleteSettingsGenerator != nil {
		return softDeleteSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSoftDeleteSettings(generators)
	softDeleteSettingsGenerator = gen.Struct(reflect.TypeOf(SoftDeleteSettings{}), generators)

	return softDeleteSettingsGenerator
}

// AddIndependentPropertyGeneratorsForSoftDeleteSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSoftDeleteSettings(gens map[string]gopter.Gen) {
	gens["RetentionDurationInDays"] = gen.PtrOf(gen.Float64())
	gens["State"] = gen.PtrOf(gen.OneConstOf(SoftDeleteSettings_State_AlwaysOn, SoftDeleteSettings_State_Off, SoftDeleteSettings_State_On))
}

func Test_StorageSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageSetting, StorageSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageSetting runs a test to see if a specific instance of StorageSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageSetting(subject StorageSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageSetting
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

// Generator of StorageSetting instances for property testing - lazily instantiated by StorageSettingGenerator()
var storageSettingGenerator gopter.Gen

// StorageSettingGenerator returns a generator of StorageSetting instances for property testing.
func StorageSettingGenerator() gopter.Gen {
	if storageSettingGenerator != nil {
		return storageSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageSetting(generators)
	storageSettingGenerator = gen.Struct(reflect.TypeOf(StorageSetting{}), generators)

	return storageSettingGenerator
}

// AddIndependentPropertyGeneratorsForStorageSetting is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageSetting(gens map[string]gopter.Gen) {
	gens["DatastoreType"] = gen.PtrOf(gen.OneConstOf(StorageSetting_DatastoreType_ArchiveStore, StorageSetting_DatastoreType_OperationalStore, StorageSetting_DatastoreType_VaultStore))
	gens["Type"] = gen.PtrOf(gen.OneConstOf(StorageSetting_Type_GeoRedundant, StorageSetting_Type_LocallyRedundant, StorageSetting_Type_ZoneRedundant))
}

func Test_UserAssignedIdentityDetails_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentityDetails via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentityDetails, UserAssignedIdentityDetailsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentityDetails runs a test to see if a specific instance of UserAssignedIdentityDetails round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentityDetails(subject UserAssignedIdentityDetails) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentityDetails
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

// Generator of UserAssignedIdentityDetails instances for property testing - lazily instantiated by
// UserAssignedIdentityDetailsGenerator()
var userAssignedIdentityDetailsGenerator gopter.Gen

// UserAssignedIdentityDetailsGenerator returns a generator of UserAssignedIdentityDetails instances for property testing.
func UserAssignedIdentityDetailsGenerator() gopter.Gen {
	if userAssignedIdentityDetailsGenerator != nil {
		return userAssignedIdentityDetailsGenerator
	}

	generators := make(map[string]gopter.Gen)
	userAssignedIdentityDetailsGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentityDetails{}), generators)

	return userAssignedIdentityDetailsGenerator
}
