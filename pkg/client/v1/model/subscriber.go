/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Subscriber


// Subscriber :
type Subscriber struct {
   // ID
   ID string `json:"id,omitempty"`
   // Type
   Type string `json:"type,omitempty"`
   // RenewInterval
   RenewInterval float64 `json:"renew_interval,omitempty"`
   // RenewResponseTimeout
   RenewResponseTimeout float64 `json:"renew_response_timeout,omitempty"`
   // IsConnected
   IsConnected bool `json:"is_connected,omitempty"`
   // Force
   Force bool `json:"force,omitempty"`
}
