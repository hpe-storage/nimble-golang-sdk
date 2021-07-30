// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmAck - Arguments acknowledge alarm.

// Export NsAlarmAckFields provides field names to use in filter parameters, for example.
var NsAlarmAckFields *NsAlarmAckStringFields

func init() {
	fieldID := "id"
	fieldRemindEvery := "remind_every"
	fieldRemindEveryUnit := "remind_every_unit"

	NsAlarmAckFields = &NsAlarmAckStringFields{
		ID:              &fieldID,
		RemindEvery:     &fieldRemindEvery,
		RemindEveryUnit: &fieldRemindEveryUnit,
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
