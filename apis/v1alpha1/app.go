// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppSpec defines the desired state of App.
type AppSpec struct {
	// The name of the app.
	// +kubebuilder:validation:Required
	AppName *string `json:"appName"`
	// The type of app. Supported apps are JupyterServer and KernelGateway. TensorBoard
	// is not supported.
	// +kubebuilder:validation:Required
	AppType *string `json:"appType"`
	// The domain ID.
	// +kubebuilder:validation:Required
	DomainID *string `json:"domainID"`
	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image
	// created on the instance.
	ResourceSpec *ResourceSpec `json:"resourceSpec,omitempty"`
	// Each tag consists of a key and an optional value. Tag keys must be unique
	// per resource.
	Tags []*Tag `json:"tags,omitempty"`
	// The user profile name.
	// +kubebuilder:validation:Required
	UserProfileName *string `json:"userProfileName"`
}

// AppStatus defines the observed state of App
type AppStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The status.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty"`
}

// App is the Schema for the Apps API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="STATUS",type=string,priority=0,JSONPath=`.status.status`
type App struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppSpec   `json:"spec,omitempty"`
	Status            AppStatus `json:"status,omitempty"`
}

// AppList contains a list of App
// +kubebuilder:object:root=true
type AppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []App `json:"items"`
}

func init() {
	SchemeBuilder.Register(&App{}, &AppList{})
}
