// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsAlarmUpdateAttrFields provides field names to use in filter parameters, for example.
var NsAlarmUpdateAttrFields *NsAlarmUpdateAttrFieldHandles

func init() {
	NsAlarmUpdateAttrFields = &NsAlarmUpdateAttrFieldHandles{
		ID:              "id",
		RemindEvery:     "remind_every",
		RemindEveryUnit: "remind_every_unit",
	}
}

// NsAlarmUpdateAttr - Alarm object used for updating alarms.
type NsAlarmUpdateAttr struct {
	// ID - Identifier for the alarm.
	ID *string `json:"id,omitempty"`
	// RemindEvery - Frequency of notification.
	RemindEvery *int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Timeunit over which to send the number of notification specified in 'remind every'.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}

// NsAlarmUpdateAttrFieldHandles provides a string representation for each AccessControlRecord field.
type NsAlarmUpdateAttrFieldHandles struct {
	ID              string
	RemindEvery     string
	RemindEveryUnit string
}
