/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsNIC


// NsNIC :
type NsNIC struct {
   // Name
   Name string `json:"name,omitempty"`
   // SubnetLabel
   SubnetLabel string `json:"subnet_label,omitempty"`
   // DataIp
   DataIp string `json:"data_ip,omitempty"`
   // Tagged
   Tagged bool `json:"tagged,omitempty"`
}
