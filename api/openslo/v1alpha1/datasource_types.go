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

// DatasourceSpec defines the desired state of Datasource
type DatasourceSpec struct {
	// Description is a human-readable description of the datasource
	Description Description `json:"description,omitempty"`
	// DisplayName is a human-readable name for the datasource
	DisplayName string `json:"displayName,omitempty"`
	// Type is the type of the datasource. It is used to determine the connection details that are required.
	Type string `json:"type,omitempty"`
	// ConnectionDetails is a map of key-value pairs that are specific to the datasource type
	ConnectionDetails map[string]string `json:"connectionDetails,omitempty"`
}

// DatasourceStatus defines the observed state of Datasource
type DatasourceStatus struct {
	LastChecked   metav1.Time `json:"lastChecked,omitempty"`
	StatusMessage string      `json:"statusMessage,omitempty"`
	IsHealthy     bool        `json:"isHealthy,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Type",type="string",JSONPath=".spec.type"
//+kubebuilder:printcolumn:name="LastChecked",type="date",JSONPath=".status.lastChecked"
//+kubebuilder:printcolumn:name="IsHealthy",type="boolean",JSONPath=".status.isHealthy"
//+kubebuilder:printcolumn:name="StatusMessage",type="string",JSONPath=".status.statusMessage"

// Datasource is the Schema for the datasources API
// It enables reusing one source between many SLOs and moving connection specific details (e.g. authentication) away from SLO definitions.
type Datasource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasourceSpec   `json:"spec,omitempty"`
	Status            DatasourceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DatasourceList contains a list of Datasource
type DatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Datasource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Datasource{}, &DatasourceList{})
}
