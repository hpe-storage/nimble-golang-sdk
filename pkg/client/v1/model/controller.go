/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Controller


// Controller :
type Controller struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // ArrayID
   ArrayID string `json:"array_id,omitempty"`
   // PartialResponseOk
   PartialResponseOk bool `json:"partial_response_ok,omitempty"`
   // Serial
   Serial string `json:"serial,omitempty"`
   // Hostname
   Hostname string `json:"hostname,omitempty"`
   // SupportAddress
   SupportAddress string `json:"support_address,omitempty"`
   // SupportNetmask
   SupportNetmask string `json:"support_netmask,omitempty"`
   // SupportNic
   SupportNic string `json:"support_nic,omitempty"`
   // AsupTime
   AsupTime float64 `json:"asup_time,omitempty"`
}
