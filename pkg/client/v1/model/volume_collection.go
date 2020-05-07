/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/VolumeCollection


// VolumeCollection :
type VolumeCollection struct {
   // ID
   ID string `json:"id,omitempty"`
   // ProttmplID
   ProttmplID string `json:"prottmpl_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // FullName
   FullName string `json:"full_name,omitempty"`
   // SearchName
   SearchName string `json:"search_name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // PolOwnerName
   PolOwnerName string `json:"pol_owner_name,omitempty"`
   // AppServer
   AppServer string `json:"app_server,omitempty"`
   // AppClusterName
   AppClusterName string `json:"app_cluster_name,omitempty"`
   // AppServiceName
   AppServiceName string `json:"app_service_name,omitempty"`
   // VcenterHostname
   VcenterHostname string `json:"vcenter_hostname,omitempty"`
   // VcenterUsername
   VcenterUsername string `json:"vcenter_username,omitempty"`
   // VcenterPassword
   VcenterPassword string `json:"vcenter_password,omitempty"`
   // AgentHostname
   AgentHostname string `json:"agent_hostname,omitempty"`
   // AgentUsername
   AgentUsername string `json:"agent_username,omitempty"`
   // AgentPassword
   AgentPassword string `json:"agent_password,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModifiedTime
   LastModifiedTime float64 `json:"last_modified_time,omitempty"`
   // VolumeCount
   VolumeCount float64 `json:"volume_count,omitempty"`
   // SnapcollCount
   SnapcollCount float64 `json:"snapcoll_count,omitempty"`
   // ReplicationPartner
   ReplicationPartner string `json:"replication_partner,omitempty"`
   // LagTime
   LagTime float64 `json:"lag_time,omitempty"`
   // IsStandaloneVolcoll
   IsStandaloneVolcoll bool `json:"is_standalone_volcoll,omitempty"`
   // TotalReplBytes
   TotalReplBytes float64 `json:"total_repl_bytes,omitempty"`
   // ReplBytesTransferred
   ReplBytesTransferred float64 `json:"repl_bytes_transferred,omitempty"`
   // IsHandingOver
   IsHandingOver bool `json:"is_handing_over,omitempty"`
   // HandoverReplicationPartner
   HandoverReplicationPartner string `json:"handover_replication_partner,omitempty"`
   // SrepLastSync
   SrepLastSync float64 `json:"srep_last_sync,omitempty"`
   // SrepResyncPercent
   SrepResyncPercent float64 `json:"srep_resync_percent,omitempty"`
}
