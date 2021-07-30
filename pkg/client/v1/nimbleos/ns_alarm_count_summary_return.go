// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmCountSummaryReturn - List of alarm count for each category.
// Export NsAlarmCountSummaryReturnFields for advance operations like search filter etc.
var NsAlarmCountSummaryReturnFields *NsAlarmCountSummaryReturnStringFields

func init() {
	AlarmSummaryfield := "alarm_summary"

	NsAlarmCountSummaryReturnFields = &NsAlarmCountSummaryReturnStringFields{
		AlarmSummary: &AlarmSummaryfield,
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
