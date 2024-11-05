// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/insights/v1api20220615/storage"
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

func Test_DiagnosticSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiagnosticSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiagnosticSetting, DiagnosticSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiagnosticSetting runs a test to see if a specific instance of DiagnosticSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForDiagnosticSetting(subject DiagnosticSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiagnosticSetting
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

// Generator of DiagnosticSetting instances for property testing - lazily instantiated by DiagnosticSettingGenerator()
var diagnosticSettingGenerator gopter.Gen

// DiagnosticSettingGenerator returns a generator of DiagnosticSetting instances for property testing.
func DiagnosticSettingGenerator() gopter.Gen {
	if diagnosticSettingGenerator != nil {
		return diagnosticSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDiagnosticSetting(generators)
	diagnosticSettingGenerator = gen.Struct(reflect.TypeOf(DiagnosticSetting{}), generators)

	return diagnosticSettingGenerator
}

// AddRelatedPropertyGeneratorsForDiagnosticSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiagnosticSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = DiagnosticSetting_SpecGenerator()
	gens["Status"] = DiagnosticSetting_STATUSGenerator()
}

func Test_DiagnosticSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiagnosticSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiagnosticSetting_STATUS, DiagnosticSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiagnosticSetting_STATUS runs a test to see if a specific instance of DiagnosticSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDiagnosticSetting_STATUS(subject DiagnosticSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiagnosticSetting_STATUS
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

// Generator of DiagnosticSetting_STATUS instances for property testing - lazily instantiated by
// DiagnosticSetting_STATUSGenerator()
var diagnosticSetting_STATUSGenerator gopter.Gen

// DiagnosticSetting_STATUSGenerator returns a generator of DiagnosticSetting_STATUS instances for property testing.
// We first initialize diagnosticSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiagnosticSetting_STATUSGenerator() gopter.Gen {
	if diagnosticSetting_STATUSGenerator != nil {
		return diagnosticSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiagnosticSetting_STATUS(generators)
	diagnosticSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DiagnosticSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiagnosticSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForDiagnosticSetting_STATUS(generators)
	diagnosticSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(DiagnosticSetting_STATUS{}), generators)

	return diagnosticSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDiagnosticSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiagnosticSetting_STATUS(gens map[string]gopter.Gen) {
	gens["EventHubAuthorizationRuleId"] = gen.PtrOf(gen.AlphaString())
	gens["EventHubName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["LogAnalyticsDestinationType"] = gen.PtrOf(gen.AlphaString())
	gens["MarketplacePartnerId"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ServiceBusRuleId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["WorkspaceId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiagnosticSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiagnosticSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Logs"] = gen.SliceOf(LogSettings_STATUSGenerator())
	gens["Metrics"] = gen.SliceOf(MetricSettings_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_DiagnosticSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DiagnosticSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDiagnosticSetting_Spec, DiagnosticSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDiagnosticSetting_Spec runs a test to see if a specific instance of DiagnosticSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDiagnosticSetting_Spec(subject DiagnosticSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DiagnosticSetting_Spec
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

// Generator of DiagnosticSetting_Spec instances for property testing - lazily instantiated by
// DiagnosticSetting_SpecGenerator()
var diagnosticSetting_SpecGenerator gopter.Gen

// DiagnosticSetting_SpecGenerator returns a generator of DiagnosticSetting_Spec instances for property testing.
// We first initialize diagnosticSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DiagnosticSetting_SpecGenerator() gopter.Gen {
	if diagnosticSetting_SpecGenerator != nil {
		return diagnosticSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiagnosticSetting_Spec(generators)
	diagnosticSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DiagnosticSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDiagnosticSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForDiagnosticSetting_Spec(generators)
	diagnosticSetting_SpecGenerator = gen.Struct(reflect.TypeOf(DiagnosticSetting_Spec{}), generators)

	return diagnosticSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDiagnosticSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDiagnosticSetting_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EventHubName"] = gen.PtrOf(gen.AlphaString())
	gens["LogAnalyticsDestinationType"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["ServiceBusRuleId"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDiagnosticSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDiagnosticSetting_Spec(gens map[string]gopter.Gen) {
	gens["Logs"] = gen.SliceOf(LogSettingsGenerator())
	gens["Metrics"] = gen.SliceOf(MetricSettingsGenerator())
}

func Test_LogSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LogSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLogSettings, LogSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLogSettings runs a test to see if a specific instance of LogSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForLogSettings(subject LogSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LogSettings
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

// Generator of LogSettings instances for property testing - lazily instantiated by LogSettingsGenerator()
var logSettingsGenerator gopter.Gen

// LogSettingsGenerator returns a generator of LogSettings instances for property testing.
// We first initialize logSettingsGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LogSettingsGenerator() gopter.Gen {
	if logSettingsGenerator != nil {
		return logSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLogSettings(generators)
	logSettingsGenerator = gen.Struct(reflect.TypeOf(LogSettings{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLogSettings(generators)
	AddRelatedPropertyGeneratorsForLogSettings(generators)
	logSettingsGenerator = gen.Struct(reflect.TypeOf(LogSettings{}), generators)

	return logSettingsGenerator
}

// AddIndependentPropertyGeneratorsForLogSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLogSettings(gens map[string]gopter.Gen) {
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["CategoryGroup"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLogSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLogSettings(gens map[string]gopter.Gen) {
	gens["RetentionPolicy"] = gen.PtrOf(RetentionPolicyGenerator())
}

func Test_LogSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LogSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLogSettings_STATUS, LogSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLogSettings_STATUS runs a test to see if a specific instance of LogSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForLogSettings_STATUS(subject LogSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LogSettings_STATUS
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

// Generator of LogSettings_STATUS instances for property testing - lazily instantiated by LogSettings_STATUSGenerator()
var logSettings_STATUSGenerator gopter.Gen

// LogSettings_STATUSGenerator returns a generator of LogSettings_STATUS instances for property testing.
// We first initialize logSettings_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LogSettings_STATUSGenerator() gopter.Gen {
	if logSettings_STATUSGenerator != nil {
		return logSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLogSettings_STATUS(generators)
	logSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(LogSettings_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLogSettings_STATUS(generators)
	AddRelatedPropertyGeneratorsForLogSettings_STATUS(generators)
	logSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(LogSettings_STATUS{}), generators)

	return logSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForLogSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLogSettings_STATUS(gens map[string]gopter.Gen) {
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["CategoryGroup"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForLogSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLogSettings_STATUS(gens map[string]gopter.Gen) {
	gens["RetentionPolicy"] = gen.PtrOf(RetentionPolicy_STATUSGenerator())
}

func Test_MetricSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricSettings, MetricSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricSettings runs a test to see if a specific instance of MetricSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricSettings(subject MetricSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricSettings
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

// Generator of MetricSettings instances for property testing - lazily instantiated by MetricSettingsGenerator()
var metricSettingsGenerator gopter.Gen

// MetricSettingsGenerator returns a generator of MetricSettings instances for property testing.
// We first initialize metricSettingsGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricSettingsGenerator() gopter.Gen {
	if metricSettingsGenerator != nil {
		return metricSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricSettings(generators)
	metricSettingsGenerator = gen.Struct(reflect.TypeOf(MetricSettings{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricSettings(generators)
	AddRelatedPropertyGeneratorsForMetricSettings(generators)
	metricSettingsGenerator = gen.Struct(reflect.TypeOf(MetricSettings{}), generators)

	return metricSettingsGenerator
}

// AddIndependentPropertyGeneratorsForMetricSettings is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricSettings(gens map[string]gopter.Gen) {
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["TimeGrain"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMetricSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricSettings(gens map[string]gopter.Gen) {
	gens["RetentionPolicy"] = gen.PtrOf(RetentionPolicyGenerator())
}

func Test_MetricSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MetricSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMetricSettings_STATUS, MetricSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMetricSettings_STATUS runs a test to see if a specific instance of MetricSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMetricSettings_STATUS(subject MetricSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MetricSettings_STATUS
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

// Generator of MetricSettings_STATUS instances for property testing - lazily instantiated by
// MetricSettings_STATUSGenerator()
var metricSettings_STATUSGenerator gopter.Gen

// MetricSettings_STATUSGenerator returns a generator of MetricSettings_STATUS instances for property testing.
// We first initialize metricSettings_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MetricSettings_STATUSGenerator() gopter.Gen {
	if metricSettings_STATUSGenerator != nil {
		return metricSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricSettings_STATUS(generators)
	metricSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(MetricSettings_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMetricSettings_STATUS(generators)
	AddRelatedPropertyGeneratorsForMetricSettings_STATUS(generators)
	metricSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(MetricSettings_STATUS{}), generators)

	return metricSettings_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMetricSettings_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMetricSettings_STATUS(gens map[string]gopter.Gen) {
	gens["Category"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["TimeGrain"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMetricSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMetricSettings_STATUS(gens map[string]gopter.Gen) {
	gens["RetentionPolicy"] = gen.PtrOf(RetentionPolicy_STATUSGenerator())
}

func Test_RetentionPolicy_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetentionPolicy via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetentionPolicy, RetentionPolicyGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetentionPolicy runs a test to see if a specific instance of RetentionPolicy round trips to JSON and back losslessly
func RunJSONSerializationTestForRetentionPolicy(subject RetentionPolicy) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetentionPolicy
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

// Generator of RetentionPolicy instances for property testing - lazily instantiated by RetentionPolicyGenerator()
var retentionPolicyGenerator gopter.Gen

// RetentionPolicyGenerator returns a generator of RetentionPolicy instances for property testing.
func RetentionPolicyGenerator() gopter.Gen {
	if retentionPolicyGenerator != nil {
		return retentionPolicyGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetentionPolicy(generators)
	retentionPolicyGenerator = gen.Struct(reflect.TypeOf(RetentionPolicy{}), generators)

	return retentionPolicyGenerator
}

// AddIndependentPropertyGeneratorsForRetentionPolicy is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetentionPolicy(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_RetentionPolicy_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RetentionPolicy_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRetentionPolicy_STATUS, RetentionPolicy_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRetentionPolicy_STATUS runs a test to see if a specific instance of RetentionPolicy_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRetentionPolicy_STATUS(subject RetentionPolicy_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RetentionPolicy_STATUS
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

// Generator of RetentionPolicy_STATUS instances for property testing - lazily instantiated by
// RetentionPolicy_STATUSGenerator()
var retentionPolicy_STATUSGenerator gopter.Gen

// RetentionPolicy_STATUSGenerator returns a generator of RetentionPolicy_STATUS instances for property testing.
func RetentionPolicy_STATUSGenerator() gopter.Gen {
	if retentionPolicy_STATUSGenerator != nil {
		return retentionPolicy_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRetentionPolicy_STATUS(generators)
	retentionPolicy_STATUSGenerator = gen.Struct(reflect.TypeOf(RetentionPolicy_STATUS{}), generators)

	return retentionPolicy_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRetentionPolicy_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRetentionPolicy_STATUS(gens map[string]gopter.Gen) {
	gens["Days"] = gen.PtrOf(gen.Int())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_SystemData_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SystemData_STATUS to SystemData_STATUS via AssignProperties_To_SystemData_STATUS & AssignProperties_From_SystemData_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSystemData_STATUS tests if a specific instance of SystemData_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SystemData_STATUS
	err := copied.AssignProperties_To_SystemData_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SystemData_STATUS
	err = actual.AssignProperties_From_SystemData_STATUS(&other)
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

func Test_SystemData_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemData_STATUS, SystemData_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemData_STATUS runs a test to see if a specific instance of SystemData_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemData_STATUS(subject SystemData_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_STATUS
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

// Generator of SystemData_STATUS instances for property testing - lazily instantiated by SystemData_STATUSGenerator()
var systemData_STATUSGenerator gopter.Gen

// SystemData_STATUSGenerator returns a generator of SystemData_STATUS instances for property testing.
func SystemData_STATUSGenerator() gopter.Gen {
	if systemData_STATUSGenerator != nil {
		return systemData_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemData_STATUS(generators)
	systemData_STATUSGenerator = gen.Struct(reflect.TypeOf(SystemData_STATUS{}), generators)

	return systemData_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSystemData_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemData_STATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.AlphaString())
}