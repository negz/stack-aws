/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"

	runtimev1alpha1 "github.com/crossplaneio/crossplane-runtime/apis/core/v1alpha1"
)

// GetBindingPhase of this KubernetesCluster.
func (cm *KubernetesCluster) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return cm.Status.GetBindingPhase()
}

// GetPortableClassReference of this KubernetesCluster.
func (cm *KubernetesCluster) GetPortableClassReference() *corev1.LocalObjectReference {
	return cm.Spec.PortableClassReference
}

// GetResourceReference of this KubernetesCluster.
func (cm *KubernetesCluster) GetResourceReference() *corev1.ObjectReference {
	return cm.Spec.ResourceReference
}

// GetWriteConnectionSecretToReference of this KubernetesCluster.
func (cm *KubernetesCluster) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return cm.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this KubernetesCluster.
func (cm *KubernetesCluster) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	cm.Status.SetBindingPhase(p)
}

// SetConditions of this KubernetesCluster.
func (cm *KubernetesCluster) SetConditions(c ...runtimev1alpha1.Condition) {
	cm.Status.SetConditions(c...)
}

// SetPortableClassReference of this KubernetesCluster.
func (cm *KubernetesCluster) SetPortableClassReference(r *corev1.LocalObjectReference) {
	cm.Spec.PortableClassReference = r
}

// SetResourceReference of this KubernetesCluster.
func (cm *KubernetesCluster) SetResourceReference(r *corev1.ObjectReference) {
	cm.Spec.ResourceReference = r
}

// SetWriteConnectionSecretToReference of this KubernetesCluster.
func (cm *KubernetesCluster) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	cm.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this MachineInstance.
func (cm *MachineInstance) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return cm.Status.GetBindingPhase()
}

// GetPortableClassReference of this MachineInstance.
func (cm *MachineInstance) GetPortableClassReference() *corev1.LocalObjectReference {
	return cm.Spec.PortableClassReference
}

// GetResourceReference of this MachineInstance.
func (cm *MachineInstance) GetResourceReference() *corev1.ObjectReference {
	return cm.Spec.ResourceReference
}

// GetWriteConnectionSecretToReference of this MachineInstance.
func (cm *MachineInstance) GetWriteConnectionSecretToReference() corev1.LocalObjectReference {
	return cm.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this MachineInstance.
func (cm *MachineInstance) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	cm.Status.SetBindingPhase(p)
}

// SetConditions of this MachineInstance.
func (cm *MachineInstance) SetConditions(c ...runtimev1alpha1.Condition) {
	cm.Status.SetConditions(c...)
}

// SetPortableClassReference of this MachineInstance.
func (cm *MachineInstance) SetPortableClassReference(r *corev1.LocalObjectReference) {
	cm.Spec.PortableClassReference = r
}

// SetResourceReference of this MachineInstance.
func (cm *MachineInstance) SetResourceReference(r *corev1.ObjectReference) {
	cm.Spec.ResourceReference = r
}

// SetWriteConnectionSecretToReference of this MachineInstance.
func (cm *MachineInstance) SetWriteConnectionSecretToReference(r corev1.LocalObjectReference) {
	cm.Spec.WriteConnectionSecretToReference = r
}