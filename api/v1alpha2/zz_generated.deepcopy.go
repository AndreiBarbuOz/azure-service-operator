// +build !ignore_autogenerated

// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ASOStatus) DeepCopyInto(out *ASOStatus) {
	*out = *in
	if in.RequestedAt != nil {
		in, out := &in.RequestedAt, &out.RequestedAt
		*out = (*in).DeepCopy()
	}
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ASOStatus.
func (in *ASOStatus) DeepCopy() *ASOStatus {
	if in == nil {
		return nil
	}
	out := new(ASOStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureDBsSQLSku) DeepCopyInto(out *AzureDBsSQLSku) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureDBsSQLSku.
func (in *AzureDBsSQLSku) DeepCopy() *AzureDBsSQLSku {
	if in == nil {
		return nil
	}
	out := new(AzureDBsSQLSku)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlobContainer) DeepCopyInto(out *BlobContainer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlobContainer.
func (in *BlobContainer) DeepCopy() *BlobContainer {
	if in == nil {
		return nil
	}
	out := new(BlobContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlobContainer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlobContainerList) DeepCopyInto(out *BlobContainerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BlobContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlobContainerList.
func (in *BlobContainerList) DeepCopy() *BlobContainerList {
	if in == nil {
		return nil
	}
	out := new(BlobContainerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlobContainerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlobContainerSpec) DeepCopyInto(out *BlobContainerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlobContainerSpec.
func (in *BlobContainerSpec) DeepCopy() *BlobContainerSpec {
	if in == nil {
		return nil
	}
	out := new(BlobContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericResource) DeepCopyInto(out *GenericResource) {
	*out = *in
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericResource.
func (in *GenericResource) DeepCopy() *GenericResource {
	if in == nil {
		return nil
	}
	out := new(GenericResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericSpec) DeepCopyInto(out *GenericSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericSpec.
func (in *GenericSpec) DeepCopy() *GenericSpec {
	if in == nil {
		return nil
	}
	out := new(GenericSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLAADUser) DeepCopyInto(out *MySQLAADUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLAADUser.
func (in *MySQLAADUser) DeepCopy() *MySQLAADUser {
	if in == nil {
		return nil
	}
	out := new(MySQLAADUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLAADUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLAADUserList) DeepCopyInto(out *MySQLAADUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLAADUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLAADUserList.
func (in *MySQLAADUserList) DeepCopy() *MySQLAADUserList {
	if in == nil {
		return nil
	}
	out := new(MySQLAADUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLAADUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLAADUserSpec) DeepCopyInto(out *MySQLAADUserSpec) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DatabaseRoles != nil {
		in, out := &in.DatabaseRoles, &out.DatabaseRoles
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLAADUserSpec.
func (in *MySQLAADUserSpec) DeepCopy() *MySQLAADUserSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLAADUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServer) DeepCopyInto(out *MySQLServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServer.
func (in *MySQLServer) DeepCopy() *MySQLServer {
	if in == nil {
		return nil
	}
	out := new(MySQLServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerList) DeepCopyInto(out *MySQLServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerList.
func (in *MySQLServerList) DeepCopy() *MySQLServerList {
	if in == nil {
		return nil
	}
	out := new(MySQLServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLServerSpec) DeepCopyInto(out *MySQLServerSpec) {
	*out = *in
	out.Sku = in.Sku
	out.ReplicaProperties = in.ReplicaProperties
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(MySQLStorageProfile)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLServerSpec.
func (in *MySQLServerSpec) DeepCopy() *MySQLServerSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLStorageProfile) DeepCopyInto(out *MySQLStorageProfile) {
	*out = *in
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(int32)
		**out = **in
	}
	if in.StorageMB != nil {
		in, out := &in.StorageMB, &out.StorageMB
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLStorageProfile.
func (in *MySQLStorageProfile) DeepCopy() *MySQLStorageProfile {
	if in == nil {
		return nil
	}
	out := new(MySQLStorageProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLUser) DeepCopyInto(out *MySQLUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLUser.
func (in *MySQLUser) DeepCopy() *MySQLUser {
	if in == nil {
		return nil
	}
	out := new(MySQLUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLUserList) DeepCopyInto(out *MySQLUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MySQLUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLUserList.
func (in *MySQLUserList) DeepCopy() *MySQLUserList {
	if in == nil {
		return nil
	}
	out := new(MySQLUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MySQLUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLUserSpec) DeepCopyInto(out *MySQLUserSpec) {
	*out = *in
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DatabaseRoles != nil {
		in, out := &in.DatabaseRoles, &out.DatabaseRoles
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLUserSpec.
func (in *MySQLUserSpec) DeepCopy() *MySQLUserSpec {
	if in == nil {
		return nil
	}
	out := new(MySQLUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PSQLStorageProfile) DeepCopyInto(out *PSQLStorageProfile) {
	*out = *in
	if in.BackupRetentionDays != nil {
		in, out := &in.BackupRetentionDays, &out.BackupRetentionDays
		*out = new(int32)
		**out = **in
	}
	if in.StorageMB != nil {
		in, out := &in.StorageMB, &out.StorageMB
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PSQLStorageProfile.
func (in *PSQLStorageProfile) DeepCopy() *PSQLStorageProfile {
	if in == nil {
		return nil
	}
	out := new(PSQLStorageProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServer) DeepCopyInto(out *PostgreSQLServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServer.
func (in *PostgreSQLServer) DeepCopy() *PostgreSQLServer {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerList) DeepCopyInto(out *PostgreSQLServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgreSQLServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerList.
func (in *PostgreSQLServerList) DeepCopy() *PostgreSQLServerList {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgreSQLServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgreSQLServerSpec) DeepCopyInto(out *PostgreSQLServerSpec) {
	*out = *in
	out.Sku = in.Sku
	if in.StorageProfile != nil {
		in, out := &in.StorageProfile, &out.StorageProfile
		*out = new(PSQLStorageProfile)
		(*in).DeepCopyInto(*out)
	}
	out.ReplicaProperties = in.ReplicaProperties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgreSQLServerSpec.
func (in *PostgreSQLServerSpec) DeepCopy() *PostgreSQLServerSpec {
	if in == nil {
		return nil
	}
	out := new(PostgreSQLServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicaProperties) DeepCopyInto(out *ReplicaProperties) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicaProperties.
func (in *ReplicaProperties) DeepCopy() *ReplicaProperties {
	if in == nil {
		return nil
	}
	out := new(ReplicaProperties)
	in.DeepCopyInto(out)
	return out
}
