// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

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

func Test_Profiles_RuleSet_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Profiles_RuleSet_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProfiles_RuleSet_STATUS_ARM, Profiles_RuleSet_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProfiles_RuleSet_STATUS_ARM runs a test to see if a specific instance of Profiles_RuleSet_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForProfiles_RuleSet_STATUS_ARM(subject Profiles_RuleSet_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Profiles_RuleSet_STATUS_ARM
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

// Generator of Profiles_RuleSet_STATUS_ARM instances for property testing - lazily instantiated by
// Profiles_RuleSet_STATUS_ARMGenerator()
var profiles_RuleSet_STATUS_ARMGenerator gopter.Gen

// Profiles_RuleSet_STATUS_ARMGenerator returns a generator of Profiles_RuleSet_STATUS_ARM instances for property testing.
// We first initialize profiles_RuleSet_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Profiles_RuleSet_STATUS_ARMGenerator() gopter.Gen {
	if profiles_RuleSet_STATUS_ARMGenerator != nil {
		return profiles_RuleSet_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_RuleSet_STATUS_ARM(generators)
	profiles_RuleSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_RuleSet_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForProfiles_RuleSet_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForProfiles_RuleSet_STATUS_ARM(generators)
	profiles_RuleSet_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Profiles_RuleSet_STATUS_ARM{}), generators)

	return profiles_RuleSet_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForProfiles_RuleSet_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForProfiles_RuleSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForProfiles_RuleSet_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProfiles_RuleSet_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RuleSetProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_RuleSetProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RuleSetProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRuleSetProperties_STATUS_ARM, RuleSetProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRuleSetProperties_STATUS_ARM runs a test to see if a specific instance of RuleSetProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRuleSetProperties_STATUS_ARM(subject RuleSetProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RuleSetProperties_STATUS_ARM
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

// Generator of RuleSetProperties_STATUS_ARM instances for property testing - lazily instantiated by
// RuleSetProperties_STATUS_ARMGenerator()
var ruleSetProperties_STATUS_ARMGenerator gopter.Gen

// RuleSetProperties_STATUS_ARMGenerator returns a generator of RuleSetProperties_STATUS_ARM instances for property testing.
func RuleSetProperties_STATUS_ARMGenerator() gopter.Gen {
	if ruleSetProperties_STATUS_ARMGenerator != nil {
		return ruleSetProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRuleSetProperties_STATUS_ARM(generators)
	ruleSetProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RuleSetProperties_STATUS_ARM{}), generators)

	return ruleSetProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRuleSetProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRuleSetProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["DeploymentStatus"] = gen.PtrOf(gen.OneConstOf(
		RuleSetProperties_DeploymentStatus_STATUS_Failed,
		RuleSetProperties_DeploymentStatus_STATUS_InProgress,
		RuleSetProperties_DeploymentStatus_STATUS_NotStarted,
		RuleSetProperties_DeploymentStatus_STATUS_Succeeded))
	gens["ProfileName"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		RuleSetProperties_ProvisioningState_STATUS_Creating,
		RuleSetProperties_ProvisioningState_STATUS_Deleting,
		RuleSetProperties_ProvisioningState_STATUS_Failed,
		RuleSetProperties_ProvisioningState_STATUS_Succeeded,
		RuleSetProperties_ProvisioningState_STATUS_Updating))
}
