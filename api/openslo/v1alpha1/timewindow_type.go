package v1alpha1

type TimeWindowSpec struct {
	Duration  Duration     `json:"duration,omitempty"`
	IsRolling bool         `json:"isRolling,omitempty"`
	Calendar  CalendarSpec `json:"calendar,omitempty"`
}

type CalendarSpec struct {
	// Date with time in 24h format, format without time zone
	// +kubebuilder:example="2020-01-21 12:30:00"
	StartTime string `json:"startTime,omitempty"`

	// Name as in IANA Time Zone Database
	// +kubebuilder:example="America/New_York"
	TimeZone string `json:"timeZone,omitempty"`
}
