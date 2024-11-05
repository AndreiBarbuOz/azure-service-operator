// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type WorkspacesConnection_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of workspace connection.
	Properties *WorkspaceConnectionProps `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &WorkspacesConnection_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (connection WorkspacesConnection_Spec) GetAPIVersion() string {
	return "2021-07-01"
}

// GetName returns the Name of the resource
func (connection *WorkspacesConnection_Spec) GetName() string {
	return connection.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces/connections"
func (connection *WorkspacesConnection_Spec) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces/connections"
}

// Workspace Connection specific properties.
type WorkspaceConnectionProps struct {
	// AuthType: Authorization type of the workspace connection.
	AuthType *string `json:"authType,omitempty"`

	// Category: Category of the workspace connection.
	Category *string `json:"category,omitempty"`

	// Target: Target of the workspace connection.
	Target *string `json:"target,omitempty"`

	// Value: Value details of the workspace connection.
	Value *string `json:"value,omitempty"`

	// ValueFormat: format for the workspace connection value
	ValueFormat *WorkspaceConnectionProps_ValueFormat `json:"valueFormat,omitempty"`
}

// +kubebuilder:validation:Enum={"JSON"}
type WorkspaceConnectionProps_ValueFormat string

const WorkspaceConnectionProps_ValueFormat_JSON = WorkspaceConnectionProps_ValueFormat("JSON")

// Mapping from string to WorkspaceConnectionProps_ValueFormat
var workspaceConnectionProps_ValueFormat_Values = map[string]WorkspaceConnectionProps_ValueFormat{
	"json": WorkspaceConnectionProps_ValueFormat_JSON,
}