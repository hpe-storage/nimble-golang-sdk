// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmCount - List of alarm count for each category.

// Export NsAlarmCountFields provides field names to use in filter parameters, for example.
var NsAlarmCountFields *NsAlarmCountStringFields

func init() {
	fieldCategory := "category"
	fieldCritical := "critical"
	fieldWarning := "warning"

	NsAlarmCountFields = &NsAlarmCountStringFields{
		Category: &fieldCategory,
		Critical: &fieldCritical,
		Warning:  &fieldWarning,
	}
}

type NsAlarmCount struct {
	// Category - Alert category.
	Category *NsEventCategory `json:"category,omitempty"`
	// Critical - Critical alarm count of a particular category.
	Critical *int64 `json:"critical,omitempty"`
	// Warning - Warning alarm count of a particular category.
	Warning *int64 `json:"warning,omitempty"`
}

// Struct for NsAlarmCountFields
type NsAlarmCountStringFields struct {
	Category *string
	Critical *string
	Warning  *string
}
