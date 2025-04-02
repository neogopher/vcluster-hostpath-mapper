/*
Copyright The Kubernetes Authors.

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
	policyv1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	internal "k8s.io/client-go/applyconfigurations/internal"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// EvictionApplyConfiguration represents a declarative configuration of the Eviction type for use
// with apply.
type EvictionApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	DeleteOptions                    *v1.DeleteOptionsApplyConfiguration `json:"deleteOptions,omitempty"`
}

// Eviction constructs a declarative configuration of the Eviction type for use with
// apply.
func Eviction(name, namespace string) *EvictionApplyConfiguration {
	b := &EvictionApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("Eviction")
	b.WithAPIVersion("policy/v1beta1")
	return b
}

// ExtractEviction extracts the applied configuration owned by fieldManager from
// eviction. If no managedFields are found in eviction for fieldManager, a
// EvictionApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// eviction must be a unmodified Eviction API object that was retrieved from the Kubernetes API.
// ExtractEviction provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractEviction(eviction *policyv1beta1.Eviction, fieldManager string) (*EvictionApplyConfiguration, error) {
	return extractEviction(eviction, fieldManager, "")
}

// ExtractEvictionStatus is the same as ExtractEviction except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractEvictionStatus(eviction *policyv1beta1.Eviction, fieldManager string) (*EvictionApplyConfiguration, error) {
	return extractEviction(eviction, fieldManager, "status")
}

func extractEviction(eviction *policyv1beta1.Eviction, fieldManager string, subresource string) (*EvictionApplyConfiguration, error) {
	b := &EvictionApplyConfiguration{}
	err := managedfields.ExtractInto(eviction, internal.Parser().Type("io.k8s.api.policy.v1beta1.Eviction"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(eviction.Name)
	b.WithNamespace(eviction.Namespace)

	b.WithKind("Eviction")
	b.WithAPIVersion("policy/v1beta1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithKind(value string) *EvictionApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithAPIVersion(value string) *EvictionApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithName(value string) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithGenerateName(value string) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithNamespace(value string) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithUID(value types.UID) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithResourceVersion(value string) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithGeneration(value int64) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithCreationTimestamp(value metav1.Time) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ObjectMetaApplyConfiguration.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *EvictionApplyConfiguration) WithLabels(entries map[string]string) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Labels == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *EvictionApplyConfiguration) WithAnnotations(entries map[string]string) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.ObjectMetaApplyConfiguration.Annotations == nil && len(entries) > 0 {
		b.ObjectMetaApplyConfiguration.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.ObjectMetaApplyConfiguration.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *EvictionApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.ObjectMetaApplyConfiguration.OwnerReferences = append(b.ObjectMetaApplyConfiguration.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *EvictionApplyConfiguration) WithFinalizers(values ...string) *EvictionApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.ObjectMetaApplyConfiguration.Finalizers = append(b.ObjectMetaApplyConfiguration.Finalizers, values[i])
	}
	return b
}

func (b *EvictionApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithDeleteOptions sets the DeleteOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeleteOptions field is set to the value of the last call.
func (b *EvictionApplyConfiguration) WithDeleteOptions(value *v1.DeleteOptionsApplyConfiguration) *EvictionApplyConfiguration {
	b.DeleteOptions = value
	return b
}

// GetName retrieves the value of the Name field in the declarative configuration.
func (b *EvictionApplyConfiguration) GetName() *string {
	b.ensureObjectMetaApplyConfigurationExists()
	return b.ObjectMetaApplyConfiguration.Name
}
