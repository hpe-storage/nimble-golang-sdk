// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmCountSummaryReturn - List of alarm count for each category.

// Export NsAlarmCountSummaryReturnFields provides field names to use in filter parameters, for example.
var NsAlarmCountSummaryReturnFields *NsAlarmCountSummaryReturnStringFields

func init() {
	fieldAlarmSummary := "alarm_summary"

	NsAlarmCountSummaryReturnFields = &NsAlarmCountSummaryReturnStringFields{
		AlarmSummary: &fieldAlarmSummary,
	}
}

type NsAlarmCountSummaryReturn struct {
	// AlarmSummary - List of alarm count for each category.
	AlarmSummary []*NsAlarmCount `json:"alarm_summary,omitempty"`
}

// Struct for NsAlarmCountSummaryReturnFields
type NsAlarmCountSummaryReturnStringFields struct {
	AlarmSummary *string
}
