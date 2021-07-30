// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsScheduleFields provides field names to use in filter parameters, for example.
var NsScheduleFields *NsScheduleFieldHandles

func init() {
	NsScheduleFields = &NsScheduleFieldHandles{
		ID:                     "id",
		ScheduleId:             "schedule_id",
		Name:                   "name",
		ScheduleName:           "schedule_name",
		Period:                 "period",
		PeriodUnit:             "period_unit",
		AtTime:                 "at_time",
		UntilTime:              "until_time",
		Days:                   "days",
		NumRetain:              "num_retain",
		DownstreamPartner:      "downstream_partner",
		DownstreamPartnerId:    "downstream_partner_id",
		DownstreamPartnerName:  "downstream_partner_name",
		ReplicateEvery:         "replicate_every",
		NumRetainReplica:       "num_retain_replica",
		ReplAlertThres:         "repl_alert_thres",
		SnapVerify:             "snap_verify",
		SkipDbConsistencyCheck: "skip_db_consistency_check",
		DisableAppsync:         "disable_appsync",
		ScheduleType:           "schedule_type",
		Active:                 "active",
	}
}

// NsSchedule - Protection schedule associated with a volume collection or protection template.
type NsSchedule struct {
	// ID - ID of protection schedule.
	ID *string `json:"id,omitempty"`
	// ScheduleId - ID of protection schedule.
	ScheduleId *string `json:"schedule_id,omitempty"`
	// Name - Name of protection schedule.
	Name *string `json:"name,omitempty"`
	// ScheduleName - Name of protection schedule.
	ScheduleName *string `json:"schedule_name,omitempty"`
	// Period - Frequency of snapshots.
	Period *int64 `json:"period,omitempty"`
	// PeriodUnit - Units for repeat frequency -- minutes, hours, days or weeks.
	PeriodUnit *NsPeriodUnit `json:"period_unit,omitempty"`
	// AtTime - Time of day when snapshot should be taken.
	AtTime *int64 `json:"at_time,omitempty"`
	// UntilTime - Time of day to stop taking snapshots.
	UntilTime *int64 `json:"until_time,omitempty"`
	// Days - Which days snapshots should be taken.
	Days *string `json:"days,omitempty"`
	// NumRetain - Number of snapshots to retain.
	NumRetain *int64 `json:"num_retain,omitempty"`
	// DownstreamPartner - Partner name if snapshots created by this schedule should be replicated.
	DownstreamPartner *string `json:"downstream_partner,omitempty"`
	// DownstreamPartnerId - Partner ID if snapshots created by this schedule should be replicated.
	DownstreamPartnerId *string `json:"downstream_partner_id,omitempty"`
	// DownstreamPartnerName - Partner name if snapshots created by this schedule should be replicated.
	DownstreamPartnerName *string `json:"downstream_partner_name,omitempty"`
	// ReplicateEvery - Which snapshots should be replicated.
	ReplicateEvery *int64 `json:"replicate_every,omitempty"`
	// NumRetainReplica - Number of snapshots to retain on the replica.
	NumRetainReplica *int64 `json:"num_retain_replica,omitempty"`
	// ReplAlertThres - Replication alert threshold.
	ReplAlertThres *int64 `json:"repl_alert_thres,omitempty"`
	// SnapVerify - Run verification tool on snapshot created by this schedule.
	SnapVerify *bool `json:"snap_verify,omitempty"`
	// SkipDbConsistencyCheck - Skip consistency check for database files on snapshots created by this schedule.
	SkipDbConsistencyCheck *bool `json:"skip_db_consistency_check,omitempty"`
	// DisableAppsync - Disables application synchronized snapshots and creates crash consistent snapshots instead.
	DisableAppsync *bool `json:"disable_appsync,omitempty"`
	// ScheduleType - Schedule type: regular or external_trigger.
	ScheduleType *NsScheduleType `json:"schedule_type,omitempty"`
	// Active - A schedule is active only if it is owned by the same owner as the volume collection. Only active schedules of a volume collection participate in the creation of snapshots and replication.
	Active *bool `json:"active,omitempty"`
}

// NsScheduleFieldHandles provides a string representation for each AccessControlRecord field.
type NsScheduleFieldHandles struct {
	ID                     string
	ScheduleId             string
	Name                   string
	ScheduleName           string
	Period                 string
	PeriodUnit             string
	AtTime                 string
	UntilTime              string
	Days                   string
	NumRetain              string
	DownstreamPartner      string
	DownstreamPartnerId    string
	DownstreamPartnerName  string
	ReplicateEvery         string
	NumRetainReplica       string
	ReplAlertThres         string
	SnapVerify             string
	SkipDbConsistencyCheck string
	DisableAppsync         string
	ScheduleType           string
	Active                 string
}
