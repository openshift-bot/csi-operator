/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
	v1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// KubevirtPersistentVolumeApplyConfiguration represents an declarative configuration of the KubevirtPersistentVolume type for use
// with apply.
type KubevirtPersistentVolumeApplyConfiguration struct {
	Size         *resource.Quantity                   `json:"size,omitempty"`
	StorageClass *string                              `json:"storageClass,omitempty"`
	AccessModes  []v1beta1.PersistentVolumeAccessMode `json:"accessModes,omitempty"`
	VolumeMode   *v1.PersistentVolumeMode             `json:"volumeMode,omitempty"`
}

// KubevirtPersistentVolumeApplyConfiguration constructs an declarative configuration of the KubevirtPersistentVolume type for use with
// apply.
func KubevirtPersistentVolume() *KubevirtPersistentVolumeApplyConfiguration {
	return &KubevirtPersistentVolumeApplyConfiguration{}
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *KubevirtPersistentVolumeApplyConfiguration) WithSize(value resource.Quantity) *KubevirtPersistentVolumeApplyConfiguration {
	b.Size = &value
	return b
}

// WithStorageClass sets the StorageClass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClass field is set to the value of the last call.
func (b *KubevirtPersistentVolumeApplyConfiguration) WithStorageClass(value string) *KubevirtPersistentVolumeApplyConfiguration {
	b.StorageClass = &value
	return b
}

// WithAccessModes adds the given value to the AccessModes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AccessModes field.
func (b *KubevirtPersistentVolumeApplyConfiguration) WithAccessModes(values ...v1beta1.PersistentVolumeAccessMode) *KubevirtPersistentVolumeApplyConfiguration {
	for i := range values {
		b.AccessModes = append(b.AccessModes, values[i])
	}
	return b
}

// WithVolumeMode sets the VolumeMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeMode field is set to the value of the last call.
func (b *KubevirtPersistentVolumeApplyConfiguration) WithVolumeMode(value v1.PersistentVolumeMode) *KubevirtPersistentVolumeApplyConfiguration {
	b.VolumeMode = &value
	return b
}
