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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubevault/operator/apis/extensions/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	testing "k8s.io/client-go/testing"
)

// FakeVaultSecrets implements VaultSecretInterface
type FakeVaultSecrets struct {
	Fake *FakeExtensionsV1alpha1
	ns   string
}

var vaultsecretsResource = schema.GroupVersionResource{Group: "extensions.kubevault.com", Version: "v1alpha1", Resource: "vaultsecrets"}

var vaultsecretsKind = schema.GroupVersionKind{Group: "extensions.kubevault.com", Version: "v1alpha1", Kind: "VaultSecret"}

// Get takes name of the vaultSecret, and returns the corresponding vaultSecret object, and an error if there is any.
func (c *FakeVaultSecrets) Get(name string, options v1.GetOptions) (result *v1alpha1.VaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vaultsecretsResource, c.ns, name), &v1alpha1.VaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultSecret), err
}

// List takes label and field selectors, and returns the list of VaultSecrets that match those selectors.
func (c *FakeVaultSecrets) List(opts v1.ListOptions) (result *v1alpha1.VaultSecretList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vaultsecretsResource, vaultsecretsKind, c.ns, opts), &v1alpha1.VaultSecretList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VaultSecretList{ListMeta: obj.(*v1alpha1.VaultSecretList).ListMeta}
	for _, item := range obj.(*v1alpha1.VaultSecretList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Create takes the representation of a vaultSecret and creates it.  Returns the server's representation of the vaultSecret, and an error, if there is any.
func (c *FakeVaultSecrets) Create(vaultSecret *v1alpha1.VaultSecret) (result *v1alpha1.VaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vaultsecretsResource, c.ns, vaultSecret), &v1alpha1.VaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultSecret), err
}

// Update takes the representation of a vaultSecret and updates it. Returns the server's representation of the vaultSecret, and an error, if there is any.
func (c *FakeVaultSecrets) Update(vaultSecret *v1alpha1.VaultSecret) (result *v1alpha1.VaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vaultsecretsResource, c.ns, vaultSecret), &v1alpha1.VaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultSecret), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVaultSecrets) UpdateStatus(vaultSecret *v1alpha1.VaultSecret) (*v1alpha1.VaultSecret, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vaultsecretsResource, "status", c.ns, vaultSecret), &v1alpha1.VaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultSecret), err
}

// Delete takes name of the vaultSecret and deletes it. Returns an error if one occurs.
func (c *FakeVaultSecrets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vaultsecretsResource, c.ns, name), &v1alpha1.VaultSecret{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVaultSecrets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vaultsecretsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VaultSecretList{})
	return err
}

// Patch applies the patch and returns the patched vaultSecret.
func (c *FakeVaultSecrets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VaultSecret, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vaultsecretsResource, c.ns, name, data, subresources...), &v1alpha1.VaultSecret{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VaultSecret), err
}
