//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20231201

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Actions) DeepCopyInto(out *Actions) {
	*out = *in
	if in.ActionGroupsReferences != nil {
		in, out := &in.ActionGroupsReferences, &out.ActionGroupsReferences
		*out = make([]genruntime.ResourceReference, len(*in))
		copy(*out, *in)
	}
	if in.ActionProperties != nil {
		in, out := &in.ActionProperties, &out.ActionProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Actions.
func (in *Actions) DeepCopy() *Actions {
	if in == nil {
		return nil
	}
	out := new(Actions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Actions_STATUS) DeepCopyInto(out *Actions_STATUS) {
	*out = *in
	if in.ActionGroups != nil {
		in, out := &in.ActionGroups, &out.ActionGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ActionProperties != nil {
		in, out := &in.ActionProperties, &out.ActionProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Actions_STATUS.
func (in *Actions_STATUS) DeepCopy() *Actions_STATUS {
	if in == nil {
		return nil
	}
	out := new(Actions_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]Dimension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailingPeriods != nil {
		in, out := &in.FailingPeriods, &out.FailingPeriods
		*out = new(Condition_FailingPeriods)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricMeasureColumn != nil {
		in, out := &in.MetricMeasureColumn, &out.MetricMeasureColumn
		*out = new(string)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(Condition_Operator)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.ResourceIdColumnReference != nil {
		in, out := &in.ResourceIdColumnReference, &out.ResourceIdColumnReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.TimeAggregation != nil {
		in, out := &in.TimeAggregation, &out.TimeAggregation
		*out = new(Condition_TimeAggregation)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition_FailingPeriods) DeepCopyInto(out *Condition_FailingPeriods) {
	*out = *in
	if in.MinFailingPeriodsToAlert != nil {
		in, out := &in.MinFailingPeriodsToAlert, &out.MinFailingPeriodsToAlert
		*out = new(int)
		**out = **in
	}
	if in.NumberOfEvaluationPeriods != nil {
		in, out := &in.NumberOfEvaluationPeriods, &out.NumberOfEvaluationPeriods
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition_FailingPeriods.
func (in *Condition_FailingPeriods) DeepCopy() *Condition_FailingPeriods {
	if in == nil {
		return nil
	}
	out := new(Condition_FailingPeriods)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition_FailingPeriods_STATUS) DeepCopyInto(out *Condition_FailingPeriods_STATUS) {
	*out = *in
	if in.MinFailingPeriodsToAlert != nil {
		in, out := &in.MinFailingPeriodsToAlert, &out.MinFailingPeriodsToAlert
		*out = new(int)
		**out = **in
	}
	if in.NumberOfEvaluationPeriods != nil {
		in, out := &in.NumberOfEvaluationPeriods, &out.NumberOfEvaluationPeriods
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition_FailingPeriods_STATUS.
func (in *Condition_FailingPeriods_STATUS) DeepCopy() *Condition_FailingPeriods_STATUS {
	if in == nil {
		return nil
	}
	out := new(Condition_FailingPeriods_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition_STATUS) DeepCopyInto(out *Condition_STATUS) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]Dimension_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailingPeriods != nil {
		in, out := &in.FailingPeriods, &out.FailingPeriods
		*out = new(Condition_FailingPeriods_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricMeasureColumn != nil {
		in, out := &in.MetricMeasureColumn, &out.MetricMeasureColumn
		*out = new(string)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(Condition_Operator_STATUS)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.ResourceIdColumn != nil {
		in, out := &in.ResourceIdColumn, &out.ResourceIdColumn
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.TimeAggregation != nil {
		in, out := &in.TimeAggregation, &out.TimeAggregation
		*out = new(Condition_TimeAggregation_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition_STATUS.
func (in *Condition_STATUS) DeepCopy() *Condition_STATUS {
	if in == nil {
		return nil
	}
	out := new(Condition_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dimension) DeepCopyInto(out *Dimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(Dimension_Operator)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dimension.
func (in *Dimension) DeepCopy() *Dimension {
	if in == nil {
		return nil
	}
	out := new(Dimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dimension_STATUS) DeepCopyInto(out *Dimension_STATUS) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(Dimension_Operator_STATUS)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dimension_STATUS.
func (in *Dimension_STATUS) DeepCopy() *Dimension_STATUS {
	if in == nil {
		return nil
	}
	out := new(Dimension_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity) DeepCopyInto(out *Identity) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(Identity_Type)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make([]UserAssignedIdentityDetails, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity.
func (in *Identity) DeepCopy() *Identity {
	if in == nil {
		return nil
	}
	out := new(Identity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identity_STATUS) DeepCopyInto(out *Identity_STATUS) {
	*out = *in
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
	if in.TenantId != nil {
		in, out := &in.TenantId, &out.TenantId
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(Identity_Type_STATUS)
		**out = **in
	}
	if in.UserAssignedIdentities != nil {
		in, out := &in.UserAssignedIdentities, &out.UserAssignedIdentities
		*out = make(map[string]UserIdentityProperties_STATUS, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identity_STATUS.
func (in *Identity_STATUS) DeepCopy() *Identity_STATUS {
	if in == nil {
		return nil
	}
	out := new(Identity_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleResolveConfiguration) DeepCopyInto(out *RuleResolveConfiguration) {
	*out = *in
	if in.AutoResolved != nil {
		in, out := &in.AutoResolved, &out.AutoResolved
		*out = new(bool)
		**out = **in
	}
	if in.TimeToResolve != nil {
		in, out := &in.TimeToResolve, &out.TimeToResolve
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleResolveConfiguration.
func (in *RuleResolveConfiguration) DeepCopy() *RuleResolveConfiguration {
	if in == nil {
		return nil
	}
	out := new(RuleResolveConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuleResolveConfiguration_STATUS) DeepCopyInto(out *RuleResolveConfiguration_STATUS) {
	*out = *in
	if in.AutoResolved != nil {
		in, out := &in.AutoResolved, &out.AutoResolved
		*out = new(bool)
		**out = **in
	}
	if in.TimeToResolve != nil {
		in, out := &in.TimeToResolve, &out.TimeToResolve
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuleResolveConfiguration_STATUS.
func (in *RuleResolveConfiguration_STATUS) DeepCopy() *RuleResolveConfiguration_STATUS {
	if in == nil {
		return nil
	}
	out := new(RuleResolveConfiguration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRule) DeepCopyInto(out *ScheduledQueryRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRule.
func (in *ScheduledQueryRule) DeepCopy() *ScheduledQueryRule {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledQueryRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRuleCriteria) DeepCopyInto(out *ScheduledQueryRuleCriteria) {
	*out = *in
	if in.AllOf != nil {
		in, out := &in.AllOf, &out.AllOf
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRuleCriteria.
func (in *ScheduledQueryRuleCriteria) DeepCopy() *ScheduledQueryRuleCriteria {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRuleCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRuleCriteria_STATUS) DeepCopyInto(out *ScheduledQueryRuleCriteria_STATUS) {
	*out = *in
	if in.AllOf != nil {
		in, out := &in.AllOf, &out.AllOf
		*out = make([]Condition_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRuleCriteria_STATUS.
func (in *ScheduledQueryRuleCriteria_STATUS) DeepCopy() *ScheduledQueryRuleCriteria_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRuleCriteria_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRuleList) DeepCopyInto(out *ScheduledQueryRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledQueryRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRuleList.
func (in *ScheduledQueryRuleList) DeepCopy() *ScheduledQueryRuleList {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledQueryRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRuleOperatorSpec) DeepCopyInto(out *ScheduledQueryRuleOperatorSpec) {
	*out = *in
	if in.ConfigMapExpressions != nil {
		in, out := &in.ConfigMapExpressions, &out.ConfigMapExpressions
		*out = make([]*core.DestinationExpression, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(core.DestinationExpression)
				**out = **in
			}
		}
	}
	if in.SecretExpressions != nil {
		in, out := &in.SecretExpressions, &out.SecretExpressions
		*out = make([]*core.DestinationExpression, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(core.DestinationExpression)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRuleOperatorSpec.
func (in *ScheduledQueryRuleOperatorSpec) DeepCopy() *ScheduledQueryRuleOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRuleOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRule_STATUS) DeepCopyInto(out *ScheduledQueryRule_STATUS) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = new(Actions_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoMitigate != nil {
		in, out := &in.AutoMitigate, &out.AutoMitigate
		*out = new(bool)
		**out = **in
	}
	if in.CheckWorkspaceAlertsStorageConfigured != nil {
		in, out := &in.CheckWorkspaceAlertsStorageConfigured, &out.CheckWorkspaceAlertsStorageConfigured
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreatedWithApiVersion != nil {
		in, out := &in.CreatedWithApiVersion, &out.CreatedWithApiVersion
		*out = new(string)
		**out = **in
	}
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = new(ScheduledQueryRuleCriteria_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.EvaluationFrequency != nil {
		in, out := &in.EvaluationFrequency, &out.EvaluationFrequency
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.IsLegacyLogAnalyticsRule != nil {
		in, out := &in.IsLegacyLogAnalyticsRule, &out.IsLegacyLogAnalyticsRule
		*out = new(bool)
		**out = **in
	}
	if in.IsWorkspaceAlertsStorageConfigured != nil {
		in, out := &in.IsWorkspaceAlertsStorageConfigured, &out.IsWorkspaceAlertsStorageConfigured
		*out = new(bool)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(ScheduledQueryRule_Kind_STATUS)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MuteActionsDuration != nil {
		in, out := &in.MuteActionsDuration, &out.MuteActionsDuration
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OverrideQueryTimeRange != nil {
		in, out := &in.OverrideQueryTimeRange, &out.OverrideQueryTimeRange
		*out = new(string)
		**out = **in
	}
	if in.ResolveConfiguration != nil {
		in, out := &in.ResolveConfiguration, &out.ResolveConfiguration
		*out = new(RuleResolveConfiguration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(ScheduledQueryRuleProperties_Severity_STATUS)
		**out = **in
	}
	if in.SkipQueryValidation != nil {
		in, out := &in.SkipQueryValidation, &out.SkipQueryValidation
		*out = new(bool)
		**out = **in
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetResourceTypes != nil {
		in, out := &in.TargetResourceTypes, &out.TargetResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.WindowSize != nil {
		in, out := &in.WindowSize, &out.WindowSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRule_STATUS.
func (in *ScheduledQueryRule_STATUS) DeepCopy() *ScheduledQueryRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRule_Spec) DeepCopyInto(out *ScheduledQueryRule_Spec) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = new(Actions)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoMitigate != nil {
		in, out := &in.AutoMitigate, &out.AutoMitigate
		*out = new(bool)
		**out = **in
	}
	if in.CheckWorkspaceAlertsStorageConfigured != nil {
		in, out := &in.CheckWorkspaceAlertsStorageConfigured, &out.CheckWorkspaceAlertsStorageConfigured
		*out = new(bool)
		**out = **in
	}
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = new(ScheduledQueryRuleCriteria)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EvaluationFrequency != nil {
		in, out := &in.EvaluationFrequency, &out.EvaluationFrequency
		*out = new(string)
		**out = **in
	}
	if in.Identity != nil {
		in, out := &in.Identity, &out.Identity
		*out = new(Identity)
		(*in).DeepCopyInto(*out)
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(ScheduledQueryRule_Kind_Spec)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MuteActionsDuration != nil {
		in, out := &in.MuteActionsDuration, &out.MuteActionsDuration
		*out = new(string)
		**out = **in
	}
	if in.OperatorSpec != nil {
		in, out := &in.OperatorSpec, &out.OperatorSpec
		*out = new(ScheduledQueryRuleOperatorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.OverrideQueryTimeRange != nil {
		in, out := &in.OverrideQueryTimeRange, &out.OverrideQueryTimeRange
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.ResolveConfiguration != nil {
		in, out := &in.ResolveConfiguration, &out.ResolveConfiguration
		*out = new(RuleResolveConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.ScopesReferences != nil {
		in, out := &in.ScopesReferences, &out.ScopesReferences
		*out = make([]genruntime.ResourceReference, len(*in))
		copy(*out, *in)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(ScheduledQueryRuleProperties_Severity)
		**out = **in
	}
	if in.SkipQueryValidation != nil {
		in, out := &in.SkipQueryValidation, &out.SkipQueryValidation
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetResourceTypes != nil {
		in, out := &in.TargetResourceTypes, &out.TargetResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WindowSize != nil {
		in, out := &in.WindowSize, &out.WindowSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRule_Spec.
func (in *ScheduledQueryRule_Spec) DeepCopy() *ScheduledQueryRule_Spec {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRule_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS) DeepCopyInto(out *SystemData_STATUS) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(SystemData_CreatedByType_STATUS)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
		*out = new(SystemData_LastModifiedByType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS.
func (in *SystemData_STATUS) DeepCopy() *SystemData_STATUS {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserAssignedIdentityDetails) DeepCopyInto(out *UserAssignedIdentityDetails) {
	*out = *in
	out.Reference = in.Reference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserAssignedIdentityDetails.
func (in *UserAssignedIdentityDetails) DeepCopy() *UserAssignedIdentityDetails {
	if in == nil {
		return nil
	}
	out := new(UserAssignedIdentityDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserIdentityProperties_STATUS) DeepCopyInto(out *UserIdentityProperties_STATUS) {
	*out = *in
	if in.ClientId != nil {
		in, out := &in.ClientId, &out.ClientId
		*out = new(string)
		**out = **in
	}
	if in.PrincipalId != nil {
		in, out := &in.PrincipalId, &out.PrincipalId
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserIdentityProperties_STATUS.
func (in *UserIdentityProperties_STATUS) DeepCopy() *UserIdentityProperties_STATUS {
	if in == nil {
		return nil
	}
	out := new(UserIdentityProperties_STATUS)
	in.DeepCopyInto(out)
	return out
}
