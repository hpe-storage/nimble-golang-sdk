// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsAlarmUpdateAttr - Alarm object used for updating alarms.
// Export NsAlarmUpdateAttrFields for advance operations like search filter etc.
var NsAlarmUpdateAttrFields *NsAlarmUpdateAttr

func init(){
	IDfield:= "id"
		
	NsAlarmUpdateAttrFields= &NsAlarmUpdateAttr{
		ID:               &IDfield,
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
