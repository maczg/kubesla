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

// Indicator is the inlined version of the SLISpec
type Indicator struct {
	Metadata metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec     SLISpec           `json:"spec,omitempty"`
}

// SLOSpec defines the desired state of SLO
type SLOSpec struct {
	Description Description `json:"description,omitempty"`
	// Service is the name of the service that the SLO is associated with
	Service string `json:"service,omitempty"`

	// Sli is the inlined version of the SLISpec. Either Indicator or IndicatorRef must be set
	Sli *SLISpec `json:"sli,omitempty"`
	// SliRef is a reference to an SLI object. Either Indicator or IndicatorRef must be set
	SliRef *string `json:"SliRef,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// TimeWindow is exactly one item; one of possible: Rolling or CalendarAligned time window
	TimeWindow []TimeWindowSpec `json:"timeWindow,omitempty"`

	// +kubebuilder:validation:Enum=Occurrences;Timeslices;RatioTimeslices
	BudgetingMethod string `json:"budgetingMethod,omitempty"`

	// Objectives is a list of ObjectivesSpec that defines the SLO
	// +kubebuilder:validation:MinItems=1
	Objectives []ObjectivesSpec `json:"objectives,omitempty"`
	//AlertPolicies   []SLOAlertPolicy `json:"alertPolicies,omitempty"`
}

// SLOStatus defines the observed state of SLO
type SLOStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SLO is the Schema for the slos API
type SLO struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SLOSpec   `json:"spec,omitempty"`
	Status SLOStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SLOList contains a list of SLO
type SLOList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SLO `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SLO{}, &SLOList{})
}
