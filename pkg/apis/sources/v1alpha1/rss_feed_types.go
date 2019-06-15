/*
Copyright 2019 The Knative Authors.

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

	"github.com/knative/pkg/apis"
	duckv1beta1 "github.com/knative/pkg/apis/duck/v1beta1"
	"github.com/knative/pkg/kmeta"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RssFeed is a steam of updates from a `.rss` url.
type RssFeed struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the RssFeed (from the client).
	// +optional
	Spec RssFeedSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the RssFeed (from the controller).
	// +optional
	Status RssFeedStatus `json:"status,omitempty"`
}

// Check that RssFeed can be validated and defaulted.
var _ apis.Validatable = (*RssFeed)(nil)
var _ apis.Defaultable = (*RssFeed)(nil)
var _ kmeta.OwnerRefable = (*RssFeed)(nil)

// RssFeedSpec holds the desired state of the RssFeed (from the client).
type RssFeedSpec struct {
	// ServiceName holds the name of the Kubernetes Service to expose as an "addressable".
	ServiceName string `json:"serviceName"`
}

const (
	// RssFeedConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	RssFeedConditionReady = apis.ConditionReady
)

// RssFeedStatus communicates the observed state of the RssFeed (from the controller).
type RssFeedStatus struct {
	duckv1beta1.Status `json:",inline"`

	// Address holds the information needed to connect this Addressable up to receive events.
	// +optional
	Address *duckv1beta1.Addressable `json:"address,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RssFeedList is a list of RssFeed resources
type RssFeedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RssFeed `json:"items"`
}
