/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model
//package nimblestorage/v1/Group


// Group :
type Group struct {
   // ID
   ID string `json:"id,omitempty"`
   // Name
   Name string `json:"name,omitempty"`
   // SmtpServer
   SmtpServer string `json:"smtp_server,omitempty"`
   // SmtpPort
   SmtpPort float64 `json:"smtp_port,omitempty"`
   // SmtpAuthEnabled
   SmtpAuthEnabled bool `json:"smtp_auth_enabled,omitempty"`
   // SmtpAuthUsername
   SmtpAuthUsername string `json:"smtp_auth_username,omitempty"`
   // SmtpAuthPassword
   SmtpAuthPassword string `json:"smtp_auth_password,omitempty"`
   // AutosupportEnabled
   AutosupportEnabled bool `json:"autosupport_enabled,omitempty"`
   // AllowSupportTunnel
   AllowSupportTunnel bool `json:"allow_support_tunnel,omitempty"`
   // ProxyServer
   ProxyServer string `json:"proxy_server,omitempty"`
   // ProxyPort
   ProxyPort float64 `json:"proxy_port,omitempty"`
   // ProxyUsername
   ProxyUsername string `json:"proxy_username,omitempty"`
   // ProxyPassword
   ProxyPassword string `json:"proxy_password,omitempty"`
   // AlertToEmailAddrs
   AlertToEmailAddrs string `json:"alert_to_email_addrs,omitempty"`
   // SendAlertToSupport
   SendAlertToSupport bool `json:"send_alert_to_support,omitempty"`
   // AlertFromEmailAddr
   AlertFromEmailAddr string `json:"alert_from_email_addr,omitempty"`
   // IsnsEnabled
   IsnsEnabled bool `json:"isns_enabled,omitempty"`
   // IsnsServer
   IsnsServer string `json:"isns_server,omitempty"`
   // IsnsPort
   IsnsPort float64 `json:"isns_port,omitempty"`
   // SnmpTrapEnabled
   SnmpTrapEnabled bool `json:"snmp_trap_enabled,omitempty"`
   // SnmpTrapHost
   SnmpTrapHost string `json:"snmp_trap_host,omitempty"`
   // SnmpTrapPort
   SnmpTrapPort float64 `json:"snmp_trap_port,omitempty"`
   // SnmpGetEnabled
   SnmpGetEnabled bool `json:"snmp_get_enabled,omitempty"`
   // SnmpCommunity
   SnmpCommunity string `json:"snmp_community,omitempty"`
   // SnmpGetPort
   SnmpGetPort float64 `json:"snmp_get_port,omitempty"`
   // SnmpSysContact
   SnmpSysContact string `json:"snmp_sys_contact,omitempty"`
   // SnmpSysLocation
   SnmpSysLocation string `json:"snmp_sys_location,omitempty"`
   // DomainName
   DomainName string `json:"domain_name,omitempty"`
   // NtpServer
   NtpServer string `json:"ntp_server,omitempty"`
   // Timezone
   Timezone string `json:"timezone,omitempty"`
   // UserInactivityTimeout
   UserInactivityTimeout float64 `json:"user_inactivity_timeout,omitempty"`
   // SyslogdEnabled
   SyslogdEnabled bool `json:"syslogd_enabled,omitempty"`
   // SyslogdServer
   SyslogdServer string `json:"syslogd_server,omitempty"`
   // SyslogdPort
   SyslogdPort float64 `json:"syslogd_port,omitempty"`
   // VvolEnabled
   VvolEnabled bool `json:"vvol_enabled,omitempty"`
   // IscsiEnabled
   IscsiEnabled bool `json:"iscsi_enabled,omitempty"`
   // FcEnabled
   FcEnabled bool `json:"fc_enabled,omitempty"`
   // UniqueNameEnabled
   UniqueNameEnabled bool `json:"unique_name_enabled,omitempty"`
   // GroupTargetEnabled
   GroupTargetEnabled bool `json:"group_target_enabled,omitempty"`
   // GroupTargetName
   GroupTargetName string `json:"group_target_name,omitempty"`
   // DefaultVolumeReserve
   DefaultVolumeReserve float64 `json:"default_volume_reserve,omitempty"`
   // DefaultVolumeWarnLevel
   DefaultVolumeWarnLevel float64 `json:"default_volume_warn_level,omitempty"`
   // DefaultVolumeLimit
   DefaultVolumeLimit float64 `json:"default_volume_limit,omitempty"`
   // DefaultSnapReserve
   DefaultSnapReserve float64 `json:"default_snap_reserve,omitempty"`
   // DefaultSnapWarnLevel
   DefaultSnapWarnLevel float64 `json:"default_snap_warn_level,omitempty"`
   // DefaultSnapLimit
   DefaultSnapLimit float64 `json:"default_snap_limit,omitempty"`
   // DefaultSnapLimitPercent
   DefaultSnapLimitPercent  int32 `json:"default_snap_limit_percent,omitempty"`
   // AlarmsEnabled
   AlarmsEnabled bool `json:"alarms_enabled,omitempty"`
   // VssValIDationTimeout
   VssValIDationTimeout float64 `json:"vss_validation_timeout,omitempty"`
   // MergeGroupName
   MergeGroupName string `json:"merge_group_name,omitempty"`
   // Tlsv1Enabled
   Tlsv1Enabled bool `json:"tlsv1_enabled,omitempty"`
   // CcModeEnabled
   CcModeEnabled bool `json:"cc_mode_enabled,omitempty"`
   // GroupSnapshotTtl
   GroupSnapshotTtl  int32 `json:"group_snapshot_ttl,omitempty"`
   // AutocleanUnmanagedSnapshotsTtlUnit
   AutocleanUnmanagedSnapshotsTtlUnit float64 `json:"autoclean_unmanaged_snapshots_ttl_unit,omitempty"`
   // AutocleanUnmanagedSnapshotsEnabled
   AutocleanUnmanagedSnapshotsEnabled bool `json:"autoclean_unmanaged_snapshots_enabled,omitempty"`
   // LeaderArrayName
   LeaderArrayName string `json:"leader_array_name,omitempty"`
   // LeaderArraySerial
   LeaderArraySerial string `json:"leader_array_serial,omitempty"`
   // ManagementServiceBackupArrayName
   ManagementServiceBackupArrayName string `json:"management_service_backup_array_name,omitempty"`
   // CompressedVolUsageBytes
   CompressedVolUsageBytes float64 `json:"compressed_vol_usage_bytes,omitempty"`
   // CompressedSnapUsageBytes
   CompressedSnapUsageBytes float64 `json:"compressed_snap_usage_bytes,omitempty"`
   // UncompressedVolUsageBytes
   UncompressedVolUsageBytes float64 `json:"uncompressed_vol_usage_bytes,omitempty"`
   // UncompressedSnapUsageBytes
   UncompressedSnapUsageBytes float64 `json:"uncompressed_snap_usage_bytes,omitempty"`
   // UsableCapacityBytes
   UsableCapacityBytes float64 `json:"usable_capacity_bytes,omitempty"`
   // Usage
   Usage float64 `json:"usage,omitempty"`
   // RawCapacity
   RawCapacity float64 `json:"raw_capacity,omitempty"`
   // UsableCacheCapacity
   UsableCacheCapacity float64 `json:"usable_cache_capacity,omitempty"`
   // RawCacheCapacity
   RawCacheCapacity float64 `json:"raw_cache_capacity,omitempty"`
   // SnapUsagePopulated
   SnapUsagePopulated float64 `json:"snap_usage_populated,omitempty"`
   // PendingDeletes
   PendingDeletes float64 `json:"pending_deletes,omitempty"`
   // NumConnections
   NumConnections float64 `json:"num_connections,omitempty"`
   // VolCompressionRatio
   VolCompressionRatio float32 `json:"vol_compression_ratio,omitempty"`
   // SnapCompressionRatio
   SnapCompressionRatio float32 `json:"snap_compression_ratio,omitempty"`
   // CompressionRatio
   CompressionRatio float32 `json:"compression_ratio,omitempty"`
   // DedupeRatio
   DedupeRatio float32 `json:"dedupe_ratio,omitempty"`
   // CloneRatio
   CloneRatio float32 `json:"clone_ratio,omitempty"`
   // VolThinProvisioningRatio
   VolThinProvisioningRatio float32 `json:"vol_thin_provisioning_ratio,omitempty"`
   // SavingsRatio
   SavingsRatio float32 `json:"savings_ratio,omitempty"`
   // DataReductionRatio
   DataReductionRatio float32 `json:"data_reduction_ratio,omitempty"`
   // SavingsDedupe
   SavingsDedupe float64 `json:"savings_dedupe,omitempty"`
   // SavingsCompression
   SavingsCompression float64 `json:"savings_compression,omitempty"`
   // SavingsClone
   SavingsClone float64 `json:"savings_clone,omitempty"`
   // SavingsVolThinProvisioning
   SavingsVolThinProvisioning float64 `json:"savings_vol_thin_provisioning,omitempty"`
   // SavingsDataReduction
   SavingsDataReduction float64 `json:"savings_data_reduction,omitempty"`
   // Savings
   Savings float64 `json:"savings,omitempty"`
   // FreeSpace
   FreeSpace float64 `json:"free_space,omitempty"`
   // UnusedReserveBytes
   UnusedReserveBytes float64 `json:"unused_reserve_bytes,omitempty"`
   // UsageValID
   UsageValID bool `json:"usage_valid,omitempty"`
   // SpaceInfoValID
   SpaceInfoValID bool `json:"space_info_valid,omitempty"`
   // VersionCurrent
   VersionCurrent string `json:"version_current,omitempty"`
   // VersionTarget
   VersionTarget string `json:"version_target,omitempty"`
   // VersionRollback
   VersionRollback string `json:"version_rollback,omitempty"`
   // UpdateStartTime
   UpdateStartTime float64 `json:"update_start_time,omitempty"`
   // UpdateEndTime
   UpdateEndTime float64 `json:"update_end_time,omitempty"`
   // UpdateArrayNames
   UpdateArrayNames string `json:"update_array_names,omitempty"`
   // UpdateProgressMsg
   UpdateProgressMsg string `json:"update_progress_msg,omitempty"`
   // UpdateErrorCode
   UpdateErrorCode string `json:"update_error_code,omitempty"`
   // UpdateDownloading
   UpdateDownloading bool `json:"update_downloading,omitempty"`
   // UpdateDownloadErrorCode
   UpdateDownloadErrorCode string `json:"update_download_error_code,omitempty"`
   // UpdateDownloadStartTime
   UpdateDownloadStartTime float64 `json:"update_download_start_time,omitempty"`
   // UpdateDownloadEndTime
   UpdateDownloadEndTime float64 `json:"update_download_end_time,omitempty"`
   // IscsiAutomaticConnectionMethod
   IscsiAutomaticConnectionMethod bool `json:"iscsi_automatic_connection_method,omitempty"`
   // IscsiConnectionRebalancing
   IscsiConnectionRebalancing bool `json:"iscsi_connection_rebalancing,omitempty"`
   // ReplThrottledBandwIDth
   ReplThrottledBandwIDth  int32 `json:"repl_throttled_bandwidth,omitempty"`
   // ReplThrottledBandwIDthKbps
   ReplThrottledBandwIDthKbps  int32 `json:"repl_throttled_bandwidth_kbps,omitempty"`
   // ScsiVendorID
   ScsiVendorID string `json:"scsi_vendor_id,omitempty"`
   // LastLogin
   LastLogin string `json:"last_login,omitempty"`
   // NumSnaps
   NumSnaps float64 `json:"num_snaps,omitempty"`
   // NumSnapcolls
   NumSnapcolls float64 `json:"num_snapcolls,omitempty"`
   // Date
   Date float64 `json:"date,omitempty"`
   // LoginBannerMessage
   LoginBannerMessage string `json:"login_banner_message,omitempty"`
   // LoginBannerAfterAuth
   LoginBannerAfterAuth bool `json:"login_banner_after_auth,omitempty"`
   // LoginBannerReset
   LoginBannerReset bool `json:"login_banner_reset,omitempty"`
   // SnapRetnMeterHigh
   SnapRetnMeterHigh float64 `json:"snap_retn_meter_high,omitempty"`
   // SnapRetnMeterVeryHigh
   SnapRetnMeterVeryHigh float64 `json:"snap_retn_meter_very_high,omitempty"`
}
