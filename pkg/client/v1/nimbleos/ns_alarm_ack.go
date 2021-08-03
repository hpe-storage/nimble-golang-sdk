// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmAckFields provides field names to use in filter parameters, for example.
var NsAlarmAckFields *NsAlarmAckFieldHandles

func init() {
	NsAlarmAckFields = &NsAlarmAckFieldHandles{
		ID:              "id",
		RemindEvery:     "remind_every",
		RemindEveryUnit: "remind_every_unit",
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

// NsAlarmAckFieldHandles provides a string representation for each NsAlarmAck field.
type NsAlarmAckFieldHandles struct {
	ID              string
	RemindEvery     string
	RemindEveryUnit string
}
