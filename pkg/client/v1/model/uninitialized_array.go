/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/UninitializedArray


// UninitializedArray :
type UninitializedArray struct {
   // ID
   ID string `json:"id,omitempty"`
   // Serial
   Serial string `json:"serial,omitempty"`
   // ArrayName
   ArrayName string `json:"array_name,omitempty"`
   // Model
   Model string `json:"model,omitempty"`
   // ModelStr
   ModelStr string `json:"model_str,omitempty"`
   // Version
   Version string `json:"version,omitempty"`
   // IpAddress
   IpAddress string `json:"ip_address,omitempty"`
   // CountOfFcPorts
   CountOfFcPorts float64 `json:"count_of_fc_ports,omitempty"`
   // AllFlash
   AllFlash bool `json:"all_flash,omitempty"`
}
