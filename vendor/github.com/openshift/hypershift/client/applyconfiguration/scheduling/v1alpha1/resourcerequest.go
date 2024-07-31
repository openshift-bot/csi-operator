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

package v1alpha1

import (
	resource "k8s.io/apimachinery/pkg/api/resource"
)

// ResourceRequestApplyConfiguration represents an declarative configuration of the ResourceRequest type for use
// with apply.
type ResourceRequestApplyConfiguration struct {
	DeploymentName *string            `json:"deploymentName,omitempty"`
	ContainerName  *string            `json:"containerName,omitempty"`
	Memory         *resource.Quantity `json:"memory,omitempty"`
	CPU            *resource.Quantity `json:"cpu,omitempty"`
}

// ResourceRequestApplyConfiguration constructs an declarative configuration of the ResourceRequest type for use with
// apply.
func ResourceRequest() *ResourceRequestApplyConfiguration {
	return &ResourceRequestApplyConfiguration{}
}

// WithDeploymentName sets the DeploymentName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeploymentName field is set to the value of the last call.
func (b *ResourceRequestApplyConfiguration) WithDeploymentName(value string) *ResourceRequestApplyConfiguration {
	b.DeploymentName = &value
	return b
}

// WithContainerName sets the ContainerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ContainerName field is set to the value of the last call.
func (b *ResourceRequestApplyConfiguration) WithContainerName(value string) *ResourceRequestApplyConfiguration {
	b.ContainerName = &value
	return b
}

// WithMemory sets the Memory field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Memory field is set to the value of the last call.
func (b *ResourceRequestApplyConfiguration) WithMemory(value resource.Quantity) *ResourceRequestApplyConfiguration {
	b.Memory = &value
	return b
}

// WithCPU sets the CPU field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CPU field is set to the value of the last call.
func (b *ResourceRequestApplyConfiguration) WithCPU(value resource.Quantity) *ResourceRequestApplyConfiguration {
	b.CPU = &value
	return b
}
