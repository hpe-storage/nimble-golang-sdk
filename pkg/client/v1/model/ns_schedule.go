/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/NsSchedule


// NsSchedule :
type NsSchedule struct {
   // ID
   ID string `json:"id,omitempty"`
   // ScheduleID
   ScheduleID string `json:"schedule_id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // ScheduleName
   ScheduleName string `json:"schedule_name,omitempty"`
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
   // DownstreamPartnerID
   DownstreamPartnerID string `json:"downstream_partner_id,omitempty"`
   // DownstreamPartnerName
   DownstreamPartnerName string `json:"downstream_partner_name,omitempty"`
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
}
