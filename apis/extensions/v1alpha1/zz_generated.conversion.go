// +build !ignore_autogenerated

/*
Copyright 2018 The Kube Vault Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	extensions "github.com/kubevault/operator/apis/extensions"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_VaultSecret_To_extensions_VaultSecret,
		Convert_extensions_VaultSecret_To_v1alpha1_VaultSecret,
		Convert_v1alpha1_VaultSecretList_To_extensions_VaultSecretList,
		Convert_extensions_VaultSecretList_To_v1alpha1_VaultSecretList,
		Convert_v1alpha1_VaultSecretStatus_To_extensions_VaultSecretStatus,
		Convert_extensions_VaultSecretStatus_To_v1alpha1_VaultSecretStatus,
	)
}

func autoConvert_v1alpha1_VaultSecret_To_extensions_VaultSecret(in *VaultSecret, out *extensions.VaultSecret, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string][]byte)(unsafe.Pointer(&in.Data))
	if err := Convert_v1alpha1_VaultSecretStatus_To_extensions_VaultSecretStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_VaultSecret_To_extensions_VaultSecret is an autogenerated conversion function.
func Convert_v1alpha1_VaultSecret_To_extensions_VaultSecret(in *VaultSecret, out *extensions.VaultSecret, s conversion.Scope) error {
	return autoConvert_v1alpha1_VaultSecret_To_extensions_VaultSecret(in, out, s)
}

func autoConvert_extensions_VaultSecret_To_v1alpha1_VaultSecret(in *extensions.VaultSecret, out *VaultSecret, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string][]byte)(unsafe.Pointer(&in.Data))
	if err := Convert_extensions_VaultSecretStatus_To_v1alpha1_VaultSecretStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_extensions_VaultSecret_To_v1alpha1_VaultSecret is an autogenerated conversion function.
func Convert_extensions_VaultSecret_To_v1alpha1_VaultSecret(in *extensions.VaultSecret, out *VaultSecret, s conversion.Scope) error {
	return autoConvert_extensions_VaultSecret_To_v1alpha1_VaultSecret(in, out, s)
}

func autoConvert_v1alpha1_VaultSecretList_To_extensions_VaultSecretList(in *VaultSecretList, out *extensions.VaultSecretList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]extensions.VaultSecret)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_VaultSecretList_To_extensions_VaultSecretList is an autogenerated conversion function.
func Convert_v1alpha1_VaultSecretList_To_extensions_VaultSecretList(in *VaultSecretList, out *extensions.VaultSecretList, s conversion.Scope) error {
	return autoConvert_v1alpha1_VaultSecretList_To_extensions_VaultSecretList(in, out, s)
}

func autoConvert_extensions_VaultSecretList_To_v1alpha1_VaultSecretList(in *extensions.VaultSecretList, out *VaultSecretList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]VaultSecret)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_extensions_VaultSecretList_To_v1alpha1_VaultSecretList is an autogenerated conversion function.
func Convert_extensions_VaultSecretList_To_v1alpha1_VaultSecretList(in *extensions.VaultSecretList, out *VaultSecretList, s conversion.Scope) error {
	return autoConvert_extensions_VaultSecretList_To_v1alpha1_VaultSecretList(in, out, s)
}

func autoConvert_v1alpha1_VaultSecretStatus_To_extensions_VaultSecretStatus(in *VaultSecretStatus, out *extensions.VaultSecretStatus, s conversion.Scope) error {
	out.Tree = in.Tree
	out.Paths = *(*[]string)(unsafe.Pointer(&in.Paths))
	out.Hostname = in.Hostname
	out.Username = in.Username
	out.UID = in.UID
	out.Gid = in.Gid
	out.Tags = *(*[]string)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1alpha1_VaultSecretStatus_To_extensions_VaultSecretStatus is an autogenerated conversion function.
func Convert_v1alpha1_VaultSecretStatus_To_extensions_VaultSecretStatus(in *VaultSecretStatus, out *extensions.VaultSecretStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_VaultSecretStatus_To_extensions_VaultSecretStatus(in, out, s)
}

func autoConvert_extensions_VaultSecretStatus_To_v1alpha1_VaultSecretStatus(in *extensions.VaultSecretStatus, out *VaultSecretStatus, s conversion.Scope) error {
	out.Tree = in.Tree
	out.Paths = *(*[]string)(unsafe.Pointer(&in.Paths))
	out.Hostname = in.Hostname
	out.Username = in.Username
	out.UID = in.UID
	out.Gid = in.Gid
	out.Tags = *(*[]string)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_extensions_VaultSecretStatus_To_v1alpha1_VaultSecretStatus is an autogenerated conversion function.
func Convert_extensions_VaultSecretStatus_To_v1alpha1_VaultSecretStatus(in *extensions.VaultSecretStatus, out *VaultSecretStatus, s conversion.Scope) error {
	return autoConvert_extensions_VaultSecretStatus_To_v1alpha1_VaultSecretStatus(in, out, s)
}
