/*
Copyright 2023 The Kubernetes Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=crc
// +kubebuilder:metadata:annotations=api-approved.kubernetes.io=unapproved
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +kubebuilder:storageversion

// ClusterReferenceConsumer identifies a consumer of a type of reference. For
// example, a consumer may support references from Gateways to Secrets for tls.
type ClusterReferenceConsumer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Subject Subject       `json:"subject"`
	From    GroupResource `json:"from"`
	To      GroupResource `json:"to"`
	For     For           `json:"for"`
}

// +kubebuilder:object:root=true

// ClusterReferenceConsumerList contains a list of ClusterReferenceConsumer
type ClusterReferenceConsumerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterReferenceConsumer `json:"items"`
}

// Subject is a copy of RBAC Subject that excludes APIGroup.
type Subject struct {
	// Kind of object being referenced. Values defined by this API group are
	// "User", "Group", and "ServiceAccount". If the Authorizer does not
	// recognized the kind value, the Authorizer should report an error.
	Kind string `json:"kind"`
	// Name of the object being referenced.
	Name string `json:"name"`
	// Namespace of the referenced object.  If the object kind is non-namespace,
	// such as "User" or "Group", and this value is not empty the Authorizer
	// should report an error. +optional
	Namespace string `json:"namespace,omitempty"`
}
