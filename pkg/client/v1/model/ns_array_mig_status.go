/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsArrayMigStatus


// NsArrayMigStatus :
type NsArrayMigStatus struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // IsDataSource
   IsDataSource bool `json:"is_data_source,omitempty"`
   // SpaceUtilization
   SpaceUtilization float64 `json:"space_utilization,omitempty"`
}
