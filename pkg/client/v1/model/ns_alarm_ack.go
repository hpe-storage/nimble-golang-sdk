/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsAlarmAck - Arguments acknowledge alarm.
// Export NsAlarmAckFields for advance operations like search filter etc.
var NsAlarmAckFields *NsAlarmAck

func init(){
	IDfield:= "id"
		
	NsAlarmAckFields= &NsAlarmAck{
		ID: &IDfield,
		
	}
}

type NsAlarmAck struct {
   
    // ID of the alarm.
    
 	ID *string `json:"id,omitempty"`
   
    // Notification frequency unit.
    
   	RemindEvery *int64 `json:"remind_every,omitempty"`
   
    // Period unit.
    
   	RemindEveryUnit *NsPeriodUnit `json:"remind_every_unit,omitempty"`
}
