// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

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

func Test_SqlTriggerGetResults_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetResults_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetResultsSTATUSARM, SqlTriggerGetResultsSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetResultsSTATUSARM runs a test to see if a specific instance of SqlTriggerGetResults_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetResultsSTATUSARM(subject SqlTriggerGetResults_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetResults_STATUSARM
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

// Generator of SqlTriggerGetResults_STATUSARM instances for property testing - lazily instantiated by
// SqlTriggerGetResultsSTATUSARMGenerator()
var sqlTriggerGetResultsSTATUSARMGenerator gopter.Gen

// SqlTriggerGetResultsSTATUSARMGenerator returns a generator of SqlTriggerGetResults_STATUSARM instances for property testing.
// We first initialize sqlTriggerGetResultsSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlTriggerGetResultsSTATUSARMGenerator() gopter.Gen {
	if sqlTriggerGetResultsSTATUSARMGenerator != nil {
		return sqlTriggerGetResultsSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResultsSTATUSARM(generators)
	sqlTriggerGetResultsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetResultsSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForSqlTriggerGetResultsSTATUSARM(generators)
	sqlTriggerGetResultsSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetResults_STATUSARM{}), generators)

	return sqlTriggerGetResultsSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetResultsSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetResultsSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlTriggerGetResultsSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerGetResultsSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SqlTriggerGetPropertiesSTATUSARMGenerator())
}

func Test_SqlTriggerGetProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetPropertiesSTATUSARM, SqlTriggerGetPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetPropertiesSTATUSARM runs a test to see if a specific instance of SqlTriggerGetProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetPropertiesSTATUSARM(subject SqlTriggerGetProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_STATUSARM
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

// Generator of SqlTriggerGetProperties_STATUSARM instances for property testing - lazily instantiated by
// SqlTriggerGetPropertiesSTATUSARMGenerator()
var sqlTriggerGetPropertiesSTATUSARMGenerator gopter.Gen

// SqlTriggerGetPropertiesSTATUSARMGenerator returns a generator of SqlTriggerGetProperties_STATUSARM instances for property testing.
func SqlTriggerGetPropertiesSTATUSARMGenerator() gopter.Gen {
	if sqlTriggerGetPropertiesSTATUSARMGenerator != nil {
		return sqlTriggerGetPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlTriggerGetPropertiesSTATUSARM(generators)
	sqlTriggerGetPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_STATUSARM{}), generators)

	return sqlTriggerGetPropertiesSTATUSARMGenerator
}

// AddRelatedPropertyGeneratorsForSqlTriggerGetPropertiesSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlTriggerGetPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(SqlTriggerGetPropertiesSTATUSResourceARMGenerator())
}

func Test_SqlTriggerGetProperties_STATUS_ResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlTriggerGetProperties_STATUS_ResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlTriggerGetPropertiesSTATUSResourceARM, SqlTriggerGetPropertiesSTATUSResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlTriggerGetPropertiesSTATUSResourceARM runs a test to see if a specific instance of SqlTriggerGetProperties_STATUS_ResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlTriggerGetPropertiesSTATUSResourceARM(subject SqlTriggerGetProperties_STATUS_ResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlTriggerGetProperties_STATUS_ResourceARM
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

// Generator of SqlTriggerGetProperties_STATUS_ResourceARM instances for property testing - lazily instantiated by
// SqlTriggerGetPropertiesSTATUSResourceARMGenerator()
var sqlTriggerGetPropertiesSTATUSResourceARMGenerator gopter.Gen

// SqlTriggerGetPropertiesSTATUSResourceARMGenerator returns a generator of SqlTriggerGetProperties_STATUS_ResourceARM instances for property testing.
func SqlTriggerGetPropertiesSTATUSResourceARMGenerator() gopter.Gen {
	if sqlTriggerGetPropertiesSTATUSResourceARMGenerator != nil {
		return sqlTriggerGetPropertiesSTATUSResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesSTATUSResourceARM(generators)
	sqlTriggerGetPropertiesSTATUSResourceARMGenerator = gen.Struct(reflect.TypeOf(SqlTriggerGetProperties_STATUS_ResourceARM{}), generators)

	return sqlTriggerGetPropertiesSTATUSResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesSTATUSResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlTriggerGetPropertiesSTATUSResourceARM(gens map[string]gopter.Gen) {
	gens["Body"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["TriggerOperation"] = gen.PtrOf(gen.OneConstOf(
		SqlTriggerGetPropertiesSTATUSResourceTriggerOperation_All,
		SqlTriggerGetPropertiesSTATUSResourceTriggerOperation_Create,
		SqlTriggerGetPropertiesSTATUSResourceTriggerOperation_Delete,
		SqlTriggerGetPropertiesSTATUSResourceTriggerOperation_Replace,
		SqlTriggerGetPropertiesSTATUSResourceTriggerOperation_Update))
	gens["TriggerType"] = gen.PtrOf(gen.OneConstOf(SqlTriggerGetPropertiesSTATUSResourceTriggerType_Post, SqlTriggerGetPropertiesSTATUSResourceTriggerType_Pre))
	gens["Ts"] = gen.PtrOf(gen.Float64())
}