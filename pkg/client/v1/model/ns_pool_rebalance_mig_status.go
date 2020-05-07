/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsPoolRebalanceMigStatus


// NsPoolRebalanceMigStatus :
type NsPoolRebalanceMigStatus struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // PoolAvgSpaceUtilization
   PoolAvgSpaceUtilization float64 `json:"pool_avg_space_utilization,omitempty"`
}
