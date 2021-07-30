// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsAlarmAckFields provides field names to use in filter parameters, for example.
var NsAlarmAckFields *NsAlarmAckFieldHandles

func init() {
	fieldID := "id"
	fieldRemindEvery := "remind_every"
	fieldRemindEveryUnit := "remind_every_unit"

	NsAlarmAckFields = &NsAlarmAckFieldHandles{
		ID:              &fieldID,
		RemindEvery:     &fieldRemindEvery,
		RemindEveryUnit: &fieldRemindEveryUnit,
	}
}

// NsAlarmAck - Arguments acknowledge alarm.
type NsAlarmAck struct {
	// ID - ID of the alarm.
	ID *string `json:"id,omitempty"`
	// RemindEvery - Notification frequency unit.
	RemindEvery *int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Period unit.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}

// NsAlarmAckFieldHandles provides a string representation for each AccessControlRecord field.
type NsAlarmAckFieldHandles struct {
	ID              *string
	RemindEvery     *string
	RemindEveryUnit *string
}
