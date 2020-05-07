/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsAlarmUpdateAttr


// NsAlarmUpdateAttr :
type NsAlarmUpdateAttr struct {
   // ID
   ID string `json:"id,omitempty"`
   // RemindEvery
   RemindEvery float64 `json:"remind_every,omitempty"`
}
