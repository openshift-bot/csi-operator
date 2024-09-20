// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apioperatorv1 "github.com/openshift/api/operator/v1"
	internal "github.com/openshift/client-go/operator/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// OpenShiftControllerManagerApplyConfiguration represents an declarative configuration of the OpenShiftControllerManager type for use
// with apply.
type OpenShiftControllerManagerApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	Spec                             *OpenShiftControllerManagerSpecApplyConfiguration   `json:"spec,omitempty"`
	Status                           *OpenShiftControllerManagerStatusApplyConfiguration `json:"status,omitempty"`
}

// OpenShiftControllerManager constructs an declarative configuration of the OpenShiftControllerManager type for use with
// apply.
func OpenShiftControllerManager(name string) *OpenShiftControllerManagerApplyConfiguration {
	b := &OpenShiftControllerManagerApplyConfiguration{}
	b.WithName(name)
	b.WithKind("OpenShiftControllerManager")
	b.WithAPIVersion("operator.openshift.io/v1")
	return b
}

// ExtractOpenShiftControllerManager extracts the applied configuration owned by fieldManager from
// openShiftControllerManager. If no managedFields are found in openShiftControllerManager for fieldManager, a
// OpenShiftControllerManagerApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// openShiftControllerManager must be a unmodified OpenShiftControllerManager API object that was retrieved from the Kubernetes API.
// ExtractOpenShiftControllerManager provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractOpenShiftControllerManager(openShiftControllerManager *apioperatorv1.OpenShiftControllerManager, fieldManager string) (*OpenShiftControllerManagerApplyConfiguration, error) {
	return extractOpenShiftControllerManager(openShiftControllerManager, fieldManager, "")
}

// ExtractOpenShiftControllerManagerStatus is the same as ExtractOpenShiftControllerManager except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractOpenShiftControllerManagerStatus(openShiftControllerManager *apioperatorv1.OpenShiftControllerManager, fieldManager string) (*OpenShiftControllerManagerApplyConfiguration, error) {
	return extractOpenShiftControllerManager(openShiftControllerManager, fieldManager, "status")
}

func extractOpenShiftControllerManager(openShiftControllerManager *apioperatorv1.OpenShiftControllerManager, fieldManager string, subresource string) (*OpenShiftControllerManagerApplyConfiguration, error) {
	b := &OpenShiftControllerManagerApplyConfiguration{}
	err := managedfields.ExtractInto(openShiftControllerManager, internal.Parser().Type("com.github.openshift.api.operator.v1.OpenShiftControllerManager"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(openShiftControllerManager.Name)

	b.WithKind("OpenShiftControllerManager")
	b.WithAPIVersion("operator.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithKind(value string) *OpenShiftControllerManagerApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithAPIVersion(value string) *OpenShiftControllerManagerApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithName(value string) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithGenerateName(value string) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithNamespace(value string) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithUID(value types.UID) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithResourceVersion(value string) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithGeneration(value int64) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithCreationTimestamp(value metav1.Time) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *OpenShiftControllerManagerApplyConfiguration) WithLabels(entries map[string]string) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *OpenShiftControllerManagerApplyConfiguration) WithAnnotations(entries map[string]string) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *OpenShiftControllerManagerApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *OpenShiftControllerManagerApplyConfiguration) WithFinalizers(values ...string) *OpenShiftControllerManagerApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *OpenShiftControllerManagerApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithSpec sets the Spec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Spec field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithSpec(value *OpenShiftControllerManagerSpecApplyConfiguration) *OpenShiftControllerManagerApplyConfiguration {
	b.Spec = value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *OpenShiftControllerManagerApplyConfiguration) WithStatus(value *OpenShiftControllerManagerStatusApplyConfiguration) *OpenShiftControllerManagerApplyConfiguration {
	b.Status = value
	return b
}
