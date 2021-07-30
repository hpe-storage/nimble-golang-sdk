// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsAlarmUpdateAttr - Alarm object used for updating alarms.
// Export NsAlarmUpdateAttrFields for advance operations like search filter etc.
var NsAlarmUpdateAttrFields *NsAlarmUpdateAttrStringFields

func init() {
	IDfield := "id"
	RemindEveryfield := "remind_every"
	RemindEveryUnitfield := "remind_every_unit"

	NsAlarmUpdateAttrFields = &NsAlarmUpdateAttrStringFields{
		ID:              &IDfield,
		RemindEvery:     &RemindEveryfield,
		RemindEveryUnit: &RemindEveryUnitfield,
	}
}

type NsAlarmUpdateAttr struct {
	// ID - Identifier for the alarm.
	ID *string `json:"id,omitempty"`
	// RemindEvery - Frequency of notification.
	RemindEvery *int64 `json:"remind_every,omitempty"`
	// RemindEveryUnit - Timeunit over which to send the number of notification specified in 'remind every'.
	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}

// Struct for NsAlarmUpdateAttrFields
type NsAlarmUpdateAttrStringFields struct {
	ID              *string
	RemindEvery     *string
	RemindEveryUnit *string
}
