/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/ProtectionSchedule


// ProtectionSchedule :
type ProtectionSchedule struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // Description
   Description string `json:"description,omitempty"`
   // VolcollOrProttmplID
   VolcollOrProttmplID string `json:"volcoll_or_prottmpl_id,omitempty"`
   // Period
   Period float64 `json:"period,omitempty"`
   // AtTime
   AtTime float64 `json:"at_time,omitempty"`
   // UntilTime
   UntilTime float64 `json:"until_time,omitempty"`
   // Days
   Days string `json:"days,omitempty"`
   // NumRetain
   NumRetain float64 `json:"num_retain,omitempty"`
   // DownstreamPartner
   DownstreamPartner string `json:"downstream_partner,omitempty"`
   // DownstreamPartnerName
   DownstreamPartnerName string `json:"downstream_partner_name,omitempty"`
   // DownstreamPartnerID
   DownstreamPartnerID string `json:"downstream_partner_id,omitempty"`
   // UpstreamPartnerName
   UpstreamPartnerName string `json:"upstream_partner_name,omitempty"`
   // UpstreamPartnerID
   UpstreamPartnerID string `json:"upstream_partner_id,omitempty"`
   // ReplicateEvery
   ReplicateEvery float64 `json:"replicate_every,omitempty"`
   // NumRetainReplica
   NumRetainReplica float64 `json:"num_retain_replica,omitempty"`
   // ReplAlertThres
   ReplAlertThres float64 `json:"repl_alert_thres,omitempty"`
   // SnapVerify
   SnapVerify bool `json:"snap_verify,omitempty"`
   // SkipDbConsistencyCheck
   SkipDbConsistencyCheck bool `json:"skip_db_consistency_check,omitempty"`
   // DisableAppsync
   DisableAppsync bool `json:"disable_appsync,omitempty"`
   // CreationTime
   CreationTime float64 `json:"creation_time,omitempty"`
   // LastModified
   LastModified float64 `json:"last_modified,omitempty"`
   // LastModSchedTime
   LastModSchedTime float64 `json:"last_mod_sched_time,omitempty"`
   // LastReplicatedSnapcollName
   LastReplicatedSnapcollName string `json:"last_replicated_snapcoll_name,omitempty"`
   // LastReplicatedSnapcollID
   LastReplicatedSnapcollID string `json:"last_replicated_snapcoll_id,omitempty"`
   // LastReplicatedAtTime
   LastReplicatedAtTime float64 `json:"last_replicated_at_time,omitempty"`
   // LastSnapTime
   LastSnapTime float64 `json:"last_snap_time,omitempty"`
   // NextSnapTime
   NextSnapTime float64 `json:"next_snap_time,omitempty"`
   // NextReplSnapTime
   NextReplSnapTime float64 `json:"next_repl_snap_time,omitempty"`
   // SnapCounter
   SnapCounter float64 `json:"snap_counter,omitempty"`
   // SchedOwnerID
   SchedOwnerID string `json:"sched_owner_id,omitempty"`
   // SchedOwnerName
   SchedOwnerName string `json:"sched_owner_name,omitempty"`
   // LastConfigChangeTime
   LastConfigChangeTime float64 `json:"last_config_change_time,omitempty"`
   // UseDownstreamForDr
   UseDownstreamForDr bool `json:"use_downstream_for_dr,omitempty"`
}
