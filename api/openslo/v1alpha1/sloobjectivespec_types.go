package v1alpha1

import "k8s.io/apimachinery/pkg/api/resource"

// +kubebuilder:validation:Pattern=`^[1-9]\d*[s m h d]$`
type Duration string

type ObjectivesSpec struct {
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// +kubebuilder:validation:Enum=lte;gte;lt;gt
	Op              string            `json:"op,omitempty"`
	Value           string            `json:"value,omitempty"`
	Target          string            `json:"target,omitempty"`
	TargetPercent   string            `json:"targetPercent,omitempty"`
	TimeSliceTarget string            `json:"timeSliceTarget,omitempty"`
	TimeSliceWindow Duration          `json:"timeSliceWindow,omitempty"`
	Indicator       *Indicator        `json:"indicator,omitempty"`
	IndicatorRef    *string           `json:"indicatorRef,omitempty"`
	CompositeWeight resource.Quantity `json:"compositeWeight,omitempty"`
}
