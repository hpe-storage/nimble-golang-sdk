// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmAck - Arguments acknowledge alarm.
// Export NsAlarmAckFields for advance operations like search filter etc.
var NsAlarmAckFields *NsAlarmAckStringFields

func init() {
	IDfield := "id"
	RemindEveryfield := "remind_every"
	RemindEveryUnitfield := "remind_every_unit"

	NsAlarmAckFields = &NsAlarmAckStringFields{
		ID:              &IDfield,
		RemindEvery:     &RemindEveryfield,
		RemindEveryUnit: &RemindEveryUnitfield,
	}
}

type NsAlarmAck struct {
	// ID - ID of the alarm.
	ID *string `json:"id,omitempty"`
	// RemindEvery - Notification frequency unit.
	RemindEvery *int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Period unit.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}

// Struct for NsAlarmAckFields
type NsAlarmAckStringFields struct {
	ID              *string
	RemindEvery     *string
	RemindEveryUnit *string
}
