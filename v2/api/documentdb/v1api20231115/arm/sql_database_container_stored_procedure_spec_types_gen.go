// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type SqlDatabaseContainerStoredProcedure_Spec struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB storedProcedure.
	Properties *SqlStoredProcedureCreateUpdateProperties `json:"properties,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &SqlDatabaseContainerStoredProcedure_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-15"
func (procedure SqlDatabaseContainerStoredProcedure_Spec) GetAPIVersion() string {
	return "2023-11-15"
}

// GetName returns the Name of the resource
func (procedure *SqlDatabaseContainerStoredProcedure_Spec) GetName() string {
	return procedure.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
func (procedure *SqlDatabaseContainerStoredProcedure_Spec) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
}

// Properties to create and update Azure Cosmos DB storedProcedure.
type SqlStoredProcedureCreateUpdateProperties struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions `json:"options,omitempty"`

	// Resource: The standard JSON format of a storedProcedure
	Resource *SqlStoredProcedureResource `json:"resource,omitempty"`
}

// Cosmos DB SQL storedProcedure resource object
type SqlStoredProcedureResource struct {
	// Body: Body of the Stored Procedure
	Body *string `json:"body,omitempty"`

	// Id: Name of the Cosmos DB SQL storedProcedure
	Id *string `json:"id,omitempty"`
}