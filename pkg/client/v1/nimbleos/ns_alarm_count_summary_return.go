// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmCountSummaryReturnFields provides field names to use in filter parameters, for example.
var NsAlarmCountSummaryReturnFields *NsAlarmCountSummaryReturnFieldHandles

func init() {
	NsAlarmCountSummaryReturnFields = &NsAlarmCountSummaryReturnFieldHandles{
		AlarmSummary: "alarm_summary",
	}
}

// NsAlarmCountSummaryReturn - List of alarm count for each category.
type NsAlarmCountSummaryReturn struct {
	// AlarmSummary - List of alarm count for each category.
	AlarmSummary []*NsAlarmCount `json:"alarm_summary,omitempty"`
}

// NsAlarmCountSummaryReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsAlarmCountSummaryReturnFieldHandles struct {
	AlarmSummary string
}
