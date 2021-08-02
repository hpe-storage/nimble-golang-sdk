// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmCountFields provides field names to use in filter parameters, for example.
var NsAlarmCountFields *NsAlarmCountFieldHandles

func init() {
	NsAlarmCountFields = &NsAlarmCountFieldHandles{
		Category: "category",
		Critical: "critical",
		Warning:  "warning",
	}
}

// NsAlarmCount - List of alarm count for each category.
type NsAlarmCount struct {
	// Category - Alert category.
	Category *NsEventCategory `json:"category,omitempty"`
	// Critical - Critical alarm count of a particular category.
	Critical *int64 `json:"critical,omitempty"`
	// Warning - Warning alarm count of a particular category.
	Warning *int64 `json:"warning,omitempty"`
}

// NsAlarmCountFieldHandles provides a string representation for each AccessControlRecord field.
type NsAlarmCountFieldHandles struct {
	Category string
	Critical string
	Warning  string
}
