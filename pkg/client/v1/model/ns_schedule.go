/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsSchedule - Protection schedule associated with a volume collection or protection template.
// Export NsScheduleFields for advance operations like search filter etc.
var NsScheduleFields *NsSchedule

func init(){
	IDfield:= "id"
	ScheduleIDfield:= "schedule_id"
	Namefield:= "name"
	ScheduleNamefield:= "schedule_name"
	Daysfield:= "days"
	DownstreamPartnerfield:= "downstream_partner"
	DownstreamPartnerIDfield:= "downstream_partner_id"
	DownstreamPartnerNamefield:= "downstream_partner_name"
		
	NsScheduleFields= &NsSchedule{
		ID: &IDfield,
		ScheduleID: &ScheduleIDfield,
		Name: &Namefield,
		ScheduleName: &ScheduleNamefield,
		Days: &Daysfield,
		DownstreamPartner: &DownstreamPartnerfield,
		DownstreamPartnerID: &DownstreamPartnerIDfield,
		DownstreamPartnerName: &DownstreamPartnerNamefield,
		
	}
}

type NsSchedule struct {
   
    // ID of protection schedule.
    
 	ID *string `json:"id,omitempty"`
   
    // ID of protection schedule.
    
 	ScheduleID *string `json:"schedule_id,omitempty"`
   
    // Name of protection schedule.
    
 	Name *string `json:"name,omitempty"`
   
    // Name of protection schedule.
    
 	ScheduleName *string `json:"schedule_name,omitempty"`
   
    // Frequency of snapshots.
    
   	Period *int64 `json:"period,omitempty"`
   
    // Units for repeat frequency -- minutes, hours, days or weeks.
    
   	PeriodUnit *NsPeriodUnit `json:"period_unit,omitempty"`
   
    // Time of day when snapshot should be taken.
    
   	AtTime *int64 `json:"at_time,omitempty"`
   
    // Time of day to stop taking snapshots.
    
   	UntilTime *int64 `json:"until_time,omitempty"`
   
    // Which days snapshots should be taken.
    
 	Days *string `json:"days,omitempty"`
   
    // Number of snapshots to retain.
    
   	NumRetain *int64 `json:"num_retain,omitempty"`
   
    // Partner name if snapshots created by this schedule should be replicated.
    
 	DownstreamPartner *string `json:"downstream_partner,omitempty"`
   
    // Partner ID if snapshots created by this schedule should be replicated.
    
 	DownstreamPartnerID *string `json:"downstream_partner_id,omitempty"`
   
    // Partner name if snapshots created by this schedule should be replicated.
    
 	DownstreamPartnerName *string `json:"downstream_partner_name,omitempty"`
   
    // Which snapshots should be replicated.
    
   	ReplicateEvery *int64 `json:"replicate_every,omitempty"`
   
    // Number of snapshots to retain on the replica.
    
   	NumRetainReplica *int64 `json:"num_retain_replica,omitempty"`
   
    // Replication alert threshold.
    
   	ReplAlertThres *int64 `json:"repl_alert_thres,omitempty"`
   
    // Run verification tool on snapshot created by this schedule.
    
 	SnapVerify *bool `json:"snap_verify,omitempty"`
   
    // Skip consistency check for database files on snapshots created by this schedule.
    
 	SkipDbConsistencyCheck *bool `json:"skip_db_consistency_check,omitempty"`
   
    // Disables application synchronized snapshots and creates crash consistent snapshots instead.
    
 	DisableAppsync *bool `json:"disable_appsync,omitempty"`
   
    // Schedule type: regular or external_trigger.
    
   	ScheduleType *NsScheduleType `json:"schedule_type,omitempty"`
   
    // A schedule is active only if it is owned by the same owner as the volume collection. Only active schedules of a volume collection participate in the creation of snapshots and replication.
    
 	Active *bool `json:"active,omitempty"`
}
