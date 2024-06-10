package v1alpha1

type MetricSource struct {
	//MetricSourceRef can be used to refer to DataSource object
	MetricSourceRef string `json:"metricSourceRef,omitempty"`

	//Type describes predefined metric source type e.g. Prometheus, Datadog, etc.
	Type string `json:"type,omitempty"`

	//Spec is arbitrary chosen fields for every data source type to make it comfortable to use
	// anything that is valid YAML can be put here.
	Spec MetricSourceSpec `json:"spec,omitempty"`
}

type MetricSourceSpec struct {
	Query string `json:"query,omitempty"`
}
