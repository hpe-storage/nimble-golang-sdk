// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export GroupFields provides field names to use in filter parameters, for example.
var GroupFields *GroupFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldSmtpServer := "smtp_server"
	fieldSmtpPort := "smtp_port"
	fieldSmtpAuthEnabled := "smtp_auth_enabled"
	fieldSmtpAuthUsername := "smtp_auth_username"
	fieldSmtpAuthPassword := "smtp_auth_password"
	fieldSmtpEncryptType := "smtp_encrypt_type"
	fieldAutosupportEnabled := "autosupport_enabled"
	fieldAllowAnalyticsGui := "allow_analytics_gui"
	fieldAllowSupportTunnel := "allow_support_tunnel"
	fieldProxyServer := "proxy_server"
	fieldProxyPort := "proxy_port"
	fieldProxyUsername := "proxy_username"
	fieldProxyPassword := "proxy_password"
	fieldAlertToEmailAddrs := "alert_to_email_addrs"
	fieldSendAlertToSupport := "send_alert_to_support"
	fieldAlertFromEmailAddr := "alert_from_email_addr"
	fieldAlertMinLevel := "alert_min_level"
	fieldIsnsEnabled := "isns_enabled"
	fieldIsnsServer := "isns_server"
	fieldIsnsPort := "isns_port"
	fieldSnmpTrapEnabled := "snmp_trap_enabled"
	fieldSnmpTrapHost := "snmp_trap_host"
	fieldSnmpTrapPort := "snmp_trap_port"
	fieldSnmpGetEnabled := "snmp_get_enabled"
	fieldSnmpCommunity := "snmp_community"
	fieldSnmpGetPort := "snmp_get_port"
	fieldSnmpSysContact := "snmp_sys_contact"
	fieldSnmpSysLocation := "snmp_sys_location"
	fieldDomainName := "domain_name"
	fieldDnsServers := "dns_servers"
	fieldNtpServer := "ntp_server"
	fieldTimezone := "timezone"
	fieldUserInactivityTimeout := "user_inactivity_timeout"
	fieldSyslogdEnabled := "syslogd_enabled"
	fieldSyslogdServer := "syslogd_server"
	fieldSyslogdPort := "syslogd_port"
	fieldSyslogdServers := "syslogd_servers"
	fieldVvolEnabled := "vvol_enabled"
	fieldSetupProgress := "setup_progress"
	fieldIscsiEnabled := "iscsi_enabled"
	fieldFcEnabled := "fc_enabled"
	fieldUniqueNameEnabled := "unique_name_enabled"
	fieldAccessProtocolList := "access_protocol_list"
	fieldGroupTargetEnabled := "group_target_enabled"
	fieldDefaultIscsiTargetScope := "default_iscsi_target_scope"
	fieldTdzEnabled := "tdz_enabled"
	fieldTdzPrefix := "tdz_prefix"
	fieldGroupTargetName := "group_target_name"
	fieldDefaultVolumeReserve := "default_volume_reserve"
	fieldDefaultVolumeWarnLevel := "default_volume_warn_level"
	fieldDefaultVolumeLimit := "default_volume_limit"
	fieldDefaultSnapReserve := "default_snap_reserve"
	fieldDefaultSnapWarnLevel := "default_snap_warn_level"
	fieldDefaultSnapLimit := "default_snap_limit"
	fieldDefaultSnapLimitPercent := "default_snap_limit_percent"
	fieldAlarmsEnabled := "alarms_enabled"
	fieldVssValidationTimeout := "vss_validation_timeout"
	fieldAutoSwitchoverEnabled := "auto_switchover_enabled"
	fieldSoftwareSubscriptionEnabled := "software_subscription_enabled"
	fieldAutoSwitchoverMessages := "auto_switchover_messages"
	fieldMergeState := "merge_state"
	fieldMergeGroupName := "merge_group_name"
	fieldTlsv1Enabled := "tlsv1_enabled"
	fieldCcModeEnabled := "cc_mode_enabled"
	fieldCertIpsIncluded := "cert_ips_included"
	fieldGroupSnapshotTtl := "group_snapshot_ttl"
	fieldAutocleanUnmanagedSnapshotsTtlUnit := "autoclean_unmanaged_snapshots_ttl_unit"
	fieldAutocleanUnmanagedSnapshotsEnabled := "autoclean_unmanaged_snapshots_enabled"
	fieldLeaderArrayName := "leader_array_name"
	fieldLeaderArraySerial := "leader_array_serial"
	fieldManagementServiceBackupArrayName := "management_service_backup_array_name"
	fieldManagementServiceBackupStatus := "management_service_backup_status"
	fieldFailoverMode := "failover_mode"
	fieldWitnessStatus := "witness_status"
	fieldMemberList := "member_list"
	fieldCompressedVolUsageBytes := "compressed_vol_usage_bytes"
	fieldCompressedSnapUsageBytes := "compressed_snap_usage_bytes"
	fieldUncompressedVolUsageBytes := "uncompressed_vol_usage_bytes"
	fieldUncompressedSnapUsageBytes := "uncompressed_snap_usage_bytes"
	fieldUsableCapacityBytes := "usable_capacity_bytes"
	fieldUsage := "usage"
	fieldRawCapacity := "raw_capacity"
	fieldUsableCacheCapacity := "usable_cache_capacity"
	fieldRawCacheCapacity := "raw_cache_capacity"
	fieldSnapUsagePopulated := "snap_usage_populated"
	fieldPendingDeletes := "pending_deletes"
	fieldNumConnections := "num_connections"
	fieldVolCompressionRatio := "vol_compression_ratio"
	fieldSnapCompressionRatio := "snap_compression_ratio"
	fieldCompressionRatio := "compression_ratio"
	fieldDedupeRatio := "dedupe_ratio"
	fieldCloneRatio := "clone_ratio"
	fieldVolThinProvisioningRatio := "vol_thin_provisioning_ratio"
	fieldSavingsRatio := "savings_ratio"
	fieldDataReductionRatio := "data_reduction_ratio"
	fieldSavingsDedupe := "savings_dedupe"
	fieldSavingsCompression := "savings_compression"
	fieldSavingsClone := "savings_clone"
	fieldSavingsVolThinProvisioning := "savings_vol_thin_provisioning"
	fieldSavingsDataReduction := "savings_data_reduction"
	fieldSavings := "savings"
	fieldFreeSpace := "free_space"
	fieldUnusedReserveBytes := "unused_reserve_bytes"
	fieldUsageValid := "usage_valid"
	fieldSpaceInfoValid := "space_info_valid"
	fieldVersionCurrent := "version_current"
	fieldVersionTarget := "version_target"
	fieldVersionRollback := "version_rollback"
	fieldUpdateState := "update_state"
	fieldUpdateStartTime := "update_start_time"
	fieldUpdateEndTime := "update_end_time"
	fieldUpdateArrayNames := "update_array_names"
	fieldUpdateProgressMsg := "update_progress_msg"
	fieldUpdateErrorCode := "update_error_code"
	fieldUpdateDownloading := "update_downloading"
	fieldUpdateDownloadErrorCode := "update_download_error_code"
	fieldUpdateDownloadStartTime := "update_download_start_time"
	fieldUpdateDownloadEndTime := "update_download_end_time"
	fieldIscsiAutomaticConnectionMethod := "iscsi_automatic_connection_method"
	fieldIscsiConnectionRebalancing := "iscsi_connection_rebalancing"
	fieldReplThrottledBandwidth := "repl_throttled_bandwidth"
	fieldReplThrottledBandwidthKbps := "repl_throttled_bandwidth_kbps"
	fieldReplThrottleList := "repl_throttle_list"
	fieldVolumeMigrationStatus := "volume_migration_status"
	fieldArrayUnassignMigrationStatus := "array_unassign_migration_status"
	fieldDataRebalanceStatus := "data_rebalance_status"
	fieldScsiVendorId := "scsi_vendor_id"
	fieldEncryptionConfig := "encryption_config"
	fieldLastLogin := "last_login"
	fieldNumSnaps := "num_snaps"
	fieldNumSnapcolls := "num_snapcolls"
	fieldDate := "date"
	fieldLoginBannerMessage := "login_banner_message"
	fieldLoginBannerAfterAuth := "login_banner_after_auth"
	fieldLoginBannerReset := "login_banner_reset"
	fieldSnapRetnMeterHigh := "snap_retn_meter_high"
	fieldSnapRetnMeterVeryHigh := "snap_retn_meter_very_high"

	GroupFields = &GroupFieldHandles{
		ID:                                 &fieldID,
		Name:                               &fieldName,
		SmtpServer:                         &fieldSmtpServer,
		SmtpPort:                           &fieldSmtpPort,
		SmtpAuthEnabled:                    &fieldSmtpAuthEnabled,
		SmtpAuthUsername:                   &fieldSmtpAuthUsername,
		SmtpAuthPassword:                   &fieldSmtpAuthPassword,
		SmtpEncryptType:                    &fieldSmtpEncryptType,
		AutosupportEnabled:                 &fieldAutosupportEnabled,
		AllowAnalyticsGui:                  &fieldAllowAnalyticsGui,
		AllowSupportTunnel:                 &fieldAllowSupportTunnel,
		ProxyServer:                        &fieldProxyServer,
		ProxyPort:                          &fieldProxyPort,
		ProxyUsername:                      &fieldProxyUsername,
		ProxyPassword:                      &fieldProxyPassword,
		AlertToEmailAddrs:                  &fieldAlertToEmailAddrs,
		SendAlertToSupport:                 &fieldSendAlertToSupport,
		AlertFromEmailAddr:                 &fieldAlertFromEmailAddr,
		AlertMinLevel:                      &fieldAlertMinLevel,
		IsnsEnabled:                        &fieldIsnsEnabled,
		IsnsServer:                         &fieldIsnsServer,
		IsnsPort:                           &fieldIsnsPort,
		SnmpTrapEnabled:                    &fieldSnmpTrapEnabled,
		SnmpTrapHost:                       &fieldSnmpTrapHost,
		SnmpTrapPort:                       &fieldSnmpTrapPort,
		SnmpGetEnabled:                     &fieldSnmpGetEnabled,
		SnmpCommunity:                      &fieldSnmpCommunity,
		SnmpGetPort:                        &fieldSnmpGetPort,
		SnmpSysContact:                     &fieldSnmpSysContact,
		SnmpSysLocation:                    &fieldSnmpSysLocation,
		DomainName:                         &fieldDomainName,
		DnsServers:                         &fieldDnsServers,
		NtpServer:                          &fieldNtpServer,
		Timezone:                           &fieldTimezone,
		UserInactivityTimeout:              &fieldUserInactivityTimeout,
		SyslogdEnabled:                     &fieldSyslogdEnabled,
		SyslogdServer:                      &fieldSyslogdServer,
		SyslogdPort:                        &fieldSyslogdPort,
		SyslogdServers:                     &fieldSyslogdServers,
		VvolEnabled:                        &fieldVvolEnabled,
		SetupProgress:                      &fieldSetupProgress,
		IscsiEnabled:                       &fieldIscsiEnabled,
		FcEnabled:                          &fieldFcEnabled,
		UniqueNameEnabled:                  &fieldUniqueNameEnabled,
		AccessProtocolList:                 &fieldAccessProtocolList,
		GroupTargetEnabled:                 &fieldGroupTargetEnabled,
		DefaultIscsiTargetScope:            &fieldDefaultIscsiTargetScope,
		TdzEnabled:                         &fieldTdzEnabled,
		TdzPrefix:                          &fieldTdzPrefix,
		GroupTargetName:                    &fieldGroupTargetName,
		DefaultVolumeReserve:               &fieldDefaultVolumeReserve,
		DefaultVolumeWarnLevel:             &fieldDefaultVolumeWarnLevel,
		DefaultVolumeLimit:                 &fieldDefaultVolumeLimit,
		DefaultSnapReserve:                 &fieldDefaultSnapReserve,
		DefaultSnapWarnLevel:               &fieldDefaultSnapWarnLevel,
		DefaultSnapLimit:                   &fieldDefaultSnapLimit,
		DefaultSnapLimitPercent:            &fieldDefaultSnapLimitPercent,
		AlarmsEnabled:                      &fieldAlarmsEnabled,
		VssValidationTimeout:               &fieldVssValidationTimeout,
		AutoSwitchoverEnabled:              &fieldAutoSwitchoverEnabled,
		SoftwareSubscriptionEnabled:        &fieldSoftwareSubscriptionEnabled,
		AutoSwitchoverMessages:             &fieldAutoSwitchoverMessages,
		MergeState:                         &fieldMergeState,
		MergeGroupName:                     &fieldMergeGroupName,
		Tlsv1Enabled:                       &fieldTlsv1Enabled,
		CcModeEnabled:                      &fieldCcModeEnabled,
		CertIpsIncluded:                    &fieldCertIpsIncluded,
		GroupSnapshotTtl:                   &fieldGroupSnapshotTtl,
		AutocleanUnmanagedSnapshotsTtlUnit: &fieldAutocleanUnmanagedSnapshotsTtlUnit,
		AutocleanUnmanagedSnapshotsEnabled: &fieldAutocleanUnmanagedSnapshotsEnabled,
		LeaderArrayName:                    &fieldLeaderArrayName,
		LeaderArraySerial:                  &fieldLeaderArraySerial,
		ManagementServiceBackupArrayName:   &fieldManagementServiceBackupArrayName,
		ManagementServiceBackupStatus:      &fieldManagementServiceBackupStatus,
		FailoverMode:                       &fieldFailoverMode,
		WitnessStatus:                      &fieldWitnessStatus,
		MemberList:                         &fieldMemberList,
		CompressedVolUsageBytes:            &fieldCompressedVolUsageBytes,
		CompressedSnapUsageBytes:           &fieldCompressedSnapUsageBytes,
		UncompressedVolUsageBytes:          &fieldUncompressedVolUsageBytes,
		UncompressedSnapUsageBytes:         &fieldUncompressedSnapUsageBytes,
		UsableCapacityBytes:                &fieldUsableCapacityBytes,
		Usage:                              &fieldUsage,
		RawCapacity:                        &fieldRawCapacity,
		UsableCacheCapacity:                &fieldUsableCacheCapacity,
		RawCacheCapacity:                   &fieldRawCacheCapacity,
		SnapUsagePopulated:                 &fieldSnapUsagePopulated,
		PendingDeletes:                     &fieldPendingDeletes,
		NumConnections:                     &fieldNumConnections,
		VolCompressionRatio:                &fieldVolCompressionRatio,
		SnapCompressionRatio:               &fieldSnapCompressionRatio,
		CompressionRatio:                   &fieldCompressionRatio,
		DedupeRatio:                        &fieldDedupeRatio,
		CloneRatio:                         &fieldCloneRatio,
		VolThinProvisioningRatio:           &fieldVolThinProvisioningRatio,
		SavingsRatio:                       &fieldSavingsRatio,
		DataReductionRatio:                 &fieldDataReductionRatio,
		SavingsDedupe:                      &fieldSavingsDedupe,
		SavingsCompression:                 &fieldSavingsCompression,
		SavingsClone:                       &fieldSavingsClone,
		SavingsVolThinProvisioning:         &fieldSavingsVolThinProvisioning,
		SavingsDataReduction:               &fieldSavingsDataReduction,
		Savings:                            &fieldSavings,
		FreeSpace:                          &fieldFreeSpace,
		UnusedReserveBytes:                 &fieldUnusedReserveBytes,
		UsageValid:                         &fieldUsageValid,
		SpaceInfoValid:                     &fieldSpaceInfoValid,
		VersionCurrent:                     &fieldVersionCurrent,
		VersionTarget:                      &fieldVersionTarget,
		VersionRollback:                    &fieldVersionRollback,
		UpdateState:                        &fieldUpdateState,
		UpdateStartTime:                    &fieldUpdateStartTime,
		UpdateEndTime:                      &fieldUpdateEndTime,
		UpdateArrayNames:                   &fieldUpdateArrayNames,
		UpdateProgressMsg:                  &fieldUpdateProgressMsg,
		UpdateErrorCode:                    &fieldUpdateErrorCode,
		UpdateDownloading:                  &fieldUpdateDownloading,
		UpdateDownloadErrorCode:            &fieldUpdateDownloadErrorCode,
		UpdateDownloadStartTime:            &fieldUpdateDownloadStartTime,
		UpdateDownloadEndTime:              &fieldUpdateDownloadEndTime,
		IscsiAutomaticConnectionMethod:     &fieldIscsiAutomaticConnectionMethod,
		IscsiConnectionRebalancing:         &fieldIscsiConnectionRebalancing,
		ReplThrottledBandwidth:             &fieldReplThrottledBandwidth,
		ReplThrottledBandwidthKbps:         &fieldReplThrottledBandwidthKbps,
		ReplThrottleList:                   &fieldReplThrottleList,
		VolumeMigrationStatus:              &fieldVolumeMigrationStatus,
		ArrayUnassignMigrationStatus:       &fieldArrayUnassignMigrationStatus,
		DataRebalanceStatus:                &fieldDataRebalanceStatus,
		ScsiVendorId:                       &fieldScsiVendorId,
		EncryptionConfig:                   &fieldEncryptionConfig,
		LastLogin:                          &fieldLastLogin,
		NumSnaps:                           &fieldNumSnaps,
		NumSnapcolls:                       &fieldNumSnapcolls,
		Date:                               &fieldDate,
		LoginBannerMessage:                 &fieldLoginBannerMessage,
		LoginBannerAfterAuth:               &fieldLoginBannerAfterAuth,
		LoginBannerReset:                   &fieldLoginBannerReset,
		SnapRetnMeterHigh:                  &fieldSnapRetnMeterHigh,
		SnapRetnMeterVeryHigh:              &fieldSnapRetnMeterVeryHigh,
	}
}

// Group - Group is a collection of arrays operating together organized into storage pools.
type Group struct {
	// ID - Identifier of the group.
	ID *string `json:"id,omitempty"`
	// Name - Name of the group.
	Name *string `json:"name,omitempty"`
	// SmtpServer - Hostname or IP Address of SMTP Server.
	SmtpServer *string `json:"smtp_server,omitempty"`
	// SmtpPort - Port number of SMTP Server.
	SmtpPort *int64 `json:"smtp_port,omitempty"`
	// SmtpAuthEnabled - Whether SMTP Server requires authentication.
	SmtpAuthEnabled *bool `json:"smtp_auth_enabled,omitempty"`
	// SmtpAuthUsername - Username to authenticate with SMTP Server.
	SmtpAuthUsername *string `json:"smtp_auth_username,omitempty"`
	// SmtpAuthPassword - Password to authenticate with SMTP Server.
	SmtpAuthPassword *string `json:"smtp_auth_password,omitempty"`
	// SmtpEncryptType - Level of encryption for SMTP. Requires use of SMTP Authentication if encryption is enabled.
	SmtpEncryptType *NsSmtpEncryptType `json:"smtp_encrypt_type,omitempty"`
	// AutosupportEnabled - Whether to send autosupport.
	AutosupportEnabled *bool `json:"autosupport_enabled,omitempty"`
	// AllowAnalyticsGui - Specify whether to use Google Analytics in the GUI. HPE Storage uses Google Analytics to gather data related to GUI usage. The data gathered is used to evaluate and improve the product.
	AllowAnalyticsGui *bool `json:"allow_analytics_gui,omitempty"`
	// AllowSupportTunnel - Whether to allow support tunnel.
	AllowSupportTunnel *bool `json:"allow_support_tunnel,omitempty"`
	// ProxyServer - Hostname or IP Address of HTTP Proxy Server. Setting this attribute to an empty string will unset all proxy settings.
	ProxyServer *string `json:"proxy_server,omitempty"`
	// ProxyPort - Proxy Port of HTTP Proxy Server.
	ProxyPort *int64 `json:"proxy_port,omitempty"`
	// ProxyUsername - Username to authenticate with HTTP Proxy Server.
	ProxyUsername *string `json:"proxy_username,omitempty"`
	// ProxyPassword - Password to authenticate with HTTP Proxy Server.
	ProxyPassword *string `json:"proxy_password,omitempty"`
	// AlertToEmailAddrs - Comma-separated list of email addresss to receive emails.
	AlertToEmailAddrs *string `json:"alert_to_email_addrs,omitempty"`
	// SendAlertToSupport - Whether to send alert to Support.
	SendAlertToSupport *bool `json:"send_alert_to_support,omitempty"`
	// AlertFromEmailAddr - From email address to use while sending emails.
	AlertFromEmailAddr *string `json:"alert_from_email_addr,omitempty"`
	// AlertMinLevel - Minimum level of alert to be notified.
	AlertMinLevel *NsSeverityLevel `json:"alert_min_level,omitempty"`
	// IsnsEnabled - Whether iSNS is enabled.
	IsnsEnabled *bool `json:"isns_enabled,omitempty"`
	// IsnsServer - Hostname or IP Address of iSNS Server.
	IsnsServer *string `json:"isns_server,omitempty"`
	// IsnsPort - Port number for iSNS Server.
	IsnsPort *int64 `json:"isns_port,omitempty"`
	// SnmpTrapEnabled - Whether to enable SNMP traps.
	SnmpTrapEnabled *bool `json:"snmp_trap_enabled,omitempty"`
	// SnmpTrapHost - Hostname or IP Address to send SNMP traps.
	SnmpTrapHost *string `json:"snmp_trap_host,omitempty"`
	// SnmpTrapPort - Port number of SNMP trap host.
	SnmpTrapPort *int64 `json:"snmp_trap_port,omitempty"`
	// SnmpGetEnabled - Whether to accept SNMP get commands.
	SnmpGetEnabled *bool `json:"snmp_get_enabled,omitempty"`
	// SnmpCommunity - Community string to be used with SNMP.
	SnmpCommunity *string `json:"snmp_community,omitempty"`
	// SnmpGetPort - Port number to which SNMP get requests should be sent.
	SnmpGetPort *int64 `json:"snmp_get_port,omitempty"`
	// SnmpSysContact - Name of the SNMP administrator.
	SnmpSysContact *string `json:"snmp_sys_contact,omitempty"`
	// SnmpSysLocation - Location of the group.
	SnmpSysLocation *string `json:"snmp_sys_location,omitempty"`
	// DomainName - Domain name for this group.
	DomainName *string `json:"domain_name,omitempty"`
	// DnsServers - IP addresses for this group's dns servers.
	DnsServers []*NsIPAddressObject `json:"dns_servers,omitempty"`
	// NtpServer - Either IP address or hostname of the NTP server for this group.
	NtpServer *string `json:"ntp_server,omitempty"`
	// Timezone - Timezone in which this group is located.
	Timezone *string `json:"timezone,omitempty"`
	// UserInactivityTimeout - The amount of time in seconds that the user session is inactive before timing out.
	UserInactivityTimeout *int64 `json:"user_inactivity_timeout,omitempty"`
	// SyslogdEnabled - Is syslogd enabled on this system.
	SyslogdEnabled *bool `json:"syslogd_enabled,omitempty"`
	// SyslogdServer - Hostname of the syslogd server.
	SyslogdServer *string `json:"syslogd_server,omitempty"`
	// SyslogdPort - Port number for syslogd server.
	SyslogdPort *int64 `json:"syslogd_port,omitempty"`
	// SyslogdServers - Hostname and/or port of the syslogd servers.
	SyslogdServers []*NsFqdnOrIpAndPortObject `json:"syslogd_servers,omitempty"`
	// VvolEnabled - Are vvols enabled on this group.
	VvolEnabled *bool `json:"vvol_enabled,omitempty"`
	// SetupProgress - Progress indicator of complete setup operation.
	SetupProgress *string `json:"setup_progress,omitempty"`
	// IscsiEnabled - Whether iSCSI is enabled on this group.
	IscsiEnabled *bool `json:"iscsi_enabled,omitempty"`
	// FcEnabled - Whether FC is enabled on this group.
	FcEnabled *bool `json:"fc_enabled,omitempty"`
	// UniqueNameEnabled - Are new volume and volume collection names transformed on this group.
	UniqueNameEnabled *bool `json:"unique_name_enabled,omitempty"`
	// AccessProtocolList - Protocol used to access this group.
	AccessProtocolList []*NsAccessProtocol `json:"access_protocol_list,omitempty"`
	// GroupTargetEnabled - Is group_target enabled on this group.
	GroupTargetEnabled *bool `json:"group_target_enabled,omitempty"`
	// DefaultIscsiTargetScope - Newly created volumes are exported under iSCSI Group Target or iSCSI Volume Target.
	DefaultIscsiTargetScope *NsTargetScope `json:"default_iscsi_target_scope,omitempty"`
	// TdzEnabled - Is Target Driven Zoning (TDZ) enabled on this group.
	TdzEnabled *bool `json:"tdz_enabled,omitempty"`
	// TdzPrefix - Target Driven Zoning (TDZ) prefix for peer zones created by TDZ.
	TdzPrefix *string `json:"tdz_prefix,omitempty"`
	// GroupTargetName - Iscsi target name for this group.
	GroupTargetName *string `json:"group_target_name,omitempty"`
	// DefaultVolumeReserve - Amount of space to reserve for a volume as a percentage of volume size.
	DefaultVolumeReserve *int64 `json:"default_volume_reserve,omitempty"`
	// DefaultVolumeWarnLevel - Default threshold for volume space usage as a percentage of volume size above which an alert is raised.
	DefaultVolumeWarnLevel *int64 `json:"default_volume_warn_level,omitempty"`
	// DefaultVolumeLimit - Default limit for a volume space usage as a percentage of volume size. Volume will be taken offline/made non-writable on exceeding its limit.
	DefaultVolumeLimit *int64 `json:"default_volume_limit,omitempty"`
	// DefaultSnapReserve - Amount of space to reserve for snapshots of a volume as a percentage of volume size.
	DefaultSnapReserve *int64 `json:"default_snap_reserve,omitempty"`
	// DefaultSnapWarnLevel - Default threshold for snapshot space usage of a volume as a percentage of volume size above which an alert is raised.
	DefaultSnapWarnLevel *int64 `json:"default_snap_warn_level,omitempty"`
	// DefaultSnapLimit - This attribute is deprecated. The array does not limit a volume's snapshot space usage. The attribute is ignored on input and returns max int64 value on output.
	DefaultSnapLimit *int64 `json:"default_snap_limit,omitempty"`
	// DefaultSnapLimitPercent - This attribute is deprecated. The array does not limit a volume's snapshot space usage. The attribute is ignored on input and returns -1 on output.
	DefaultSnapLimitPercent *int64 `json:"default_snap_limit_percent,omitempty"`
	// AlarmsEnabled - Whether alarm feature is enabled.
	AlarmsEnabled *bool `json:"alarms_enabled,omitempty"`
	// VssValidationTimeout - The amount of time in seconds to validate Microsoft VSS application synchronization before timing out.
	VssValidationTimeout *int64 `json:"vss_validation_timeout,omitempty"`
	// AutoSwitchoverEnabled - Whether automatic switchover of Group management services feature is enabled.
	AutoSwitchoverEnabled *bool `json:"auto_switchover_enabled,omitempty"`
	// SoftwareSubscriptionEnabled - Whether software subscription of Group management services feature is enabled.
	SoftwareSubscriptionEnabled *bool `json:"software_subscription_enabled,omitempty"`
	// AutoSwitchoverMessages - List of validation messages for automatic switchover of Group Management. This will be empty when there are no conflicts found.
	AutoSwitchoverMessages []*NsErrorWithArguments `json:"auto_switchover_messages,omitempty"`
	// MergeState - State of group merge.
	MergeState *NsMergeState `json:"merge_state,omitempty"`
	// MergeGroupName - Group that we're being merged with.
	MergeGroupName *string `json:"merge_group_name,omitempty"`
	// Tlsv1Enabled - Enable or disable TLSv1.0 and TLSv1.1.
	Tlsv1Enabled *bool `json:"tlsv1_enabled,omitempty"`
	// CcModeEnabled - Enable or disable Common Criteria mode.
	CcModeEnabled *bool `json:"cc_mode_enabled,omitempty"`
	// CertIpsIncluded - Whether to include IPs in group certificate.
	CertIpsIncluded *bool `json:"cert_ips_included,omitempty"`
	// GroupSnapshotTtl - Snapshot Time-to-live(TTL) configured at group level for automatic deletion of unmanaged snapshots. Value 0 indicates unlimited TTL.
	GroupSnapshotTtl *int64 `json:"group_snapshot_ttl,omitempty"`
	// AutocleanUnmanagedSnapshotsTtlUnit - Unit for unmanaged snapshot time to live.
	AutocleanUnmanagedSnapshotsTtlUnit *int64 `json:"autoclean_unmanaged_snapshots_ttl_unit,omitempty"`
	// AutocleanUnmanagedSnapshotsEnabled - Whether autoclean unmanaged snapshots feature is enabled.
	AutocleanUnmanagedSnapshotsEnabled *bool `json:"autoclean_unmanaged_snapshots_enabled,omitempty"`
	// LeaderArrayName - Name of the array where the group Management Service is running.
	LeaderArrayName *string `json:"leader_array_name,omitempty"`
	// LeaderArraySerial - Serial number of the array where the group Management Service is running.
	LeaderArraySerial *string `json:"leader_array_serial,omitempty"`
	// ManagementServiceBackupArrayName - Name of the array where backup the group Management Service is running.
	ManagementServiceBackupArrayName *string `json:"management_service_backup_array_name,omitempty"`
	// ManagementServiceBackupStatus - HA status of the group Management Service.
	ManagementServiceBackupStatus *NsGroupHAStatus `json:"management_service_backup_status,omitempty"`
	// FailoverMode - Failover mode of the group Management Service.
	FailoverMode *NsGroupFailoverMode `json:"failover_mode,omitempty"`
	// WitnessStatus - Witness status from group Management Service array and group Management Service backup array.
	WitnessStatus []*NsWitnessTestResponse `json:"witness_status,omitempty"`
	// MemberList - Members of this group.
	MemberList []*string `json:"member_list,omitempty"`
	// CompressedVolUsageBytes - Compressed usage of volumes in the group.
	CompressedVolUsageBytes *int64 `json:"compressed_vol_usage_bytes,omitempty"`
	// CompressedSnapUsageBytes - Compressed usage of snapshots in the group.
	CompressedSnapUsageBytes *int64 `json:"compressed_snap_usage_bytes,omitempty"`
	// UncompressedVolUsageBytes - Uncompressed usage of volumes in the group.
	UncompressedVolUsageBytes *int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
	// UncompressedSnapUsageBytes - Uncompressed usage of snapshots in the group.
	UncompressedSnapUsageBytes *int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
	// UsableCapacityBytes - Usable capacity bytes of the group.
	UsableCapacityBytes *int64 `json:"usable_capacity_bytes,omitempty"`
	// Usage - Used space of the group in bytes.
	Usage *int64 `json:"usage,omitempty"`
	// RawCapacity - Total capacity of the group.
	RawCapacity *int64 `json:"raw_capacity,omitempty"`
	// UsableCacheCapacity - Usable cache capacity of the group.
	UsableCacheCapacity *int64 `json:"usable_cache_capacity,omitempty"`
	// RawCacheCapacity - Total cache capacity of the group.
	RawCacheCapacity *int64 `json:"raw_cache_capacity,omitempty"`
	// SnapUsagePopulated - Total snapshot usage as if each snapshot is deep copy of the volume.
	SnapUsagePopulated *int64 `json:"snap_usage_populated,omitempty"`
	// PendingDeletes - Usage for blocks that are not yet deleted.
	PendingDeletes *int64 `json:"pending_deletes,omitempty"`
	// NumConnections - Number of connections to the group.
	NumConnections *int64 `json:"num_connections,omitempty"`
	// VolCompressionRatio - Compression ratio of volumes in the group.
	VolCompressionRatio *float64 `json:"vol_compression_ratio,omitempty"`
	// SnapCompressionRatio - Compression ratio of snapshots in the group.
	SnapCompressionRatio *float64 `json:"snap_compression_ratio,omitempty"`
	// CompressionRatio - Compression savings for the group expressed as ratio.
	CompressionRatio *float64 `json:"compression_ratio,omitempty"`
	// DedupeRatio - Dedupe savings for the group expressed as ratio.
	DedupeRatio *float64 `json:"dedupe_ratio,omitempty"`
	// CloneRatio - Clone savings for the group expressed as ratio.
	CloneRatio *float64 `json:"clone_ratio,omitempty"`
	// VolThinProvisioningRatio - Thin provisioning savings for volumes in the group expressed as ratio.
	VolThinProvisioningRatio *float64 `json:"vol_thin_provisioning_ratio,omitempty"`
	// SavingsRatio - Overall savings in the group expressed as ratio.
	SavingsRatio *float64 `json:"savings_ratio,omitempty"`
	// DataReductionRatio - Space savings in the group that does not include thin-provisioning savings expressed as ratio.
	DataReductionRatio *float64 `json:"data_reduction_ratio,omitempty"`
	// SavingsDedupe - Space usage savings in the group due to deduplication.
	SavingsDedupe *int64 `json:"savings_dedupe,omitempty"`
	// SavingsCompression - Space usage savings in the group due to compression.
	SavingsCompression *int64 `json:"savings_compression,omitempty"`
	// SavingsClone - Space usage savings in the group due to cloning of volumes.
	SavingsClone *int64 `json:"savings_clone,omitempty"`
	// SavingsVolThinProvisioning - Space usage savings in the group due to thin provisioning of volumes.
	SavingsVolThinProvisioning *int64 `json:"savings_vol_thin_provisioning,omitempty"`
	// SavingsDataReduction - Space usage savings in the group that does not include thin-provisioning savings.
	SavingsDataReduction *int64 `json:"savings_data_reduction,omitempty"`
	// Savings - Overall space usage savings in the group.
	Savings *int64 `json:"savings,omitempty"`
	// FreeSpace - Free space of the pool in bytes.
	FreeSpace *int64 `json:"free_space,omitempty"`
	// UnusedReserveBytes - Reserved space that is not utilized.
	UnusedReserveBytes *int64 `json:"unused_reserve_bytes,omitempty"`
	// UsageValid - Indicates whether the usage of group is valid.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// SpaceInfoValid - Is space info for this group valid.
	SpaceInfoValid *bool `json:"space_info_valid,omitempty"`
	// VersionCurrent - Version of software running on the group.
	VersionCurrent *string `json:"version_current,omitempty"`
	// VersionTarget - Desired software version for the group.
	VersionTarget *string `json:"version_target,omitempty"`
	// VersionRollback - Rollback software version for the group.
	VersionRollback *string `json:"version_rollback,omitempty"`
	// UpdateState - Group update state.
	UpdateState *NsUpdateState `json:"update_state,omitempty"`
	// UpdateStartTime - Start time of last update.
	UpdateStartTime *int64 `json:"update_start_time,omitempty"`
	// UpdateEndTime - End time of last update.
	UpdateEndTime *int64 `json:"update_end_time,omitempty"`
	// UpdateArrayNames - Arrays in the group undergoing update.
	UpdateArrayNames *string `json:"update_array_names,omitempty"`
	// UpdateProgressMsg - Group update detailed progress message.
	UpdateProgressMsg *string `json:"update_progress_msg,omitempty"`
	// UpdateErrorCode - If the software update has failed, this indicates the error code corresponding to the failure.
	UpdateErrorCode *string `json:"update_error_code,omitempty"`
	// UpdateDownloading - Is software update package currently downloading.
	UpdateDownloading *bool `json:"update_downloading,omitempty"`
	// UpdateDownloadErrorCode - If the software download has failed, this indicates the error code corresponding to the failure.
	UpdateDownloadErrorCode *string `json:"update_download_error_code,omitempty"`
	// UpdateDownloadStartTime - Start time of last update.
	UpdateDownloadStartTime *int64 `json:"update_download_start_time,omitempty"`
	// UpdateDownloadEndTime - End time of last update.
	UpdateDownloadEndTime *int64 `json:"update_download_end_time,omitempty"`
	// IscsiAutomaticConnectionMethod - Is iscsi reconnection automatic.
	IscsiAutomaticConnectionMethod *bool `json:"iscsi_automatic_connection_method,omitempty"`
	// IscsiConnectionRebalancing - Does iscsi automatically rebalance connections.
	IscsiConnectionRebalancing *bool `json:"iscsi_connection_rebalancing,omitempty"`
	// ReplThrottledBandwidth - Current bandwidth throttle for replication, expressed either as megabits per second or as -1 to indicate that there is no throttle.
	ReplThrottledBandwidth *int64 `json:"repl_throttled_bandwidth,omitempty"`
	// ReplThrottledBandwidthKbps - Current bandwidth throttle for replication, expressed either as kilobits per second or as -1 to indicate that there is no throttle.
	ReplThrottledBandwidthKbps *int64 `json:"repl_throttled_bandwidth_kbps,omitempty"`
	// ReplThrottleList - All the replication bandwidth limits on the system.
	ReplThrottleList []*NsThrottle `json:"repl_throttle_list,omitempty"`
	// VolumeMigrationStatus - Status of data migration activity related to volumes being relocated to different pools.
	VolumeMigrationStatus []*NsVolFamMigStatus `json:"volume_migration_status,omitempty"`
	// ArrayUnassignMigrationStatus - Data migration status for arrays being removed from their pool.
	ArrayUnassignMigrationStatus []*NsArrayUnassignMigStatus `json:"array_unassign_migration_status,omitempty"`
	// DataRebalanceStatus - Status of data rebalancing operations for pools in the group.
	DataRebalanceStatus []*NsPoolRebalanceMigStatus `json:"data_rebalance_status,omitempty"`
	// ScsiVendorId - SCSI vendor ID.
	ScsiVendorId *string `json:"scsi_vendor_id,omitempty"`
	// EncryptionConfig - How encryption is configured for this group.
	EncryptionConfig *NsEncryptionSettings `json:"encryption_config,omitempty"`
	// LastLogin - Time and user of last login to this group.
	LastLogin *string `json:"last_login,omitempty"`
	// NumSnaps - Number of snapshots in the group.
	NumSnaps *int64 `json:"num_snaps,omitempty"`
	// NumSnapcolls - Number of snapshot collections in this group.
	NumSnapcolls *int64 `json:"num_snapcolls,omitempty"`
	// Date - Unix epoch time local to the group.
	Date *int64 `json:"date,omitempty"`
	// LoginBannerMessage - The message for the login banner that is displayed during user login activity.
	LoginBannerMessage *string `json:"login_banner_message,omitempty"`
	// LoginBannerAfterAuth - Should the banner be displayed before the user credentials are prompted or after prompting the user credentials.
	LoginBannerAfterAuth *bool `json:"login_banner_after_auth,omitempty"`
	// LoginBannerReset - This will reset the banner to the version of the installed NOS. When login_banner_after_auth is specified, login_banner_reset can not be set to true.
	LoginBannerReset *bool `json:"login_banner_reset,omitempty"`
	// SnapRetnMeterHigh - Threshold for considering a volume as high retention.
	SnapRetnMeterHigh *int64 `json:"snap_retn_meter_high,omitempty"`
	// SnapRetnMeterVeryHigh - Threshold for considering a volume as very high retention.
	SnapRetnMeterVeryHigh *int64 `json:"snap_retn_meter_very_high,omitempty"`
}

// GroupFieldHandles provides a string representation for each AccessControlRecord field.
type GroupFieldHandles struct {
	ID                                 *string
	Name                               *string
	SmtpServer                         *string
	SmtpPort                           *string
	SmtpAuthEnabled                    *string
	SmtpAuthUsername                   *string
	SmtpAuthPassword                   *string
	SmtpEncryptType                    *string
	AutosupportEnabled                 *string
	AllowAnalyticsGui                  *string
	AllowSupportTunnel                 *string
	ProxyServer                        *string
	ProxyPort                          *string
	ProxyUsername                      *string
	ProxyPassword                      *string
	AlertToEmailAddrs                  *string
	SendAlertToSupport                 *string
	AlertFromEmailAddr                 *string
	AlertMinLevel                      *string
	IsnsEnabled                        *string
	IsnsServer                         *string
	IsnsPort                           *string
	SnmpTrapEnabled                    *string
	SnmpTrapHost                       *string
	SnmpTrapPort                       *string
	SnmpGetEnabled                     *string
	SnmpCommunity                      *string
	SnmpGetPort                        *string
	SnmpSysContact                     *string
	SnmpSysLocation                    *string
	DomainName                         *string
	DnsServers                         *string
	NtpServer                          *string
	Timezone                           *string
	UserInactivityTimeout              *string
	SyslogdEnabled                     *string
	SyslogdServer                      *string
	SyslogdPort                        *string
	SyslogdServers                     *string
	VvolEnabled                        *string
	SetupProgress                      *string
	IscsiEnabled                       *string
	FcEnabled                          *string
	UniqueNameEnabled                  *string
	AccessProtocolList                 *string
	GroupTargetEnabled                 *string
	DefaultIscsiTargetScope            *string
	TdzEnabled                         *string
	TdzPrefix                          *string
	GroupTargetName                    *string
	DefaultVolumeReserve               *string
	DefaultVolumeWarnLevel             *string
	DefaultVolumeLimit                 *string
	DefaultSnapReserve                 *string
	DefaultSnapWarnLevel               *string
	DefaultSnapLimit                   *string
	DefaultSnapLimitPercent            *string
	AlarmsEnabled                      *string
	VssValidationTimeout               *string
	AutoSwitchoverEnabled              *string
	SoftwareSubscriptionEnabled        *string
	AutoSwitchoverMessages             *string
	MergeState                         *string
	MergeGroupName                     *string
	Tlsv1Enabled                       *string
	CcModeEnabled                      *string
	CertIpsIncluded                    *string
	GroupSnapshotTtl                   *string
	AutocleanUnmanagedSnapshotsTtlUnit *string
	AutocleanUnmanagedSnapshotsEnabled *string
	LeaderArrayName                    *string
	LeaderArraySerial                  *string
	ManagementServiceBackupArrayName   *string
	ManagementServiceBackupStatus      *string
	FailoverMode                       *string
	WitnessStatus                      *string
	MemberList                         *string
	CompressedVolUsageBytes            *string
	CompressedSnapUsageBytes           *string
	UncompressedVolUsageBytes          *string
	UncompressedSnapUsageBytes         *string
	UsableCapacityBytes                *string
	Usage                              *string
	RawCapacity                        *string
	UsableCacheCapacity                *string
	RawCacheCapacity                   *string
	SnapUsagePopulated                 *string
	PendingDeletes                     *string
	NumConnections                     *string
	VolCompressionRatio                *string
	SnapCompressionRatio               *string
	CompressionRatio                   *string
	DedupeRatio                        *string
	CloneRatio                         *string
	VolThinProvisioningRatio           *string
	SavingsRatio                       *string
	DataReductionRatio                 *string
	SavingsDedupe                      *string
	SavingsCompression                 *string
	SavingsClone                       *string
	SavingsVolThinProvisioning         *string
	SavingsDataReduction               *string
	Savings                            *string
	FreeSpace                          *string
	UnusedReserveBytes                 *string
	UsageValid                         *string
	SpaceInfoValid                     *string
	VersionCurrent                     *string
	VersionTarget                      *string
	VersionRollback                    *string
	UpdateState                        *string
	UpdateStartTime                    *string
	UpdateEndTime                      *string
	UpdateArrayNames                   *string
	UpdateProgressMsg                  *string
	UpdateErrorCode                    *string
	UpdateDownloading                  *string
	UpdateDownloadErrorCode            *string
	UpdateDownloadStartTime            *string
	UpdateDownloadEndTime              *string
	IscsiAutomaticConnectionMethod     *string
	IscsiConnectionRebalancing         *string
	ReplThrottledBandwidth             *string
	ReplThrottledBandwidthKbps         *string
	ReplThrottleList                   *string
	VolumeMigrationStatus              *string
	ArrayUnassignMigrationStatus       *string
	DataRebalanceStatus                *string
	ScsiVendorId                       *string
	EncryptionConfig                   *string
	LastLogin                          *string
	NumSnaps                           *string
	NumSnapcolls                       *string
	Date                               *string
	LoginBannerMessage                 *string
	LoginBannerAfterAuth               *string
	LoginBannerReset                   *string
	SnapRetnMeterHigh                  *string
	SnapRetnMeterVeryHigh              *string
}
