/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Group - Group is a collection of arrays operating together organized into storage pools.
// Export GroupFields for advance operations like search filter etc.
var GroupFields *Group

func init(){
	IDfield:= "id"
	Namefield:= "name"
	SmtpServerfield:= "smtp_server"
	SmtpAuthUsernamefield:= "smtp_auth_username"
	SmtpAuthPasswordfield:= "smtp_auth_password"
	ProxyServerfield:= "proxy_server"
	ProxyUsernamefield:= "proxy_username"
	ProxyPasswordfield:= "proxy_password"
	AlertToEmailAddrsfield:= "alert_to_email_addrs"
	AlertFromEmailAddrfield:= "alert_from_email_addr"
	IsnsServerfield:= "isns_server"
	SnmpTrapHostfield:= "snmp_trap_host"
	SnmpCommunityfield:= "snmp_community"
	SnmpSysContactfield:= "snmp_sys_contact"
	SnmpSysLocationfield:= "snmp_sys_location"
	DomainNamefield:= "domain_name"
	NtpServerfield:= "ntp_server"
	Timezonefield:= "timezone"
	SyslogdServerfield:= "syslogd_server"
	TdzPrefixfield:= "tdz_prefix"
	GroupTargetNamefield:= "group_target_name"
	MergeGroupNamefield:= "merge_group_name"
	LeaderArrayNamefield:= "leader_array_name"
	LeaderArraySerialfield:= "leader_array_serial"
	ManagementServiceBackupArrayNamefield:= "management_service_backup_array_name"
	VersionCurrentfield:= "version_current"
	VersionTargetfield:= "version_target"
	VersionRollbackfield:= "version_rollback"
	UpdateArrayNamesfield:= "update_array_names"
	UpdateProgressMsgfield:= "update_progress_msg"
	UpdateErrorCodefield:= "update_error_code"
	UpdateDownloadErrorCodefield:= "update_download_error_code"
	ScsiVendorIDfield:= "scsi_vendor_id"
	LastLoginfield:= "last_login"
	LoginBannerMessagefield:= "login_banner_message"
		
	GroupFields= &Group{
		ID: &IDfield,
		Name: &Namefield,
		SmtpServer: &SmtpServerfield,
		SmtpAuthUsername: &SmtpAuthUsernamefield,
		SmtpAuthPassword: &SmtpAuthPasswordfield,
		ProxyServer: &ProxyServerfield,
		ProxyUsername: &ProxyUsernamefield,
		ProxyPassword: &ProxyPasswordfield,
		AlertToEmailAddrs: &AlertToEmailAddrsfield,
		AlertFromEmailAddr: &AlertFromEmailAddrfield,
		IsnsServer: &IsnsServerfield,
		SnmpTrapHost: &SnmpTrapHostfield,
		SnmpCommunity: &SnmpCommunityfield,
		SnmpSysContact: &SnmpSysContactfield,
		SnmpSysLocation: &SnmpSysLocationfield,
		DomainName: &DomainNamefield,
		NtpServer: &NtpServerfield,
		Timezone: &Timezonefield,
		SyslogdServer: &SyslogdServerfield,
		TdzPrefix: &TdzPrefixfield,
		GroupTargetName: &GroupTargetNamefield,
		MergeGroupName: &MergeGroupNamefield,
		LeaderArrayName: &LeaderArrayNamefield,
		LeaderArraySerial: &LeaderArraySerialfield,
		ManagementServiceBackupArrayName: &ManagementServiceBackupArrayNamefield,
		VersionCurrent: &VersionCurrentfield,
		VersionTarget: &VersionTargetfield,
		VersionRollback: &VersionRollbackfield,
		UpdateArrayNames: &UpdateArrayNamesfield,
		UpdateProgressMsg: &UpdateProgressMsgfield,
		UpdateErrorCode: &UpdateErrorCodefield,
		UpdateDownloadErrorCode: &UpdateDownloadErrorCodefield,
		ScsiVendorID: &ScsiVendorIDfield,
		LastLogin: &LastLoginfield,
		LoginBannerMessage: &LoginBannerMessagefield,
		
	}
}

type Group struct {
   
    // Identifier of the group.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the group.
    
 	Name *string `json:"name,omitempty"`
   
    // Hostname or IP Address of SMTP Server.
    
 	SmtpServer *string `json:"smtp_server,omitempty"`
   
    // Port number of SMTP Server.
    
   	SmtpPort *int64 `json:"smtp_port,omitempty"`
   
    // Whether SMTP Server requires authentication.
    
 	SmtpAuthEnabled *bool `json:"smtp_auth_enabled,omitempty"`
   
    // Username to authenticate with SMTP Server.
    
 	SmtpAuthUsername *string `json:"smtp_auth_username,omitempty"`
   
    // Password to authenticate with SMTP Server.
    
 	SmtpAuthPassword *string `json:"smtp_auth_password,omitempty"`
   
    // Level of encryption for SMTP. Requires use of SMTP Authentication if encryption is enabled.
    
   	SmtpEncryptType *NsSmtpEncryptType `json:"smtp_encrypt_type,omitempty"`
   
    // Whether to send autosupport.
    
 	AutosupportEnabled *bool `json:"autosupport_enabled,omitempty"`
   
    // Specify whether to allow HPE Nimble Storage to use Google Analytics in the GUI.  HPE Nimble Storage uses Google Analytics to gather data related to GUI usage.  The data gathered is used to evaluate and improve the product.
    
 	AllowAnalyticsGui *bool `json:"allow_analytics_gui,omitempty"`
   
    // Whether to allow support tunnel.
    
 	AllowSupportTunnel *bool `json:"allow_support_tunnel,omitempty"`
   
    // Hostname or IP Address of HTTP Proxy Server. Setting this attribute to an empty string will unset all proxy settings.
    
 	ProxyServer *string `json:"proxy_server,omitempty"`
   
    // Proxy Port of HTTP Proxy Server.
    
   	ProxyPort *int64 `json:"proxy_port,omitempty"`
   
    // Username to authenticate with HTTP Proxy Server.
    
 	ProxyUsername *string `json:"proxy_username,omitempty"`
   
    // Password to authenticate with HTTP Proxy Server.
    
 	ProxyPassword *string `json:"proxy_password,omitempty"`
   
    // Comma-separated list of email addresss to receive emails.
    
 	AlertToEmailAddrs *string `json:"alert_to_email_addrs,omitempty"`
   
    // Whether to send alert to Support.
    
 	SendAlertToSupport *bool `json:"send_alert_to_support,omitempty"`
   
    // From email address to use while sending emails.
    
 	AlertFromEmailAddr *string `json:"alert_from_email_addr,omitempty"`
   
    // Minimum level of alert to be notified.
    
   	AlertMinLevel *NsSeverityLevel `json:"alert_min_level,omitempty"`
   
    // Whether iSNS is enabled.
    
 	IsnsEnabled *bool `json:"isns_enabled,omitempty"`
   
    // Hostname or IP Address of iSNS Server.
    
 	IsnsServer *string `json:"isns_server,omitempty"`
   
    // Port number for iSNS Server.
    
   	IsnsPort *int64 `json:"isns_port,omitempty"`
   
    // Whether to enable SNMP traps.
    
 	SnmpTrapEnabled *bool `json:"snmp_trap_enabled,omitempty"`
   
    // Hostname or IP Address to send SNMP traps.
    
 	SnmpTrapHost *string `json:"snmp_trap_host,omitempty"`
   
    // Port number of SNMP trap host.
    
   	SnmpTrapPort *int64 `json:"snmp_trap_port,omitempty"`
   
    // Whether to accept SNMP get commands.
    
 	SnmpGetEnabled *bool `json:"snmp_get_enabled,omitempty"`
   
    // Community string to be used with SNMP.
    
 	SnmpCommunity *string `json:"snmp_community,omitempty"`
   
    // Port number to which SNMP get requests should be sent.
    
   	SnmpGetPort *int64 `json:"snmp_get_port,omitempty"`
   
    // Name of the SNMP administrator.
    
 	SnmpSysContact *string `json:"snmp_sys_contact,omitempty"`
   
    // Location of the group.
    
 	SnmpSysLocation *string `json:"snmp_sys_location,omitempty"`
   
    // Domain name for this group.
    
 	DomainName *string `json:"domain_name,omitempty"`
   
    // IP addresses for this group's dns servers.
    
   	DnsServers []*NsIPAddressObject `json:"dns_servers,omitempty"`
   
    // Either IP address or hostname of the NTP server for this group.
    
 	NtpServer *string `json:"ntp_server,omitempty"`
   
    // Timezone in which this group is located.
    
 	Timezone *string `json:"timezone,omitempty"`
   
    // The amount of time in seconds that the user session is inactive before timing out.
    
   	UserInactivityTimeout *int64 `json:"user_inactivity_timeout,omitempty"`
   
    // Is syslogd enabled on this system.
    
 	SyslogdEnabled *bool `json:"syslogd_enabled,omitempty"`
   
    // Hostname of the syslogd server.
    
 	SyslogdServer *string `json:"syslogd_server,omitempty"`
   
    // Port number for syslogd server.
    
   	SyslogdPort *int64 `json:"syslogd_port,omitempty"`
   
    // Are vvols enabled on this group.
    
 	VvolEnabled *bool `json:"vvol_enabled,omitempty"`
   
    // Whether iSCSI is enabled on this group.
    
 	IscsiEnabled *bool `json:"iscsi_enabled,omitempty"`
   
    // Whether FC is enabled on this group.
    
 	FcEnabled *bool `json:"fc_enabled,omitempty"`
   
    // Are new volume and volume collection names transformed on this group.
    
 	UniqueNameEnabled *bool `json:"unique_name_enabled,omitempty"`
   
    // Protocol used to access this group.
    
	AccessProtocolList []*NsAccessProtocol `json:"access_protocol_list,omitempty"`
   
    // Is group_target enabled on this group.
    
 	GroupTargetEnabled *bool `json:"group_target_enabled,omitempty"`
   
    // Newly created volumes are exported under iSCSI Group Target or iSCSI Volume Target.
    
   	DefaultIscsiTargetScope *NsTargetScope `json:"default_iscsi_target_scope,omitempty"`
   
    // Is Target Driven Zoning (TDZ) enabled on this group.
    
 	TdzEnabled *bool `json:"tdz_enabled,omitempty"`
   
    // Target Driven Zoning (TDZ) prefix for peer zones created by TDZ.
    
 	TdzPrefix *string `json:"tdz_prefix,omitempty"`
   
    // Iscsi target name for this group.
    
 	GroupTargetName *string `json:"group_target_name,omitempty"`
   
    // Amount of space to reserve for a volume as a percentage of volume size.
    
   	DefaultVolumeReserve *int64 `json:"default_volume_reserve,omitempty"`
   
    // Default threshold for volume space usage as a percentage of volume size above which an alert is raised.
    
   	DefaultVolumeWarnLevel *int64 `json:"default_volume_warn_level,omitempty"`
   
    // Default limit for a volume space usage as a percentage of volume size. Volume will be taken offline/made non-writable on exceeding its limit.
    
   	DefaultVolumeLimit *int64 `json:"default_volume_limit,omitempty"`
   
    // Amount of space to reserve for snapshots of a volume as a percentage of volume size.
    
   	DefaultSnapReserve *int64 `json:"default_snap_reserve,omitempty"`
   
    // Default threshold for snapshot space usage of a volume as a percentage of volume size above which an alert is raised.
    
   	DefaultSnapWarnLevel *int64 `json:"default_snap_warn_level,omitempty"`
   
    // This attribute is deprecated. The array does not limit a volume's snapshot space usage. The attribute is ignored on input and returns max int64 value on output.
    
   	DefaultSnapLimit *int64 `json:"default_snap_limit,omitempty"`
   
    // This attribute is deprecated. The array does not limit a volume's snapshot space usage. The attribute is ignored on input and returns -1 on output.
    
  	DefaultSnapLimitPercent  *int64 `json:"default_snap_limit_percent,omitempty"`
   
    // Whether alarm feature is enabled.
    
 	AlarmsEnabled *bool `json:"alarms_enabled,omitempty"`
   
    // The amount of time in seconds to validate Microsoft VSS application synchronization before timing out.
    
   	VssValIDationTimeout *int64 `json:"vss_validation_timeout,omitempty"`
   
    // Whether automatic switchover of Group management services feature is enabled.
    
 	AutoSwitchoverEnabled *bool `json:"auto_switchover_enabled,omitempty"`
   
    // List of validation messages for automatic switchover of Group Management. This will be empty when there are no conflicts found.
    
   	AutoSwitchoverMessages []*NsErrorWithArguments `json:"auto_switchover_messages,omitempty"`
   
    // State of group merge.
    
   	MergeState *NsMergeState `json:"merge_state,omitempty"`
   
    // Group that we're being merged with.
    
 	MergeGroupName *string `json:"merge_group_name,omitempty"`
   
    // Enable or disable TLSv1.0 and TLSv1.1.
    
 	Tlsv1Enabled *bool `json:"tlsv1_enabled,omitempty"`
   
    // Enable or disable Common Criteria mode.
    
 	CcModeEnabled *bool `json:"cc_mode_enabled,omitempty"`
   
    // Snapshot Time-to-live(TTL) configured at group level for automatic deletion of unmanaged snapshots. Value 0 indicates unlimited TTL.
    
  	GroupSnapshotTtl  *int64 `json:"group_snapshot_ttl,omitempty"`
   
    // Unit for unmanaged snapshot time to live.
    
   	AutocleanUnmanagedSnapshotsTtlUnit *int64 `json:"autoclean_unmanaged_snapshots_ttl_unit,omitempty"`
   
    // Whether autoclean unmanaged snapshots feature is enabled.
    
 	AutocleanUnmanagedSnapshotsEnabled *bool `json:"autoclean_unmanaged_snapshots_enabled,omitempty"`
   
    // Name of the array where the group Management Service is running.
    
 	LeaderArrayName *string `json:"leader_array_name,omitempty"`
   
    // Serial number of the array where the group Management Service is running.
    
 	LeaderArraySerial *string `json:"leader_array_serial,omitempty"`
   
    // Name of the array where backup the group Management Service is running.
    
 	ManagementServiceBackupArrayName *string `json:"management_service_backup_array_name,omitempty"`
   
    // HA status of the group Management Service.
    
   	ManagementServiceBackupStatus *NsGroupHAStatus `json:"management_service_backup_status,omitempty"`
   
    // Failover mode of the group Management Service.
    
   	FailoverMode *NsGroupFailoverMode `json:"failover_mode,omitempty"`
   
    // Witness status from group Management Service array and group Management Service backup array.
    
   	WitnessStatus []*NsWitnessTestResponse `json:"witness_status,omitempty"`
   
    // Members of this group.
    
	MemberList []*string `json:"member_list,omitempty"`
   
    // Compressed usage of volumes in the group.
    
   	CompressedVolUsageBytes *int64 `json:"compressed_vol_usage_bytes,omitempty"`
   
    // Compressed usage of snapshots in the group.
    
   	CompressedSnapUsageBytes *int64 `json:"compressed_snap_usage_bytes,omitempty"`
   
    // Uncompressed usage of volumes in the group.
    
   	UncompressedVolUsageBytes *int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
   
    // Uncompressed usage of snapshots in the group.
    
   	UncompressedSnapUsageBytes *int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
   
    // Usable capacity bytes of the group.
    
   	UsableCapacityBytes *int64 `json:"usable_capacity_bytes,omitempty"`
   
    // Used space of the group in bytes.
    
   	Usage *int64 `json:"usage,omitempty"`
   
    // Total capacity of the group.
    
   	RawCapacity *int64 `json:"raw_capacity,omitempty"`
   
    // Usable cache capacity of the group.
    
   	UsableCacheCapacity *int64 `json:"usable_cache_capacity,omitempty"`
   
    // Total cache capacity of the group.
    
   	RawCacheCapacity *int64 `json:"raw_cache_capacity,omitempty"`
   
    // Total snapshot usage as if each snapshot is deep copy of the volume.
    
   	SnapUsagePopulated *int64 `json:"snap_usage_populated,omitempty"`
   
    // Usage for blocks that are not yet deleted.
    
   	PendingDeletes *int64 `json:"pending_deletes,omitempty"`
   
    // Number of connections to the group.
    
   	NumConnections *int64 `json:"num_connections,omitempty"`
   
    // Compression ratio of volumes in the group.
    
  	VolCompressionRatio *float32 `json:"vol_compression_ratio,omitempty"`
   
    // Compression ratio of snapshots in the group.
    
  	SnapCompressionRatio *float32 `json:"snap_compression_ratio,omitempty"`
   
    // Compression savings for the group expressed as ratio.
    
  	CompressionRatio *float32 `json:"compression_ratio,omitempty"`
   
    // Dedupe savings for the group expressed as ratio.
    
  	DedupeRatio *float32 `json:"dedupe_ratio,omitempty"`
   
    // Clone savings for the group expressed as ratio.
    
  	CloneRatio *float32 `json:"clone_ratio,omitempty"`
   
    // Thin provisioning savings for volumes in the group expressed as ratio.
    
  	VolThinProvisioningRatio *float32 `json:"vol_thin_provisioning_ratio,omitempty"`
   
    // Overall savings in the group expressed as ratio.
    
  	SavingsRatio *float32 `json:"savings_ratio,omitempty"`
   
    // Space savings in the group that does not include thin-provisioning savings expressed as ratio.
    
  	DataReductionRatio *float32 `json:"data_reduction_ratio,omitempty"`
   
    // Space usage savings in the group due to deduplication.
    
   	SavingsDedupe *int64 `json:"savings_dedupe,omitempty"`
   
    // Space usage savings in the group due to compression.
    
   	SavingsCompression *int64 `json:"savings_compression,omitempty"`
   
    // Space usage savings in the group due to cloning of volumes.
    
   	SavingsClone *int64 `json:"savings_clone,omitempty"`
   
    // Space usage savings in the group due to thin provisioning of volumes.
    
   	SavingsVolThinProvisioning *int64 `json:"savings_vol_thin_provisioning,omitempty"`
   
    // Space usage savings in the group that does not include thin-provisioning savings.
    
   	SavingsDataReduction *int64 `json:"savings_data_reduction,omitempty"`
   
    // Overall space usage savings in the group.
    
   	Savings *int64 `json:"savings,omitempty"`
   
    // Free space of the pool in bytes.
    
   	FreeSpace *int64 `json:"free_space,omitempty"`
   
    // Reserved space that is not utilized.
    
   	UnusedReserveBytes *int64 `json:"unused_reserve_bytes,omitempty"`
   
    // Indicates whether the usage of group is valid.
    
 	UsageValID *bool `json:"usage_valid,omitempty"`
   
    // Is space info for this group valid.
    
 	SpaceInfoValID *bool `json:"space_info_valid,omitempty"`
   
    // Version of software running on the group.
    
 	VersionCurrent *string `json:"version_current,omitempty"`
   
    // Desired software version for the group.
    
 	VersionTarget *string `json:"version_target,omitempty"`
   
    // Rollback software version for the group.
    
 	VersionRollback *string `json:"version_rollback,omitempty"`
   
    // Group update state.
    
   	UpdateState *NsUpdateState `json:"update_state,omitempty"`
   
    // Start time of last update.
    
   	UpdateStartTime *int64 `json:"update_start_time,omitempty"`
   
    // End time of last update.
    
   	UpdateEndTime *int64 `json:"update_end_time,omitempty"`
   
    // Arrays in the group undergoing update.
    
 	UpdateArrayNames *string `json:"update_array_names,omitempty"`
   
    // Group update detailed progress message.
    
 	UpdateProgressMsg *string `json:"update_progress_msg,omitempty"`
   
    // If the software update has failed, this indicates the error code corresponding to the failure.
    
 	UpdateErrorCode *string `json:"update_error_code,omitempty"`
   
    // Is software update package currently downloading.
    
 	UpdateDownloading *bool `json:"update_downloading,omitempty"`
   
    // If the software download has failed, this indicates the error code corresponding to the failure.
    
 	UpdateDownloadErrorCode *string `json:"update_download_error_code,omitempty"`
   
    // Start time of last update.
    
   	UpdateDownloadStartTime *int64 `json:"update_download_start_time,omitempty"`
   
    // End time of last update.
    
   	UpdateDownloadEndTime *int64 `json:"update_download_end_time,omitempty"`
   
    // Is iscsi reconnection automatic.
    
 	IscsiAutomaticConnectionMethod *bool `json:"iscsi_automatic_connection_method,omitempty"`
   
    // Does iscsi automatically rebalance connections.
    
 	IscsiConnectionRebalancing *bool `json:"iscsi_connection_rebalancing,omitempty"`
   
    // Current bandwidth throttle for replication, expressed either as megabits per second or as -1 to indicate that there is no throttle.
    
  	ReplThrottledBandwIDth  *int64 `json:"repl_throttled_bandwidth,omitempty"`
   
    // Current bandwidth throttle for replication, expressed either as kilobits per second or as -1 to indicate that there is no throttle.
    
  	ReplThrottledBandwIDthKbps  *int64 `json:"repl_throttled_bandwidth_kbps,omitempty"`
   
    // All the replication bandwidth limits on the system.
    
   	ReplThrottleList []*NsThrottle `json:"repl_throttle_list,omitempty"`
   
    // Status of data migration activity related to volumes being relocated to different pools.
    
   	VolumeMigrationStatus []*NsVolFamMigStatus `json:"volume_migration_status,omitempty"`
   
    // Data migration status for arrays being removed from their pool.
    
   	ArrayUnassignMigrationStatus []*NsArrayUnassignMigStatus `json:"array_unassign_migration_status,omitempty"`
   
    // Status of data rebalancing operations for pools in the group.
    
   	DataRebalanceStatus []*NsPoolRebalanceMigStatus `json:"data_rebalance_status,omitempty"`
   
    // SCSI vendor ID.
    
 	ScsiVendorID *string `json:"scsi_vendor_id,omitempty"`
   
    // How encryption is configured for this group.
    
	EncryptionConfig *NsEncryptionSettings `json:"encryption_config,omitempty"`
   
    // Time and user of last login to this group.
    
 	LastLogin *string `json:"last_login,omitempty"`
   
    // Number of snapshots in the group.
    
   	NumSnaps *int64 `json:"num_snaps,omitempty"`
   
    // Number of snapshot collections in this group.
    
   	NumSnapcolls *int64 `json:"num_snapcolls,omitempty"`
   
    // Unix epoch time local to the group.
    
   	Date *int64 `json:"date,omitempty"`
   
    // The message for the login banner that is displayed during user login activity.
    
 	LoginBannerMessage *string `json:"login_banner_message,omitempty"`
   
    // Should the banner be displayed before the user credentials are prompted or after prompting the user credentials.
    
 	LoginBannerAfterAuth *bool `json:"login_banner_after_auth,omitempty"`
   
    // This will reset the banner to the version of the installed NOS. When login_banner_after_auth is specified, login_banner_reset can not be set to true.
    
 	LoginBannerReset *bool `json:"login_banner_reset,omitempty"`
   
    // Threshold for considering a volume as high retention.
    
   	SnapRetnMeterHigh *int64 `json:"snap_retn_meter_high,omitempty"`
   
    // Threshold for considering a volume as very high retention.
    
   	SnapRetnMeterVeryHigh *int64 `json:"snap_retn_meter_very_high,omitempty"`
}
