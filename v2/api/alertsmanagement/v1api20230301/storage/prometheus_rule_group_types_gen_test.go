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

func Test_PrometheusRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRule, PrometheusRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRule runs a test to see if a specific instance of PrometheusRule round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRule(subject PrometheusRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRule
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

// Generator of PrometheusRule instances for property testing - lazily instantiated by PrometheusRuleGenerator()
var prometheusRuleGenerator gopter.Gen

// PrometheusRuleGenerator returns a generator of PrometheusRule instances for property testing.
// We first initialize prometheusRuleGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrometheusRuleGenerator() gopter.Gen {
	if prometheusRuleGenerator != nil {
		return prometheusRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRule(generators)
	prometheusRuleGenerator = gen.Struct(reflect.TypeOf(PrometheusRule{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRule(generators)
	AddRelatedPropertyGeneratorsForPrometheusRule(generators)
	prometheusRuleGenerator = gen.Struct(reflect.TypeOf(PrometheusRule{}), generators)

	return prometheusRuleGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRule is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRule(gens map[string]gopter.Gen) {
	gens["Alert"] = gen.PtrOf(gen.AlphaString())
	gens["Annotations"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Expression"] = gen.PtrOf(gen.AlphaString())
	gens["For"] = gen.PtrOf(gen.AlphaString())
	gens["Labels"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Record"] = gen.PtrOf(gen.AlphaString())
	gens["Severity"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrometheusRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrometheusRule(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.SliceOf(PrometheusRuleGroupActionGenerator())
	gens["ResolveConfiguration"] = gen.PtrOf(PrometheusRuleResolveConfigurationGenerator())
}

func Test_PrometheusRuleGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleGroup, PrometheusRuleGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleGroup runs a test to see if a specific instance of PrometheusRuleGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleGroup(subject PrometheusRuleGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleGroup
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

// Generator of PrometheusRuleGroup instances for property testing - lazily instantiated by
// PrometheusRuleGroupGenerator()
var prometheusRuleGroupGenerator gopter.Gen

// PrometheusRuleGroupGenerator returns a generator of PrometheusRuleGroup instances for property testing.
func PrometheusRuleGroupGenerator() gopter.Gen {
	if prometheusRuleGroupGenerator != nil {
		return prometheusRuleGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrometheusRuleGroup(generators)
	prometheusRuleGroupGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroup{}), generators)

	return prometheusRuleGroupGenerator
}

// AddRelatedPropertyGeneratorsForPrometheusRuleGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrometheusRuleGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = PrometheusRuleGroup_SpecGenerator()
	gens["Status"] = PrometheusRuleGroup_STATUSGenerator()
}

func Test_PrometheusRuleGroupAction_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleGroupAction via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleGroupAction, PrometheusRuleGroupActionGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleGroupAction runs a test to see if a specific instance of PrometheusRuleGroupAction round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleGroupAction(subject PrometheusRuleGroupAction) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleGroupAction
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

// Generator of PrometheusRuleGroupAction instances for property testing - lazily instantiated by
// PrometheusRuleGroupActionGenerator()
var prometheusRuleGroupActionGenerator gopter.Gen

// PrometheusRuleGroupActionGenerator returns a generator of PrometheusRuleGroupAction instances for property testing.
func PrometheusRuleGroupActionGenerator() gopter.Gen {
	if prometheusRuleGroupActionGenerator != nil {
		return prometheusRuleGroupActionGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction(generators)
	prometheusRuleGroupActionGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroupAction{}), generators)

	return prometheusRuleGroupActionGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction(gens map[string]gopter.Gen) {
	gens["ActionProperties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_PrometheusRuleGroupAction_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleGroupAction_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleGroupAction_STATUS, PrometheusRuleGroupAction_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleGroupAction_STATUS runs a test to see if a specific instance of PrometheusRuleGroupAction_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleGroupAction_STATUS(subject PrometheusRuleGroupAction_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleGroupAction_STATUS
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

// Generator of PrometheusRuleGroupAction_STATUS instances for property testing - lazily instantiated by
// PrometheusRuleGroupAction_STATUSGenerator()
var prometheusRuleGroupAction_STATUSGenerator gopter.Gen

// PrometheusRuleGroupAction_STATUSGenerator returns a generator of PrometheusRuleGroupAction_STATUS instances for property testing.
func PrometheusRuleGroupAction_STATUSGenerator() gopter.Gen {
	if prometheusRuleGroupAction_STATUSGenerator != nil {
		return prometheusRuleGroupAction_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction_STATUS(generators)
	prometheusRuleGroupAction_STATUSGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroupAction_STATUS{}), generators)

	return prometheusRuleGroupAction_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleGroupAction_STATUS(gens map[string]gopter.Gen) {
	gens["ActionGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["ActionProperties"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_PrometheusRuleGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleGroup_STATUS, PrometheusRuleGroup_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleGroup_STATUS runs a test to see if a specific instance of PrometheusRuleGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleGroup_STATUS(subject PrometheusRuleGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleGroup_STATUS
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

// Generator of PrometheusRuleGroup_STATUS instances for property testing - lazily instantiated by
// PrometheusRuleGroup_STATUSGenerator()
var prometheusRuleGroup_STATUSGenerator gopter.Gen

// PrometheusRuleGroup_STATUSGenerator returns a generator of PrometheusRuleGroup_STATUS instances for property testing.
// We first initialize prometheusRuleGroup_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrometheusRuleGroup_STATUSGenerator() gopter.Gen {
	if prometheusRuleGroup_STATUSGenerator != nil {
		return prometheusRuleGroup_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroup_STATUS(generators)
	prometheusRuleGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroup_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrometheusRuleGroup_STATUS(generators)
	prometheusRuleGroup_STATUSGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroup_STATUS{}), generators)

	return prometheusRuleGroup_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleGroup_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleGroup_STATUS(gens map[string]gopter.Gen) {
	gens["ClusterName"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Interval"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Scopes"] = gen.SliceOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrometheusRuleGroup_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrometheusRuleGroup_STATUS(gens map[string]gopter.Gen) {
	gens["Rules"] = gen.SliceOf(PrometheusRule_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_PrometheusRuleGroup_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleGroup_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleGroup_Spec, PrometheusRuleGroup_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleGroup_Spec runs a test to see if a specific instance of PrometheusRuleGroup_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleGroup_Spec(subject PrometheusRuleGroup_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleGroup_Spec
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

// Generator of PrometheusRuleGroup_Spec instances for property testing - lazily instantiated by
// PrometheusRuleGroup_SpecGenerator()
var prometheusRuleGroup_SpecGenerator gopter.Gen

// PrometheusRuleGroup_SpecGenerator returns a generator of PrometheusRuleGroup_Spec instances for property testing.
// We first initialize prometheusRuleGroup_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrometheusRuleGroup_SpecGenerator() gopter.Gen {
	if prometheusRuleGroup_SpecGenerator != nil {
		return prometheusRuleGroup_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroup_Spec(generators)
	prometheusRuleGroup_SpecGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroup_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleGroup_Spec(generators)
	AddRelatedPropertyGeneratorsForPrometheusRuleGroup_Spec(generators)
	prometheusRuleGroup_SpecGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleGroup_Spec{}), generators)

	return prometheusRuleGroup_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleGroup_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleGroup_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["ClusterName"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Interval"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrometheusRuleGroup_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrometheusRuleGroup_Spec(gens map[string]gopter.Gen) {
	gens["Rules"] = gen.SliceOf(PrometheusRuleGenerator())
}

func Test_PrometheusRuleResolveConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleResolveConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleResolveConfiguration, PrometheusRuleResolveConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleResolveConfiguration runs a test to see if a specific instance of PrometheusRuleResolveConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleResolveConfiguration(subject PrometheusRuleResolveConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleResolveConfiguration
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

// Generator of PrometheusRuleResolveConfiguration instances for property testing - lazily instantiated by
// PrometheusRuleResolveConfigurationGenerator()
var prometheusRuleResolveConfigurationGenerator gopter.Gen

// PrometheusRuleResolveConfigurationGenerator returns a generator of PrometheusRuleResolveConfiguration instances for property testing.
func PrometheusRuleResolveConfigurationGenerator() gopter.Gen {
	if prometheusRuleResolveConfigurationGenerator != nil {
		return prometheusRuleResolveConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration(generators)
	prometheusRuleResolveConfigurationGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleResolveConfiguration{}), generators)

	return prometheusRuleResolveConfigurationGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration(gens map[string]gopter.Gen) {
	gens["AutoResolved"] = gen.PtrOf(gen.Bool())
	gens["TimeToResolve"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrometheusRuleResolveConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRuleResolveConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRuleResolveConfiguration_STATUS, PrometheusRuleResolveConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRuleResolveConfiguration_STATUS runs a test to see if a specific instance of PrometheusRuleResolveConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRuleResolveConfiguration_STATUS(subject PrometheusRuleResolveConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRuleResolveConfiguration_STATUS
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

// Generator of PrometheusRuleResolveConfiguration_STATUS instances for property testing - lazily instantiated by
// PrometheusRuleResolveConfiguration_STATUSGenerator()
var prometheusRuleResolveConfiguration_STATUSGenerator gopter.Gen

// PrometheusRuleResolveConfiguration_STATUSGenerator returns a generator of PrometheusRuleResolveConfiguration_STATUS instances for property testing.
func PrometheusRuleResolveConfiguration_STATUSGenerator() gopter.Gen {
	if prometheusRuleResolveConfiguration_STATUSGenerator != nil {
		return prometheusRuleResolveConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration_STATUS(generators)
	prometheusRuleResolveConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(PrometheusRuleResolveConfiguration_STATUS{}), generators)

	return prometheusRuleResolveConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRuleResolveConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["AutoResolved"] = gen.PtrOf(gen.Bool())
	gens["TimeToResolve"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrometheusRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrometheusRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrometheusRule_STATUS, PrometheusRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrometheusRule_STATUS runs a test to see if a specific instance of PrometheusRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForPrometheusRule_STATUS(subject PrometheusRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrometheusRule_STATUS
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

// Generator of PrometheusRule_STATUS instances for property testing - lazily instantiated by
// PrometheusRule_STATUSGenerator()
var prometheusRule_STATUSGenerator gopter.Gen

// PrometheusRule_STATUSGenerator returns a generator of PrometheusRule_STATUS instances for property testing.
// We first initialize prometheusRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrometheusRule_STATUSGenerator() gopter.Gen {
	if prometheusRule_STATUSGenerator != nil {
		return prometheusRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRule_STATUS(generators)
	prometheusRule_STATUSGenerator = gen.Struct(reflect.TypeOf(PrometheusRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrometheusRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForPrometheusRule_STATUS(generators)
	prometheusRule_STATUSGenerator = gen.Struct(reflect.TypeOf(PrometheusRule_STATUS{}), generators)

	return prometheusRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForPrometheusRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrometheusRule_STATUS(gens map[string]gopter.Gen) {
	gens["Alert"] = gen.PtrOf(gen.AlphaString())
	gens["Annotations"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Expression"] = gen.PtrOf(gen.AlphaString())
	gens["For"] = gen.PtrOf(gen.AlphaString())
	gens["Labels"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Record"] = gen.PtrOf(gen.AlphaString())
	gens["Severity"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForPrometheusRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrometheusRule_STATUS(gens map[string]gopter.Gen) {
	gens["Actions"] = gen.SliceOf(PrometheusRuleGroupAction_STATUSGenerator())
	gens["ResolveConfiguration"] = gen.PtrOf(PrometheusRuleResolveConfiguration_STATUSGenerator())
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