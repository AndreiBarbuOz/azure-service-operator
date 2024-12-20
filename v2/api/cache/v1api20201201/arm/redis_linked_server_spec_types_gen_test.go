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

func Test_RedisLinkedServerCreateProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerCreateProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerCreateProperties, RedisLinkedServerCreatePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerCreateProperties runs a test to see if a specific instance of RedisLinkedServerCreateProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerCreateProperties(subject RedisLinkedServerCreateProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerCreateProperties
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

// Generator of RedisLinkedServerCreateProperties instances for property testing - lazily instantiated by
// RedisLinkedServerCreatePropertiesGenerator()
var redisLinkedServerCreatePropertiesGenerator gopter.Gen

// RedisLinkedServerCreatePropertiesGenerator returns a generator of RedisLinkedServerCreateProperties instances for property testing.
func RedisLinkedServerCreatePropertiesGenerator() gopter.Gen {
	if redisLinkedServerCreatePropertiesGenerator != nil {
		return redisLinkedServerCreatePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerCreateProperties(generators)
	redisLinkedServerCreatePropertiesGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerCreateProperties{}), generators)

	return redisLinkedServerCreatePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerCreateProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerCreateProperties(gens map[string]gopter.Gen) {
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.OneConstOf(RedisLinkedServerCreateProperties_ServerRole_Primary, RedisLinkedServerCreateProperties_ServerRole_Secondary))
}

func Test_RedisLinkedServer_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServer_Spec, RedisLinkedServer_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServer_Spec runs a test to see if a specific instance of RedisLinkedServer_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServer_Spec(subject RedisLinkedServer_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_Spec
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

// Generator of RedisLinkedServer_Spec instances for property testing - lazily instantiated by
// RedisLinkedServer_SpecGenerator()
var redisLinkedServer_SpecGenerator gopter.Gen

// RedisLinkedServer_SpecGenerator returns a generator of RedisLinkedServer_Spec instances for property testing.
// We first initialize redisLinkedServer_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisLinkedServer_SpecGenerator() gopter.Gen {
	if redisLinkedServer_SpecGenerator != nil {
		return redisLinkedServer_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(generators)
	redisLinkedServer_SpecGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(generators)
	AddRelatedPropertyGeneratorsForRedisLinkedServer_Spec(generators)
	redisLinkedServer_SpecGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_Spec{}), generators)

	return redisLinkedServer_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServer_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedisLinkedServer_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServer_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisLinkedServerCreatePropertiesGenerator())
}
