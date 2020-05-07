/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Alarm


// Alarm :
type Alarm struct {
   // ID
   ID string `json:"id,omitempty"`
   // Type
   Type float64 `json:"type,omitempty"`
   // Array
   Array string `json:"array,omitempty"`
   // CurrOnsetEventID
   CurrOnsetEventID string `json:"curr_onset_event_id,omitempty"`
   // ObjectID
   ObjectID string `json:"object_id,omitempty"`
   // ObjectName
   ObjectName string `json:"object_name,omitempty"`
   // OnsetTime
   OnsetTime float64 `json:"onset_time,omitempty"`
   // AckTime
   AckTime float64 `json:"ack_time,omitempty"`
   // UserID
   UserID string `json:"user_id,omitempty"`
   // UserName
   UserName string `json:"user_name,omitempty"`
   // UserFullName
   UserFullName string `json:"user_full_name,omitempty"`
   // RemindEvery
   RemindEvery float64 `json:"remind_every,omitempty"`
   // Activity
   Activity string `json:"activity,omitempty"`
   // NextNotificationTime
   NextNotificationTime  int32 `json:"next_notification_time,omitempty"`
}
