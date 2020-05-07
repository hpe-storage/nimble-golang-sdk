/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/SnapshotCollection


// SnapshotCollection :
type SnapshotCollection struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // VolcollName
   VolcollName string `json:"volcoll_name,omitempty"`
   // VolcollID
   VolcollID string `json:"volcoll_id,omitempty"`
   // OriginName
   OriginName string `json:"origin_name,omitempty"`
   // IsReplica
   IsReplica bool `json:"is_replica,omitempty"`
   // SrepOwnerName
   SrepOwnerName string `json:"srep_owner_name,omitempty"`
   // SrepOwnerID
   SrepOwnerID string `json:"srep_owner_id,omitempty"`
   // PeerSnapcollID
   PeerSnapcollID string `json:"peer_snapcoll_id,omitempty"`
   // IsComplete
   IsComplete bool `json:"is_complete,omitempty"`
   // IsManual
   IsManual bool `json:"is_manual,omitempty"`
   // IsExternalTrigger
   IsExternalTrigger bool `json:"is_external_trigger,omitempty"`
   // IsUnmanaged
   IsUnmanaged bool `json:"is_unmanaged,omitempty"`
   // IsManuallyManaged
   IsManuallyManaged bool `json:"is_manually_managed,omitempty"`
   // ReplStartTime
   ReplStartTime float64 `json:"repl_start_time,omitempty"`
   // ReplCompleteTime
   ReplCompleteTime float64 `json:"repl_complete_time,omitempty"`
   // ReplBytesTransferred
   ReplBytesTransferred float64 `json:"repl_bytes_transferred,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // Replicate
   Replicate bool `json:"replicate,omitempty"`
   // StartOnline
   StartOnline bool `json:"start_online,omitempty"`
   // AllowWrites
   AllowWrites bool `json:"allow_writes,omitempty"`
   // DisableAppsync
   DisableAppsync bool `json:"disable_appsync,omitempty"`
   // SnapVerify
   SnapVerify bool `json:"snap_verify,omitempty"`
   // SkipDbConsistencyCheck
   SkipDbConsistencyCheck bool `json:"skip_db_consistency_check,omitempty"`
   // SchedID
   SchedID string `json:"sched_id,omitempty"`
   // SchedName
   SchedName string `json:"sched_name,omitempty"`
   // InvokeOnUpstreamPartner
   InvokeOnUpstreamPartner bool `json:"invoke_on_upstream_partner,omitempty"`
   // ExpiryAfter
   ExpiryAfter  int32 `json:"expiry_after,omitempty"`
}
