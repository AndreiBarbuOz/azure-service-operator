// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Database_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a database.
	Properties *DatabaseProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Database_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (database Database_Spec) GetAPIVersion() string {
	return "2018-06-01"
}

// GetName returns the Name of the resource
func (database *Database_Spec) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers/databases"
func (database *Database_Spec) GetType() string {
	return "Microsoft.DBforMariaDB/servers/databases"
}

// The properties of a database.
type DatabaseProperties struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`
}