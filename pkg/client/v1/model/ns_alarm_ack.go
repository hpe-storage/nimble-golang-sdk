/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsAlarmAck


// NsAlarmAck :
type NsAlarmAck struct {
   // ID
   ID string `json:"id,omitempty"`
   // RemindEvery
   RemindEvery float64 `json:"remind_every,omitempty"`
}
