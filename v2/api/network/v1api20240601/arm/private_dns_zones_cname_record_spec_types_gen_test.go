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

func Test_PrivateDnsZonesCNAMERecord_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonesCNAMERecord_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec, PrivateDnsZonesCNAMERecord_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec runs a test to see if a specific instance of PrivateDnsZonesCNAMERecord_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonesCNAMERecord_Spec(subject PrivateDnsZonesCNAMERecord_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonesCNAMERecord_Spec
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

// Generator of PrivateDnsZonesCNAMERecord_Spec instances for property testing - lazily instantiated by
// PrivateDnsZonesCNAMERecord_SpecGenerator()
var privateDnsZonesCNAMERecord_SpecGenerator gopter.Gen

// PrivateDnsZonesCNAMERecord_SpecGenerator returns a generator of PrivateDnsZonesCNAMERecord_Spec instances for property testing.
// We first initialize privateDnsZonesCNAMERecord_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZonesCNAMERecord_SpecGenerator() gopter.Gen {
	if privateDnsZonesCNAMERecord_SpecGenerator != nil {
		return privateDnsZonesCNAMERecord_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(generators)
	privateDnsZonesCNAMERecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(generators)
	privateDnsZonesCNAMERecord_SpecGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonesCNAMERecord_Spec{}), generators)

	return privateDnsZonesCNAMERecord_SpecGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZonesCNAMERecord_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetPropertiesGenerator())
}