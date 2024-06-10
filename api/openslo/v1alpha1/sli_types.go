/*
Copyright 2024.

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

// SLISpec defines the desired state of SLI
type SLISpec struct {
	//Description of the SLI
	Description     Description     `json:"description,omitempty"`
	ThresholdMetric ThresholdMetric `json:"thresholdMetric,omitempty"`
	RatioMetric     RatioMetric     `json:"ratioMetric,omitempty"`
}

// SLIStatus defines the observed state of SLI
type SLIStatus struct {
	LastValue float64 `json:"lastValue,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SLI is the Schema for the slis API
type SLI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SLISpec   `json:"spec,omitempty"`
	Status            SLIStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SLIList contains a list of SLI
type SLIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SLI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SLI{}, &SLIList{})
}
