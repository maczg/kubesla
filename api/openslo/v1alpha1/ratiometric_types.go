package v1alpha1

// RatioMetric either ThresholdMetric or RatioMetric must be provided
type RatioMetric struct {
	// Counter is true if the metric is a monotonically increasing counter,
	// or false, if it is a single number that can arbitrarily go up or down
	// ignored when using "raw"
	Counter bool `json:"counter,omitempty"`

	// Good the numerator, either "good" or "bad" must be provided if "total" is used
	Good MetricSource `json:"good,omitempty"`

	// Bad the numerator, either "good" or "bad" must be provided if "total" is used
	Bad MetricSource `json:"bad,omitempty"`

	// Total the denominator, either "good" or "bad" must be provided if "total" is used
	Total MetricSource `json:"total,omitempty"`

	// RawType required with "raw", indicates how the stored ratio was calculated:
	// success – good/total
	// failure – bad/total
	RawType bool `json:"rawType,omitempty"`
	//Raw the raw ratio value, ignored when using "good", "bad", "total"
	Raw MetricSource `json:"raw,omitempty"`
}
