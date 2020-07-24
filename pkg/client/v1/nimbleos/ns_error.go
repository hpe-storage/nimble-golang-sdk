// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsError Enum.

type NsError string

const (
	cNsErrorSmReplVolcollDeletion                         NsError = "SM_repl_volcoll_deletion"
	cNsErrorSmEncryptionGroupCipherOverride               NsError = "SM_encryption_group_cipher_override"
	cNsErrorSmEncryptionMustBeEnabled                     NsError = "SM_encryption_must_be_enabled"
	cNsErrorSmExtTrigSchedNotPresent                      NsError = "SM_ext_trig_sched_not_present"
	cNsErrorSmAppserverNotFound                           NsError = "SM_appserver_not_found"
	cNsErrorSmFolderReplPartner                           NsError = "SM_folder_repl_partner"
	cNsErrorSmArrayPoolMember                             NsError = "SM_array_pool_member"
	cNsErrorSmErrShelfCreateRfail                         NsError = "SM_err_shelf_create_rfail"
	cNsErrorSmStarterVolCreate                            NsError = "SM_starter_vol_create"
	cNsErrorSmInvalidNetconfigName                        NsError = "SM_invalid_netconfig_name"
	cNsErrorSmInvalidVolaclScope                          NsError = "SM_invalid_volacl_scope"
	cNsErrorSmErrShelfNoElocId                            NsError = "SM_err_shelf_no_eloc_id"
	cNsErrorSmInvalidNetzoneType                          NsError = "SM_invalid_netzone_type"
	cNsErrorSmErrCannotModifyTdz                          NsError = "SM_err_cannot_modify_tdz"
	cNsErrorSmControllerActive                            NsError = "SM_controller_active"
	cNsErrorSmPoolHasVolume                               NsError = "SM_pool_has_volume"
	cNsErrorSmExtTrigSchedAttrib                          NsError = "SM_ext_trig_sched_attrib"
	cNsErrorSmMissingCriteriaFieldname                    NsError = "SM_missing_criteria_fieldName"
	cNsErrorSmNetconfigAlreadyActive                      NsError = "SM_netconfig_already_active"
	cNsErrorSmStartRowBeyondEndRow                        NsError = "SM_start_row_beyond_end_row"
	cNsErrorSmInsufficientFcInitiatorInput                NsError = "SM_insufficient_fc_initiator_input"
	cNsErrorSmSecondMgmtSubnet                            NsError = "SM_second_mgmt_subnet"
	cNsErrorSmInvalidInitiatorWwpn                        NsError = "SM_invalid_initiator_wwpn"
	cNsErrorSmArrayFlashMismatch                          NsError = "SM_array_flash_mismatch"
	cNsErrorSmErrShelfPmdState                            NsError = "SM_err_shelf_pmd_state"
	cNsErrorSmInvalidFcConfig                             NsError = "SM_invalid_fc_config"
	cNsErrorSmSrepDownstreamIsUpstream                    NsError = "SM_srep_downstream_is_upstream"
	cNsErrorSmKeymgrLeave                                 NsError = "SM_keymgr_leave"
	cNsErrorSmVolmvVolEinprog                             NsError = "SM_volmv_vol_einprog"
	cNsErrorSmSrepNameConflictVol                         NsError = "SM_srep_name_conflict_vol"
	cNsErrorSmMultiArrayWithoutAutomaticConnectionMethod  NsError = "SM_multi_array_without_automatic_connection_method"
	cNsErrorSmFolderUsageLimitOverPoolCapacity            NsError = "SM_folder_usage_limit_over_pool_capacity"
	cNsErrorSmEnomem                                      NsError = "SM_enomem"
	cNsErrorSmErrShelfNotInuse                            NsError = "SM_err_shelf_not_inuse"
	cNsErrorSmErrShelfPreactivationMfrErr                 NsError = "SM_err_shelf_preactivation_mfr_err"
	cNsErrorSmUpdateBusy                                  NsError = "SM_update_busy"
	cNsErrorSmProtpolInvalidAppSync                       NsError = "SM_protpol_invalid_app_sync"
	cNsErrorSmErrShelfRaidDegraded                        NsError = "SM_err_shelf_raid_degraded"
	cNsErrorSmLimitFreqSchedGroup                         NsError = "SM_limit_freq_sched_group"
	cNsErrorSmSwupdateEinprog                             NsError = "SM_swupdate_einprog"
	cNsErrorSmMissingCriteriaParam                        NsError = "SM_missing_criteria_param"
	cNsErrorSmNoPathFound                                 NsError = "SM_no_path_found"
	cNsErrorSmIncompatibleAppCategory                     NsError = "SM_incompatible_app_category"
	cNsErrorSmVolmvAbortEnomove                           NsError = "SM_volmv_abort_enomove"
	cNsErrorSmLimitSnapRetentionVolcoll                   NsError = "SM_limit_snap_retention_volcoll"
	cNsErrorSmKeymgrRemove                                NsError = "SM_keymgr_remove"
	cNsErrorSmArrayNotReachable                           NsError = "SM_array_not_reachable"
	cNsErrorSmErrGroupMergeEventsPending                  NsError = "SM_err_group_merge_events_pending"
	cNsErrorSmSrepNameConflictVols                        NsError = "SM_srep_name_conflict_vols"
	cNsErrorSmPoolPartnerResumeUnsup                      NsError = "SM_pool_partner_resume_unsup"
	cNsErrorSmArrayNotFound                               NsError = "SM_array_not_found"
	cNsErrorSmNoIscsiHardware                             NsError = "SM_no_iscsi_hardware"
	cNsErrorSmEnospc                                      NsError = "SM_enospc"
	cNsErrorSmReservedPerfpolName                         NsError = "SM_reserved_perfpol_name"
	cNsErrorSmInvalidRoute                                NsError = "SM_invalid_route"
	cNsErrorSmVolDedupeMoveInvalid                        NsError = "SM_vol_dedupe_move_invalid"
	cNsErrorSmKeymgrJoin                                  NsError = "SM_keymgr_join"
	cNsErrorSmErrSrvUnreach                               NsError = "SM_err_srv_unreach"
	cNsErrorSmErrShelfDiskSasLinkDegraded                 NsError = "SM_err_shelf_disk_sas_link_degraded"
	cNsErrorSmErrShelfInvalidEeprom                       NsError = "SM_err_shelf_invalid_eeprom"
	cNsErrorSmInvalidKeyValue                             NsError = "SM_invalid_key_value"
	cNsErrorSmNoIscsiLunAssignment                        NsError = "SM_no_iscsi_lun_assignment"
	cNsErrorSmSnapshotOffline                             NsError = "SM_snapshot_offline"
	cNsErrorSmDefaultGatewayNotInMgmtSubnet               NsError = "SM_default_gateway_not_in_mgmt_subnet"
	cNsErrorSmErrPassphraseAuth                           NsError = "SM_err_passphrase_auth"
	cNsErrorSmReplEexist                                  NsError = "SM_repl_eexist"
	cNsErrorSmDisableLastProtocol                         NsError = "SM_disable_last_protocol"
	cNsErrorSmSecondUntaggedAssignment                    NsError = "SM_second_untagged_assignment"
	cNsErrorSmPoolHasPe                                   NsError = "SM_pool_has_pe"
	cNsErrorSmErrShelfInvalidCount                        NsError = "SM_err_shelf_invalid_count"
	cNsErrorSmLimitSnapRetentionGroup                     NsError = "SM_limit_snap_retention_group"
	cNsErrorSmVolDedupeThickProvInvalid                   NsError = "SM_vol_dedupe_thick_prov_invalid"
	cNsErrorSmUnknown                                     NsError = "SM_unknown"
	cNsErrorSmErrShelfDedupeImpact                        NsError = "SM_err_shelf_dedupe_impact"
	cNsErrorSmInvalidInitiatorAccessProtocol              NsError = "SM_invalid_initiator_access_protocol"
	cNsErrorSmInternal                                    NsError = "SM_internal"
	cNsErrorSmAsupEbusy                                   NsError = "SM_asup_ebusy"
	cNsErrorSmInvalidVolMbpsLimit                         NsError = "SM_invalid_vol_mbps_limit"
	cNsErrorSmInfoConfigSyncInprogress                    NsError = "SM_info_config_sync_inprogress"
	cNsErrorSmErrPoolStillMerging                         NsError = "SM_err_pool_still_merging"
	cNsErrorSmFolderPerfpolAgentType                      NsError = "SM_folder_perfpol_agent_type"
	cNsErrorSmEinval                                      NsError = "SM_einval"
	cNsErrorSmNoAssocVols                                 NsError = "SM_no_assoc_vols"
	cNsErrorSmShelfNotInuse                               NsError = "SM_shelf_not_inuse"
	cNsErrorSmErrProtpolSettingsNotAllowed                NsError = "SM_err_protpol_settings_not_allowed"
	cNsErrorSmFcRegenerate                                NsError = "SM_fc_regenerate"
	cNsErrorSmErrMultiArrayGroup                          NsError = "SM_err_multi_array_group"
	cNsErrorSmReplRemoteNoBaseSnap                        NsError = "SM_repl_remote_no_base_snap"
	cNsErrorSmUnsupportedQueryOperator                    NsError = "SM_unsupported_query_operator"
	cNsErrorSmSrepSizeMismatchDownstreamVol               NsError = "SM_srep_size_mismatch_downstream_vol"
	cNsErrorSmNoSuchPartner                               NsError = "SM_no_such_partner"
	cNsErrorSmIscsiAllAccessNotAvailable                  NsError = "SM_iscsi_all_access_not_available"
	cNsErrorSmVvolAlreadyEnabled                          NsError = "SM_vvol_already_enabled"
	cNsErrorSmUsageUnavaiable                             NsError = "SM_usage_unavaiable"
	cNsErrorSmEconnrefused                                NsError = "SM_econnrefused"
	cNsErrorSmReplRenameNotsup                            NsError = "SM_repl_rename_notsup"
	cNsErrorSmSessionCreate                               NsError = "SM_session_create"
	cNsErrorSmConflictingAclVol                           NsError = "SM_conflicting_acl_vol"
	cNsErrorSmNospace                                     NsError = "SM_nospace"
	cNsErrorSmReservedVolcollName                         NsError = "SM_reserved_volcoll_name"
	cNsErrorSmPoolHasFolder                               NsError = "SM_pool_has_folder"
	cNsErrorSmPartnerOffline                              NsError = "SM_partner_offline"
	cNsErrorSmConflictingInitiatorAliasWithArgs           NsError = "SM_conflicting_initiator_alias_with_args"
	cNsErrorSmErrShelfNotReady                            NsError = "SM_err_shelf_not_ready"
	cNsErrorSmReplCmprVersionUnsup                        NsError = "SM_repl_cmpr_version_unsup"
	cNsErrorSmErrShelfPreactivationIoErr                  NsError = "SM_err_shelf_preactivation_io_err"
	cNsErrorSmVolmvEinprog                                NsError = "SM_volmv_einprog"
	cNsErrorSmPerfpolIncompatibleAppCategory              NsError = "SM_perfpol_incompatible_app_category"
	cNsErrorSmInvalidArgValue                             NsError = "SM_invalid_arg_value"
	cNsErrorSmErrShelfDedupeBelowFdr                      NsError = "SM_err_shelf_dedupe_below_fdr"
	cNsErrorSmErrShelfInvalidAfsCount                     NsError = "SM_err_shelf_invalid_afs_count"
	cNsErrorSmSerialNotAvail                              NsError = "SM_serial_not_avail"
	cNsErrorSmErrShelfSesMshipErr                         NsError = "SM_err_shelf_ses_mship_err"
	cNsErrorSmFolderProvisionedLimitBelowCurrentUsage     NsError = "SM_folder_provisioned_limit_below_current_usage"
	cNsErrorSmPeMultiProtocolAccessNotSupported           NsError = "SM_pe_multi_protocol_access_not_supported"
	cNsErrorSmLimitHretSnapRetentionPool                  NsError = "SM_limit_hret_snap_retention_pool"
	cNsErrorSmPeConflictingAcl                            NsError = "SM_pe_conflicting_acl"
	cNsErrorSmFolderOverUsageLimit                        NsError = "SM_folder_over_usage_limit"
	cNsErrorSmErrEncryptionMasterKeyMissing               NsError = "SM_err_encryption_master_key_missing"
	cNsErrorSmProtpolInvalidDomainName                    NsError = "SM_protpol_invalid_domain_name"
	cNsErrorSmTooMany                                     NsError = "SM_too_many"
	cNsErrorSmErrShelfExprRevIncompatible                 NsError = "SM_err_shelf_expr_rev_incompatible"
	cNsErrorSmSrvUnreachable                              NsError = "SM_srv_unreachable"
	cNsErrorSmVolumeConflict                              NsError = "SM_volume_conflict"
	cNsErrorSmInvalidArrayName                            NsError = "SM_invalid_array_name"
	cNsErrorSmCannotReadObject                            NsError = "SM_cannot_read_object"
	cNsErrorSmReservedVolName                             NsError = "SM_reserved_vol_name"
	cNsErrorSmPoolPartnerInUse                            NsError = "SM_pool_partner_in_use"
	cNsErrorSmInvalidInitiatorgrpName                     NsError = "SM_invalid_initiatorgrp_name"
	cNsErrorSmAsupValidateError                           NsError = "SM_asup_validate_error"
	cNsErrorSmVersionName                                 NsError = "SM_version_name"
	cNsErrorSmVvolAlreadyDisabled                         NsError = "SM_vvol_already_disabled"
	cNsErrorSmUnexpectedArg                               NsError = "SM_unexpected_arg"
	cNsErrorSmErrShelfInvalidAfs                          NsError = "SM_err_shelf_invalid_afs"
	cNsErrorSmUnexpectedChild                             NsError = "SM_unexpected_child"
	cNsErrorSmFolderOverProvisionedLimit                  NsError = "SM_folder_over_provisioned_limit"
	cNsErrorSmPoolFlashMismatch                           NsError = "SM_pool_flash_mismatch"
	cNsErrorSmLimitScope                                  NsError = "SM_limit_scope"
	cNsErrorSmSrepSizeMismatchDownstreamVols              NsError = "SM_srep_size_mismatch_downstream_vols"
	cNsErrorSmFolderAppsrvrInconsistent                   NsError = "SM_folder_appsrvr_inconsistent"
	cNsErrorSmErrShelfDedupeReduction                     NsError = "SM_err_shelf_dedupe_reduction"
	cNsErrorSmIscsiSvcNotAvailable                        NsError = "SM_iscsi_svc_not_available"
	cNsErrorSmInvalidNetconfigToDelete                    NsError = "SM_invalid_netconfig_to_delete"
	cNsErrorSmErrUnknown                                  NsError = "SM_err_unknown"
	cNsErrorSmMissingCriteriaOperator                     NsError = "SM_missing_criteria_operator"
	cNsErrorSmInvalidAppUuid                              NsError = "SM_invalid_app_uuid"
	cNsErrorSmArrayMemberOrphanWithArgs                   NsError = "SM_array_member_orphan_with_args"
	cNsErrorSmEmptyVol                                    NsError = "SM_empty_vol"
	cNsErrorSmDuplicateSubnetLabel                        NsError = "SM_duplicate_subnet_label"
	cNsErrorSmZeroVlanIdForTaggedAssignment               NsError = "SM_zero_vlan_id_for_tagged_assignment"
	cNsErrorSmNoActionFound                               NsError = "SM_no_action_found"
	cNsErrorSmSyncReplUnconfigureInProgress               NsError = "SM_sync_repl_unconfigure_in_progress"
	cNsErrorSmVolThickProvMoveInvalid                     NsError = "SM_vol_thick_prov_move_invalid"
	cNsErrorSmMissingAdvancedCriteriaConstructor          NsError = "SM_missing_advanced_criteria_constructor"
	cNsErrorSmPoolUnknown                                 NsError = "SM_pool_unknown"
	cNsErrorSmHaltGlWithLiveMemberArray                   NsError = "SM_halt_gl_with_live_member_array"
	cNsErrorSmReplThrottleOverlap                         NsError = "SM_repl_throttle_overlap"
	cNsErrorSmNetconfigUpdateMismatch                     NsError = "SM_netconfig_update_mismatch"
	cNsErrorSmSrepDownstreamVolsAcl                       NsError = "SM_srep_downstream_vols_acl"
	cNsErrorSmStatsFrequencyInvalid                       NsError = "SM_stats_frequency_invalid"
	cNsErrorSmDdupFolderMerge                             NsError = "SM_ddup_folder_merge"
	cNsErrorSmInvalidSubnet                               NsError = "SM_invalid_subnet"
	cNsErrorSmReplEnabled                                 NsError = "SM_repl_enabled"
	cNsErrorSmPoolPartnerNameConflict                     NsError = "SM_pool_partner_name_conflict"
	cNsErrorSmAsupPingfromMgmtipError                     NsError = "SM_asup_pingfrom_mgmtip_error"
	cNsErrorSmEncryptionGroupScopeOverride                NsError = "SM_encryption_group_scope_override"
	cNsErrorSmErrMaxSessionsReached                       NsError = "SM_err_max_sessions_reached"
	cNsErrorSmErrReplMultiplePartners                     NsError = "SM_err_repl_multiple_partners"
	cNsErrorSmStatsNoSensors                              NsError = "SM_stats_no_sensors"
	cNsErrorSmFolderNeedsLimit                            NsError = "SM_folder_needs_limit"
	cNsErrorSmVolmvVvol                                   NsError = "SM_volmv_vvol"
	cNsErrorSmErrGroupMergeInprogress                     NsError = "SM_err_group_merge_inprogress"
	cNsErrorSmFolderNotFound                              NsError = "SM_folder_not_found"
	cNsErrorSmRouteExists                                 NsError = "SM_route_exists"
	cNsErrorSmInvalidDiscoveryIp                          NsError = "SM_invalid_discovery_ip"
	cNsErrorSmErrVolmvCachePinDedupeNotsupp               NsError = "SM_err_volmv_cache_pin_dedupe_notsupp"
	cNsErrorSmMissingCriteria                             NsError = "SM_missing_criteria"
	cNsErrorSmRootBranchPinned                            NsError = "SM_root_branch_pinned"
	cNsErrorSmInvalidMtu                                  NsError = "SM_invalid_mtu"
	cNsErrorSmNoFcHardware                                NsError = "SM_no_fc_hardware"
	cNsErrorSmPoolNotFound                                NsError = "SM_pool_not_found"
	cNsErrorSmDuplicateIp                                 NsError = "SM_duplicate_ip"
	cNsErrorSmNoAction                                    NsError = "SM_no_action"
	cNsErrorSmProtpolAppSyncOracleParams                  NsError = "SM_protpol_app_sync_oracle_params"
	cNsErrorSmProtpolInvalidVssSettings                   NsError = "SM_protpol_invalid_vss_settings"
	cNsErrorSmEncryptionInvalidScope                      NsError = "SM_encryption_invalid_scope"
	cNsErrorSmLimitSnapcollVolcoll                        NsError = "SM_limit_snapcoll_volcoll"
	cNsErrorSmShelfRaidDegraded                           NsError = "SM_shelf_raid_degraded"
	cNsErrorSmInvalidNonIscsiDataSubnetType               NsError = "SM_invalid_non_iscsi_data_subnet_type"
	cNsErrorSmVolmvIncompatibleAgentType                  NsError = "SM_volmv_incompatible_agent_type"
	cNsErrorSmReplPartnerVersionUnknown                   NsError = "SM_repl_partner_version_unknown"
	cNsErrorSmAsupNameresError                            NsError = "SM_asup_nameres_error"
	cNsErrorSmReplApiUnsup                                NsError = "SM_repl_api_unsup"
	cNsErrorSmEperm                                       NsError = "SM_eperm"
	cNsErrorSmEnoentEnoent                                NsError = "SM_enoent,ENOENT"
	cNsErrorSmArrayGroupLeader                            NsError = "SM_array_group_leader"
	cNsErrorSmInvalidInitiatorIqn                         NsError = "SM_invalid_initiator_iqn"
	cNsErrorSmReplIntragroup                              NsError = "SM_repl_intragroup"
	cNsErrorSmRemoveNonemptyFolder                        NsError = "SM_remove_nonempty_folder"
	cNsErrorSmAddNoniscsiToIscsiGroup                     NsError = "SM_add_noniscsi_to_iscsi_group"
	cNsErrorSmPoolDedupeIncapable                         NsError = "SM_pool_dedupe_incapable"
	cNsErrorSmNoVvolSupport                               NsError = "SM_no_vvol_support"
	cNsErrorSmNoInitiatorgrp                              NsError = "SM_no_initiatorgrp"
	cNsErrorSmInvalidNetmask                              NsError = "SM_invalid_netmask"
	cNsErrorSmNoDiscoveryIpInManualMode                   NsError = "SM_no_discovery_ip_in_manual_mode"
	cNsErrorSmInvalidVolReserveValues                     NsError = "SM_invalid_vol_reserve_values"
	cNsErrorSmSupportIpInvalid                            NsError = "SM_support_ip_invalid"
	cNsErrorSmErrShelfForeignDisk                         NsError = "SM_err_shelf_foreign_disk"
	cNsErrorSmVolsnapAlreadyExported                      NsError = "SM_volsnap_already_exported"
	cNsErrorSmFcInitiatorgrpSubnetNotSupported            NsError = "SM_fc_initiatorgrp_subnet_not_supported"
	cNsErrorSmIncompatibleInitiatorAccessProtocol         NsError = "SM_incompatible_initiator_access_protocol"
	cNsErrorSmInvalidInitiatorgrpAccessProtocol           NsError = "SM_invalid_initiatorgrp_access_protocol"
	cNsErrorSmLimitPeacl                                  NsError = "SM_limit_peacl"
	cNsErrorSmEtimedout                                   NsError = "SM_etimedout"
	cNsErrorSmInitiatorgrpSubnetDoesNotExist              NsError = "SM_initiatorgrp_subnet_does_not_exist"
	cNsErrorSmVolOffline                                  NsError = "SM_vol_offline"
	cNsErrorSmErrSreplMgmtOpInProgress                    NsError = "SM_err_srepl_mgmt_op_in_progress"
	cNsErrorSmReplFanoutMaximumCloudPartnersExceeded      NsError = "SM_repl_fanout_maximum_cloud_partners_exceeded"
	cNsErrorSmMalformedUrl                                NsError = "SM_malformed_url"
	cNsErrorSmErrSessionCreate                            NsError = "SM_err_session_create"
	cNsErrorSmErrShelfCtrlrLoop                           NsError = "SM_err_shelf_ctrlr_loop"
	cNsErrorSmPeConflictingAclLun                         NsError = "SM_pe_conflicting_acl_lun"
	cNsErrorSmInvalidCtrlrName                            NsError = "SM_invalid_ctrlr_name"
	cNsErrorSmBackupNetconfigReadonly                     NsError = "SM_backup_netconfig_readonly"
	cNsErrorSmLimitSnapGroup                              NsError = "SM_limit_snap_group"
	cNsErrorSmEncryptionInvalidCipher                     NsError = "SM_encryption_invalid_cipher"
	cNsErrorSmPerfpolInvalidAppCategory                   NsError = "SM_perfpol_invalid_app_category"
	cNsErrorSmArrayPoolFlashMismatch                      NsError = "SM_array_pool_flash_mismatch"
	cNsErrorSmFolderLimitInval                            NsError = "SM_folder_limit_inval"
	cNsErrorSmShelfSsdDegraded                            NsError = "SM_shelf_ssd_degraded"
	cNsErrorSmFolderOverdraftLimitNeedsUsageLimit         NsError = "SM_folder_overdraft_limit_needs_usage_limit"
	cNsErrorSmEncryptionMasterKeyMissing                  NsError = "SM_encryption_master_key_missing"
	cNsErrorSmSrepAgentTypeMismatchDownstreamVols         NsError = "SM_srep_agent_type_mismatch_downstream_vols"
	cNsErrorSmErrVolNotOfflineOnRestore                   NsError = "SM_err_vol_not_offline_on_restore"
	cNsErrorSmReplHandoverBusy                            NsError = "SM_repl_handover_busy"
	cNsErrorSmNotOwner                                    NsError = "SM_not_owner"
	cNsErrorSmSrepDownstreamAcl                           NsError = "SM_srep_downstream_acl"
	cNsErrorSmNoSubnetWithLabel                           NsError = "SM_no_subnet_with_label"
	cNsErrorSmPoolUpdateInvalArrays                       NsError = "SM_pool_update_inval_arrays"
	cNsErrorSmVolHasOnlineSnap                            NsError = "SM_vol_has_online_snap"
	cNsErrorSmInvalidDataIp                               NsError = "SM_invalid_data_ip"
	cNsErrorSmPoolVolmvEinprog                            NsError = "SM_pool_volmv_einprog"
	cNsErrorSmNoDataIpOnMgmtPlusData                      NsError = "SM_no_data_ip_on_mgmt_plus_data"
	cNsErrorSmConflictingAclLun                           NsError = "SM_conflicting_acl_lun"
	cNsErrorSmVpCreatedIgrp                               NsError = "SM_vp_created_igrp"
	cNsErrorSmStarterVolAclCreate                         NsError = "SM_starter_vol_acl_create"
	cNsErrorSmBadPkg                                      NsError = "SM_bad_pkg"
	cNsErrorSmPasswdSameAsCurrent                         NsError = "SM_passwd_same_as_current"
	cNsErrorSmSnaplunsOutOfSync                           NsError = "SM_snapluns_out_of_sync"
	cNsErrorSmLimitSnapPool                               NsError = "SM_limit_snap_pool"
	cNsErrorSmMultiProtocolAccessNotSupported             NsError = "SM_multi_protocol_access_not_supported"
	cNsErrorSmVolHasClone                                 NsError = "SM_vol_has_clone"
	cNsErrorSmDedupeSingleArray                           NsError = "SM_dedupe_single_array"
	cNsErrorSmEncryptionMasterKeyInactive                 NsError = "SM_encryption_master_key_inactive"
	cNsErrorSmSnapHasClone                                NsError = "SM_snap_has_clone"
	cNsErrorSmErrEncryptionNeeded                         NsError = "SM_err_encryption_needed"
	cNsErrorSmMismatchingDuplicateSubnet                  NsError = "SM_mismatching_duplicate_subnet"
	cNsErrorSmSrepNotGroupScopedVol                       NsError = "SM_srep_not_group_scoped_vol"
	cNsErrorSmOnlyVvolFolderFolset                        NsError = "SM_only_vvol_folder_folset"
	cNsErrorSmReplSnapshotSync                            NsError = "SM_repl_snapshot_sync"
	cNsErrorSmStatsNoSuchObject                           NsError = "SM_stats_no_such_object"
	cNsErrorSmDefaultRouteMissing                         NsError = "SM_default_route_missing"
	cNsErrorSmReplVasa3ApiUnsup                           NsError = "SM_repl_vasa3_api_unsup"
	cNsErrorSmOverlappingSubnets                          NsError = "SM_overlapping_subnets"
	cNsErrorSmMissingArg                                  NsError = "SM_missing_arg"
	cNsErrorSmVolmvCachePinDedupeNotsupp                  NsError = "SM_volmv_cache_pin_dedupe_notsupp"
	cNsErrorSmErrMergeConflict                            NsError = "SM_err_merge_conflict"
	cNsErrorSmTooSmall                                    NsError = "SM_too_small"
	cNsErrorSmVolWarnGreaterThanLimit                     NsError = "SM_vol_warn_greater_than_limit"
	cNsErrorSmKeymgrUnreach                               NsError = "SM_keymgr_unreach"
	cNsErrorSmErrFcAsymmetryNotSupported                  NsError = "SM_err_fc_asymmetry_not_supported"
	cNsErrorSmLunMismatch                                 NsError = "SM_lun_mismatch"
	cNsErrorSmErrShelfInvalidDiskCount                    NsError = "SM_err_shelf_invalid_disk_count"
	cNsErrorSmLimitVolacl                                 NsError = "SM_limit_volacl"
	cNsErrorSmControllerNotActive                         NsError = "SM_controller_not_active"
	cNsErrorSmReplNoPartnerAvail                          NsError = "SM_repl_no_partner_avail"
	cNsErrorSmSreplMgmtOpDisallowedWhenSolo               NsError = "SM_srepl_mgmt_op_disallowed_when_solo"
	cNsErrorSmNoMgmtSubnetSpecified                       NsError = "SM_no_mgmt_subnet_specified"
	cNsErrorSmFolderConflict                              NsError = "SM_folder_conflict"
	cNsErrorSmInvalidDataSubnet                           NsError = "SM_invalid_data_subnet"
	cNsErrorSmVolUsageUnavailable                         NsError = "SM_vol_usage_unavailable"
	cNsErrorSmAclScopeOverlap                             NsError = "SM_acl_scope_overlap"
	cNsErrorSmErrVvolAclGrpMerge                          NsError = "SM_err_vvol_acl_grp_merge"
	cNsErrorSmDisabledProtocolArtifacts                   NsError = "SM_disabled_protocol_artifacts"
	cNsErrorSmDuplicateSubnetVlanId                       NsError = "SM_duplicate_subnet_vlan_id"
	cNsErrorSmReplOpenstackUnsup                          NsError = "SM_repl_openstack_unsup"
	cNsErrorSmNoAvailableLun                              NsError = "SM_no_available_lun"
	cNsErrorSmGroupPartnerNameConflict                    NsError = "SM_group_partner_name_conflict"
	cNsErrorSmMgmtIpNotOnMgmt                             NsError = "SM_mgmt_ip_not_on_mgmt"
	cNsErrorSmInUseLun                                    NsError = "SM_in_use_lun"
	cNsErrorSmNotFcInitiatorgrp                           NsError = "SM_not_fc_initiatorgrp"
	cNsErrorSmErrShelfInvalidCsz                          NsError = "SM_err_shelf_invalid_csz"
	cNsErrorSmVolDedupeInvalidPerfPolicy                  NsError = "SM_vol_dedupe_invalid_perf_policy"
	cNsErrorSmEncryptionInvalidMode                       NsError = "SM_encryption_invalid_mode"
	cNsErrorSmAsupDisabled                                NsError = "SM_asup_disabled"
	cNsErrorSmErrShelfConnectedOnlyOneSide                NsError = "SM_err_shelf_connected_only_one_side"
	cNsErrorSmVolmvEnospace                               NsError = "SM_volmv_enospace"
	cNsErrorSmIncompatibleVolumes                         NsError = "SM_incompatible_volumes"
	cNsErrorSmComplexTypeQueryParam                       NsError = "SM_complex_type_query_param"
	cNsErrorSmAsupError                                   NsError = "SM_asup_error"
	cNsErrorSmReplUnassignedAppcat                        NsError = "SM_repl_unassigned_appcat"
	cNsErrorSmErrFcRegenerateInvalidOperationTdz          NsError = "SM_err_fc_regenerate_invalid_operation_tdz"
	cNsErrorSmInvalidPathVariable                         NsError = "SM_invalid_path_variable"
	cNsErrorSmErrArgChangeNotAllowed                      NsError = "SM_err_arg_change_not_allowed"
	cNsErrorSmSpecifiedSnapshotLun                        NsError = "SM_specified_snapshot_lun"
	cNsErrorSmSrepDownstreamAssocedVol                    NsError = "SM_srep_downstream_assoced_vol"
	cNsErrorSmLimitScopeValues                            NsError = "SM_limit_scope_values"
	cNsErrorSmErrPassphraseInval                          NsError = "SM_err_passphrase_inval"
	cNsErrorSmInitiatorgroupsOutOfSync                    NsError = "SM_initiatorgroups_out_of_sync"
	cNsErrorSmCachePinnedNotsup                           NsError = "SM_cache_pinned_notsup"
	cNsErrorSmNoOperationFound                            NsError = "SM_no_operation_found"
	cNsErrorSmInvalidDataForPartnerType                   NsError = "SM_invalid_data_for_partner_type"
	cNsErrorSmSrvUpdatePrecheck                           NsError = "SM_srv_update_precheck"
	cNsErrorSmVolDedupeEncryptionInvalid                  NsError = "SM_vol_dedupe_encryption_invalid"
	cNsErrorSmVvolFolderNoAppsrvr                         NsError = "SM_vvol_folder_no_appsrvr"
	cNsErrorSmHttpConflict                                NsError = "SM_http_conflict"
	cNsErrorSmArrayRenameInNetconfigFailed                NsError = "SM_array_rename_in_netconfig_failed"
	cNsErrorSmInvalidInitiatorIp                          NsError = "SM_invalid_initiator_ip"
	cNsErrorSmDuplicateVol                                NsError = "SM_duplicate_vol"
	cNsErrorSmErrShelfInvalidModel                        NsError = "SM_err_shelf_invalid_model"
	cNsErrorSmReplProtectLastSnap                         NsError = "SM_repl_protect_last_snap"
	cNsErrorSmErrGroupMergeBusy                           NsError = "SM_err_group_merge_busy"
	cNsErrorSmErrVolmvPoolMerging                         NsError = "SM_err_volmv_pool_merging"
	cNsErrorSmArrayNotFoundWithArgs                       NsError = "SM_array_not_found_with_args"
	cNsErrorSmConnectionRebalancingWithoutAutomaticMethod NsError = "SM_connection_rebalancing_without_automatic_method"
	cNsErrorSmSrepDownstreamNoCommonSnapVols              NsError = "SM_srep_downstream_no_common_snap_vols"
	cNsErrorSmErrShelfWrongSasPort                        NsError = "SM_err_shelf_wrong_sas_port"
	cNsErrorSmFolderUsageLimitBelowCurrentUsage           NsError = "SM_folder_usage_limit_below_current_usage"
	cNsErrorSmLimitSnapRetentionPool                      NsError = "SM_limit_snap_retention_pool"
	cNsErrorSmErrShelfExprFwVerInval                      NsError = "SM_err_shelf_expr_fw_ver_inval"
	cNsErrorSmSrepVolcollUnsup                            NsError = "SM_srep_volcoll_unsup"
	cNsErrorSmSrepNotGroupScopedVols                      NsError = "SM_srep_not_group_scoped_vols"
	cNsErrorSmFcIntfNotFound                              NsError = "SM_fc_intf_not_found"
	cNsErrorSmVolUnknown                                  NsError = "SM_vol_unknown"
	cNsErrorSmErrShelfEbusy                               NsError = "SM_err_shelf_ebusy"
	cNsErrorSmAppserverInUse                              NsError = "SM_appserver_in_use"
	cNsErrorSmPoolDoesNotHaveFolder                       NsError = "SM_pool_does_not_have_folder"
	cNsErrorSmPerfpolNotFound                             NsError = "SM_perfpol_not_found"
	cNsErrorSmPerfpolOob                                  NsError = "SM_perfpol_oob"
	cNsErrorSmDuplicateInitiator                          NsError = "SM_duplicate_initiator"
	cNsErrorSmInUseAppUuid                                NsError = "SM_in_use_app_uuid"
	cNsErrorSmPartnerCfgSync                              NsError = "SM_partner_cfg_sync"
	cNsErrorSmNotDownloadingSw                            NsError = "SM_not_downloading_sw"
	cNsErrorSmMissingMgmtIp                               NsError = "SM_missing_mgmt_ip"
	cNsErrorSmSrepDownstreamNoCommonSnapVol               NsError = "SM_srep_downstream_no_common_snap_vol"
	cNsErrorSmEbusy                                       NsError = "SM_ebusy"
	cNsErrorSmErrReplCantConnect                          NsError = "SM_err_repl_cant_connect"
	cNsErrorSmErrShelfExpanderErr                         NsError = "SM_err_shelf_expander_err"
	cNsErrorSmPeFailAclRemoval                            NsError = "SM_pe_fail_acl_removal"
	cNsErrorSmSupportIpNotOnMgmt                          NsError = "SM_support_ip_not_on_mgmt"
	cNsErrorSmAtLeastOneGroupSubnet                       NsError = "SM_at_least_one_group_subnet"
	cNsErrorSmUnsupportedAccessProtocol                   NsError = "SM_unsupported_access_protocol"
	cNsErrorSmSpaceInfoUnavail                            NsError = "SM_space_info_unavail"
	cNsErrorSmLimitVolmvHretSnapPool                      NsError = "SM_limit_volmv_hret_snap_pool"
	cNsErrorSmStartTimeBeyondEndTime                      NsError = "SM_start_time_beyond_end_time"
	cNsErrorSmReplRemotePaused                            NsError = "SM_repl_remote_paused"
	cNsErrorSmDataIpNotOnSubnet                           NsError = "SM_data_ip_not_on_subnet"
	cNsErrorSmConflictingInitiatorAlias                   NsError = "SM_conflicting_initiator_alias"
	cNsErrorSmReplEditPartnerNameNotPaused                NsError = "SM_repl_edit_partner_name_not_paused"
	cNsErrorSmShelfForeignDisk                            NsError = "SM_shelf_foreign_disk"
	cNsErrorSmQosLimitNotInRange                          NsError = "SM_qos_limit_not_in_range"
	cNsErrorSmErrShelfNoRserial                           NsError = "SM_err_shelf_no_rserial"
	cNsErrorSmUntaggedMtuNotLargest                       NsError = "SM_untagged_mtu_not_largest"
	cNsErrorSmErrShelfDisconnected                        NsError = "SM_err_shelf_disconnected"
	cNsErrorSmSubnetAlreadyAssignedOnNic                  NsError = "SM_subnet_already_assigned_on_nic"
	cNsErrorSmFcSessionExist                              NsError = "SM_fc_session_exist"
	cNsErrorSmVvolCannotOfflineBoundSnap                  NsError = "SM_vvol_cannot_offline_bound_snap"
	cNsErrorSmArrayMemberOrphan                           NsError = "SM_array_member_orphan"
	cNsErrorSmArrayMissingSubnet                          NsError = "SM_array_missing_subnet"
	cNsErrorSmDisableVvolWithPe                           NsError = "SM_disable_vvol_with_pe"
	cNsErrorSmLimitSnapVol                                NsError = "SM_limit_snap_vol"
	cNsErrorSmPoolUsageUnavailable                        NsError = "SM_pool_usage_unavailable"
	cNsErrorSmInvalidQueryParam                           NsError = "SM_invalid_query_param"
	cNsErrorSmErrGroupMergeDbLoad                         NsError = "SM_err_group_merge_db_load"
	cNsErrorSmEio                                         NsError = "SM_eio"
	cNsErrorSmSrepDownstreamOnlineVol                     NsError = "SM_srep_downstream_online_vol"
	cNsErrorSmPoolNoSrcArray                              NsError = "SM_pool_no_src_array"
	cNsErrorSmInvalidKeyvalue                             NsError = "SM_invalid_keyvalue"
	cNsErrorSmProtpolMaxLength                            NsError = "SM_protpol_max_length"
	cNsErrorSmVolAssocVolcoll                             NsError = "SM_vol_assoc_volcoll"
	cNsErrorSmMissingDiscoveryIp                          NsError = "SM_missing_discovery_ip"
	cNsErrorSmReplDeleteReplicaUnsup                      NsError = "SM_repl_delete_replica_unsup"
	cNsErrorSmVolSizeDecreased                            NsError = "SM_vol_size_decreased"
	cNsErrorSmSrepDownstreamAssocedVols                   NsError = "SM_srep_downstream_assoced_vols"
	cNsErrorSmAddNonfcToFcGroup                           NsError = "SM_add_nonfc_to_fc_group"
	cNsErrorSmNetconfigDoesNotExist                       NsError = "SM_netconfig_does_not_exist"
	cNsErrorSmFolderVolmvEnospace                         NsError = "SM_folder_volmv_enospace"
	cNsErrorSmNetconfigExistNoForce                       NsError = "SM_netconfig_exist_no_force"
	cNsErrorSmPoolUpdateInval                             NsError = "SM_pool_update_inval"
	cNsErrorSmVlanSubnetInManual                          NsError = "SM_vlan_subnet_in_manual"
	cNsErrorSmReplVvolUnsup                               NsError = "SM_repl_vvol_unsup"
	cNsErrorSmErrTdzNotSupported                          NsError = "SM_err_tdz_not_supported"
	cNsErrorSmPoolExists                                  NsError = "SM_pool_exists"
	cNsErrorSmAuth                                        NsError = "SM_auth"
	cNsErrorSmInvalidObjectSetQuery                       NsError = "SM_invalid_object_set_query"
	cNsErrorSmSrepDownstreamOnlineVols                    NsError = "SM_srep_downstream_online_vols"
	cNsErrorSmNoMethodForUrlPattern                       NsError = "SM_no_method_for_URL_pattern"
	cNsErrorSmVolNotOnline                                NsError = "SM_vol_not_online"
	cNsErrorSmVolDedupeVolfamAppcat                       NsError = "SM_vol_dedupe_volfam_appcat"
	cNsErrorSmInvalidNic                                  NsError = "SM_invalid_nic"
	cNsErrorSmArrayNotGroupLeader                         NsError = "SM_array_not_group_leader"
	cNsErrorSmInvalidVlanId                               NsError = "SM_invalid_vlan_id"
	cNsErrorSmLimitSnaplun                                NsError = "SM_limit_snaplun"
	cNsErrorSmIncompatibleCachePolicy                     NsError = "SM_incompatible_cache_policy"
	cNsErrorSmVolmvAbortEnospace                          NsError = "SM_volmv_abort_enospace"
	cNsErrorSmLimitHretSnapGroup                          NsError = "SM_limit_hret_snap_group"
	cNsErrorSmVolmvEalready                               NsError = "SM_volmv_ealready"
	cNsErrorSmAsupPingfromCtrlrbError                     NsError = "SM_asup_pingfrom_ctrlrB_error"
	cNsErrorSmVolDedupeUnassignedAppCategory              NsError = "SM_vol_dedupe_unassigned_app_category"
	cNsErrorSmEnoent                                      NsError = "SM_enoent"
	cNsErrorSmPerfpolDedupeUnassignedAppCategory          NsError = "SM_perfpol_dedupe_unassigned_app_category"
	cNsErrorSmInvalidInitiatorLabel                       NsError = "SM_invalid_initiator_label"
	cNsErrorSmDuplicateInitiatorWithArgs                  NsError = "SM_duplicate_initiator_with_args"
	cNsErrorSmErrShelfForeign                             NsError = "SM_err_shelf_foreign"
	cNsErrorSmInvalidAgentType                            NsError = "SM_invalid_agent_type"
	cNsErrorSmEinprogress                                 NsError = "SM_einprogress"
	cNsErrorSmNotEnoughCache                              NsError = "SM_not_enough_cache"
	cNsErrorSmEexist                                      NsError = "SM_eexist"
	cNsErrorSmMissingArrayNetconfig                       NsError = "SM_missing_array_netconfig"
	cNsErrorSmInvalidInitiatorAlias                       NsError = "SM_invalid_initiator_alias"
	cNsErrorSmProtpolAppSyncOracle                        NsError = "SM_protpol_app_sync_oracle"
	cNsErrorSmNoSupport                                   NsError = "SM_no_support"
	cNsErrorSmDataIpMissingSubnet                         NsError = "SM_data_ip_missing_subnet"
	cNsErrorSmErrShelfHddsInAfs                           NsError = "SM_err_shelf_hdds_in_afs"
	cNsErrorSmStartRowBeyondTotalRows                     NsError = "SM_start_row_beyond_total_rows"
	cNsErrorSmExtraneousArrayNetconfig                    NsError = "SM_extraneous_array_netconfig"
	cNsErrorSmPoolCachePinNotsupp                         NsError = "SM_pool_cache_pin_notsupp"
	cNsErrorSmUsageUnavailable                            NsError = "SM_usage_unavailable"
	cNsErrorSmReplAgentTypeUnsup                          NsError = "SM_repl_agent_type_unsup"
	cNsErrorSmReplFanoutMaximumPartnersExceeded           NsError = "SM_repl_fanout_maximum_partners_exceeded"
	cNsErrorSmAsupPingfromCtrlraError                     NsError = "SM_asup_pingfrom_ctrlrA_error"
	cNsErrorSmErrShelfInvalidLoc                          NsError = "SM_err_shelf_invalid_loc"
	cNsErrorSmErrShelfSsdDegraded                         NsError = "SM_err_shelf_ssd_degraded"
	cNsErrorSmSyncReplConfigure                           NsError = "SM_sync_repl_configure"
	cNsErrorSmAsupHeartbeatError                          NsError = "SM_asup_heartbeat_error"
	cNsErrorSmLimitHretSnapRetentionPoolWarn              NsError = "SM_limit_hret_snap_retention_pool_warn"
	cNsErrorSmErrAuth                                     NsError = "SM_err_auth"
	cNsErrorSmPartnerSubnetDoesNotExist                   NsError = "SM_partner_subnet_does_not_exist"
	cNsErrorSmErrShelfLocOrder                            NsError = "SM_err_shelf_loc_order"
	cNsErrorSmFolderEnospace                              NsError = "SM_folder_enospace"
	cNsErrorSmErrPoolHasGroupPartners                     NsError = "SM_err_pool_has_group_partners"
	cNsErrorSmProtpolNotSpecified                         NsError = "SM_protpol_not_specified"
	cNsErrorSmUnexpectedQueryParam                        NsError = "SM_unexpected_query_param"
	cNsErrorSmVolmvAbortEalready                          NsError = "SM_volmv_abort_ealready"
	cNsErrorSmFcIntfAlreadyInState                        NsError = "SM_fc_intf_already_in_state"
	cNsErrorSmLimitHretSnapPool                           NsError = "SM_limit_hret_snap_pool"
	cNsErrorSmIpUpdateNoForce                             NsError = "SM_ip_update_no_force"
	cNsErrorSmSrepAgentTypeMismatchDownstreamVol          NsError = "SM_srep_agent_type_mismatch_downstream_vol"
	cNsErrorSmVssValidationTimedout                       NsError = "SM_vss_validation_timedout"
	cNsErrorSmConfigSyncInprogress                        NsError = "SM_config_sync_inprogress"
	cNsErrorSmAsyncJobId                                  NsError = "SM_async_job_id"
	cNsErrorSmEagain                                      NsError = "SM_eagain"
	cNsErrorSmPerfpolVolMoveAppCategory                   NsError = "SM_perfpol_vol_move_app_category"
	cNsErrorSmLimitHretSnapRetentionPoolMax               NsError = "SM_limit_hret_snap_retention_pool_max"
	cNsErrorSmVolHasConnections                           NsError = "SM_vol_has_connections"
	cNsErrorSmNoCommonLun                                 NsError = "SM_no_common_lun"
	cNsErrorSmErrShelfSasLanesDegraded                    NsError = "SM_err_shelf_sas_lanes_degraded"
	cNsErrorSmVolAppCategoryMoveInvalid                   NsError = "SM_vol_app_category_move_invalid"
	cNsErrorSmExtTrigSchedAlreadyPresent                  NsError = "SM_ext_trig_sched_already_present"
	cNsErrorSmNoDataIpSpecified                           NsError = "SM_no_data_ip_specified"
	cNsErrorSmInvalidVolAssoc                             NsError = "SM_invalid_vol_assoc"
	cNsErrorSmReplObjectBusy                              NsError = "SM_repl_object_busy"
	cNsErrorSmVolcollOwner                                NsError = "SM_volcoll_owner"
	cNsErrorSmReservedUsername                            NsError = "SM_reserved_username"
	cNsErrorSmVvolsnapOnline                              NsError = "SM_vvolsnap_online"
	cNsErrorSmFolderVolmvEinprog                          NsError = "SM_folder_volmv_einprog"
	cNsErrorSmUnexpectedObjectSetQuery                    NsError = "SM_unexpected_object_set_query"
	cNsErrorSmProtpolInvalidValue                         NsError = "SM_protpol_invalid_value"
	cNsErrorSmFolderIncompatibleAgentType                 NsError = "SM_folder_incompatible_agent_type"
	cNsErrorSmProtpolVmwareInvalidVcenterHostname         NsError = "SM_protpol_vmware_invalid_vcenter_hostname"
	cNsErrorSmErrVolCollMultipleSchedules                 NsError = "SM_err_vol_coll_multiple_schedules"
	cNsErrorSmReplPartnerNameMismatch                     NsError = "SM_repl_partner_name_mismatch"
	cNsErrorSmInvalidFolder                               NsError = "SM_invalid_folder"
	cNsErrorSmSrvUpdatePrecheckArray                      NsError = "SM_srv_update_precheck_array"
	cNsErrorSmGatewayNotInSubnets                         NsError = "SM_gateway_not_in_subnets"
	cNsErrorSmDeprecatedPerfpol                           NsError = "SM_deprecated_perfpol"
	cNsErrorSmTakeoverSplitBrain                          NsError = "SM_takeover_split_brain"
	cNsErrorSmPeIgroupProtocolMismatched                  NsError = "SM_pe_igroup_protocol_mismatched"
	cNsErrorSmOnlyVvolFolderAppsrvr                       NsError = "SM_only_vvol_folder_appsrvr"
	cNsErrorSmVersionMismatch                             NsError = "SM_version_mismatch"
	cNsErrorSmPoolLastArray                               NsError = "SM_pool_last_array"
	cNsErrorSmEaccess                                     NsError = "SM_eaccess"
	cNsErrorSmInvalidSubnetLabel                          NsError = "SM_invalid_subnet_label"
	cNsErrorSmInvalidArg                                  NsError = "SM_invalid_arg"
	cNsErrorSmDedupeVolfamAppcat                          NsError = "SM_dedupe_volfam_appcat"
	cNsErrorSmSrvUnreach                                  NsError = "SM_srv_unreach"
	cNsErrorSmPoolPartnerPauseUnsup                       NsError = "SM_pool_partner_pause_unsup"
	cNsErrorSmNetconfigCreateDraftOnly                    NsError = "SM_netconfig_create_draft_only"
	cNsErrorSmErrArrayNotFound                            NsError = "SM_err_array_not_found"
	cNsErrorSmFolsetInUse                                 NsError = "SM_folset_in_use"
	cNsErrorSmErrShelfSesDeviceNotReady                   NsError = "SM_err_shelf_ses_device_not_ready"
	cNsErrorSmInvalidIp                                   NsError = "SM_invalid_ip"
	cNsErrorSmEalready                                    NsError = "SM_ealready"
	cNsErrorSmInvalidNonIscsiDataSubnet                   NsError = "SM_invalid_non_iscsi_data_subnet"
	cNsErrorSmReplHandoverUnsupPtype                      NsError = "SM_repl_handover_unsup_ptype"
	cNsErrorSmArrayNotAssigned                            NsError = "SM_array_not_assigned"
	cNsErrorSmEpartial                                    NsError = "SM_epartial"
	cNsErrorSmErrProtpolMissingName                       NsError = "SM_err_protpol_missing_name"
	cNsErrorSmVfVolCachePinned                            NsError = "SM_vf_vol_cache_pinned"
	cNsErrorSmAtLeastOneIscsiCluster                      NsError = "SM_at_least_one_iscsi_cluster"
	cNsErrorSmErrTooMany                                  NsError = "SM_err_too_many"
	cNsErrorSmErrShelfExprMfgVerInval                     NsError = "SM_err_shelf_expr_mfg_ver_inval"
	cNsErrorSmMgmtIpInvalid                               NsError = "SM_mgmt_ip_invalid"
	cNsErrorSmErrShelfWrongCtrlrSide                      NsError = "SM_err_shelf_wrong_ctrlr_side"
)

var pNsErrorSmReplVolcollDeletion NsError
var pNsErrorSmEncryptionGroupCipherOverride NsError
var pNsErrorSmEncryptionMustBeEnabled NsError
var pNsErrorSmExtTrigSchedNotPresent NsError
var pNsErrorSmAppserverNotFound NsError
var pNsErrorSmFolderReplPartner NsError
var pNsErrorSmArrayPoolMember NsError
var pNsErrorSmErrShelfCreateRfail NsError
var pNsErrorSmStarterVolCreate NsError
var pNsErrorSmInvalidNetconfigName NsError
var pNsErrorSmInvalidVolaclScope NsError
var pNsErrorSmErrShelfNoElocId NsError
var pNsErrorSmInvalidNetzoneType NsError
var pNsErrorSmErrCannotModifyTdz NsError
var pNsErrorSmControllerActive NsError
var pNsErrorSmPoolHasVolume NsError
var pNsErrorSmExtTrigSchedAttrib NsError
var pNsErrorSmMissingCriteriaFieldname NsError
var pNsErrorSmNetconfigAlreadyActive NsError
var pNsErrorSmStartRowBeyondEndRow NsError
var pNsErrorSmInsufficientFcInitiatorInput NsError
var pNsErrorSmSecondMgmtSubnet NsError
var pNsErrorSmInvalidInitiatorWwpn NsError
var pNsErrorSmArrayFlashMismatch NsError
var pNsErrorSmErrShelfPmdState NsError
var pNsErrorSmInvalidFcConfig NsError
var pNsErrorSmSrepDownstreamIsUpstream NsError
var pNsErrorSmKeymgrLeave NsError
var pNsErrorSmVolmvVolEinprog NsError
var pNsErrorSmSrepNameConflictVol NsError
var pNsErrorSmMultiArrayWithoutAutomaticConnectionMethod NsError
var pNsErrorSmFolderUsageLimitOverPoolCapacity NsError
var pNsErrorSmEnomem NsError
var pNsErrorSmErrShelfNotInuse NsError
var pNsErrorSmErrShelfPreactivationMfrErr NsError
var pNsErrorSmUpdateBusy NsError
var pNsErrorSmProtpolInvalidAppSync NsError
var pNsErrorSmErrShelfRaidDegraded NsError
var pNsErrorSmLimitFreqSchedGroup NsError
var pNsErrorSmSwupdateEinprog NsError
var pNsErrorSmMissingCriteriaParam NsError
var pNsErrorSmNoPathFound NsError
var pNsErrorSmIncompatibleAppCategory NsError
var pNsErrorSmVolmvAbortEnomove NsError
var pNsErrorSmLimitSnapRetentionVolcoll NsError
var pNsErrorSmKeymgrRemove NsError
var pNsErrorSmArrayNotReachable NsError
var pNsErrorSmErrGroupMergeEventsPending NsError
var pNsErrorSmSrepNameConflictVols NsError
var pNsErrorSmPoolPartnerResumeUnsup NsError
var pNsErrorSmArrayNotFound NsError
var pNsErrorSmNoIscsiHardware NsError
var pNsErrorSmEnospc NsError
var pNsErrorSmReservedPerfpolName NsError
var pNsErrorSmInvalidRoute NsError
var pNsErrorSmVolDedupeMoveInvalid NsError
var pNsErrorSmKeymgrJoin NsError
var pNsErrorSmErrSrvUnreach NsError
var pNsErrorSmErrShelfDiskSasLinkDegraded NsError
var pNsErrorSmErrShelfInvalidEeprom NsError
var pNsErrorSmInvalidKeyValue NsError
var pNsErrorSmNoIscsiLunAssignment NsError
var pNsErrorSmSnapshotOffline NsError
var pNsErrorSmDefaultGatewayNotInMgmtSubnet NsError
var pNsErrorSmErrPassphraseAuth NsError
var pNsErrorSmReplEexist NsError
var pNsErrorSmDisableLastProtocol NsError
var pNsErrorSmSecondUntaggedAssignment NsError
var pNsErrorSmPoolHasPe NsError
var pNsErrorSmErrShelfInvalidCount NsError
var pNsErrorSmLimitSnapRetentionGroup NsError
var pNsErrorSmVolDedupeThickProvInvalid NsError
var pNsErrorSmUnknown NsError
var pNsErrorSmErrShelfDedupeImpact NsError
var pNsErrorSmInvalidInitiatorAccessProtocol NsError
var pNsErrorSmInternal NsError
var pNsErrorSmAsupEbusy NsError
var pNsErrorSmInvalidVolMbpsLimit NsError
var pNsErrorSmInfoConfigSyncInprogress NsError
var pNsErrorSmErrPoolStillMerging NsError
var pNsErrorSmFolderPerfpolAgentType NsError
var pNsErrorSmEinval NsError
var pNsErrorSmNoAssocVols NsError
var pNsErrorSmShelfNotInuse NsError
var pNsErrorSmErrProtpolSettingsNotAllowed NsError
var pNsErrorSmFcRegenerate NsError
var pNsErrorSmErrMultiArrayGroup NsError
var pNsErrorSmReplRemoteNoBaseSnap NsError
var pNsErrorSmUnsupportedQueryOperator NsError
var pNsErrorSmSrepSizeMismatchDownstreamVol NsError
var pNsErrorSmNoSuchPartner NsError
var pNsErrorSmIscsiAllAccessNotAvailable NsError
var pNsErrorSmVvolAlreadyEnabled NsError
var pNsErrorSmUsageUnavaiable NsError
var pNsErrorSmEconnrefused NsError
var pNsErrorSmReplRenameNotsup NsError
var pNsErrorSmSessionCreate NsError
var pNsErrorSmConflictingAclVol NsError
var pNsErrorSmNospace NsError
var pNsErrorSmReservedVolcollName NsError
var pNsErrorSmPoolHasFolder NsError
var pNsErrorSmPartnerOffline NsError
var pNsErrorSmConflictingInitiatorAliasWithArgs NsError
var pNsErrorSmErrShelfNotReady NsError
var pNsErrorSmReplCmprVersionUnsup NsError
var pNsErrorSmErrShelfPreactivationIoErr NsError
var pNsErrorSmVolmvEinprog NsError
var pNsErrorSmPerfpolIncompatibleAppCategory NsError
var pNsErrorSmInvalidArgValue NsError
var pNsErrorSmErrShelfDedupeBelowFdr NsError
var pNsErrorSmErrShelfInvalidAfsCount NsError
var pNsErrorSmSerialNotAvail NsError
var pNsErrorSmErrShelfSesMshipErr NsError
var pNsErrorSmFolderProvisionedLimitBelowCurrentUsage NsError
var pNsErrorSmPeMultiProtocolAccessNotSupported NsError
var pNsErrorSmLimitHretSnapRetentionPool NsError
var pNsErrorSmPeConflictingAcl NsError
var pNsErrorSmFolderOverUsageLimit NsError
var pNsErrorSmErrEncryptionMasterKeyMissing NsError
var pNsErrorSmProtpolInvalidDomainName NsError
var pNsErrorSmTooMany NsError
var pNsErrorSmErrShelfExprRevIncompatible NsError
var pNsErrorSmSrvUnreachable NsError
var pNsErrorSmVolumeConflict NsError
var pNsErrorSmInvalidArrayName NsError
var pNsErrorSmCannotReadObject NsError
var pNsErrorSmReservedVolName NsError
var pNsErrorSmPoolPartnerInUse NsError
var pNsErrorSmInvalidInitiatorgrpName NsError
var pNsErrorSmAsupValidateError NsError
var pNsErrorSmVersionName NsError
var pNsErrorSmVvolAlreadyDisabled NsError
var pNsErrorSmUnexpectedArg NsError
var pNsErrorSmErrShelfInvalidAfs NsError
var pNsErrorSmUnexpectedChild NsError
var pNsErrorSmFolderOverProvisionedLimit NsError
var pNsErrorSmPoolFlashMismatch NsError
var pNsErrorSmLimitScope NsError
var pNsErrorSmSrepSizeMismatchDownstreamVols NsError
var pNsErrorSmFolderAppsrvrInconsistent NsError
var pNsErrorSmErrShelfDedupeReduction NsError
var pNsErrorSmIscsiSvcNotAvailable NsError
var pNsErrorSmInvalidNetconfigToDelete NsError
var pNsErrorSmErrUnknown NsError
var pNsErrorSmMissingCriteriaOperator NsError
var pNsErrorSmInvalidAppUuid NsError
var pNsErrorSmArrayMemberOrphanWithArgs NsError
var pNsErrorSmEmptyVol NsError
var pNsErrorSmDuplicateSubnetLabel NsError
var pNsErrorSmZeroVlanIdForTaggedAssignment NsError
var pNsErrorSmNoActionFound NsError
var pNsErrorSmSyncReplUnconfigureInProgress NsError
var pNsErrorSmVolThickProvMoveInvalid NsError
var pNsErrorSmMissingAdvancedCriteriaConstructor NsError
var pNsErrorSmPoolUnknown NsError
var pNsErrorSmHaltGlWithLiveMemberArray NsError
var pNsErrorSmReplThrottleOverlap NsError
var pNsErrorSmNetconfigUpdateMismatch NsError
var pNsErrorSmSrepDownstreamVolsAcl NsError
var pNsErrorSmStatsFrequencyInvalid NsError
var pNsErrorSmDdupFolderMerge NsError
var pNsErrorSmInvalidSubnet NsError
var pNsErrorSmReplEnabled NsError
var pNsErrorSmPoolPartnerNameConflict NsError
var pNsErrorSmAsupPingfromMgmtipError NsError
var pNsErrorSmEncryptionGroupScopeOverride NsError
var pNsErrorSmErrMaxSessionsReached NsError
var pNsErrorSmErrReplMultiplePartners NsError
var pNsErrorSmStatsNoSensors NsError
var pNsErrorSmFolderNeedsLimit NsError
var pNsErrorSmVolmvVvol NsError
var pNsErrorSmErrGroupMergeInprogress NsError
var pNsErrorSmFolderNotFound NsError
var pNsErrorSmRouteExists NsError
var pNsErrorSmInvalidDiscoveryIp NsError
var pNsErrorSmErrVolmvCachePinDedupeNotsupp NsError
var pNsErrorSmMissingCriteria NsError
var pNsErrorSmRootBranchPinned NsError
var pNsErrorSmInvalidMtu NsError
var pNsErrorSmNoFcHardware NsError
var pNsErrorSmPoolNotFound NsError
var pNsErrorSmDuplicateIp NsError
var pNsErrorSmNoAction NsError
var pNsErrorSmProtpolAppSyncOracleParams NsError
var pNsErrorSmProtpolInvalidVssSettings NsError
var pNsErrorSmEncryptionInvalidScope NsError
var pNsErrorSmLimitSnapcollVolcoll NsError
var pNsErrorSmShelfRaidDegraded NsError
var pNsErrorSmInvalidNonIscsiDataSubnetType NsError
var pNsErrorSmVolmvIncompatibleAgentType NsError
var pNsErrorSmReplPartnerVersionUnknown NsError
var pNsErrorSmAsupNameresError NsError
var pNsErrorSmReplApiUnsup NsError
var pNsErrorSmEperm NsError
var pNsErrorSmEnoentEnoent NsError
var pNsErrorSmArrayGroupLeader NsError
var pNsErrorSmInvalidInitiatorIqn NsError
var pNsErrorSmReplIntragroup NsError
var pNsErrorSmRemoveNonemptyFolder NsError
var pNsErrorSmAddNoniscsiToIscsiGroup NsError
var pNsErrorSmPoolDedupeIncapable NsError
var pNsErrorSmNoVvolSupport NsError
var pNsErrorSmNoInitiatorgrp NsError
var pNsErrorSmInvalidNetmask NsError
var pNsErrorSmNoDiscoveryIpInManualMode NsError
var pNsErrorSmInvalidVolReserveValues NsError
var pNsErrorSmSupportIpInvalid NsError
var pNsErrorSmErrShelfForeignDisk NsError
var pNsErrorSmVolsnapAlreadyExported NsError
var pNsErrorSmFcInitiatorgrpSubnetNotSupported NsError
var pNsErrorSmIncompatibleInitiatorAccessProtocol NsError
var pNsErrorSmInvalidInitiatorgrpAccessProtocol NsError
var pNsErrorSmLimitPeacl NsError
var pNsErrorSmEtimedout NsError
var pNsErrorSmInitiatorgrpSubnetDoesNotExist NsError
var pNsErrorSmVolOffline NsError
var pNsErrorSmErrSreplMgmtOpInProgress NsError
var pNsErrorSmReplFanoutMaximumCloudPartnersExceeded NsError
var pNsErrorSmMalformedUrl NsError
var pNsErrorSmErrSessionCreate NsError
var pNsErrorSmErrShelfCtrlrLoop NsError
var pNsErrorSmPeConflictingAclLun NsError
var pNsErrorSmInvalidCtrlrName NsError
var pNsErrorSmBackupNetconfigReadonly NsError
var pNsErrorSmLimitSnapGroup NsError
var pNsErrorSmEncryptionInvalidCipher NsError
var pNsErrorSmPerfpolInvalidAppCategory NsError
var pNsErrorSmArrayPoolFlashMismatch NsError
var pNsErrorSmFolderLimitInval NsError
var pNsErrorSmShelfSsdDegraded NsError
var pNsErrorSmFolderOverdraftLimitNeedsUsageLimit NsError
var pNsErrorSmEncryptionMasterKeyMissing NsError
var pNsErrorSmSrepAgentTypeMismatchDownstreamVols NsError
var pNsErrorSmErrVolNotOfflineOnRestore NsError
var pNsErrorSmReplHandoverBusy NsError
var pNsErrorSmNotOwner NsError
var pNsErrorSmSrepDownstreamAcl NsError
var pNsErrorSmNoSubnetWithLabel NsError
var pNsErrorSmPoolUpdateInvalArrays NsError
var pNsErrorSmVolHasOnlineSnap NsError
var pNsErrorSmInvalidDataIp NsError
var pNsErrorSmPoolVolmvEinprog NsError
var pNsErrorSmNoDataIpOnMgmtPlusData NsError
var pNsErrorSmConflictingAclLun NsError
var pNsErrorSmVpCreatedIgrp NsError
var pNsErrorSmStarterVolAclCreate NsError
var pNsErrorSmBadPkg NsError
var pNsErrorSmPasswdSameAsCurrent NsError
var pNsErrorSmSnaplunsOutOfSync NsError
var pNsErrorSmLimitSnapPool NsError
var pNsErrorSmMultiProtocolAccessNotSupported NsError
var pNsErrorSmVolHasClone NsError
var pNsErrorSmDedupeSingleArray NsError
var pNsErrorSmEncryptionMasterKeyInactive NsError
var pNsErrorSmSnapHasClone NsError
var pNsErrorSmErrEncryptionNeeded NsError
var pNsErrorSmMismatchingDuplicateSubnet NsError
var pNsErrorSmSrepNotGroupScopedVol NsError
var pNsErrorSmOnlyVvolFolderFolset NsError
var pNsErrorSmReplSnapshotSync NsError
var pNsErrorSmStatsNoSuchObject NsError
var pNsErrorSmDefaultRouteMissing NsError
var pNsErrorSmReplVasa3ApiUnsup NsError
var pNsErrorSmOverlappingSubnets NsError
var pNsErrorSmMissingArg NsError
var pNsErrorSmVolmvCachePinDedupeNotsupp NsError
var pNsErrorSmErrMergeConflict NsError
var pNsErrorSmTooSmall NsError
var pNsErrorSmVolWarnGreaterThanLimit NsError
var pNsErrorSmKeymgrUnreach NsError
var pNsErrorSmErrFcAsymmetryNotSupported NsError
var pNsErrorSmLunMismatch NsError
var pNsErrorSmErrShelfInvalidDiskCount NsError
var pNsErrorSmLimitVolacl NsError
var pNsErrorSmControllerNotActive NsError
var pNsErrorSmReplNoPartnerAvail NsError
var pNsErrorSmSreplMgmtOpDisallowedWhenSolo NsError
var pNsErrorSmNoMgmtSubnetSpecified NsError
var pNsErrorSmFolderConflict NsError
var pNsErrorSmInvalidDataSubnet NsError
var pNsErrorSmVolUsageUnavailable NsError
var pNsErrorSmAclScopeOverlap NsError
var pNsErrorSmErrVvolAclGrpMerge NsError
var pNsErrorSmDisabledProtocolArtifacts NsError
var pNsErrorSmDuplicateSubnetVlanId NsError
var pNsErrorSmReplOpenstackUnsup NsError
var pNsErrorSmNoAvailableLun NsError
var pNsErrorSmGroupPartnerNameConflict NsError
var pNsErrorSmMgmtIpNotOnMgmt NsError
var pNsErrorSmInUseLun NsError
var pNsErrorSmNotFcInitiatorgrp NsError
var pNsErrorSmErrShelfInvalidCsz NsError
var pNsErrorSmVolDedupeInvalidPerfPolicy NsError
var pNsErrorSmEncryptionInvalidMode NsError
var pNsErrorSmAsupDisabled NsError
var pNsErrorSmErrShelfConnectedOnlyOneSide NsError
var pNsErrorSmVolmvEnospace NsError
var pNsErrorSmIncompatibleVolumes NsError
var pNsErrorSmComplexTypeQueryParam NsError
var pNsErrorSmAsupError NsError
var pNsErrorSmReplUnassignedAppcat NsError
var pNsErrorSmErrFcRegenerateInvalidOperationTdz NsError
var pNsErrorSmInvalidPathVariable NsError
var pNsErrorSmErrArgChangeNotAllowed NsError
var pNsErrorSmSpecifiedSnapshotLun NsError
var pNsErrorSmSrepDownstreamAssocedVol NsError
var pNsErrorSmLimitScopeValues NsError
var pNsErrorSmErrPassphraseInval NsError
var pNsErrorSmInitiatorgroupsOutOfSync NsError
var pNsErrorSmCachePinnedNotsup NsError
var pNsErrorSmNoOperationFound NsError
var pNsErrorSmInvalidDataForPartnerType NsError
var pNsErrorSmSrvUpdatePrecheck NsError
var pNsErrorSmVolDedupeEncryptionInvalid NsError
var pNsErrorSmVvolFolderNoAppsrvr NsError
var pNsErrorSmHttpConflict NsError
var pNsErrorSmArrayRenameInNetconfigFailed NsError
var pNsErrorSmInvalidInitiatorIp NsError
var pNsErrorSmDuplicateVol NsError
var pNsErrorSmErrShelfInvalidModel NsError
var pNsErrorSmReplProtectLastSnap NsError
var pNsErrorSmErrGroupMergeBusy NsError
var pNsErrorSmErrVolmvPoolMerging NsError
var pNsErrorSmArrayNotFoundWithArgs NsError
var pNsErrorSmConnectionRebalancingWithoutAutomaticMethod NsError
var pNsErrorSmSrepDownstreamNoCommonSnapVols NsError
var pNsErrorSmErrShelfWrongSasPort NsError
var pNsErrorSmFolderUsageLimitBelowCurrentUsage NsError
var pNsErrorSmLimitSnapRetentionPool NsError
var pNsErrorSmErrShelfExprFwVerInval NsError
var pNsErrorSmSrepVolcollUnsup NsError
var pNsErrorSmSrepNotGroupScopedVols NsError
var pNsErrorSmFcIntfNotFound NsError
var pNsErrorSmVolUnknown NsError
var pNsErrorSmErrShelfEbusy NsError
var pNsErrorSmAppserverInUse NsError
var pNsErrorSmPoolDoesNotHaveFolder NsError
var pNsErrorSmPerfpolNotFound NsError
var pNsErrorSmPerfpolOob NsError
var pNsErrorSmDuplicateInitiator NsError
var pNsErrorSmInUseAppUuid NsError
var pNsErrorSmPartnerCfgSync NsError
var pNsErrorSmNotDownloadingSw NsError
var pNsErrorSmMissingMgmtIp NsError
var pNsErrorSmSrepDownstreamNoCommonSnapVol NsError
var pNsErrorSmEbusy NsError
var pNsErrorSmErrReplCantConnect NsError
var pNsErrorSmErrShelfExpanderErr NsError
var pNsErrorSmPeFailAclRemoval NsError
var pNsErrorSmSupportIpNotOnMgmt NsError
var pNsErrorSmAtLeastOneGroupSubnet NsError
var pNsErrorSmUnsupportedAccessProtocol NsError
var pNsErrorSmSpaceInfoUnavail NsError
var pNsErrorSmLimitVolmvHretSnapPool NsError
var pNsErrorSmStartTimeBeyondEndTime NsError
var pNsErrorSmReplRemotePaused NsError
var pNsErrorSmDataIpNotOnSubnet NsError
var pNsErrorSmConflictingInitiatorAlias NsError
var pNsErrorSmReplEditPartnerNameNotPaused NsError
var pNsErrorSmShelfForeignDisk NsError
var pNsErrorSmQosLimitNotInRange NsError
var pNsErrorSmErrShelfNoRserial NsError
var pNsErrorSmUntaggedMtuNotLargest NsError
var pNsErrorSmErrShelfDisconnected NsError
var pNsErrorSmSubnetAlreadyAssignedOnNic NsError
var pNsErrorSmFcSessionExist NsError
var pNsErrorSmVvolCannotOfflineBoundSnap NsError
var pNsErrorSmArrayMemberOrphan NsError
var pNsErrorSmArrayMissingSubnet NsError
var pNsErrorSmDisableVvolWithPe NsError
var pNsErrorSmLimitSnapVol NsError
var pNsErrorSmPoolUsageUnavailable NsError
var pNsErrorSmInvalidQueryParam NsError
var pNsErrorSmErrGroupMergeDbLoad NsError
var pNsErrorSmEio NsError
var pNsErrorSmSrepDownstreamOnlineVol NsError
var pNsErrorSmPoolNoSrcArray NsError
var pNsErrorSmInvalidKeyvalue NsError
var pNsErrorSmProtpolMaxLength NsError
var pNsErrorSmVolAssocVolcoll NsError
var pNsErrorSmMissingDiscoveryIp NsError
var pNsErrorSmReplDeleteReplicaUnsup NsError
var pNsErrorSmVolSizeDecreased NsError
var pNsErrorSmSrepDownstreamAssocedVols NsError
var pNsErrorSmAddNonfcToFcGroup NsError
var pNsErrorSmNetconfigDoesNotExist NsError
var pNsErrorSmFolderVolmvEnospace NsError
var pNsErrorSmNetconfigExistNoForce NsError
var pNsErrorSmPoolUpdateInval NsError
var pNsErrorSmVlanSubnetInManual NsError
var pNsErrorSmReplVvolUnsup NsError
var pNsErrorSmErrTdzNotSupported NsError
var pNsErrorSmPoolExists NsError
var pNsErrorSmAuth NsError
var pNsErrorSmInvalidObjectSetQuery NsError
var pNsErrorSmSrepDownstreamOnlineVols NsError
var pNsErrorSmNoMethodForUrlPattern NsError
var pNsErrorSmVolNotOnline NsError
var pNsErrorSmVolDedupeVolfamAppcat NsError
var pNsErrorSmInvalidNic NsError
var pNsErrorSmArrayNotGroupLeader NsError
var pNsErrorSmInvalidVlanId NsError
var pNsErrorSmLimitSnaplun NsError
var pNsErrorSmIncompatibleCachePolicy NsError
var pNsErrorSmVolmvAbortEnospace NsError
var pNsErrorSmLimitHretSnapGroup NsError
var pNsErrorSmVolmvEalready NsError
var pNsErrorSmAsupPingfromCtrlrbError NsError
var pNsErrorSmVolDedupeUnassignedAppCategory NsError
var pNsErrorSmEnoent NsError
var pNsErrorSmPerfpolDedupeUnassignedAppCategory NsError
var pNsErrorSmInvalidInitiatorLabel NsError
var pNsErrorSmDuplicateInitiatorWithArgs NsError
var pNsErrorSmErrShelfForeign NsError
var pNsErrorSmInvalidAgentType NsError
var pNsErrorSmEinprogress NsError
var pNsErrorSmNotEnoughCache NsError
var pNsErrorSmEexist NsError
var pNsErrorSmMissingArrayNetconfig NsError
var pNsErrorSmInvalidInitiatorAlias NsError
var pNsErrorSmProtpolAppSyncOracle NsError
var pNsErrorSmNoSupport NsError
var pNsErrorSmDataIpMissingSubnet NsError
var pNsErrorSmErrShelfHddsInAfs NsError
var pNsErrorSmStartRowBeyondTotalRows NsError
var pNsErrorSmExtraneousArrayNetconfig NsError
var pNsErrorSmPoolCachePinNotsupp NsError
var pNsErrorSmUsageUnavailable NsError
var pNsErrorSmReplAgentTypeUnsup NsError
var pNsErrorSmReplFanoutMaximumPartnersExceeded NsError
var pNsErrorSmAsupPingfromCtrlraError NsError
var pNsErrorSmErrShelfInvalidLoc NsError
var pNsErrorSmErrShelfSsdDegraded NsError
var pNsErrorSmSyncReplConfigure NsError
var pNsErrorSmAsupHeartbeatError NsError
var pNsErrorSmLimitHretSnapRetentionPoolWarn NsError
var pNsErrorSmErrAuth NsError
var pNsErrorSmPartnerSubnetDoesNotExist NsError
var pNsErrorSmErrShelfLocOrder NsError
var pNsErrorSmFolderEnospace NsError
var pNsErrorSmErrPoolHasGroupPartners NsError
var pNsErrorSmProtpolNotSpecified NsError
var pNsErrorSmUnexpectedQueryParam NsError
var pNsErrorSmVolmvAbortEalready NsError
var pNsErrorSmFcIntfAlreadyInState NsError
var pNsErrorSmLimitHretSnapPool NsError
var pNsErrorSmIpUpdateNoForce NsError
var pNsErrorSmSrepAgentTypeMismatchDownstreamVol NsError
var pNsErrorSmVssValidationTimedout NsError
var pNsErrorSmConfigSyncInprogress NsError
var pNsErrorSmAsyncJobId NsError
var pNsErrorSmEagain NsError
var pNsErrorSmPerfpolVolMoveAppCategory NsError
var pNsErrorSmLimitHretSnapRetentionPoolMax NsError
var pNsErrorSmVolHasConnections NsError
var pNsErrorSmNoCommonLun NsError
var pNsErrorSmErrShelfSasLanesDegraded NsError
var pNsErrorSmVolAppCategoryMoveInvalid NsError
var pNsErrorSmExtTrigSchedAlreadyPresent NsError
var pNsErrorSmNoDataIpSpecified NsError
var pNsErrorSmInvalidVolAssoc NsError
var pNsErrorSmReplObjectBusy NsError
var pNsErrorSmVolcollOwner NsError
var pNsErrorSmReservedUsername NsError
var pNsErrorSmVvolsnapOnline NsError
var pNsErrorSmFolderVolmvEinprog NsError
var pNsErrorSmUnexpectedObjectSetQuery NsError
var pNsErrorSmProtpolInvalidValue NsError
var pNsErrorSmFolderIncompatibleAgentType NsError
var pNsErrorSmProtpolVmwareInvalidVcenterHostname NsError
var pNsErrorSmErrVolCollMultipleSchedules NsError
var pNsErrorSmReplPartnerNameMismatch NsError
var pNsErrorSmInvalidFolder NsError
var pNsErrorSmSrvUpdatePrecheckArray NsError
var pNsErrorSmGatewayNotInSubnets NsError
var pNsErrorSmDeprecatedPerfpol NsError
var pNsErrorSmTakeoverSplitBrain NsError
var pNsErrorSmPeIgroupProtocolMismatched NsError
var pNsErrorSmOnlyVvolFolderAppsrvr NsError
var pNsErrorSmVersionMismatch NsError
var pNsErrorSmPoolLastArray NsError
var pNsErrorSmEaccess NsError
var pNsErrorSmInvalidSubnetLabel NsError
var pNsErrorSmInvalidArg NsError
var pNsErrorSmDedupeVolfamAppcat NsError
var pNsErrorSmSrvUnreach NsError
var pNsErrorSmPoolPartnerPauseUnsup NsError
var pNsErrorSmNetconfigCreateDraftOnly NsError
var pNsErrorSmErrArrayNotFound NsError
var pNsErrorSmFolsetInUse NsError
var pNsErrorSmErrShelfSesDeviceNotReady NsError
var pNsErrorSmInvalidIp NsError
var pNsErrorSmEalready NsError
var pNsErrorSmInvalidNonIscsiDataSubnet NsError
var pNsErrorSmReplHandoverUnsupPtype NsError
var pNsErrorSmArrayNotAssigned NsError
var pNsErrorSmEpartial NsError
var pNsErrorSmErrProtpolMissingName NsError
var pNsErrorSmVfVolCachePinned NsError
var pNsErrorSmAtLeastOneIscsiCluster NsError
var pNsErrorSmErrTooMany NsError
var pNsErrorSmErrShelfExprMfgVerInval NsError
var pNsErrorSmMgmtIpInvalid NsError
var pNsErrorSmErrShelfWrongCtrlrSide NsError

// NsErrorSmReplVolcollDeletion enum exports
var NsErrorSmReplVolcollDeletion *NsError

// NsErrorSmEncryptionGroupCipherOverride enum exports
var NsErrorSmEncryptionGroupCipherOverride *NsError

// NsErrorSmEncryptionMustBeEnabled enum exports
var NsErrorSmEncryptionMustBeEnabled *NsError

// NsErrorSmExtTrigSchedNotPresent enum exports
var NsErrorSmExtTrigSchedNotPresent *NsError

// NsErrorSmAppserverNotFound enum exports
var NsErrorSmAppserverNotFound *NsError

// NsErrorSmFolderReplPartner enum exports
var NsErrorSmFolderReplPartner *NsError

// NsErrorSmArrayPoolMember enum exports
var NsErrorSmArrayPoolMember *NsError

// NsErrorSmErrShelfCreateRfail enum exports
var NsErrorSmErrShelfCreateRfail *NsError

// NsErrorSmStarterVolCreate enum exports
var NsErrorSmStarterVolCreate *NsError

// NsErrorSmInvalidNetconfigName enum exports
var NsErrorSmInvalidNetconfigName *NsError

// NsErrorSmInvalidVolaclScope enum exports
var NsErrorSmInvalidVolaclScope *NsError

// NsErrorSmErrShelfNoElocId enum exports
var NsErrorSmErrShelfNoElocId *NsError

// NsErrorSmInvalidNetzoneType enum exports
var NsErrorSmInvalidNetzoneType *NsError

// NsErrorSmErrCannotModifyTdz enum exports
var NsErrorSmErrCannotModifyTdz *NsError

// NsErrorSmControllerActive enum exports
var NsErrorSmControllerActive *NsError

// NsErrorSmPoolHasVolume enum exports
var NsErrorSmPoolHasVolume *NsError

// NsErrorSmExtTrigSchedAttrib enum exports
var NsErrorSmExtTrigSchedAttrib *NsError

// NsErrorSmMissingCriteriaFieldname enum exports
var NsErrorSmMissingCriteriaFieldname *NsError

// NsErrorSmNetconfigAlreadyActive enum exports
var NsErrorSmNetconfigAlreadyActive *NsError

// NsErrorSmStartRowBeyondEndRow enum exports
var NsErrorSmStartRowBeyondEndRow *NsError

// NsErrorSmInsufficientFcInitiatorInput enum exports
var NsErrorSmInsufficientFcInitiatorInput *NsError

// NsErrorSmSecondMgmtSubnet enum exports
var NsErrorSmSecondMgmtSubnet *NsError

// NsErrorSmInvalidInitiatorWwpn enum exports
var NsErrorSmInvalidInitiatorWwpn *NsError

// NsErrorSmArrayFlashMismatch enum exports
var NsErrorSmArrayFlashMismatch *NsError

// NsErrorSmErrShelfPmdState enum exports
var NsErrorSmErrShelfPmdState *NsError

// NsErrorSmInvalidFcConfig enum exports
var NsErrorSmInvalidFcConfig *NsError

// NsErrorSmSrepDownstreamIsUpstream enum exports
var NsErrorSmSrepDownstreamIsUpstream *NsError

// NsErrorSmKeymgrLeave enum exports
var NsErrorSmKeymgrLeave *NsError

// NsErrorSmVolmvVolEinprog enum exports
var NsErrorSmVolmvVolEinprog *NsError

// NsErrorSmSrepNameConflictVol enum exports
var NsErrorSmSrepNameConflictVol *NsError

// NsErrorSmMultiArrayWithoutAutomaticConnectionMethod enum exports
var NsErrorSmMultiArrayWithoutAutomaticConnectionMethod *NsError

// NsErrorSmFolderUsageLimitOverPoolCapacity enum exports
var NsErrorSmFolderUsageLimitOverPoolCapacity *NsError

// NsErrorSmEnomem enum exports
var NsErrorSmEnomem *NsError

// NsErrorSmErrShelfNotInuse enum exports
var NsErrorSmErrShelfNotInuse *NsError

// NsErrorSmErrShelfPreactivationMfrErr enum exports
var NsErrorSmErrShelfPreactivationMfrErr *NsError

// NsErrorSmUpdateBusy enum exports
var NsErrorSmUpdateBusy *NsError

// NsErrorSmProtpolInvalidAppSync enum exports
var NsErrorSmProtpolInvalidAppSync *NsError

// NsErrorSmErrShelfRaidDegraded enum exports
var NsErrorSmErrShelfRaidDegraded *NsError

// NsErrorSmLimitFreqSchedGroup enum exports
var NsErrorSmLimitFreqSchedGroup *NsError

// NsErrorSmSwupdateEinprog enum exports
var NsErrorSmSwupdateEinprog *NsError

// NsErrorSmMissingCriteriaParam enum exports
var NsErrorSmMissingCriteriaParam *NsError

// NsErrorSmNoPathFound enum exports
var NsErrorSmNoPathFound *NsError

// NsErrorSmIncompatibleAppCategory enum exports
var NsErrorSmIncompatibleAppCategory *NsError

// NsErrorSmVolmvAbortEnomove enum exports
var NsErrorSmVolmvAbortEnomove *NsError

// NsErrorSmLimitSnapRetentionVolcoll enum exports
var NsErrorSmLimitSnapRetentionVolcoll *NsError

// NsErrorSmKeymgrRemove enum exports
var NsErrorSmKeymgrRemove *NsError

// NsErrorSmArrayNotReachable enum exports
var NsErrorSmArrayNotReachable *NsError

// NsErrorSmErrGroupMergeEventsPending enum exports
var NsErrorSmErrGroupMergeEventsPending *NsError

// NsErrorSmSrepNameConflictVols enum exports
var NsErrorSmSrepNameConflictVols *NsError

// NsErrorSmPoolPartnerResumeUnsup enum exports
var NsErrorSmPoolPartnerResumeUnsup *NsError

// NsErrorSmArrayNotFound enum exports
var NsErrorSmArrayNotFound *NsError

// NsErrorSmNoIscsiHardware enum exports
var NsErrorSmNoIscsiHardware *NsError

// NsErrorSmEnospc enum exports
var NsErrorSmEnospc *NsError

// NsErrorSmReservedPerfpolName enum exports
var NsErrorSmReservedPerfpolName *NsError

// NsErrorSmInvalidRoute enum exports
var NsErrorSmInvalidRoute *NsError

// NsErrorSmVolDedupeMoveInvalid enum exports
var NsErrorSmVolDedupeMoveInvalid *NsError

// NsErrorSmKeymgrJoin enum exports
var NsErrorSmKeymgrJoin *NsError

// NsErrorSmErrSrvUnreach enum exports
var NsErrorSmErrSrvUnreach *NsError

// NsErrorSmErrShelfDiskSasLinkDegraded enum exports
var NsErrorSmErrShelfDiskSasLinkDegraded *NsError

// NsErrorSmErrShelfInvalidEeprom enum exports
var NsErrorSmErrShelfInvalidEeprom *NsError

// NsErrorSmInvalidKeyValue enum exports
var NsErrorSmInvalidKeyValue *NsError

// NsErrorSmNoIscsiLunAssignment enum exports
var NsErrorSmNoIscsiLunAssignment *NsError

// NsErrorSmSnapshotOffline enum exports
var NsErrorSmSnapshotOffline *NsError

// NsErrorSmDefaultGatewayNotInMgmtSubnet enum exports
var NsErrorSmDefaultGatewayNotInMgmtSubnet *NsError

// NsErrorSmErrPassphraseAuth enum exports
var NsErrorSmErrPassphraseAuth *NsError

// NsErrorSmReplEexist enum exports
var NsErrorSmReplEexist *NsError

// NsErrorSmDisableLastProtocol enum exports
var NsErrorSmDisableLastProtocol *NsError

// NsErrorSmSecondUntaggedAssignment enum exports
var NsErrorSmSecondUntaggedAssignment *NsError

// NsErrorSmPoolHasPe enum exports
var NsErrorSmPoolHasPe *NsError

// NsErrorSmErrShelfInvalidCount enum exports
var NsErrorSmErrShelfInvalidCount *NsError

// NsErrorSmLimitSnapRetentionGroup enum exports
var NsErrorSmLimitSnapRetentionGroup *NsError

// NsErrorSmVolDedupeThickProvInvalid enum exports
var NsErrorSmVolDedupeThickProvInvalid *NsError

// NsErrorSmUnknown enum exports
var NsErrorSmUnknown *NsError

// NsErrorSmErrShelfDedupeImpact enum exports
var NsErrorSmErrShelfDedupeImpact *NsError

// NsErrorSmInvalidInitiatorAccessProtocol enum exports
var NsErrorSmInvalidInitiatorAccessProtocol *NsError

// NsErrorSmInternal enum exports
var NsErrorSmInternal *NsError

// NsErrorSmAsupEbusy enum exports
var NsErrorSmAsupEbusy *NsError

// NsErrorSmInvalidVolMbpsLimit enum exports
var NsErrorSmInvalidVolMbpsLimit *NsError

// NsErrorSmInfoConfigSyncInprogress enum exports
var NsErrorSmInfoConfigSyncInprogress *NsError

// NsErrorSmErrPoolStillMerging enum exports
var NsErrorSmErrPoolStillMerging *NsError

// NsErrorSmFolderPerfpolAgentType enum exports
var NsErrorSmFolderPerfpolAgentType *NsError

// NsErrorSmEinval enum exports
var NsErrorSmEinval *NsError

// NsErrorSmNoAssocVols enum exports
var NsErrorSmNoAssocVols *NsError

// NsErrorSmShelfNotInuse enum exports
var NsErrorSmShelfNotInuse *NsError

// NsErrorSmErrProtpolSettingsNotAllowed enum exports
var NsErrorSmErrProtpolSettingsNotAllowed *NsError

// NsErrorSmFcRegenerate enum exports
var NsErrorSmFcRegenerate *NsError

// NsErrorSmErrMultiArrayGroup enum exports
var NsErrorSmErrMultiArrayGroup *NsError

// NsErrorSmReplRemoteNoBaseSnap enum exports
var NsErrorSmReplRemoteNoBaseSnap *NsError

// NsErrorSmUnsupportedQueryOperator enum exports
var NsErrorSmUnsupportedQueryOperator *NsError

// NsErrorSmSrepSizeMismatchDownstreamVol enum exports
var NsErrorSmSrepSizeMismatchDownstreamVol *NsError

// NsErrorSmNoSuchPartner enum exports
var NsErrorSmNoSuchPartner *NsError

// NsErrorSmIscsiAllAccessNotAvailable enum exports
var NsErrorSmIscsiAllAccessNotAvailable *NsError

// NsErrorSmVvolAlreadyEnabled enum exports
var NsErrorSmVvolAlreadyEnabled *NsError

// NsErrorSmUsageUnavaiable enum exports
var NsErrorSmUsageUnavaiable *NsError

// NsErrorSmEconnrefused enum exports
var NsErrorSmEconnrefused *NsError

// NsErrorSmReplRenameNotsup enum exports
var NsErrorSmReplRenameNotsup *NsError

// NsErrorSmSessionCreate enum exports
var NsErrorSmSessionCreate *NsError

// NsErrorSmConflictingAclVol enum exports
var NsErrorSmConflictingAclVol *NsError

// NsErrorSmNospace enum exports
var NsErrorSmNospace *NsError

// NsErrorSmReservedVolcollName enum exports
var NsErrorSmReservedVolcollName *NsError

// NsErrorSmPoolHasFolder enum exports
var NsErrorSmPoolHasFolder *NsError

// NsErrorSmPartnerOffline enum exports
var NsErrorSmPartnerOffline *NsError

// NsErrorSmConflictingInitiatorAliasWithArgs enum exports
var NsErrorSmConflictingInitiatorAliasWithArgs *NsError

// NsErrorSmErrShelfNotReady enum exports
var NsErrorSmErrShelfNotReady *NsError

// NsErrorSmReplCmprVersionUnsup enum exports
var NsErrorSmReplCmprVersionUnsup *NsError

// NsErrorSmErrShelfPreactivationIoErr enum exports
var NsErrorSmErrShelfPreactivationIoErr *NsError

// NsErrorSmVolmvEinprog enum exports
var NsErrorSmVolmvEinprog *NsError

// NsErrorSmPerfpolIncompatibleAppCategory enum exports
var NsErrorSmPerfpolIncompatibleAppCategory *NsError

// NsErrorSmInvalidArgValue enum exports
var NsErrorSmInvalidArgValue *NsError

// NsErrorSmErrShelfDedupeBelowFdr enum exports
var NsErrorSmErrShelfDedupeBelowFdr *NsError

// NsErrorSmErrShelfInvalidAfsCount enum exports
var NsErrorSmErrShelfInvalidAfsCount *NsError

// NsErrorSmSerialNotAvail enum exports
var NsErrorSmSerialNotAvail *NsError

// NsErrorSmErrShelfSesMshipErr enum exports
var NsErrorSmErrShelfSesMshipErr *NsError

// NsErrorSmFolderProvisionedLimitBelowCurrentUsage enum exports
var NsErrorSmFolderProvisionedLimitBelowCurrentUsage *NsError

// NsErrorSmPeMultiProtocolAccessNotSupported enum exports
var NsErrorSmPeMultiProtocolAccessNotSupported *NsError

// NsErrorSmLimitHretSnapRetentionPool enum exports
var NsErrorSmLimitHretSnapRetentionPool *NsError

// NsErrorSmPeConflictingAcl enum exports
var NsErrorSmPeConflictingAcl *NsError

// NsErrorSmFolderOverUsageLimit enum exports
var NsErrorSmFolderOverUsageLimit *NsError

// NsErrorSmErrEncryptionMasterKeyMissing enum exports
var NsErrorSmErrEncryptionMasterKeyMissing *NsError

// NsErrorSmProtpolInvalidDomainName enum exports
var NsErrorSmProtpolInvalidDomainName *NsError

// NsErrorSmTooMany enum exports
var NsErrorSmTooMany *NsError

// NsErrorSmErrShelfExprRevIncompatible enum exports
var NsErrorSmErrShelfExprRevIncompatible *NsError

// NsErrorSmSrvUnreachable enum exports
var NsErrorSmSrvUnreachable *NsError

// NsErrorSmVolumeConflict enum exports
var NsErrorSmVolumeConflict *NsError

// NsErrorSmInvalidArrayName enum exports
var NsErrorSmInvalidArrayName *NsError

// NsErrorSmCannotReadObject enum exports
var NsErrorSmCannotReadObject *NsError

// NsErrorSmReservedVolName enum exports
var NsErrorSmReservedVolName *NsError

// NsErrorSmPoolPartnerInUse enum exports
var NsErrorSmPoolPartnerInUse *NsError

// NsErrorSmInvalidInitiatorgrpName enum exports
var NsErrorSmInvalidInitiatorgrpName *NsError

// NsErrorSmAsupValidateError enum exports
var NsErrorSmAsupValidateError *NsError

// NsErrorSmVersionName enum exports
var NsErrorSmVersionName *NsError

// NsErrorSmVvolAlreadyDisabled enum exports
var NsErrorSmVvolAlreadyDisabled *NsError

// NsErrorSmUnexpectedArg enum exports
var NsErrorSmUnexpectedArg *NsError

// NsErrorSmErrShelfInvalidAfs enum exports
var NsErrorSmErrShelfInvalidAfs *NsError

// NsErrorSmUnexpectedChild enum exports
var NsErrorSmUnexpectedChild *NsError

// NsErrorSmFolderOverProvisionedLimit enum exports
var NsErrorSmFolderOverProvisionedLimit *NsError

// NsErrorSmPoolFlashMismatch enum exports
var NsErrorSmPoolFlashMismatch *NsError

// NsErrorSmLimitScope enum exports
var NsErrorSmLimitScope *NsError

// NsErrorSmSrepSizeMismatchDownstreamVols enum exports
var NsErrorSmSrepSizeMismatchDownstreamVols *NsError

// NsErrorSmFolderAppsrvrInconsistent enum exports
var NsErrorSmFolderAppsrvrInconsistent *NsError

// NsErrorSmErrShelfDedupeReduction enum exports
var NsErrorSmErrShelfDedupeReduction *NsError

// NsErrorSmIscsiSvcNotAvailable enum exports
var NsErrorSmIscsiSvcNotAvailable *NsError

// NsErrorSmInvalidNetconfigToDelete enum exports
var NsErrorSmInvalidNetconfigToDelete *NsError

// NsErrorSmErrUnknown enum exports
var NsErrorSmErrUnknown *NsError

// NsErrorSmMissingCriteriaOperator enum exports
var NsErrorSmMissingCriteriaOperator *NsError

// NsErrorSmInvalidAppUuid enum exports
var NsErrorSmInvalidAppUuid *NsError

// NsErrorSmArrayMemberOrphanWithArgs enum exports
var NsErrorSmArrayMemberOrphanWithArgs *NsError

// NsErrorSmEmptyVol enum exports
var NsErrorSmEmptyVol *NsError

// NsErrorSmDuplicateSubnetLabel enum exports
var NsErrorSmDuplicateSubnetLabel *NsError

// NsErrorSmZeroVlanIdForTaggedAssignment enum exports
var NsErrorSmZeroVlanIdForTaggedAssignment *NsError

// NsErrorSmNoActionFound enum exports
var NsErrorSmNoActionFound *NsError

// NsErrorSmSyncReplUnconfigureInProgress enum exports
var NsErrorSmSyncReplUnconfigureInProgress *NsError

// NsErrorSmVolThickProvMoveInvalid enum exports
var NsErrorSmVolThickProvMoveInvalid *NsError

// NsErrorSmMissingAdvancedCriteriaConstructor enum exports
var NsErrorSmMissingAdvancedCriteriaConstructor *NsError

// NsErrorSmPoolUnknown enum exports
var NsErrorSmPoolUnknown *NsError

// NsErrorSmHaltGlWithLiveMemberArray enum exports
var NsErrorSmHaltGlWithLiveMemberArray *NsError

// NsErrorSmReplThrottleOverlap enum exports
var NsErrorSmReplThrottleOverlap *NsError

// NsErrorSmNetconfigUpdateMismatch enum exports
var NsErrorSmNetconfigUpdateMismatch *NsError

// NsErrorSmSrepDownstreamVolsAcl enum exports
var NsErrorSmSrepDownstreamVolsAcl *NsError

// NsErrorSmStatsFrequencyInvalid enum exports
var NsErrorSmStatsFrequencyInvalid *NsError

// NsErrorSmDdupFolderMerge enum exports
var NsErrorSmDdupFolderMerge *NsError

// NsErrorSmInvalidSubnet enum exports
var NsErrorSmInvalidSubnet *NsError

// NsErrorSmReplEnabled enum exports
var NsErrorSmReplEnabled *NsError

// NsErrorSmPoolPartnerNameConflict enum exports
var NsErrorSmPoolPartnerNameConflict *NsError

// NsErrorSmAsupPingfromMgmtipError enum exports
var NsErrorSmAsupPingfromMgmtipError *NsError

// NsErrorSmEncryptionGroupScopeOverride enum exports
var NsErrorSmEncryptionGroupScopeOverride *NsError

// NsErrorSmErrMaxSessionsReached enum exports
var NsErrorSmErrMaxSessionsReached *NsError

// NsErrorSmErrReplMultiplePartners enum exports
var NsErrorSmErrReplMultiplePartners *NsError

// NsErrorSmStatsNoSensors enum exports
var NsErrorSmStatsNoSensors *NsError

// NsErrorSmFolderNeedsLimit enum exports
var NsErrorSmFolderNeedsLimit *NsError

// NsErrorSmVolmvVvol enum exports
var NsErrorSmVolmvVvol *NsError

// NsErrorSmErrGroupMergeInprogress enum exports
var NsErrorSmErrGroupMergeInprogress *NsError

// NsErrorSmFolderNotFound enum exports
var NsErrorSmFolderNotFound *NsError

// NsErrorSmRouteExists enum exports
var NsErrorSmRouteExists *NsError

// NsErrorSmInvalidDiscoveryIp enum exports
var NsErrorSmInvalidDiscoveryIp *NsError

// NsErrorSmErrVolmvCachePinDedupeNotsupp enum exports
var NsErrorSmErrVolmvCachePinDedupeNotsupp *NsError

// NsErrorSmMissingCriteria enum exports
var NsErrorSmMissingCriteria *NsError

// NsErrorSmRootBranchPinned enum exports
var NsErrorSmRootBranchPinned *NsError

// NsErrorSmInvalidMtu enum exports
var NsErrorSmInvalidMtu *NsError

// NsErrorSmNoFcHardware enum exports
var NsErrorSmNoFcHardware *NsError

// NsErrorSmPoolNotFound enum exports
var NsErrorSmPoolNotFound *NsError

// NsErrorSmDuplicateIp enum exports
var NsErrorSmDuplicateIp *NsError

// NsErrorSmNoAction enum exports
var NsErrorSmNoAction *NsError

// NsErrorSmProtpolAppSyncOracleParams enum exports
var NsErrorSmProtpolAppSyncOracleParams *NsError

// NsErrorSmProtpolInvalidVssSettings enum exports
var NsErrorSmProtpolInvalidVssSettings *NsError

// NsErrorSmEncryptionInvalidScope enum exports
var NsErrorSmEncryptionInvalidScope *NsError

// NsErrorSmLimitSnapcollVolcoll enum exports
var NsErrorSmLimitSnapcollVolcoll *NsError

// NsErrorSmShelfRaidDegraded enum exports
var NsErrorSmShelfRaidDegraded *NsError

// NsErrorSmInvalidNonIscsiDataSubnetType enum exports
var NsErrorSmInvalidNonIscsiDataSubnetType *NsError

// NsErrorSmVolmvIncompatibleAgentType enum exports
var NsErrorSmVolmvIncompatibleAgentType *NsError

// NsErrorSmReplPartnerVersionUnknown enum exports
var NsErrorSmReplPartnerVersionUnknown *NsError

// NsErrorSmAsupNameresError enum exports
var NsErrorSmAsupNameresError *NsError

// NsErrorSmReplApiUnsup enum exports
var NsErrorSmReplApiUnsup *NsError

// NsErrorSmEperm enum exports
var NsErrorSmEperm *NsError

// NsErrorSmEnoentEnoent enum exports
var NsErrorSmEnoentEnoent *NsError

// NsErrorSmArrayGroupLeader enum exports
var NsErrorSmArrayGroupLeader *NsError

// NsErrorSmInvalidInitiatorIqn enum exports
var NsErrorSmInvalidInitiatorIqn *NsError

// NsErrorSmReplIntragroup enum exports
var NsErrorSmReplIntragroup *NsError

// NsErrorSmRemoveNonemptyFolder enum exports
var NsErrorSmRemoveNonemptyFolder *NsError

// NsErrorSmAddNoniscsiToIscsiGroup enum exports
var NsErrorSmAddNoniscsiToIscsiGroup *NsError

// NsErrorSmPoolDedupeIncapable enum exports
var NsErrorSmPoolDedupeIncapable *NsError

// NsErrorSmNoVvolSupport enum exports
var NsErrorSmNoVvolSupport *NsError

// NsErrorSmNoInitiatorgrp enum exports
var NsErrorSmNoInitiatorgrp *NsError

// NsErrorSmInvalidNetmask enum exports
var NsErrorSmInvalidNetmask *NsError

// NsErrorSmNoDiscoveryIpInManualMode enum exports
var NsErrorSmNoDiscoveryIpInManualMode *NsError

// NsErrorSmInvalidVolReserveValues enum exports
var NsErrorSmInvalidVolReserveValues *NsError

// NsErrorSmSupportIpInvalid enum exports
var NsErrorSmSupportIpInvalid *NsError

// NsErrorSmErrShelfForeignDisk enum exports
var NsErrorSmErrShelfForeignDisk *NsError

// NsErrorSmVolsnapAlreadyExported enum exports
var NsErrorSmVolsnapAlreadyExported *NsError

// NsErrorSmFcInitiatorgrpSubnetNotSupported enum exports
var NsErrorSmFcInitiatorgrpSubnetNotSupported *NsError

// NsErrorSmIncompatibleInitiatorAccessProtocol enum exports
var NsErrorSmIncompatibleInitiatorAccessProtocol *NsError

// NsErrorSmInvalidInitiatorgrpAccessProtocol enum exports
var NsErrorSmInvalidInitiatorgrpAccessProtocol *NsError

// NsErrorSmLimitPeacl enum exports
var NsErrorSmLimitPeacl *NsError

// NsErrorSmEtimedout enum exports
var NsErrorSmEtimedout *NsError

// NsErrorSmInitiatorgrpSubnetDoesNotExist enum exports
var NsErrorSmInitiatorgrpSubnetDoesNotExist *NsError

// NsErrorSmVolOffline enum exports
var NsErrorSmVolOffline *NsError

// NsErrorSmErrSreplMgmtOpInProgress enum exports
var NsErrorSmErrSreplMgmtOpInProgress *NsError

// NsErrorSmReplFanoutMaximumCloudPartnersExceeded enum exports
var NsErrorSmReplFanoutMaximumCloudPartnersExceeded *NsError

// NsErrorSmMalformedUrl enum exports
var NsErrorSmMalformedUrl *NsError

// NsErrorSmErrSessionCreate enum exports
var NsErrorSmErrSessionCreate *NsError

// NsErrorSmErrShelfCtrlrLoop enum exports
var NsErrorSmErrShelfCtrlrLoop *NsError

// NsErrorSmPeConflictingAclLun enum exports
var NsErrorSmPeConflictingAclLun *NsError

// NsErrorSmInvalidCtrlrName enum exports
var NsErrorSmInvalidCtrlrName *NsError

// NsErrorSmBackupNetconfigReadonly enum exports
var NsErrorSmBackupNetconfigReadonly *NsError

// NsErrorSmLimitSnapGroup enum exports
var NsErrorSmLimitSnapGroup *NsError

// NsErrorSmEncryptionInvalidCipher enum exports
var NsErrorSmEncryptionInvalidCipher *NsError

// NsErrorSmPerfpolInvalidAppCategory enum exports
var NsErrorSmPerfpolInvalidAppCategory *NsError

// NsErrorSmArrayPoolFlashMismatch enum exports
var NsErrorSmArrayPoolFlashMismatch *NsError

// NsErrorSmFolderLimitInval enum exports
var NsErrorSmFolderLimitInval *NsError

// NsErrorSmShelfSsdDegraded enum exports
var NsErrorSmShelfSsdDegraded *NsError

// NsErrorSmFolderOverdraftLimitNeedsUsageLimit enum exports
var NsErrorSmFolderOverdraftLimitNeedsUsageLimit *NsError

// NsErrorSmEncryptionMasterKeyMissing enum exports
var NsErrorSmEncryptionMasterKeyMissing *NsError

// NsErrorSmSrepAgentTypeMismatchDownstreamVols enum exports
var NsErrorSmSrepAgentTypeMismatchDownstreamVols *NsError

// NsErrorSmErrVolNotOfflineOnRestore enum exports
var NsErrorSmErrVolNotOfflineOnRestore *NsError

// NsErrorSmReplHandoverBusy enum exports
var NsErrorSmReplHandoverBusy *NsError

// NsErrorSmNotOwner enum exports
var NsErrorSmNotOwner *NsError

// NsErrorSmSrepDownstreamAcl enum exports
var NsErrorSmSrepDownstreamAcl *NsError

// NsErrorSmNoSubnetWithLabel enum exports
var NsErrorSmNoSubnetWithLabel *NsError

// NsErrorSmPoolUpdateInvalArrays enum exports
var NsErrorSmPoolUpdateInvalArrays *NsError

// NsErrorSmVolHasOnlineSnap enum exports
var NsErrorSmVolHasOnlineSnap *NsError

// NsErrorSmInvalidDataIp enum exports
var NsErrorSmInvalidDataIp *NsError

// NsErrorSmPoolVolmvEinprog enum exports
var NsErrorSmPoolVolmvEinprog *NsError

// NsErrorSmNoDataIpOnMgmtPlusData enum exports
var NsErrorSmNoDataIpOnMgmtPlusData *NsError

// NsErrorSmConflictingAclLun enum exports
var NsErrorSmConflictingAclLun *NsError

// NsErrorSmVpCreatedIgrp enum exports
var NsErrorSmVpCreatedIgrp *NsError

// NsErrorSmStarterVolAclCreate enum exports
var NsErrorSmStarterVolAclCreate *NsError

// NsErrorSmBadPkg enum exports
var NsErrorSmBadPkg *NsError

// NsErrorSmPasswdSameAsCurrent enum exports
var NsErrorSmPasswdSameAsCurrent *NsError

// NsErrorSmSnaplunsOutOfSync enum exports
var NsErrorSmSnaplunsOutOfSync *NsError

// NsErrorSmLimitSnapPool enum exports
var NsErrorSmLimitSnapPool *NsError

// NsErrorSmMultiProtocolAccessNotSupported enum exports
var NsErrorSmMultiProtocolAccessNotSupported *NsError

// NsErrorSmVolHasClone enum exports
var NsErrorSmVolHasClone *NsError

// NsErrorSmDedupeSingleArray enum exports
var NsErrorSmDedupeSingleArray *NsError

// NsErrorSmEncryptionMasterKeyInactive enum exports
var NsErrorSmEncryptionMasterKeyInactive *NsError

// NsErrorSmSnapHasClone enum exports
var NsErrorSmSnapHasClone *NsError

// NsErrorSmErrEncryptionNeeded enum exports
var NsErrorSmErrEncryptionNeeded *NsError

// NsErrorSmMismatchingDuplicateSubnet enum exports
var NsErrorSmMismatchingDuplicateSubnet *NsError

// NsErrorSmSrepNotGroupScopedVol enum exports
var NsErrorSmSrepNotGroupScopedVol *NsError

// NsErrorSmOnlyVvolFolderFolset enum exports
var NsErrorSmOnlyVvolFolderFolset *NsError

// NsErrorSmReplSnapshotSync enum exports
var NsErrorSmReplSnapshotSync *NsError

// NsErrorSmStatsNoSuchObject enum exports
var NsErrorSmStatsNoSuchObject *NsError

// NsErrorSmDefaultRouteMissing enum exports
var NsErrorSmDefaultRouteMissing *NsError

// NsErrorSmReplVasa3ApiUnsup enum exports
var NsErrorSmReplVasa3ApiUnsup *NsError

// NsErrorSmOverlappingSubnets enum exports
var NsErrorSmOverlappingSubnets *NsError

// NsErrorSmMissingArg enum exports
var NsErrorSmMissingArg *NsError

// NsErrorSmVolmvCachePinDedupeNotsupp enum exports
var NsErrorSmVolmvCachePinDedupeNotsupp *NsError

// NsErrorSmErrMergeConflict enum exports
var NsErrorSmErrMergeConflict *NsError

// NsErrorSmTooSmall enum exports
var NsErrorSmTooSmall *NsError

// NsErrorSmVolWarnGreaterThanLimit enum exports
var NsErrorSmVolWarnGreaterThanLimit *NsError

// NsErrorSmKeymgrUnreach enum exports
var NsErrorSmKeymgrUnreach *NsError

// NsErrorSmErrFcAsymmetryNotSupported enum exports
var NsErrorSmErrFcAsymmetryNotSupported *NsError

// NsErrorSmLunMismatch enum exports
var NsErrorSmLunMismatch *NsError

// NsErrorSmErrShelfInvalidDiskCount enum exports
var NsErrorSmErrShelfInvalidDiskCount *NsError

// NsErrorSmLimitVolacl enum exports
var NsErrorSmLimitVolacl *NsError

// NsErrorSmControllerNotActive enum exports
var NsErrorSmControllerNotActive *NsError

// NsErrorSmReplNoPartnerAvail enum exports
var NsErrorSmReplNoPartnerAvail *NsError

// NsErrorSmSreplMgmtOpDisallowedWhenSolo enum exports
var NsErrorSmSreplMgmtOpDisallowedWhenSolo *NsError

// NsErrorSmNoMgmtSubnetSpecified enum exports
var NsErrorSmNoMgmtSubnetSpecified *NsError

// NsErrorSmFolderConflict enum exports
var NsErrorSmFolderConflict *NsError

// NsErrorSmInvalidDataSubnet enum exports
var NsErrorSmInvalidDataSubnet *NsError

// NsErrorSmVolUsageUnavailable enum exports
var NsErrorSmVolUsageUnavailable *NsError

// NsErrorSmAclScopeOverlap enum exports
var NsErrorSmAclScopeOverlap *NsError

// NsErrorSmErrVvolAclGrpMerge enum exports
var NsErrorSmErrVvolAclGrpMerge *NsError

// NsErrorSmDisabledProtocolArtifacts enum exports
var NsErrorSmDisabledProtocolArtifacts *NsError

// NsErrorSmDuplicateSubnetVlanId enum exports
var NsErrorSmDuplicateSubnetVlanId *NsError

// NsErrorSmReplOpenstackUnsup enum exports
var NsErrorSmReplOpenstackUnsup *NsError

// NsErrorSmNoAvailableLun enum exports
var NsErrorSmNoAvailableLun *NsError

// NsErrorSmGroupPartnerNameConflict enum exports
var NsErrorSmGroupPartnerNameConflict *NsError

// NsErrorSmMgmtIpNotOnMgmt enum exports
var NsErrorSmMgmtIpNotOnMgmt *NsError

// NsErrorSmInUseLun enum exports
var NsErrorSmInUseLun *NsError

// NsErrorSmNotFcInitiatorgrp enum exports
var NsErrorSmNotFcInitiatorgrp *NsError

// NsErrorSmErrShelfInvalidCsz enum exports
var NsErrorSmErrShelfInvalidCsz *NsError

// NsErrorSmVolDedupeInvalidPerfPolicy enum exports
var NsErrorSmVolDedupeInvalidPerfPolicy *NsError

// NsErrorSmEncryptionInvalidMode enum exports
var NsErrorSmEncryptionInvalidMode *NsError

// NsErrorSmAsupDisabled enum exports
var NsErrorSmAsupDisabled *NsError

// NsErrorSmErrShelfConnectedOnlyOneSide enum exports
var NsErrorSmErrShelfConnectedOnlyOneSide *NsError

// NsErrorSmVolmvEnospace enum exports
var NsErrorSmVolmvEnospace *NsError

// NsErrorSmIncompatibleVolumes enum exports
var NsErrorSmIncompatibleVolumes *NsError

// NsErrorSmComplexTypeQueryParam enum exports
var NsErrorSmComplexTypeQueryParam *NsError

// NsErrorSmAsupError enum exports
var NsErrorSmAsupError *NsError

// NsErrorSmReplUnassignedAppcat enum exports
var NsErrorSmReplUnassignedAppcat *NsError

// NsErrorSmErrFcRegenerateInvalidOperationTdz enum exports
var NsErrorSmErrFcRegenerateInvalidOperationTdz *NsError

// NsErrorSmInvalidPathVariable enum exports
var NsErrorSmInvalidPathVariable *NsError

// NsErrorSmErrArgChangeNotAllowed enum exports
var NsErrorSmErrArgChangeNotAllowed *NsError

// NsErrorSmSpecifiedSnapshotLun enum exports
var NsErrorSmSpecifiedSnapshotLun *NsError

// NsErrorSmSrepDownstreamAssocedVol enum exports
var NsErrorSmSrepDownstreamAssocedVol *NsError

// NsErrorSmLimitScopeValues enum exports
var NsErrorSmLimitScopeValues *NsError

// NsErrorSmErrPassphraseInval enum exports
var NsErrorSmErrPassphraseInval *NsError

// NsErrorSmInitiatorgroupsOutOfSync enum exports
var NsErrorSmInitiatorgroupsOutOfSync *NsError

// NsErrorSmCachePinnedNotsup enum exports
var NsErrorSmCachePinnedNotsup *NsError

// NsErrorSmNoOperationFound enum exports
var NsErrorSmNoOperationFound *NsError

// NsErrorSmInvalidDataForPartnerType enum exports
var NsErrorSmInvalidDataForPartnerType *NsError

// NsErrorSmSrvUpdatePrecheck enum exports
var NsErrorSmSrvUpdatePrecheck *NsError

// NsErrorSmVolDedupeEncryptionInvalid enum exports
var NsErrorSmVolDedupeEncryptionInvalid *NsError

// NsErrorSmVvolFolderNoAppsrvr enum exports
var NsErrorSmVvolFolderNoAppsrvr *NsError

// NsErrorSmHttpConflict enum exports
var NsErrorSmHttpConflict *NsError

// NsErrorSmArrayRenameInNetconfigFailed enum exports
var NsErrorSmArrayRenameInNetconfigFailed *NsError

// NsErrorSmInvalidInitiatorIp enum exports
var NsErrorSmInvalidInitiatorIp *NsError

// NsErrorSmDuplicateVol enum exports
var NsErrorSmDuplicateVol *NsError

// NsErrorSmErrShelfInvalidModel enum exports
var NsErrorSmErrShelfInvalidModel *NsError

// NsErrorSmReplProtectLastSnap enum exports
var NsErrorSmReplProtectLastSnap *NsError

// NsErrorSmErrGroupMergeBusy enum exports
var NsErrorSmErrGroupMergeBusy *NsError

// NsErrorSmErrVolmvPoolMerging enum exports
var NsErrorSmErrVolmvPoolMerging *NsError

// NsErrorSmArrayNotFoundWithArgs enum exports
var NsErrorSmArrayNotFoundWithArgs *NsError

// NsErrorSmConnectionRebalancingWithoutAutomaticMethod enum exports
var NsErrorSmConnectionRebalancingWithoutAutomaticMethod *NsError

// NsErrorSmSrepDownstreamNoCommonSnapVols enum exports
var NsErrorSmSrepDownstreamNoCommonSnapVols *NsError

// NsErrorSmErrShelfWrongSasPort enum exports
var NsErrorSmErrShelfWrongSasPort *NsError

// NsErrorSmFolderUsageLimitBelowCurrentUsage enum exports
var NsErrorSmFolderUsageLimitBelowCurrentUsage *NsError

// NsErrorSmLimitSnapRetentionPool enum exports
var NsErrorSmLimitSnapRetentionPool *NsError

// NsErrorSmErrShelfExprFwVerInval enum exports
var NsErrorSmErrShelfExprFwVerInval *NsError

// NsErrorSmSrepVolcollUnsup enum exports
var NsErrorSmSrepVolcollUnsup *NsError

// NsErrorSmSrepNotGroupScopedVols enum exports
var NsErrorSmSrepNotGroupScopedVols *NsError

// NsErrorSmFcIntfNotFound enum exports
var NsErrorSmFcIntfNotFound *NsError

// NsErrorSmVolUnknown enum exports
var NsErrorSmVolUnknown *NsError

// NsErrorSmErrShelfEbusy enum exports
var NsErrorSmErrShelfEbusy *NsError

// NsErrorSmAppserverInUse enum exports
var NsErrorSmAppserverInUse *NsError

// NsErrorSmPoolDoesNotHaveFolder enum exports
var NsErrorSmPoolDoesNotHaveFolder *NsError

// NsErrorSmPerfpolNotFound enum exports
var NsErrorSmPerfpolNotFound *NsError

// NsErrorSmPerfpolOob enum exports
var NsErrorSmPerfpolOob *NsError

// NsErrorSmDuplicateInitiator enum exports
var NsErrorSmDuplicateInitiator *NsError

// NsErrorSmInUseAppUuid enum exports
var NsErrorSmInUseAppUuid *NsError

// NsErrorSmPartnerCfgSync enum exports
var NsErrorSmPartnerCfgSync *NsError

// NsErrorSmNotDownloadingSw enum exports
var NsErrorSmNotDownloadingSw *NsError

// NsErrorSmMissingMgmtIp enum exports
var NsErrorSmMissingMgmtIp *NsError

// NsErrorSmSrepDownstreamNoCommonSnapVol enum exports
var NsErrorSmSrepDownstreamNoCommonSnapVol *NsError

// NsErrorSmEbusy enum exports
var NsErrorSmEbusy *NsError

// NsErrorSmErrReplCantConnect enum exports
var NsErrorSmErrReplCantConnect *NsError

// NsErrorSmErrShelfExpanderErr enum exports
var NsErrorSmErrShelfExpanderErr *NsError

// NsErrorSmPeFailAclRemoval enum exports
var NsErrorSmPeFailAclRemoval *NsError

// NsErrorSmSupportIpNotOnMgmt enum exports
var NsErrorSmSupportIpNotOnMgmt *NsError

// NsErrorSmAtLeastOneGroupSubnet enum exports
var NsErrorSmAtLeastOneGroupSubnet *NsError

// NsErrorSmUnsupportedAccessProtocol enum exports
var NsErrorSmUnsupportedAccessProtocol *NsError

// NsErrorSmSpaceInfoUnavail enum exports
var NsErrorSmSpaceInfoUnavail *NsError

// NsErrorSmLimitVolmvHretSnapPool enum exports
var NsErrorSmLimitVolmvHretSnapPool *NsError

// NsErrorSmStartTimeBeyondEndTime enum exports
var NsErrorSmStartTimeBeyondEndTime *NsError

// NsErrorSmReplRemotePaused enum exports
var NsErrorSmReplRemotePaused *NsError

// NsErrorSmDataIpNotOnSubnet enum exports
var NsErrorSmDataIpNotOnSubnet *NsError

// NsErrorSmConflictingInitiatorAlias enum exports
var NsErrorSmConflictingInitiatorAlias *NsError

// NsErrorSmReplEditPartnerNameNotPaused enum exports
var NsErrorSmReplEditPartnerNameNotPaused *NsError

// NsErrorSmShelfForeignDisk enum exports
var NsErrorSmShelfForeignDisk *NsError

// NsErrorSmQosLimitNotInRange enum exports
var NsErrorSmQosLimitNotInRange *NsError

// NsErrorSmErrShelfNoRserial enum exports
var NsErrorSmErrShelfNoRserial *NsError

// NsErrorSmUntaggedMtuNotLargest enum exports
var NsErrorSmUntaggedMtuNotLargest *NsError

// NsErrorSmErrShelfDisconnected enum exports
var NsErrorSmErrShelfDisconnected *NsError

// NsErrorSmSubnetAlreadyAssignedOnNic enum exports
var NsErrorSmSubnetAlreadyAssignedOnNic *NsError

// NsErrorSmFcSessionExist enum exports
var NsErrorSmFcSessionExist *NsError

// NsErrorSmVvolCannotOfflineBoundSnap enum exports
var NsErrorSmVvolCannotOfflineBoundSnap *NsError

// NsErrorSmArrayMemberOrphan enum exports
var NsErrorSmArrayMemberOrphan *NsError

// NsErrorSmArrayMissingSubnet enum exports
var NsErrorSmArrayMissingSubnet *NsError

// NsErrorSmDisableVvolWithPe enum exports
var NsErrorSmDisableVvolWithPe *NsError

// NsErrorSmLimitSnapVol enum exports
var NsErrorSmLimitSnapVol *NsError

// NsErrorSmPoolUsageUnavailable enum exports
var NsErrorSmPoolUsageUnavailable *NsError

// NsErrorSmInvalidQueryParam enum exports
var NsErrorSmInvalidQueryParam *NsError

// NsErrorSmErrGroupMergeDbLoad enum exports
var NsErrorSmErrGroupMergeDbLoad *NsError

// NsErrorSmEio enum exports
var NsErrorSmEio *NsError

// NsErrorSmSrepDownstreamOnlineVol enum exports
var NsErrorSmSrepDownstreamOnlineVol *NsError

// NsErrorSmPoolNoSrcArray enum exports
var NsErrorSmPoolNoSrcArray *NsError

// NsErrorSmInvalidKeyvalue enum exports
var NsErrorSmInvalidKeyvalue *NsError

// NsErrorSmProtpolMaxLength enum exports
var NsErrorSmProtpolMaxLength *NsError

// NsErrorSmVolAssocVolcoll enum exports
var NsErrorSmVolAssocVolcoll *NsError

// NsErrorSmMissingDiscoveryIp enum exports
var NsErrorSmMissingDiscoveryIp *NsError

// NsErrorSmReplDeleteReplicaUnsup enum exports
var NsErrorSmReplDeleteReplicaUnsup *NsError

// NsErrorSmVolSizeDecreased enum exports
var NsErrorSmVolSizeDecreased *NsError

// NsErrorSmSrepDownstreamAssocedVols enum exports
var NsErrorSmSrepDownstreamAssocedVols *NsError

// NsErrorSmAddNonfcToFcGroup enum exports
var NsErrorSmAddNonfcToFcGroup *NsError

// NsErrorSmNetconfigDoesNotExist enum exports
var NsErrorSmNetconfigDoesNotExist *NsError

// NsErrorSmFolderVolmvEnospace enum exports
var NsErrorSmFolderVolmvEnospace *NsError

// NsErrorSmNetconfigExistNoForce enum exports
var NsErrorSmNetconfigExistNoForce *NsError

// NsErrorSmPoolUpdateInval enum exports
var NsErrorSmPoolUpdateInval *NsError

// NsErrorSmVlanSubnetInManual enum exports
var NsErrorSmVlanSubnetInManual *NsError

// NsErrorSmReplVvolUnsup enum exports
var NsErrorSmReplVvolUnsup *NsError

// NsErrorSmErrTdzNotSupported enum exports
var NsErrorSmErrTdzNotSupported *NsError

// NsErrorSmPoolExists enum exports
var NsErrorSmPoolExists *NsError

// NsErrorSmAuth enum exports
var NsErrorSmAuth *NsError

// NsErrorSmInvalidObjectSetQuery enum exports
var NsErrorSmInvalidObjectSetQuery *NsError

// NsErrorSmSrepDownstreamOnlineVols enum exports
var NsErrorSmSrepDownstreamOnlineVols *NsError

// NsErrorSmNoMethodForUrlPattern enum exports
var NsErrorSmNoMethodForUrlPattern *NsError

// NsErrorSmVolNotOnline enum exports
var NsErrorSmVolNotOnline *NsError

// NsErrorSmVolDedupeVolfamAppcat enum exports
var NsErrorSmVolDedupeVolfamAppcat *NsError

// NsErrorSmInvalidNic enum exports
var NsErrorSmInvalidNic *NsError

// NsErrorSmArrayNotGroupLeader enum exports
var NsErrorSmArrayNotGroupLeader *NsError

// NsErrorSmInvalidVlanId enum exports
var NsErrorSmInvalidVlanId *NsError

// NsErrorSmLimitSnaplun enum exports
var NsErrorSmLimitSnaplun *NsError

// NsErrorSmIncompatibleCachePolicy enum exports
var NsErrorSmIncompatibleCachePolicy *NsError

// NsErrorSmVolmvAbortEnospace enum exports
var NsErrorSmVolmvAbortEnospace *NsError

// NsErrorSmLimitHretSnapGroup enum exports
var NsErrorSmLimitHretSnapGroup *NsError

// NsErrorSmVolmvEalready enum exports
var NsErrorSmVolmvEalready *NsError

// NsErrorSmAsupPingfromCtrlrbError enum exports
var NsErrorSmAsupPingfromCtrlrbError *NsError

// NsErrorSmVolDedupeUnassignedAppCategory enum exports
var NsErrorSmVolDedupeUnassignedAppCategory *NsError

// NsErrorSmEnoent enum exports
var NsErrorSmEnoent *NsError

// NsErrorSmPerfpolDedupeUnassignedAppCategory enum exports
var NsErrorSmPerfpolDedupeUnassignedAppCategory *NsError

// NsErrorSmInvalidInitiatorLabel enum exports
var NsErrorSmInvalidInitiatorLabel *NsError

// NsErrorSmDuplicateInitiatorWithArgs enum exports
var NsErrorSmDuplicateInitiatorWithArgs *NsError

// NsErrorSmErrShelfForeign enum exports
var NsErrorSmErrShelfForeign *NsError

// NsErrorSmInvalidAgentType enum exports
var NsErrorSmInvalidAgentType *NsError

// NsErrorSmEinprogress enum exports
var NsErrorSmEinprogress *NsError

// NsErrorSmNotEnoughCache enum exports
var NsErrorSmNotEnoughCache *NsError

// NsErrorSmEexist enum exports
var NsErrorSmEexist *NsError

// NsErrorSmMissingArrayNetconfig enum exports
var NsErrorSmMissingArrayNetconfig *NsError

// NsErrorSmInvalidInitiatorAlias enum exports
var NsErrorSmInvalidInitiatorAlias *NsError

// NsErrorSmProtpolAppSyncOracle enum exports
var NsErrorSmProtpolAppSyncOracle *NsError

// NsErrorSmNoSupport enum exports
var NsErrorSmNoSupport *NsError

// NsErrorSmDataIpMissingSubnet enum exports
var NsErrorSmDataIpMissingSubnet *NsError

// NsErrorSmErrShelfHddsInAfs enum exports
var NsErrorSmErrShelfHddsInAfs *NsError

// NsErrorSmStartRowBeyondTotalRows enum exports
var NsErrorSmStartRowBeyondTotalRows *NsError

// NsErrorSmExtraneousArrayNetconfig enum exports
var NsErrorSmExtraneousArrayNetconfig *NsError

// NsErrorSmPoolCachePinNotsupp enum exports
var NsErrorSmPoolCachePinNotsupp *NsError

// NsErrorSmUsageUnavailable enum exports
var NsErrorSmUsageUnavailable *NsError

// NsErrorSmReplAgentTypeUnsup enum exports
var NsErrorSmReplAgentTypeUnsup *NsError

// NsErrorSmReplFanoutMaximumPartnersExceeded enum exports
var NsErrorSmReplFanoutMaximumPartnersExceeded *NsError

// NsErrorSmAsupPingfromCtrlraError enum exports
var NsErrorSmAsupPingfromCtrlraError *NsError

// NsErrorSmErrShelfInvalidLoc enum exports
var NsErrorSmErrShelfInvalidLoc *NsError

// NsErrorSmErrShelfSsdDegraded enum exports
var NsErrorSmErrShelfSsdDegraded *NsError

// NsErrorSmSyncReplConfigure enum exports
var NsErrorSmSyncReplConfigure *NsError

// NsErrorSmAsupHeartbeatError enum exports
var NsErrorSmAsupHeartbeatError *NsError

// NsErrorSmLimitHretSnapRetentionPoolWarn enum exports
var NsErrorSmLimitHretSnapRetentionPoolWarn *NsError

// NsErrorSmErrAuth enum exports
var NsErrorSmErrAuth *NsError

// NsErrorSmPartnerSubnetDoesNotExist enum exports
var NsErrorSmPartnerSubnetDoesNotExist *NsError

// NsErrorSmErrShelfLocOrder enum exports
var NsErrorSmErrShelfLocOrder *NsError

// NsErrorSmFolderEnospace enum exports
var NsErrorSmFolderEnospace *NsError

// NsErrorSmErrPoolHasGroupPartners enum exports
var NsErrorSmErrPoolHasGroupPartners *NsError

// NsErrorSmProtpolNotSpecified enum exports
var NsErrorSmProtpolNotSpecified *NsError

// NsErrorSmUnexpectedQueryParam enum exports
var NsErrorSmUnexpectedQueryParam *NsError

// NsErrorSmVolmvAbortEalready enum exports
var NsErrorSmVolmvAbortEalready *NsError

// NsErrorSmFcIntfAlreadyInState enum exports
var NsErrorSmFcIntfAlreadyInState *NsError

// NsErrorSmLimitHretSnapPool enum exports
var NsErrorSmLimitHretSnapPool *NsError

// NsErrorSmIpUpdateNoForce enum exports
var NsErrorSmIpUpdateNoForce *NsError

// NsErrorSmSrepAgentTypeMismatchDownstreamVol enum exports
var NsErrorSmSrepAgentTypeMismatchDownstreamVol *NsError

// NsErrorSmVssValidationTimedout enum exports
var NsErrorSmVssValidationTimedout *NsError

// NsErrorSmConfigSyncInprogress enum exports
var NsErrorSmConfigSyncInprogress *NsError

// NsErrorSmAsyncJobId enum exports
var NsErrorSmAsyncJobId *NsError

// NsErrorSmEagain enum exports
var NsErrorSmEagain *NsError

// NsErrorSmPerfpolVolMoveAppCategory enum exports
var NsErrorSmPerfpolVolMoveAppCategory *NsError

// NsErrorSmLimitHretSnapRetentionPoolMax enum exports
var NsErrorSmLimitHretSnapRetentionPoolMax *NsError

// NsErrorSmVolHasConnections enum exports
var NsErrorSmVolHasConnections *NsError

// NsErrorSmNoCommonLun enum exports
var NsErrorSmNoCommonLun *NsError

// NsErrorSmErrShelfSasLanesDegraded enum exports
var NsErrorSmErrShelfSasLanesDegraded *NsError

// NsErrorSmVolAppCategoryMoveInvalid enum exports
var NsErrorSmVolAppCategoryMoveInvalid *NsError

// NsErrorSmExtTrigSchedAlreadyPresent enum exports
var NsErrorSmExtTrigSchedAlreadyPresent *NsError

// NsErrorSmNoDataIpSpecified enum exports
var NsErrorSmNoDataIpSpecified *NsError

// NsErrorSmInvalidVolAssoc enum exports
var NsErrorSmInvalidVolAssoc *NsError

// NsErrorSmReplObjectBusy enum exports
var NsErrorSmReplObjectBusy *NsError

// NsErrorSmVolcollOwner enum exports
var NsErrorSmVolcollOwner *NsError

// NsErrorSmReservedUsername enum exports
var NsErrorSmReservedUsername *NsError

// NsErrorSmVvolsnapOnline enum exports
var NsErrorSmVvolsnapOnline *NsError

// NsErrorSmFolderVolmvEinprog enum exports
var NsErrorSmFolderVolmvEinprog *NsError

// NsErrorSmUnexpectedObjectSetQuery enum exports
var NsErrorSmUnexpectedObjectSetQuery *NsError

// NsErrorSmProtpolInvalidValue enum exports
var NsErrorSmProtpolInvalidValue *NsError

// NsErrorSmFolderIncompatibleAgentType enum exports
var NsErrorSmFolderIncompatibleAgentType *NsError

// NsErrorSmProtpolVmwareInvalidVcenterHostname enum exports
var NsErrorSmProtpolVmwareInvalidVcenterHostname *NsError

// NsErrorSmErrVolCollMultipleSchedules enum exports
var NsErrorSmErrVolCollMultipleSchedules *NsError

// NsErrorSmReplPartnerNameMismatch enum exports
var NsErrorSmReplPartnerNameMismatch *NsError

// NsErrorSmInvalidFolder enum exports
var NsErrorSmInvalidFolder *NsError

// NsErrorSmSrvUpdatePrecheckArray enum exports
var NsErrorSmSrvUpdatePrecheckArray *NsError

// NsErrorSmGatewayNotInSubnets enum exports
var NsErrorSmGatewayNotInSubnets *NsError

// NsErrorSmDeprecatedPerfpol enum exports
var NsErrorSmDeprecatedPerfpol *NsError

// NsErrorSmTakeoverSplitBrain enum exports
var NsErrorSmTakeoverSplitBrain *NsError

// NsErrorSmPeIgroupProtocolMismatched enum exports
var NsErrorSmPeIgroupProtocolMismatched *NsError

// NsErrorSmOnlyVvolFolderAppsrvr enum exports
var NsErrorSmOnlyVvolFolderAppsrvr *NsError

// NsErrorSmVersionMismatch enum exports
var NsErrorSmVersionMismatch *NsError

// NsErrorSmPoolLastArray enum exports
var NsErrorSmPoolLastArray *NsError

// NsErrorSmEaccess enum exports
var NsErrorSmEaccess *NsError

// NsErrorSmInvalidSubnetLabel enum exports
var NsErrorSmInvalidSubnetLabel *NsError

// NsErrorSmInvalidArg enum exports
var NsErrorSmInvalidArg *NsError

// NsErrorSmDedupeVolfamAppcat enum exports
var NsErrorSmDedupeVolfamAppcat *NsError

// NsErrorSmSrvUnreach enum exports
var NsErrorSmSrvUnreach *NsError

// NsErrorSmPoolPartnerPauseUnsup enum exports
var NsErrorSmPoolPartnerPauseUnsup *NsError

// NsErrorSmNetconfigCreateDraftOnly enum exports
var NsErrorSmNetconfigCreateDraftOnly *NsError

// NsErrorSmErrArrayNotFound enum exports
var NsErrorSmErrArrayNotFound *NsError

// NsErrorSmFolsetInUse enum exports
var NsErrorSmFolsetInUse *NsError

// NsErrorSmErrShelfSesDeviceNotReady enum exports
var NsErrorSmErrShelfSesDeviceNotReady *NsError

// NsErrorSmInvalidIp enum exports
var NsErrorSmInvalidIp *NsError

// NsErrorSmEalready enum exports
var NsErrorSmEalready *NsError

// NsErrorSmInvalidNonIscsiDataSubnet enum exports
var NsErrorSmInvalidNonIscsiDataSubnet *NsError

// NsErrorSmReplHandoverUnsupPtype enum exports
var NsErrorSmReplHandoverUnsupPtype *NsError

// NsErrorSmArrayNotAssigned enum exports
var NsErrorSmArrayNotAssigned *NsError

// NsErrorSmEpartial enum exports
var NsErrorSmEpartial *NsError

// NsErrorSmErrProtpolMissingName enum exports
var NsErrorSmErrProtpolMissingName *NsError

// NsErrorSmVfVolCachePinned enum exports
var NsErrorSmVfVolCachePinned *NsError

// NsErrorSmAtLeastOneIscsiCluster enum exports
var NsErrorSmAtLeastOneIscsiCluster *NsError

// NsErrorSmErrTooMany enum exports
var NsErrorSmErrTooMany *NsError

// NsErrorSmErrShelfExprMfgVerInval enum exports
var NsErrorSmErrShelfExprMfgVerInval *NsError

// NsErrorSmMgmtIpInvalid enum exports
var NsErrorSmMgmtIpInvalid *NsError

// NsErrorSmErrShelfWrongCtrlrSide enum exports
var NsErrorSmErrShelfWrongCtrlrSide *NsError

func init() {
	pNsErrorSmReplVolcollDeletion = cNsErrorSmReplVolcollDeletion
	NsErrorSmReplVolcollDeletion = &pNsErrorSmReplVolcollDeletion

	pNsErrorSmEncryptionGroupCipherOverride = cNsErrorSmEncryptionGroupCipherOverride
	NsErrorSmEncryptionGroupCipherOverride = &pNsErrorSmEncryptionGroupCipherOverride

	pNsErrorSmEncryptionMustBeEnabled = cNsErrorSmEncryptionMustBeEnabled
	NsErrorSmEncryptionMustBeEnabled = &pNsErrorSmEncryptionMustBeEnabled

	pNsErrorSmExtTrigSchedNotPresent = cNsErrorSmExtTrigSchedNotPresent
	NsErrorSmExtTrigSchedNotPresent = &pNsErrorSmExtTrigSchedNotPresent

	pNsErrorSmAppserverNotFound = cNsErrorSmAppserverNotFound
	NsErrorSmAppserverNotFound = &pNsErrorSmAppserverNotFound

	pNsErrorSmFolderReplPartner = cNsErrorSmFolderReplPartner
	NsErrorSmFolderReplPartner = &pNsErrorSmFolderReplPartner

	pNsErrorSmArrayPoolMember = cNsErrorSmArrayPoolMember
	NsErrorSmArrayPoolMember = &pNsErrorSmArrayPoolMember

	pNsErrorSmErrShelfCreateRfail = cNsErrorSmErrShelfCreateRfail
	NsErrorSmErrShelfCreateRfail = &pNsErrorSmErrShelfCreateRfail

	pNsErrorSmStarterVolCreate = cNsErrorSmStarterVolCreate
	NsErrorSmStarterVolCreate = &pNsErrorSmStarterVolCreate

	pNsErrorSmInvalidNetconfigName = cNsErrorSmInvalidNetconfigName
	NsErrorSmInvalidNetconfigName = &pNsErrorSmInvalidNetconfigName

	pNsErrorSmInvalidVolaclScope = cNsErrorSmInvalidVolaclScope
	NsErrorSmInvalidVolaclScope = &pNsErrorSmInvalidVolaclScope

	pNsErrorSmErrShelfNoElocId = cNsErrorSmErrShelfNoElocId
	NsErrorSmErrShelfNoElocId = &pNsErrorSmErrShelfNoElocId

	pNsErrorSmInvalidNetzoneType = cNsErrorSmInvalidNetzoneType
	NsErrorSmInvalidNetzoneType = &pNsErrorSmInvalidNetzoneType

	pNsErrorSmErrCannotModifyTdz = cNsErrorSmErrCannotModifyTdz
	NsErrorSmErrCannotModifyTdz = &pNsErrorSmErrCannotModifyTdz

	pNsErrorSmControllerActive = cNsErrorSmControllerActive
	NsErrorSmControllerActive = &pNsErrorSmControllerActive

	pNsErrorSmPoolHasVolume = cNsErrorSmPoolHasVolume
	NsErrorSmPoolHasVolume = &pNsErrorSmPoolHasVolume

	pNsErrorSmExtTrigSchedAttrib = cNsErrorSmExtTrigSchedAttrib
	NsErrorSmExtTrigSchedAttrib = &pNsErrorSmExtTrigSchedAttrib

	pNsErrorSmMissingCriteriaFieldname = cNsErrorSmMissingCriteriaFieldname
	NsErrorSmMissingCriteriaFieldname = &pNsErrorSmMissingCriteriaFieldname

	pNsErrorSmNetconfigAlreadyActive = cNsErrorSmNetconfigAlreadyActive
	NsErrorSmNetconfigAlreadyActive = &pNsErrorSmNetconfigAlreadyActive

	pNsErrorSmStartRowBeyondEndRow = cNsErrorSmStartRowBeyondEndRow
	NsErrorSmStartRowBeyondEndRow = &pNsErrorSmStartRowBeyondEndRow

	pNsErrorSmInsufficientFcInitiatorInput = cNsErrorSmInsufficientFcInitiatorInput
	NsErrorSmInsufficientFcInitiatorInput = &pNsErrorSmInsufficientFcInitiatorInput

	pNsErrorSmSecondMgmtSubnet = cNsErrorSmSecondMgmtSubnet
	NsErrorSmSecondMgmtSubnet = &pNsErrorSmSecondMgmtSubnet

	pNsErrorSmInvalidInitiatorWwpn = cNsErrorSmInvalidInitiatorWwpn
	NsErrorSmInvalidInitiatorWwpn = &pNsErrorSmInvalidInitiatorWwpn

	pNsErrorSmArrayFlashMismatch = cNsErrorSmArrayFlashMismatch
	NsErrorSmArrayFlashMismatch = &pNsErrorSmArrayFlashMismatch

	pNsErrorSmErrShelfPmdState = cNsErrorSmErrShelfPmdState
	NsErrorSmErrShelfPmdState = &pNsErrorSmErrShelfPmdState

	pNsErrorSmInvalidFcConfig = cNsErrorSmInvalidFcConfig
	NsErrorSmInvalidFcConfig = &pNsErrorSmInvalidFcConfig

	pNsErrorSmSrepDownstreamIsUpstream = cNsErrorSmSrepDownstreamIsUpstream
	NsErrorSmSrepDownstreamIsUpstream = &pNsErrorSmSrepDownstreamIsUpstream

	pNsErrorSmKeymgrLeave = cNsErrorSmKeymgrLeave
	NsErrorSmKeymgrLeave = &pNsErrorSmKeymgrLeave

	pNsErrorSmVolmvVolEinprog = cNsErrorSmVolmvVolEinprog
	NsErrorSmVolmvVolEinprog = &pNsErrorSmVolmvVolEinprog

	pNsErrorSmSrepNameConflictVol = cNsErrorSmSrepNameConflictVol
	NsErrorSmSrepNameConflictVol = &pNsErrorSmSrepNameConflictVol

	pNsErrorSmMultiArrayWithoutAutomaticConnectionMethod = cNsErrorSmMultiArrayWithoutAutomaticConnectionMethod
	NsErrorSmMultiArrayWithoutAutomaticConnectionMethod = &pNsErrorSmMultiArrayWithoutAutomaticConnectionMethod

	pNsErrorSmFolderUsageLimitOverPoolCapacity = cNsErrorSmFolderUsageLimitOverPoolCapacity
	NsErrorSmFolderUsageLimitOverPoolCapacity = &pNsErrorSmFolderUsageLimitOverPoolCapacity

	pNsErrorSmEnomem = cNsErrorSmEnomem
	NsErrorSmEnomem = &pNsErrorSmEnomem

	pNsErrorSmErrShelfNotInuse = cNsErrorSmErrShelfNotInuse
	NsErrorSmErrShelfNotInuse = &pNsErrorSmErrShelfNotInuse

	pNsErrorSmErrShelfPreactivationMfrErr = cNsErrorSmErrShelfPreactivationMfrErr
	NsErrorSmErrShelfPreactivationMfrErr = &pNsErrorSmErrShelfPreactivationMfrErr

	pNsErrorSmUpdateBusy = cNsErrorSmUpdateBusy
	NsErrorSmUpdateBusy = &pNsErrorSmUpdateBusy

	pNsErrorSmProtpolInvalidAppSync = cNsErrorSmProtpolInvalidAppSync
	NsErrorSmProtpolInvalidAppSync = &pNsErrorSmProtpolInvalidAppSync

	pNsErrorSmErrShelfRaidDegraded = cNsErrorSmErrShelfRaidDegraded
	NsErrorSmErrShelfRaidDegraded = &pNsErrorSmErrShelfRaidDegraded

	pNsErrorSmLimitFreqSchedGroup = cNsErrorSmLimitFreqSchedGroup
	NsErrorSmLimitFreqSchedGroup = &pNsErrorSmLimitFreqSchedGroup

	pNsErrorSmSwupdateEinprog = cNsErrorSmSwupdateEinprog
	NsErrorSmSwupdateEinprog = &pNsErrorSmSwupdateEinprog

	pNsErrorSmMissingCriteriaParam = cNsErrorSmMissingCriteriaParam
	NsErrorSmMissingCriteriaParam = &pNsErrorSmMissingCriteriaParam

	pNsErrorSmNoPathFound = cNsErrorSmNoPathFound
	NsErrorSmNoPathFound = &pNsErrorSmNoPathFound

	pNsErrorSmIncompatibleAppCategory = cNsErrorSmIncompatibleAppCategory
	NsErrorSmIncompatibleAppCategory = &pNsErrorSmIncompatibleAppCategory

	pNsErrorSmVolmvAbortEnomove = cNsErrorSmVolmvAbortEnomove
	NsErrorSmVolmvAbortEnomove = &pNsErrorSmVolmvAbortEnomove

	pNsErrorSmLimitSnapRetentionVolcoll = cNsErrorSmLimitSnapRetentionVolcoll
	NsErrorSmLimitSnapRetentionVolcoll = &pNsErrorSmLimitSnapRetentionVolcoll

	pNsErrorSmKeymgrRemove = cNsErrorSmKeymgrRemove
	NsErrorSmKeymgrRemove = &pNsErrorSmKeymgrRemove

	pNsErrorSmArrayNotReachable = cNsErrorSmArrayNotReachable
	NsErrorSmArrayNotReachable = &pNsErrorSmArrayNotReachable

	pNsErrorSmErrGroupMergeEventsPending = cNsErrorSmErrGroupMergeEventsPending
	NsErrorSmErrGroupMergeEventsPending = &pNsErrorSmErrGroupMergeEventsPending

	pNsErrorSmSrepNameConflictVols = cNsErrorSmSrepNameConflictVols
	NsErrorSmSrepNameConflictVols = &pNsErrorSmSrepNameConflictVols

	pNsErrorSmPoolPartnerResumeUnsup = cNsErrorSmPoolPartnerResumeUnsup
	NsErrorSmPoolPartnerResumeUnsup = &pNsErrorSmPoolPartnerResumeUnsup

	pNsErrorSmArrayNotFound = cNsErrorSmArrayNotFound
	NsErrorSmArrayNotFound = &pNsErrorSmArrayNotFound

	pNsErrorSmNoIscsiHardware = cNsErrorSmNoIscsiHardware
	NsErrorSmNoIscsiHardware = &pNsErrorSmNoIscsiHardware

	pNsErrorSmEnospc = cNsErrorSmEnospc
	NsErrorSmEnospc = &pNsErrorSmEnospc

	pNsErrorSmReservedPerfpolName = cNsErrorSmReservedPerfpolName
	NsErrorSmReservedPerfpolName = &pNsErrorSmReservedPerfpolName

	pNsErrorSmInvalidRoute = cNsErrorSmInvalidRoute
	NsErrorSmInvalidRoute = &pNsErrorSmInvalidRoute

	pNsErrorSmVolDedupeMoveInvalid = cNsErrorSmVolDedupeMoveInvalid
	NsErrorSmVolDedupeMoveInvalid = &pNsErrorSmVolDedupeMoveInvalid

	pNsErrorSmKeymgrJoin = cNsErrorSmKeymgrJoin
	NsErrorSmKeymgrJoin = &pNsErrorSmKeymgrJoin

	pNsErrorSmErrSrvUnreach = cNsErrorSmErrSrvUnreach
	NsErrorSmErrSrvUnreach = &pNsErrorSmErrSrvUnreach

	pNsErrorSmErrShelfDiskSasLinkDegraded = cNsErrorSmErrShelfDiskSasLinkDegraded
	NsErrorSmErrShelfDiskSasLinkDegraded = &pNsErrorSmErrShelfDiskSasLinkDegraded

	pNsErrorSmErrShelfInvalidEeprom = cNsErrorSmErrShelfInvalidEeprom
	NsErrorSmErrShelfInvalidEeprom = &pNsErrorSmErrShelfInvalidEeprom

	pNsErrorSmInvalidKeyValue = cNsErrorSmInvalidKeyValue
	NsErrorSmInvalidKeyValue = &pNsErrorSmInvalidKeyValue

	pNsErrorSmNoIscsiLunAssignment = cNsErrorSmNoIscsiLunAssignment
	NsErrorSmNoIscsiLunAssignment = &pNsErrorSmNoIscsiLunAssignment

	pNsErrorSmSnapshotOffline = cNsErrorSmSnapshotOffline
	NsErrorSmSnapshotOffline = &pNsErrorSmSnapshotOffline

	pNsErrorSmDefaultGatewayNotInMgmtSubnet = cNsErrorSmDefaultGatewayNotInMgmtSubnet
	NsErrorSmDefaultGatewayNotInMgmtSubnet = &pNsErrorSmDefaultGatewayNotInMgmtSubnet

	pNsErrorSmErrPassphraseAuth = cNsErrorSmErrPassphraseAuth
	NsErrorSmErrPassphraseAuth = &pNsErrorSmErrPassphraseAuth

	pNsErrorSmReplEexist = cNsErrorSmReplEexist
	NsErrorSmReplEexist = &pNsErrorSmReplEexist

	pNsErrorSmDisableLastProtocol = cNsErrorSmDisableLastProtocol
	NsErrorSmDisableLastProtocol = &pNsErrorSmDisableLastProtocol

	pNsErrorSmSecondUntaggedAssignment = cNsErrorSmSecondUntaggedAssignment
	NsErrorSmSecondUntaggedAssignment = &pNsErrorSmSecondUntaggedAssignment

	pNsErrorSmPoolHasPe = cNsErrorSmPoolHasPe
	NsErrorSmPoolHasPe = &pNsErrorSmPoolHasPe

	pNsErrorSmErrShelfInvalidCount = cNsErrorSmErrShelfInvalidCount
	NsErrorSmErrShelfInvalidCount = &pNsErrorSmErrShelfInvalidCount

	pNsErrorSmLimitSnapRetentionGroup = cNsErrorSmLimitSnapRetentionGroup
	NsErrorSmLimitSnapRetentionGroup = &pNsErrorSmLimitSnapRetentionGroup

	pNsErrorSmVolDedupeThickProvInvalid = cNsErrorSmVolDedupeThickProvInvalid
	NsErrorSmVolDedupeThickProvInvalid = &pNsErrorSmVolDedupeThickProvInvalid

	pNsErrorSmUnknown = cNsErrorSmUnknown
	NsErrorSmUnknown = &pNsErrorSmUnknown

	pNsErrorSmErrShelfDedupeImpact = cNsErrorSmErrShelfDedupeImpact
	NsErrorSmErrShelfDedupeImpact = &pNsErrorSmErrShelfDedupeImpact

	pNsErrorSmInvalidInitiatorAccessProtocol = cNsErrorSmInvalidInitiatorAccessProtocol
	NsErrorSmInvalidInitiatorAccessProtocol = &pNsErrorSmInvalidInitiatorAccessProtocol

	pNsErrorSmInternal = cNsErrorSmInternal
	NsErrorSmInternal = &pNsErrorSmInternal

	pNsErrorSmAsupEbusy = cNsErrorSmAsupEbusy
	NsErrorSmAsupEbusy = &pNsErrorSmAsupEbusy

	pNsErrorSmInvalidVolMbpsLimit = cNsErrorSmInvalidVolMbpsLimit
	NsErrorSmInvalidVolMbpsLimit = &pNsErrorSmInvalidVolMbpsLimit

	pNsErrorSmInfoConfigSyncInprogress = cNsErrorSmInfoConfigSyncInprogress
	NsErrorSmInfoConfigSyncInprogress = &pNsErrorSmInfoConfigSyncInprogress

	pNsErrorSmErrPoolStillMerging = cNsErrorSmErrPoolStillMerging
	NsErrorSmErrPoolStillMerging = &pNsErrorSmErrPoolStillMerging

	pNsErrorSmFolderPerfpolAgentType = cNsErrorSmFolderPerfpolAgentType
	NsErrorSmFolderPerfpolAgentType = &pNsErrorSmFolderPerfpolAgentType

	pNsErrorSmEinval = cNsErrorSmEinval
	NsErrorSmEinval = &pNsErrorSmEinval

	pNsErrorSmNoAssocVols = cNsErrorSmNoAssocVols
	NsErrorSmNoAssocVols = &pNsErrorSmNoAssocVols

	pNsErrorSmShelfNotInuse = cNsErrorSmShelfNotInuse
	NsErrorSmShelfNotInuse = &pNsErrorSmShelfNotInuse

	pNsErrorSmErrProtpolSettingsNotAllowed = cNsErrorSmErrProtpolSettingsNotAllowed
	NsErrorSmErrProtpolSettingsNotAllowed = &pNsErrorSmErrProtpolSettingsNotAllowed

	pNsErrorSmFcRegenerate = cNsErrorSmFcRegenerate
	NsErrorSmFcRegenerate = &pNsErrorSmFcRegenerate

	pNsErrorSmErrMultiArrayGroup = cNsErrorSmErrMultiArrayGroup
	NsErrorSmErrMultiArrayGroup = &pNsErrorSmErrMultiArrayGroup

	pNsErrorSmReplRemoteNoBaseSnap = cNsErrorSmReplRemoteNoBaseSnap
	NsErrorSmReplRemoteNoBaseSnap = &pNsErrorSmReplRemoteNoBaseSnap

	pNsErrorSmUnsupportedQueryOperator = cNsErrorSmUnsupportedQueryOperator
	NsErrorSmUnsupportedQueryOperator = &pNsErrorSmUnsupportedQueryOperator

	pNsErrorSmSrepSizeMismatchDownstreamVol = cNsErrorSmSrepSizeMismatchDownstreamVol
	NsErrorSmSrepSizeMismatchDownstreamVol = &pNsErrorSmSrepSizeMismatchDownstreamVol

	pNsErrorSmNoSuchPartner = cNsErrorSmNoSuchPartner
	NsErrorSmNoSuchPartner = &pNsErrorSmNoSuchPartner

	pNsErrorSmIscsiAllAccessNotAvailable = cNsErrorSmIscsiAllAccessNotAvailable
	NsErrorSmIscsiAllAccessNotAvailable = &pNsErrorSmIscsiAllAccessNotAvailable

	pNsErrorSmVvolAlreadyEnabled = cNsErrorSmVvolAlreadyEnabled
	NsErrorSmVvolAlreadyEnabled = &pNsErrorSmVvolAlreadyEnabled

	pNsErrorSmUsageUnavaiable = cNsErrorSmUsageUnavaiable
	NsErrorSmUsageUnavaiable = &pNsErrorSmUsageUnavaiable

	pNsErrorSmEconnrefused = cNsErrorSmEconnrefused
	NsErrorSmEconnrefused = &pNsErrorSmEconnrefused

	pNsErrorSmReplRenameNotsup = cNsErrorSmReplRenameNotsup
	NsErrorSmReplRenameNotsup = &pNsErrorSmReplRenameNotsup

	pNsErrorSmSessionCreate = cNsErrorSmSessionCreate
	NsErrorSmSessionCreate = &pNsErrorSmSessionCreate

	pNsErrorSmConflictingAclVol = cNsErrorSmConflictingAclVol
	NsErrorSmConflictingAclVol = &pNsErrorSmConflictingAclVol

	pNsErrorSmNospace = cNsErrorSmNospace
	NsErrorSmNospace = &pNsErrorSmNospace

	pNsErrorSmReservedVolcollName = cNsErrorSmReservedVolcollName
	NsErrorSmReservedVolcollName = &pNsErrorSmReservedVolcollName

	pNsErrorSmPoolHasFolder = cNsErrorSmPoolHasFolder
	NsErrorSmPoolHasFolder = &pNsErrorSmPoolHasFolder

	pNsErrorSmPartnerOffline = cNsErrorSmPartnerOffline
	NsErrorSmPartnerOffline = &pNsErrorSmPartnerOffline

	pNsErrorSmConflictingInitiatorAliasWithArgs = cNsErrorSmConflictingInitiatorAliasWithArgs
	NsErrorSmConflictingInitiatorAliasWithArgs = &pNsErrorSmConflictingInitiatorAliasWithArgs

	pNsErrorSmErrShelfNotReady = cNsErrorSmErrShelfNotReady
	NsErrorSmErrShelfNotReady = &pNsErrorSmErrShelfNotReady

	pNsErrorSmReplCmprVersionUnsup = cNsErrorSmReplCmprVersionUnsup
	NsErrorSmReplCmprVersionUnsup = &pNsErrorSmReplCmprVersionUnsup

	pNsErrorSmErrShelfPreactivationIoErr = cNsErrorSmErrShelfPreactivationIoErr
	NsErrorSmErrShelfPreactivationIoErr = &pNsErrorSmErrShelfPreactivationIoErr

	pNsErrorSmVolmvEinprog = cNsErrorSmVolmvEinprog
	NsErrorSmVolmvEinprog = &pNsErrorSmVolmvEinprog

	pNsErrorSmPerfpolIncompatibleAppCategory = cNsErrorSmPerfpolIncompatibleAppCategory
	NsErrorSmPerfpolIncompatibleAppCategory = &pNsErrorSmPerfpolIncompatibleAppCategory

	pNsErrorSmInvalidArgValue = cNsErrorSmInvalidArgValue
	NsErrorSmInvalidArgValue = &pNsErrorSmInvalidArgValue

	pNsErrorSmErrShelfDedupeBelowFdr = cNsErrorSmErrShelfDedupeBelowFdr
	NsErrorSmErrShelfDedupeBelowFdr = &pNsErrorSmErrShelfDedupeBelowFdr

	pNsErrorSmErrShelfInvalidAfsCount = cNsErrorSmErrShelfInvalidAfsCount
	NsErrorSmErrShelfInvalidAfsCount = &pNsErrorSmErrShelfInvalidAfsCount

	pNsErrorSmSerialNotAvail = cNsErrorSmSerialNotAvail
	NsErrorSmSerialNotAvail = &pNsErrorSmSerialNotAvail

	pNsErrorSmErrShelfSesMshipErr = cNsErrorSmErrShelfSesMshipErr
	NsErrorSmErrShelfSesMshipErr = &pNsErrorSmErrShelfSesMshipErr

	pNsErrorSmFolderProvisionedLimitBelowCurrentUsage = cNsErrorSmFolderProvisionedLimitBelowCurrentUsage
	NsErrorSmFolderProvisionedLimitBelowCurrentUsage = &pNsErrorSmFolderProvisionedLimitBelowCurrentUsage

	pNsErrorSmPeMultiProtocolAccessNotSupported = cNsErrorSmPeMultiProtocolAccessNotSupported
	NsErrorSmPeMultiProtocolAccessNotSupported = &pNsErrorSmPeMultiProtocolAccessNotSupported

	pNsErrorSmLimitHretSnapRetentionPool = cNsErrorSmLimitHretSnapRetentionPool
	NsErrorSmLimitHretSnapRetentionPool = &pNsErrorSmLimitHretSnapRetentionPool

	pNsErrorSmPeConflictingAcl = cNsErrorSmPeConflictingAcl
	NsErrorSmPeConflictingAcl = &pNsErrorSmPeConflictingAcl

	pNsErrorSmFolderOverUsageLimit = cNsErrorSmFolderOverUsageLimit
	NsErrorSmFolderOverUsageLimit = &pNsErrorSmFolderOverUsageLimit

	pNsErrorSmErrEncryptionMasterKeyMissing = cNsErrorSmErrEncryptionMasterKeyMissing
	NsErrorSmErrEncryptionMasterKeyMissing = &pNsErrorSmErrEncryptionMasterKeyMissing

	pNsErrorSmProtpolInvalidDomainName = cNsErrorSmProtpolInvalidDomainName
	NsErrorSmProtpolInvalidDomainName = &pNsErrorSmProtpolInvalidDomainName

	pNsErrorSmTooMany = cNsErrorSmTooMany
	NsErrorSmTooMany = &pNsErrorSmTooMany

	pNsErrorSmErrShelfExprRevIncompatible = cNsErrorSmErrShelfExprRevIncompatible
	NsErrorSmErrShelfExprRevIncompatible = &pNsErrorSmErrShelfExprRevIncompatible

	pNsErrorSmSrvUnreachable = cNsErrorSmSrvUnreachable
	NsErrorSmSrvUnreachable = &pNsErrorSmSrvUnreachable

	pNsErrorSmVolumeConflict = cNsErrorSmVolumeConflict
	NsErrorSmVolumeConflict = &pNsErrorSmVolumeConflict

	pNsErrorSmInvalidArrayName = cNsErrorSmInvalidArrayName
	NsErrorSmInvalidArrayName = &pNsErrorSmInvalidArrayName

	pNsErrorSmCannotReadObject = cNsErrorSmCannotReadObject
	NsErrorSmCannotReadObject = &pNsErrorSmCannotReadObject

	pNsErrorSmReservedVolName = cNsErrorSmReservedVolName
	NsErrorSmReservedVolName = &pNsErrorSmReservedVolName

	pNsErrorSmPoolPartnerInUse = cNsErrorSmPoolPartnerInUse
	NsErrorSmPoolPartnerInUse = &pNsErrorSmPoolPartnerInUse

	pNsErrorSmInvalidInitiatorgrpName = cNsErrorSmInvalidInitiatorgrpName
	NsErrorSmInvalidInitiatorgrpName = &pNsErrorSmInvalidInitiatorgrpName

	pNsErrorSmAsupValidateError = cNsErrorSmAsupValidateError
	NsErrorSmAsupValidateError = &pNsErrorSmAsupValidateError

	pNsErrorSmVersionName = cNsErrorSmVersionName
	NsErrorSmVersionName = &pNsErrorSmVersionName

	pNsErrorSmVvolAlreadyDisabled = cNsErrorSmVvolAlreadyDisabled
	NsErrorSmVvolAlreadyDisabled = &pNsErrorSmVvolAlreadyDisabled

	pNsErrorSmUnexpectedArg = cNsErrorSmUnexpectedArg
	NsErrorSmUnexpectedArg = &pNsErrorSmUnexpectedArg

	pNsErrorSmErrShelfInvalidAfs = cNsErrorSmErrShelfInvalidAfs
	NsErrorSmErrShelfInvalidAfs = &pNsErrorSmErrShelfInvalidAfs

	pNsErrorSmUnexpectedChild = cNsErrorSmUnexpectedChild
	NsErrorSmUnexpectedChild = &pNsErrorSmUnexpectedChild

	pNsErrorSmFolderOverProvisionedLimit = cNsErrorSmFolderOverProvisionedLimit
	NsErrorSmFolderOverProvisionedLimit = &pNsErrorSmFolderOverProvisionedLimit

	pNsErrorSmPoolFlashMismatch = cNsErrorSmPoolFlashMismatch
	NsErrorSmPoolFlashMismatch = &pNsErrorSmPoolFlashMismatch

	pNsErrorSmLimitScope = cNsErrorSmLimitScope
	NsErrorSmLimitScope = &pNsErrorSmLimitScope

	pNsErrorSmSrepSizeMismatchDownstreamVols = cNsErrorSmSrepSizeMismatchDownstreamVols
	NsErrorSmSrepSizeMismatchDownstreamVols = &pNsErrorSmSrepSizeMismatchDownstreamVols

	pNsErrorSmFolderAppsrvrInconsistent = cNsErrorSmFolderAppsrvrInconsistent
	NsErrorSmFolderAppsrvrInconsistent = &pNsErrorSmFolderAppsrvrInconsistent

	pNsErrorSmErrShelfDedupeReduction = cNsErrorSmErrShelfDedupeReduction
	NsErrorSmErrShelfDedupeReduction = &pNsErrorSmErrShelfDedupeReduction

	pNsErrorSmIscsiSvcNotAvailable = cNsErrorSmIscsiSvcNotAvailable
	NsErrorSmIscsiSvcNotAvailable = &pNsErrorSmIscsiSvcNotAvailable

	pNsErrorSmInvalidNetconfigToDelete = cNsErrorSmInvalidNetconfigToDelete
	NsErrorSmInvalidNetconfigToDelete = &pNsErrorSmInvalidNetconfigToDelete

	pNsErrorSmErrUnknown = cNsErrorSmErrUnknown
	NsErrorSmErrUnknown = &pNsErrorSmErrUnknown

	pNsErrorSmMissingCriteriaOperator = cNsErrorSmMissingCriteriaOperator
	NsErrorSmMissingCriteriaOperator = &pNsErrorSmMissingCriteriaOperator

	pNsErrorSmInvalidAppUuid = cNsErrorSmInvalidAppUuid
	NsErrorSmInvalidAppUuid = &pNsErrorSmInvalidAppUuid

	pNsErrorSmArrayMemberOrphanWithArgs = cNsErrorSmArrayMemberOrphanWithArgs
	NsErrorSmArrayMemberOrphanWithArgs = &pNsErrorSmArrayMemberOrphanWithArgs

	pNsErrorSmEmptyVol = cNsErrorSmEmptyVol
	NsErrorSmEmptyVol = &pNsErrorSmEmptyVol

	pNsErrorSmDuplicateSubnetLabel = cNsErrorSmDuplicateSubnetLabel
	NsErrorSmDuplicateSubnetLabel = &pNsErrorSmDuplicateSubnetLabel

	pNsErrorSmZeroVlanIdForTaggedAssignment = cNsErrorSmZeroVlanIdForTaggedAssignment
	NsErrorSmZeroVlanIdForTaggedAssignment = &pNsErrorSmZeroVlanIdForTaggedAssignment

	pNsErrorSmNoActionFound = cNsErrorSmNoActionFound
	NsErrorSmNoActionFound = &pNsErrorSmNoActionFound

	pNsErrorSmSyncReplUnconfigureInProgress = cNsErrorSmSyncReplUnconfigureInProgress
	NsErrorSmSyncReplUnconfigureInProgress = &pNsErrorSmSyncReplUnconfigureInProgress

	pNsErrorSmVolThickProvMoveInvalid = cNsErrorSmVolThickProvMoveInvalid
	NsErrorSmVolThickProvMoveInvalid = &pNsErrorSmVolThickProvMoveInvalid

	pNsErrorSmMissingAdvancedCriteriaConstructor = cNsErrorSmMissingAdvancedCriteriaConstructor
	NsErrorSmMissingAdvancedCriteriaConstructor = &pNsErrorSmMissingAdvancedCriteriaConstructor

	pNsErrorSmPoolUnknown = cNsErrorSmPoolUnknown
	NsErrorSmPoolUnknown = &pNsErrorSmPoolUnknown

	pNsErrorSmHaltGlWithLiveMemberArray = cNsErrorSmHaltGlWithLiveMemberArray
	NsErrorSmHaltGlWithLiveMemberArray = &pNsErrorSmHaltGlWithLiveMemberArray

	pNsErrorSmReplThrottleOverlap = cNsErrorSmReplThrottleOverlap
	NsErrorSmReplThrottleOverlap = &pNsErrorSmReplThrottleOverlap

	pNsErrorSmNetconfigUpdateMismatch = cNsErrorSmNetconfigUpdateMismatch
	NsErrorSmNetconfigUpdateMismatch = &pNsErrorSmNetconfigUpdateMismatch

	pNsErrorSmSrepDownstreamVolsAcl = cNsErrorSmSrepDownstreamVolsAcl
	NsErrorSmSrepDownstreamVolsAcl = &pNsErrorSmSrepDownstreamVolsAcl

	pNsErrorSmStatsFrequencyInvalid = cNsErrorSmStatsFrequencyInvalid
	NsErrorSmStatsFrequencyInvalid = &pNsErrorSmStatsFrequencyInvalid

	pNsErrorSmDdupFolderMerge = cNsErrorSmDdupFolderMerge
	NsErrorSmDdupFolderMerge = &pNsErrorSmDdupFolderMerge

	pNsErrorSmInvalidSubnet = cNsErrorSmInvalidSubnet
	NsErrorSmInvalidSubnet = &pNsErrorSmInvalidSubnet

	pNsErrorSmReplEnabled = cNsErrorSmReplEnabled
	NsErrorSmReplEnabled = &pNsErrorSmReplEnabled

	pNsErrorSmPoolPartnerNameConflict = cNsErrorSmPoolPartnerNameConflict
	NsErrorSmPoolPartnerNameConflict = &pNsErrorSmPoolPartnerNameConflict

	pNsErrorSmAsupPingfromMgmtipError = cNsErrorSmAsupPingfromMgmtipError
	NsErrorSmAsupPingfromMgmtipError = &pNsErrorSmAsupPingfromMgmtipError

	pNsErrorSmEncryptionGroupScopeOverride = cNsErrorSmEncryptionGroupScopeOverride
	NsErrorSmEncryptionGroupScopeOverride = &pNsErrorSmEncryptionGroupScopeOverride

	pNsErrorSmErrMaxSessionsReached = cNsErrorSmErrMaxSessionsReached
	NsErrorSmErrMaxSessionsReached = &pNsErrorSmErrMaxSessionsReached

	pNsErrorSmErrReplMultiplePartners = cNsErrorSmErrReplMultiplePartners
	NsErrorSmErrReplMultiplePartners = &pNsErrorSmErrReplMultiplePartners

	pNsErrorSmStatsNoSensors = cNsErrorSmStatsNoSensors
	NsErrorSmStatsNoSensors = &pNsErrorSmStatsNoSensors

	pNsErrorSmFolderNeedsLimit = cNsErrorSmFolderNeedsLimit
	NsErrorSmFolderNeedsLimit = &pNsErrorSmFolderNeedsLimit

	pNsErrorSmVolmvVvol = cNsErrorSmVolmvVvol
	NsErrorSmVolmvVvol = &pNsErrorSmVolmvVvol

	pNsErrorSmErrGroupMergeInprogress = cNsErrorSmErrGroupMergeInprogress
	NsErrorSmErrGroupMergeInprogress = &pNsErrorSmErrGroupMergeInprogress

	pNsErrorSmFolderNotFound = cNsErrorSmFolderNotFound
	NsErrorSmFolderNotFound = &pNsErrorSmFolderNotFound

	pNsErrorSmRouteExists = cNsErrorSmRouteExists
	NsErrorSmRouteExists = &pNsErrorSmRouteExists

	pNsErrorSmInvalidDiscoveryIp = cNsErrorSmInvalidDiscoveryIp
	NsErrorSmInvalidDiscoveryIp = &pNsErrorSmInvalidDiscoveryIp

	pNsErrorSmErrVolmvCachePinDedupeNotsupp = cNsErrorSmErrVolmvCachePinDedupeNotsupp
	NsErrorSmErrVolmvCachePinDedupeNotsupp = &pNsErrorSmErrVolmvCachePinDedupeNotsupp

	pNsErrorSmMissingCriteria = cNsErrorSmMissingCriteria
	NsErrorSmMissingCriteria = &pNsErrorSmMissingCriteria

	pNsErrorSmRootBranchPinned = cNsErrorSmRootBranchPinned
	NsErrorSmRootBranchPinned = &pNsErrorSmRootBranchPinned

	pNsErrorSmInvalidMtu = cNsErrorSmInvalidMtu
	NsErrorSmInvalidMtu = &pNsErrorSmInvalidMtu

	pNsErrorSmNoFcHardware = cNsErrorSmNoFcHardware
	NsErrorSmNoFcHardware = &pNsErrorSmNoFcHardware

	pNsErrorSmPoolNotFound = cNsErrorSmPoolNotFound
	NsErrorSmPoolNotFound = &pNsErrorSmPoolNotFound

	pNsErrorSmDuplicateIp = cNsErrorSmDuplicateIp
	NsErrorSmDuplicateIp = &pNsErrorSmDuplicateIp

	pNsErrorSmNoAction = cNsErrorSmNoAction
	NsErrorSmNoAction = &pNsErrorSmNoAction

	pNsErrorSmProtpolAppSyncOracleParams = cNsErrorSmProtpolAppSyncOracleParams
	NsErrorSmProtpolAppSyncOracleParams = &pNsErrorSmProtpolAppSyncOracleParams

	pNsErrorSmProtpolInvalidVssSettings = cNsErrorSmProtpolInvalidVssSettings
	NsErrorSmProtpolInvalidVssSettings = &pNsErrorSmProtpolInvalidVssSettings

	pNsErrorSmEncryptionInvalidScope = cNsErrorSmEncryptionInvalidScope
	NsErrorSmEncryptionInvalidScope = &pNsErrorSmEncryptionInvalidScope

	pNsErrorSmLimitSnapcollVolcoll = cNsErrorSmLimitSnapcollVolcoll
	NsErrorSmLimitSnapcollVolcoll = &pNsErrorSmLimitSnapcollVolcoll

	pNsErrorSmShelfRaidDegraded = cNsErrorSmShelfRaidDegraded
	NsErrorSmShelfRaidDegraded = &pNsErrorSmShelfRaidDegraded

	pNsErrorSmInvalidNonIscsiDataSubnetType = cNsErrorSmInvalidNonIscsiDataSubnetType
	NsErrorSmInvalidNonIscsiDataSubnetType = &pNsErrorSmInvalidNonIscsiDataSubnetType

	pNsErrorSmVolmvIncompatibleAgentType = cNsErrorSmVolmvIncompatibleAgentType
	NsErrorSmVolmvIncompatibleAgentType = &pNsErrorSmVolmvIncompatibleAgentType

	pNsErrorSmReplPartnerVersionUnknown = cNsErrorSmReplPartnerVersionUnknown
	NsErrorSmReplPartnerVersionUnknown = &pNsErrorSmReplPartnerVersionUnknown

	pNsErrorSmAsupNameresError = cNsErrorSmAsupNameresError
	NsErrorSmAsupNameresError = &pNsErrorSmAsupNameresError

	pNsErrorSmReplApiUnsup = cNsErrorSmReplApiUnsup
	NsErrorSmReplApiUnsup = &pNsErrorSmReplApiUnsup

	pNsErrorSmEperm = cNsErrorSmEperm
	NsErrorSmEperm = &pNsErrorSmEperm

	pNsErrorSmEnoentEnoent = cNsErrorSmEnoentEnoent
	NsErrorSmEnoentEnoent = &pNsErrorSmEnoentEnoent

	pNsErrorSmArrayGroupLeader = cNsErrorSmArrayGroupLeader
	NsErrorSmArrayGroupLeader = &pNsErrorSmArrayGroupLeader

	pNsErrorSmInvalidInitiatorIqn = cNsErrorSmInvalidInitiatorIqn
	NsErrorSmInvalidInitiatorIqn = &pNsErrorSmInvalidInitiatorIqn

	pNsErrorSmReplIntragroup = cNsErrorSmReplIntragroup
	NsErrorSmReplIntragroup = &pNsErrorSmReplIntragroup

	pNsErrorSmRemoveNonemptyFolder = cNsErrorSmRemoveNonemptyFolder
	NsErrorSmRemoveNonemptyFolder = &pNsErrorSmRemoveNonemptyFolder

	pNsErrorSmAddNoniscsiToIscsiGroup = cNsErrorSmAddNoniscsiToIscsiGroup
	NsErrorSmAddNoniscsiToIscsiGroup = &pNsErrorSmAddNoniscsiToIscsiGroup

	pNsErrorSmPoolDedupeIncapable = cNsErrorSmPoolDedupeIncapable
	NsErrorSmPoolDedupeIncapable = &pNsErrorSmPoolDedupeIncapable

	pNsErrorSmNoVvolSupport = cNsErrorSmNoVvolSupport
	NsErrorSmNoVvolSupport = &pNsErrorSmNoVvolSupport

	pNsErrorSmNoInitiatorgrp = cNsErrorSmNoInitiatorgrp
	NsErrorSmNoInitiatorgrp = &pNsErrorSmNoInitiatorgrp

	pNsErrorSmInvalidNetmask = cNsErrorSmInvalidNetmask
	NsErrorSmInvalidNetmask = &pNsErrorSmInvalidNetmask

	pNsErrorSmNoDiscoveryIpInManualMode = cNsErrorSmNoDiscoveryIpInManualMode
	NsErrorSmNoDiscoveryIpInManualMode = &pNsErrorSmNoDiscoveryIpInManualMode

	pNsErrorSmInvalidVolReserveValues = cNsErrorSmInvalidVolReserveValues
	NsErrorSmInvalidVolReserveValues = &pNsErrorSmInvalidVolReserveValues

	pNsErrorSmSupportIpInvalid = cNsErrorSmSupportIpInvalid
	NsErrorSmSupportIpInvalid = &pNsErrorSmSupportIpInvalid

	pNsErrorSmErrShelfForeignDisk = cNsErrorSmErrShelfForeignDisk
	NsErrorSmErrShelfForeignDisk = &pNsErrorSmErrShelfForeignDisk

	pNsErrorSmVolsnapAlreadyExported = cNsErrorSmVolsnapAlreadyExported
	NsErrorSmVolsnapAlreadyExported = &pNsErrorSmVolsnapAlreadyExported

	pNsErrorSmFcInitiatorgrpSubnetNotSupported = cNsErrorSmFcInitiatorgrpSubnetNotSupported
	NsErrorSmFcInitiatorgrpSubnetNotSupported = &pNsErrorSmFcInitiatorgrpSubnetNotSupported

	pNsErrorSmIncompatibleInitiatorAccessProtocol = cNsErrorSmIncompatibleInitiatorAccessProtocol
	NsErrorSmIncompatibleInitiatorAccessProtocol = &pNsErrorSmIncompatibleInitiatorAccessProtocol

	pNsErrorSmInvalidInitiatorgrpAccessProtocol = cNsErrorSmInvalidInitiatorgrpAccessProtocol
	NsErrorSmInvalidInitiatorgrpAccessProtocol = &pNsErrorSmInvalidInitiatorgrpAccessProtocol

	pNsErrorSmLimitPeacl = cNsErrorSmLimitPeacl
	NsErrorSmLimitPeacl = &pNsErrorSmLimitPeacl

	pNsErrorSmEtimedout = cNsErrorSmEtimedout
	NsErrorSmEtimedout = &pNsErrorSmEtimedout

	pNsErrorSmInitiatorgrpSubnetDoesNotExist = cNsErrorSmInitiatorgrpSubnetDoesNotExist
	NsErrorSmInitiatorgrpSubnetDoesNotExist = &pNsErrorSmInitiatorgrpSubnetDoesNotExist

	pNsErrorSmVolOffline = cNsErrorSmVolOffline
	NsErrorSmVolOffline = &pNsErrorSmVolOffline

	pNsErrorSmErrSreplMgmtOpInProgress = cNsErrorSmErrSreplMgmtOpInProgress
	NsErrorSmErrSreplMgmtOpInProgress = &pNsErrorSmErrSreplMgmtOpInProgress

	pNsErrorSmReplFanoutMaximumCloudPartnersExceeded = cNsErrorSmReplFanoutMaximumCloudPartnersExceeded
	NsErrorSmReplFanoutMaximumCloudPartnersExceeded = &pNsErrorSmReplFanoutMaximumCloudPartnersExceeded

	pNsErrorSmMalformedUrl = cNsErrorSmMalformedUrl
	NsErrorSmMalformedUrl = &pNsErrorSmMalformedUrl

	pNsErrorSmErrSessionCreate = cNsErrorSmErrSessionCreate
	NsErrorSmErrSessionCreate = &pNsErrorSmErrSessionCreate

	pNsErrorSmErrShelfCtrlrLoop = cNsErrorSmErrShelfCtrlrLoop
	NsErrorSmErrShelfCtrlrLoop = &pNsErrorSmErrShelfCtrlrLoop

	pNsErrorSmPeConflictingAclLun = cNsErrorSmPeConflictingAclLun
	NsErrorSmPeConflictingAclLun = &pNsErrorSmPeConflictingAclLun

	pNsErrorSmInvalidCtrlrName = cNsErrorSmInvalidCtrlrName
	NsErrorSmInvalidCtrlrName = &pNsErrorSmInvalidCtrlrName

	pNsErrorSmBackupNetconfigReadonly = cNsErrorSmBackupNetconfigReadonly
	NsErrorSmBackupNetconfigReadonly = &pNsErrorSmBackupNetconfigReadonly

	pNsErrorSmLimitSnapGroup = cNsErrorSmLimitSnapGroup
	NsErrorSmLimitSnapGroup = &pNsErrorSmLimitSnapGroup

	pNsErrorSmEncryptionInvalidCipher = cNsErrorSmEncryptionInvalidCipher
	NsErrorSmEncryptionInvalidCipher = &pNsErrorSmEncryptionInvalidCipher

	pNsErrorSmPerfpolInvalidAppCategory = cNsErrorSmPerfpolInvalidAppCategory
	NsErrorSmPerfpolInvalidAppCategory = &pNsErrorSmPerfpolInvalidAppCategory

	pNsErrorSmArrayPoolFlashMismatch = cNsErrorSmArrayPoolFlashMismatch
	NsErrorSmArrayPoolFlashMismatch = &pNsErrorSmArrayPoolFlashMismatch

	pNsErrorSmFolderLimitInval = cNsErrorSmFolderLimitInval
	NsErrorSmFolderLimitInval = &pNsErrorSmFolderLimitInval

	pNsErrorSmShelfSsdDegraded = cNsErrorSmShelfSsdDegraded
	NsErrorSmShelfSsdDegraded = &pNsErrorSmShelfSsdDegraded

	pNsErrorSmFolderOverdraftLimitNeedsUsageLimit = cNsErrorSmFolderOverdraftLimitNeedsUsageLimit
	NsErrorSmFolderOverdraftLimitNeedsUsageLimit = &pNsErrorSmFolderOverdraftLimitNeedsUsageLimit

	pNsErrorSmEncryptionMasterKeyMissing = cNsErrorSmEncryptionMasterKeyMissing
	NsErrorSmEncryptionMasterKeyMissing = &pNsErrorSmEncryptionMasterKeyMissing

	pNsErrorSmSrepAgentTypeMismatchDownstreamVols = cNsErrorSmSrepAgentTypeMismatchDownstreamVols
	NsErrorSmSrepAgentTypeMismatchDownstreamVols = &pNsErrorSmSrepAgentTypeMismatchDownstreamVols

	pNsErrorSmErrVolNotOfflineOnRestore = cNsErrorSmErrVolNotOfflineOnRestore
	NsErrorSmErrVolNotOfflineOnRestore = &pNsErrorSmErrVolNotOfflineOnRestore

	pNsErrorSmReplHandoverBusy = cNsErrorSmReplHandoverBusy
	NsErrorSmReplHandoverBusy = &pNsErrorSmReplHandoverBusy

	pNsErrorSmNotOwner = cNsErrorSmNotOwner
	NsErrorSmNotOwner = &pNsErrorSmNotOwner

	pNsErrorSmSrepDownstreamAcl = cNsErrorSmSrepDownstreamAcl
	NsErrorSmSrepDownstreamAcl = &pNsErrorSmSrepDownstreamAcl

	pNsErrorSmNoSubnetWithLabel = cNsErrorSmNoSubnetWithLabel
	NsErrorSmNoSubnetWithLabel = &pNsErrorSmNoSubnetWithLabel

	pNsErrorSmPoolUpdateInvalArrays = cNsErrorSmPoolUpdateInvalArrays
	NsErrorSmPoolUpdateInvalArrays = &pNsErrorSmPoolUpdateInvalArrays

	pNsErrorSmVolHasOnlineSnap = cNsErrorSmVolHasOnlineSnap
	NsErrorSmVolHasOnlineSnap = &pNsErrorSmVolHasOnlineSnap

	pNsErrorSmInvalidDataIp = cNsErrorSmInvalidDataIp
	NsErrorSmInvalidDataIp = &pNsErrorSmInvalidDataIp

	pNsErrorSmPoolVolmvEinprog = cNsErrorSmPoolVolmvEinprog
	NsErrorSmPoolVolmvEinprog = &pNsErrorSmPoolVolmvEinprog

	pNsErrorSmNoDataIpOnMgmtPlusData = cNsErrorSmNoDataIpOnMgmtPlusData
	NsErrorSmNoDataIpOnMgmtPlusData = &pNsErrorSmNoDataIpOnMgmtPlusData

	pNsErrorSmConflictingAclLun = cNsErrorSmConflictingAclLun
	NsErrorSmConflictingAclLun = &pNsErrorSmConflictingAclLun

	pNsErrorSmVpCreatedIgrp = cNsErrorSmVpCreatedIgrp
	NsErrorSmVpCreatedIgrp = &pNsErrorSmVpCreatedIgrp

	pNsErrorSmStarterVolAclCreate = cNsErrorSmStarterVolAclCreate
	NsErrorSmStarterVolAclCreate = &pNsErrorSmStarterVolAclCreate

	pNsErrorSmBadPkg = cNsErrorSmBadPkg
	NsErrorSmBadPkg = &pNsErrorSmBadPkg

	pNsErrorSmPasswdSameAsCurrent = cNsErrorSmPasswdSameAsCurrent
	NsErrorSmPasswdSameAsCurrent = &pNsErrorSmPasswdSameAsCurrent

	pNsErrorSmSnaplunsOutOfSync = cNsErrorSmSnaplunsOutOfSync
	NsErrorSmSnaplunsOutOfSync = &pNsErrorSmSnaplunsOutOfSync

	pNsErrorSmLimitSnapPool = cNsErrorSmLimitSnapPool
	NsErrorSmLimitSnapPool = &pNsErrorSmLimitSnapPool

	pNsErrorSmMultiProtocolAccessNotSupported = cNsErrorSmMultiProtocolAccessNotSupported
	NsErrorSmMultiProtocolAccessNotSupported = &pNsErrorSmMultiProtocolAccessNotSupported

	pNsErrorSmVolHasClone = cNsErrorSmVolHasClone
	NsErrorSmVolHasClone = &pNsErrorSmVolHasClone

	pNsErrorSmDedupeSingleArray = cNsErrorSmDedupeSingleArray
	NsErrorSmDedupeSingleArray = &pNsErrorSmDedupeSingleArray

	pNsErrorSmEncryptionMasterKeyInactive = cNsErrorSmEncryptionMasterKeyInactive
	NsErrorSmEncryptionMasterKeyInactive = &pNsErrorSmEncryptionMasterKeyInactive

	pNsErrorSmSnapHasClone = cNsErrorSmSnapHasClone
	NsErrorSmSnapHasClone = &pNsErrorSmSnapHasClone

	pNsErrorSmErrEncryptionNeeded = cNsErrorSmErrEncryptionNeeded
	NsErrorSmErrEncryptionNeeded = &pNsErrorSmErrEncryptionNeeded

	pNsErrorSmMismatchingDuplicateSubnet = cNsErrorSmMismatchingDuplicateSubnet
	NsErrorSmMismatchingDuplicateSubnet = &pNsErrorSmMismatchingDuplicateSubnet

	pNsErrorSmSrepNotGroupScopedVol = cNsErrorSmSrepNotGroupScopedVol
	NsErrorSmSrepNotGroupScopedVol = &pNsErrorSmSrepNotGroupScopedVol

	pNsErrorSmOnlyVvolFolderFolset = cNsErrorSmOnlyVvolFolderFolset
	NsErrorSmOnlyVvolFolderFolset = &pNsErrorSmOnlyVvolFolderFolset

	pNsErrorSmReplSnapshotSync = cNsErrorSmReplSnapshotSync
	NsErrorSmReplSnapshotSync = &pNsErrorSmReplSnapshotSync

	pNsErrorSmStatsNoSuchObject = cNsErrorSmStatsNoSuchObject
	NsErrorSmStatsNoSuchObject = &pNsErrorSmStatsNoSuchObject

	pNsErrorSmDefaultRouteMissing = cNsErrorSmDefaultRouteMissing
	NsErrorSmDefaultRouteMissing = &pNsErrorSmDefaultRouteMissing

	pNsErrorSmReplVasa3ApiUnsup = cNsErrorSmReplVasa3ApiUnsup
	NsErrorSmReplVasa3ApiUnsup = &pNsErrorSmReplVasa3ApiUnsup

	pNsErrorSmOverlappingSubnets = cNsErrorSmOverlappingSubnets
	NsErrorSmOverlappingSubnets = &pNsErrorSmOverlappingSubnets

	pNsErrorSmMissingArg = cNsErrorSmMissingArg
	NsErrorSmMissingArg = &pNsErrorSmMissingArg

	pNsErrorSmVolmvCachePinDedupeNotsupp = cNsErrorSmVolmvCachePinDedupeNotsupp
	NsErrorSmVolmvCachePinDedupeNotsupp = &pNsErrorSmVolmvCachePinDedupeNotsupp

	pNsErrorSmErrMergeConflict = cNsErrorSmErrMergeConflict
	NsErrorSmErrMergeConflict = &pNsErrorSmErrMergeConflict

	pNsErrorSmTooSmall = cNsErrorSmTooSmall
	NsErrorSmTooSmall = &pNsErrorSmTooSmall

	pNsErrorSmVolWarnGreaterThanLimit = cNsErrorSmVolWarnGreaterThanLimit
	NsErrorSmVolWarnGreaterThanLimit = &pNsErrorSmVolWarnGreaterThanLimit

	pNsErrorSmKeymgrUnreach = cNsErrorSmKeymgrUnreach
	NsErrorSmKeymgrUnreach = &pNsErrorSmKeymgrUnreach

	pNsErrorSmErrFcAsymmetryNotSupported = cNsErrorSmErrFcAsymmetryNotSupported
	NsErrorSmErrFcAsymmetryNotSupported = &pNsErrorSmErrFcAsymmetryNotSupported

	pNsErrorSmLunMismatch = cNsErrorSmLunMismatch
	NsErrorSmLunMismatch = &pNsErrorSmLunMismatch

	pNsErrorSmErrShelfInvalidDiskCount = cNsErrorSmErrShelfInvalidDiskCount
	NsErrorSmErrShelfInvalidDiskCount = &pNsErrorSmErrShelfInvalidDiskCount

	pNsErrorSmLimitVolacl = cNsErrorSmLimitVolacl
	NsErrorSmLimitVolacl = &pNsErrorSmLimitVolacl

	pNsErrorSmControllerNotActive = cNsErrorSmControllerNotActive
	NsErrorSmControllerNotActive = &pNsErrorSmControllerNotActive

	pNsErrorSmReplNoPartnerAvail = cNsErrorSmReplNoPartnerAvail
	NsErrorSmReplNoPartnerAvail = &pNsErrorSmReplNoPartnerAvail

	pNsErrorSmSreplMgmtOpDisallowedWhenSolo = cNsErrorSmSreplMgmtOpDisallowedWhenSolo
	NsErrorSmSreplMgmtOpDisallowedWhenSolo = &pNsErrorSmSreplMgmtOpDisallowedWhenSolo

	pNsErrorSmNoMgmtSubnetSpecified = cNsErrorSmNoMgmtSubnetSpecified
	NsErrorSmNoMgmtSubnetSpecified = &pNsErrorSmNoMgmtSubnetSpecified

	pNsErrorSmFolderConflict = cNsErrorSmFolderConflict
	NsErrorSmFolderConflict = &pNsErrorSmFolderConflict

	pNsErrorSmInvalidDataSubnet = cNsErrorSmInvalidDataSubnet
	NsErrorSmInvalidDataSubnet = &pNsErrorSmInvalidDataSubnet

	pNsErrorSmVolUsageUnavailable = cNsErrorSmVolUsageUnavailable
	NsErrorSmVolUsageUnavailable = &pNsErrorSmVolUsageUnavailable

	pNsErrorSmAclScopeOverlap = cNsErrorSmAclScopeOverlap
	NsErrorSmAclScopeOverlap = &pNsErrorSmAclScopeOverlap

	pNsErrorSmErrVvolAclGrpMerge = cNsErrorSmErrVvolAclGrpMerge
	NsErrorSmErrVvolAclGrpMerge = &pNsErrorSmErrVvolAclGrpMerge

	pNsErrorSmDisabledProtocolArtifacts = cNsErrorSmDisabledProtocolArtifacts
	NsErrorSmDisabledProtocolArtifacts = &pNsErrorSmDisabledProtocolArtifacts

	pNsErrorSmDuplicateSubnetVlanId = cNsErrorSmDuplicateSubnetVlanId
	NsErrorSmDuplicateSubnetVlanId = &pNsErrorSmDuplicateSubnetVlanId

	pNsErrorSmReplOpenstackUnsup = cNsErrorSmReplOpenstackUnsup
	NsErrorSmReplOpenstackUnsup = &pNsErrorSmReplOpenstackUnsup

	pNsErrorSmNoAvailableLun = cNsErrorSmNoAvailableLun
	NsErrorSmNoAvailableLun = &pNsErrorSmNoAvailableLun

	pNsErrorSmGroupPartnerNameConflict = cNsErrorSmGroupPartnerNameConflict
	NsErrorSmGroupPartnerNameConflict = &pNsErrorSmGroupPartnerNameConflict

	pNsErrorSmMgmtIpNotOnMgmt = cNsErrorSmMgmtIpNotOnMgmt
	NsErrorSmMgmtIpNotOnMgmt = &pNsErrorSmMgmtIpNotOnMgmt

	pNsErrorSmInUseLun = cNsErrorSmInUseLun
	NsErrorSmInUseLun = &pNsErrorSmInUseLun

	pNsErrorSmNotFcInitiatorgrp = cNsErrorSmNotFcInitiatorgrp
	NsErrorSmNotFcInitiatorgrp = &pNsErrorSmNotFcInitiatorgrp

	pNsErrorSmErrShelfInvalidCsz = cNsErrorSmErrShelfInvalidCsz
	NsErrorSmErrShelfInvalidCsz = &pNsErrorSmErrShelfInvalidCsz

	pNsErrorSmVolDedupeInvalidPerfPolicy = cNsErrorSmVolDedupeInvalidPerfPolicy
	NsErrorSmVolDedupeInvalidPerfPolicy = &pNsErrorSmVolDedupeInvalidPerfPolicy

	pNsErrorSmEncryptionInvalidMode = cNsErrorSmEncryptionInvalidMode
	NsErrorSmEncryptionInvalidMode = &pNsErrorSmEncryptionInvalidMode

	pNsErrorSmAsupDisabled = cNsErrorSmAsupDisabled
	NsErrorSmAsupDisabled = &pNsErrorSmAsupDisabled

	pNsErrorSmErrShelfConnectedOnlyOneSide = cNsErrorSmErrShelfConnectedOnlyOneSide
	NsErrorSmErrShelfConnectedOnlyOneSide = &pNsErrorSmErrShelfConnectedOnlyOneSide

	pNsErrorSmVolmvEnospace = cNsErrorSmVolmvEnospace
	NsErrorSmVolmvEnospace = &pNsErrorSmVolmvEnospace

	pNsErrorSmIncompatibleVolumes = cNsErrorSmIncompatibleVolumes
	NsErrorSmIncompatibleVolumes = &pNsErrorSmIncompatibleVolumes

	pNsErrorSmComplexTypeQueryParam = cNsErrorSmComplexTypeQueryParam
	NsErrorSmComplexTypeQueryParam = &pNsErrorSmComplexTypeQueryParam

	pNsErrorSmAsupError = cNsErrorSmAsupError
	NsErrorSmAsupError = &pNsErrorSmAsupError

	pNsErrorSmReplUnassignedAppcat = cNsErrorSmReplUnassignedAppcat
	NsErrorSmReplUnassignedAppcat = &pNsErrorSmReplUnassignedAppcat

	pNsErrorSmErrFcRegenerateInvalidOperationTdz = cNsErrorSmErrFcRegenerateInvalidOperationTdz
	NsErrorSmErrFcRegenerateInvalidOperationTdz = &pNsErrorSmErrFcRegenerateInvalidOperationTdz

	pNsErrorSmInvalidPathVariable = cNsErrorSmInvalidPathVariable
	NsErrorSmInvalidPathVariable = &pNsErrorSmInvalidPathVariable

	pNsErrorSmErrArgChangeNotAllowed = cNsErrorSmErrArgChangeNotAllowed
	NsErrorSmErrArgChangeNotAllowed = &pNsErrorSmErrArgChangeNotAllowed

	pNsErrorSmSpecifiedSnapshotLun = cNsErrorSmSpecifiedSnapshotLun
	NsErrorSmSpecifiedSnapshotLun = &pNsErrorSmSpecifiedSnapshotLun

	pNsErrorSmSrepDownstreamAssocedVol = cNsErrorSmSrepDownstreamAssocedVol
	NsErrorSmSrepDownstreamAssocedVol = &pNsErrorSmSrepDownstreamAssocedVol

	pNsErrorSmLimitScopeValues = cNsErrorSmLimitScopeValues
	NsErrorSmLimitScopeValues = &pNsErrorSmLimitScopeValues

	pNsErrorSmErrPassphraseInval = cNsErrorSmErrPassphraseInval
	NsErrorSmErrPassphraseInval = &pNsErrorSmErrPassphraseInval

	pNsErrorSmInitiatorgroupsOutOfSync = cNsErrorSmInitiatorgroupsOutOfSync
	NsErrorSmInitiatorgroupsOutOfSync = &pNsErrorSmInitiatorgroupsOutOfSync

	pNsErrorSmCachePinnedNotsup = cNsErrorSmCachePinnedNotsup
	NsErrorSmCachePinnedNotsup = &pNsErrorSmCachePinnedNotsup

	pNsErrorSmNoOperationFound = cNsErrorSmNoOperationFound
	NsErrorSmNoOperationFound = &pNsErrorSmNoOperationFound

	pNsErrorSmInvalidDataForPartnerType = cNsErrorSmInvalidDataForPartnerType
	NsErrorSmInvalidDataForPartnerType = &pNsErrorSmInvalidDataForPartnerType

	pNsErrorSmSrvUpdatePrecheck = cNsErrorSmSrvUpdatePrecheck
	NsErrorSmSrvUpdatePrecheck = &pNsErrorSmSrvUpdatePrecheck

	pNsErrorSmVolDedupeEncryptionInvalid = cNsErrorSmVolDedupeEncryptionInvalid
	NsErrorSmVolDedupeEncryptionInvalid = &pNsErrorSmVolDedupeEncryptionInvalid

	pNsErrorSmVvolFolderNoAppsrvr = cNsErrorSmVvolFolderNoAppsrvr
	NsErrorSmVvolFolderNoAppsrvr = &pNsErrorSmVvolFolderNoAppsrvr

	pNsErrorSmHttpConflict = cNsErrorSmHttpConflict
	NsErrorSmHttpConflict = &pNsErrorSmHttpConflict

	pNsErrorSmArrayRenameInNetconfigFailed = cNsErrorSmArrayRenameInNetconfigFailed
	NsErrorSmArrayRenameInNetconfigFailed = &pNsErrorSmArrayRenameInNetconfigFailed

	pNsErrorSmInvalidInitiatorIp = cNsErrorSmInvalidInitiatorIp
	NsErrorSmInvalidInitiatorIp = &pNsErrorSmInvalidInitiatorIp

	pNsErrorSmDuplicateVol = cNsErrorSmDuplicateVol
	NsErrorSmDuplicateVol = &pNsErrorSmDuplicateVol

	pNsErrorSmErrShelfInvalidModel = cNsErrorSmErrShelfInvalidModel
	NsErrorSmErrShelfInvalidModel = &pNsErrorSmErrShelfInvalidModel

	pNsErrorSmReplProtectLastSnap = cNsErrorSmReplProtectLastSnap
	NsErrorSmReplProtectLastSnap = &pNsErrorSmReplProtectLastSnap

	pNsErrorSmErrGroupMergeBusy = cNsErrorSmErrGroupMergeBusy
	NsErrorSmErrGroupMergeBusy = &pNsErrorSmErrGroupMergeBusy

	pNsErrorSmErrVolmvPoolMerging = cNsErrorSmErrVolmvPoolMerging
	NsErrorSmErrVolmvPoolMerging = &pNsErrorSmErrVolmvPoolMerging

	pNsErrorSmArrayNotFoundWithArgs = cNsErrorSmArrayNotFoundWithArgs
	NsErrorSmArrayNotFoundWithArgs = &pNsErrorSmArrayNotFoundWithArgs

	pNsErrorSmConnectionRebalancingWithoutAutomaticMethod = cNsErrorSmConnectionRebalancingWithoutAutomaticMethod
	NsErrorSmConnectionRebalancingWithoutAutomaticMethod = &pNsErrorSmConnectionRebalancingWithoutAutomaticMethod

	pNsErrorSmSrepDownstreamNoCommonSnapVols = cNsErrorSmSrepDownstreamNoCommonSnapVols
	NsErrorSmSrepDownstreamNoCommonSnapVols = &pNsErrorSmSrepDownstreamNoCommonSnapVols

	pNsErrorSmErrShelfWrongSasPort = cNsErrorSmErrShelfWrongSasPort
	NsErrorSmErrShelfWrongSasPort = &pNsErrorSmErrShelfWrongSasPort

	pNsErrorSmFolderUsageLimitBelowCurrentUsage = cNsErrorSmFolderUsageLimitBelowCurrentUsage
	NsErrorSmFolderUsageLimitBelowCurrentUsage = &pNsErrorSmFolderUsageLimitBelowCurrentUsage

	pNsErrorSmLimitSnapRetentionPool = cNsErrorSmLimitSnapRetentionPool
	NsErrorSmLimitSnapRetentionPool = &pNsErrorSmLimitSnapRetentionPool

	pNsErrorSmErrShelfExprFwVerInval = cNsErrorSmErrShelfExprFwVerInval
	NsErrorSmErrShelfExprFwVerInval = &pNsErrorSmErrShelfExprFwVerInval

	pNsErrorSmSrepVolcollUnsup = cNsErrorSmSrepVolcollUnsup
	NsErrorSmSrepVolcollUnsup = &pNsErrorSmSrepVolcollUnsup

	pNsErrorSmSrepNotGroupScopedVols = cNsErrorSmSrepNotGroupScopedVols
	NsErrorSmSrepNotGroupScopedVols = &pNsErrorSmSrepNotGroupScopedVols

	pNsErrorSmFcIntfNotFound = cNsErrorSmFcIntfNotFound
	NsErrorSmFcIntfNotFound = &pNsErrorSmFcIntfNotFound

	pNsErrorSmVolUnknown = cNsErrorSmVolUnknown
	NsErrorSmVolUnknown = &pNsErrorSmVolUnknown

	pNsErrorSmErrShelfEbusy = cNsErrorSmErrShelfEbusy
	NsErrorSmErrShelfEbusy = &pNsErrorSmErrShelfEbusy

	pNsErrorSmAppserverInUse = cNsErrorSmAppserverInUse
	NsErrorSmAppserverInUse = &pNsErrorSmAppserverInUse

	pNsErrorSmPoolDoesNotHaveFolder = cNsErrorSmPoolDoesNotHaveFolder
	NsErrorSmPoolDoesNotHaveFolder = &pNsErrorSmPoolDoesNotHaveFolder

	pNsErrorSmPerfpolNotFound = cNsErrorSmPerfpolNotFound
	NsErrorSmPerfpolNotFound = &pNsErrorSmPerfpolNotFound

	pNsErrorSmPerfpolOob = cNsErrorSmPerfpolOob
	NsErrorSmPerfpolOob = &pNsErrorSmPerfpolOob

	pNsErrorSmDuplicateInitiator = cNsErrorSmDuplicateInitiator
	NsErrorSmDuplicateInitiator = &pNsErrorSmDuplicateInitiator

	pNsErrorSmInUseAppUuid = cNsErrorSmInUseAppUuid
	NsErrorSmInUseAppUuid = &pNsErrorSmInUseAppUuid

	pNsErrorSmPartnerCfgSync = cNsErrorSmPartnerCfgSync
	NsErrorSmPartnerCfgSync = &pNsErrorSmPartnerCfgSync

	pNsErrorSmNotDownloadingSw = cNsErrorSmNotDownloadingSw
	NsErrorSmNotDownloadingSw = &pNsErrorSmNotDownloadingSw

	pNsErrorSmMissingMgmtIp = cNsErrorSmMissingMgmtIp
	NsErrorSmMissingMgmtIp = &pNsErrorSmMissingMgmtIp

	pNsErrorSmSrepDownstreamNoCommonSnapVol = cNsErrorSmSrepDownstreamNoCommonSnapVol
	NsErrorSmSrepDownstreamNoCommonSnapVol = &pNsErrorSmSrepDownstreamNoCommonSnapVol

	pNsErrorSmEbusy = cNsErrorSmEbusy
	NsErrorSmEbusy = &pNsErrorSmEbusy

	pNsErrorSmErrReplCantConnect = cNsErrorSmErrReplCantConnect
	NsErrorSmErrReplCantConnect = &pNsErrorSmErrReplCantConnect

	pNsErrorSmErrShelfExpanderErr = cNsErrorSmErrShelfExpanderErr
	NsErrorSmErrShelfExpanderErr = &pNsErrorSmErrShelfExpanderErr

	pNsErrorSmPeFailAclRemoval = cNsErrorSmPeFailAclRemoval
	NsErrorSmPeFailAclRemoval = &pNsErrorSmPeFailAclRemoval

	pNsErrorSmSupportIpNotOnMgmt = cNsErrorSmSupportIpNotOnMgmt
	NsErrorSmSupportIpNotOnMgmt = &pNsErrorSmSupportIpNotOnMgmt

	pNsErrorSmAtLeastOneGroupSubnet = cNsErrorSmAtLeastOneGroupSubnet
	NsErrorSmAtLeastOneGroupSubnet = &pNsErrorSmAtLeastOneGroupSubnet

	pNsErrorSmUnsupportedAccessProtocol = cNsErrorSmUnsupportedAccessProtocol
	NsErrorSmUnsupportedAccessProtocol = &pNsErrorSmUnsupportedAccessProtocol

	pNsErrorSmSpaceInfoUnavail = cNsErrorSmSpaceInfoUnavail
	NsErrorSmSpaceInfoUnavail = &pNsErrorSmSpaceInfoUnavail

	pNsErrorSmLimitVolmvHretSnapPool = cNsErrorSmLimitVolmvHretSnapPool
	NsErrorSmLimitVolmvHretSnapPool = &pNsErrorSmLimitVolmvHretSnapPool

	pNsErrorSmStartTimeBeyondEndTime = cNsErrorSmStartTimeBeyondEndTime
	NsErrorSmStartTimeBeyondEndTime = &pNsErrorSmStartTimeBeyondEndTime

	pNsErrorSmReplRemotePaused = cNsErrorSmReplRemotePaused
	NsErrorSmReplRemotePaused = &pNsErrorSmReplRemotePaused

	pNsErrorSmDataIpNotOnSubnet = cNsErrorSmDataIpNotOnSubnet
	NsErrorSmDataIpNotOnSubnet = &pNsErrorSmDataIpNotOnSubnet

	pNsErrorSmConflictingInitiatorAlias = cNsErrorSmConflictingInitiatorAlias
	NsErrorSmConflictingInitiatorAlias = &pNsErrorSmConflictingInitiatorAlias

	pNsErrorSmReplEditPartnerNameNotPaused = cNsErrorSmReplEditPartnerNameNotPaused
	NsErrorSmReplEditPartnerNameNotPaused = &pNsErrorSmReplEditPartnerNameNotPaused

	pNsErrorSmShelfForeignDisk = cNsErrorSmShelfForeignDisk
	NsErrorSmShelfForeignDisk = &pNsErrorSmShelfForeignDisk

	pNsErrorSmQosLimitNotInRange = cNsErrorSmQosLimitNotInRange
	NsErrorSmQosLimitNotInRange = &pNsErrorSmQosLimitNotInRange

	pNsErrorSmErrShelfNoRserial = cNsErrorSmErrShelfNoRserial
	NsErrorSmErrShelfNoRserial = &pNsErrorSmErrShelfNoRserial

	pNsErrorSmUntaggedMtuNotLargest = cNsErrorSmUntaggedMtuNotLargest
	NsErrorSmUntaggedMtuNotLargest = &pNsErrorSmUntaggedMtuNotLargest

	pNsErrorSmErrShelfDisconnected = cNsErrorSmErrShelfDisconnected
	NsErrorSmErrShelfDisconnected = &pNsErrorSmErrShelfDisconnected

	pNsErrorSmSubnetAlreadyAssignedOnNic = cNsErrorSmSubnetAlreadyAssignedOnNic
	NsErrorSmSubnetAlreadyAssignedOnNic = &pNsErrorSmSubnetAlreadyAssignedOnNic

	pNsErrorSmFcSessionExist = cNsErrorSmFcSessionExist
	NsErrorSmFcSessionExist = &pNsErrorSmFcSessionExist

	pNsErrorSmVvolCannotOfflineBoundSnap = cNsErrorSmVvolCannotOfflineBoundSnap
	NsErrorSmVvolCannotOfflineBoundSnap = &pNsErrorSmVvolCannotOfflineBoundSnap

	pNsErrorSmArrayMemberOrphan = cNsErrorSmArrayMemberOrphan
	NsErrorSmArrayMemberOrphan = &pNsErrorSmArrayMemberOrphan

	pNsErrorSmArrayMissingSubnet = cNsErrorSmArrayMissingSubnet
	NsErrorSmArrayMissingSubnet = &pNsErrorSmArrayMissingSubnet

	pNsErrorSmDisableVvolWithPe = cNsErrorSmDisableVvolWithPe
	NsErrorSmDisableVvolWithPe = &pNsErrorSmDisableVvolWithPe

	pNsErrorSmLimitSnapVol = cNsErrorSmLimitSnapVol
	NsErrorSmLimitSnapVol = &pNsErrorSmLimitSnapVol

	pNsErrorSmPoolUsageUnavailable = cNsErrorSmPoolUsageUnavailable
	NsErrorSmPoolUsageUnavailable = &pNsErrorSmPoolUsageUnavailable

	pNsErrorSmInvalidQueryParam = cNsErrorSmInvalidQueryParam
	NsErrorSmInvalidQueryParam = &pNsErrorSmInvalidQueryParam

	pNsErrorSmErrGroupMergeDbLoad = cNsErrorSmErrGroupMergeDbLoad
	NsErrorSmErrGroupMergeDbLoad = &pNsErrorSmErrGroupMergeDbLoad

	pNsErrorSmEio = cNsErrorSmEio
	NsErrorSmEio = &pNsErrorSmEio

	pNsErrorSmSrepDownstreamOnlineVol = cNsErrorSmSrepDownstreamOnlineVol
	NsErrorSmSrepDownstreamOnlineVol = &pNsErrorSmSrepDownstreamOnlineVol

	pNsErrorSmPoolNoSrcArray = cNsErrorSmPoolNoSrcArray
	NsErrorSmPoolNoSrcArray = &pNsErrorSmPoolNoSrcArray

	pNsErrorSmInvalidKeyvalue = cNsErrorSmInvalidKeyvalue
	NsErrorSmInvalidKeyvalue = &pNsErrorSmInvalidKeyvalue

	pNsErrorSmProtpolMaxLength = cNsErrorSmProtpolMaxLength
	NsErrorSmProtpolMaxLength = &pNsErrorSmProtpolMaxLength

	pNsErrorSmVolAssocVolcoll = cNsErrorSmVolAssocVolcoll
	NsErrorSmVolAssocVolcoll = &pNsErrorSmVolAssocVolcoll

	pNsErrorSmMissingDiscoveryIp = cNsErrorSmMissingDiscoveryIp
	NsErrorSmMissingDiscoveryIp = &pNsErrorSmMissingDiscoveryIp

	pNsErrorSmReplDeleteReplicaUnsup = cNsErrorSmReplDeleteReplicaUnsup
	NsErrorSmReplDeleteReplicaUnsup = &pNsErrorSmReplDeleteReplicaUnsup

	pNsErrorSmVolSizeDecreased = cNsErrorSmVolSizeDecreased
	NsErrorSmVolSizeDecreased = &pNsErrorSmVolSizeDecreased

	pNsErrorSmSrepDownstreamAssocedVols = cNsErrorSmSrepDownstreamAssocedVols
	NsErrorSmSrepDownstreamAssocedVols = &pNsErrorSmSrepDownstreamAssocedVols

	pNsErrorSmAddNonfcToFcGroup = cNsErrorSmAddNonfcToFcGroup
	NsErrorSmAddNonfcToFcGroup = &pNsErrorSmAddNonfcToFcGroup

	pNsErrorSmNetconfigDoesNotExist = cNsErrorSmNetconfigDoesNotExist
	NsErrorSmNetconfigDoesNotExist = &pNsErrorSmNetconfigDoesNotExist

	pNsErrorSmFolderVolmvEnospace = cNsErrorSmFolderVolmvEnospace
	NsErrorSmFolderVolmvEnospace = &pNsErrorSmFolderVolmvEnospace

	pNsErrorSmNetconfigExistNoForce = cNsErrorSmNetconfigExistNoForce
	NsErrorSmNetconfigExistNoForce = &pNsErrorSmNetconfigExistNoForce

	pNsErrorSmPoolUpdateInval = cNsErrorSmPoolUpdateInval
	NsErrorSmPoolUpdateInval = &pNsErrorSmPoolUpdateInval

	pNsErrorSmVlanSubnetInManual = cNsErrorSmVlanSubnetInManual
	NsErrorSmVlanSubnetInManual = &pNsErrorSmVlanSubnetInManual

	pNsErrorSmReplVvolUnsup = cNsErrorSmReplVvolUnsup
	NsErrorSmReplVvolUnsup = &pNsErrorSmReplVvolUnsup

	pNsErrorSmErrTdzNotSupported = cNsErrorSmErrTdzNotSupported
	NsErrorSmErrTdzNotSupported = &pNsErrorSmErrTdzNotSupported

	pNsErrorSmPoolExists = cNsErrorSmPoolExists
	NsErrorSmPoolExists = &pNsErrorSmPoolExists

	pNsErrorSmAuth = cNsErrorSmAuth
	NsErrorSmAuth = &pNsErrorSmAuth

	pNsErrorSmInvalidObjectSetQuery = cNsErrorSmInvalidObjectSetQuery
	NsErrorSmInvalidObjectSetQuery = &pNsErrorSmInvalidObjectSetQuery

	pNsErrorSmSrepDownstreamOnlineVols = cNsErrorSmSrepDownstreamOnlineVols
	NsErrorSmSrepDownstreamOnlineVols = &pNsErrorSmSrepDownstreamOnlineVols

	pNsErrorSmNoMethodForUrlPattern = cNsErrorSmNoMethodForUrlPattern
	NsErrorSmNoMethodForUrlPattern = &pNsErrorSmNoMethodForUrlPattern

	pNsErrorSmVolNotOnline = cNsErrorSmVolNotOnline
	NsErrorSmVolNotOnline = &pNsErrorSmVolNotOnline

	pNsErrorSmVolDedupeVolfamAppcat = cNsErrorSmVolDedupeVolfamAppcat
	NsErrorSmVolDedupeVolfamAppcat = &pNsErrorSmVolDedupeVolfamAppcat

	pNsErrorSmInvalidNic = cNsErrorSmInvalidNic
	NsErrorSmInvalidNic = &pNsErrorSmInvalidNic

	pNsErrorSmArrayNotGroupLeader = cNsErrorSmArrayNotGroupLeader
	NsErrorSmArrayNotGroupLeader = &pNsErrorSmArrayNotGroupLeader

	pNsErrorSmInvalidVlanId = cNsErrorSmInvalidVlanId
	NsErrorSmInvalidVlanId = &pNsErrorSmInvalidVlanId

	pNsErrorSmLimitSnaplun = cNsErrorSmLimitSnaplun
	NsErrorSmLimitSnaplun = &pNsErrorSmLimitSnaplun

	pNsErrorSmIncompatibleCachePolicy = cNsErrorSmIncompatibleCachePolicy
	NsErrorSmIncompatibleCachePolicy = &pNsErrorSmIncompatibleCachePolicy

	pNsErrorSmVolmvAbortEnospace = cNsErrorSmVolmvAbortEnospace
	NsErrorSmVolmvAbortEnospace = &pNsErrorSmVolmvAbortEnospace

	pNsErrorSmLimitHretSnapGroup = cNsErrorSmLimitHretSnapGroup
	NsErrorSmLimitHretSnapGroup = &pNsErrorSmLimitHretSnapGroup

	pNsErrorSmVolmvEalready = cNsErrorSmVolmvEalready
	NsErrorSmVolmvEalready = &pNsErrorSmVolmvEalready

	pNsErrorSmAsupPingfromCtrlrbError = cNsErrorSmAsupPingfromCtrlrbError
	NsErrorSmAsupPingfromCtrlrbError = &pNsErrorSmAsupPingfromCtrlrbError

	pNsErrorSmVolDedupeUnassignedAppCategory = cNsErrorSmVolDedupeUnassignedAppCategory
	NsErrorSmVolDedupeUnassignedAppCategory = &pNsErrorSmVolDedupeUnassignedAppCategory

	pNsErrorSmEnoent = cNsErrorSmEnoent
	NsErrorSmEnoent = &pNsErrorSmEnoent

	pNsErrorSmPerfpolDedupeUnassignedAppCategory = cNsErrorSmPerfpolDedupeUnassignedAppCategory
	NsErrorSmPerfpolDedupeUnassignedAppCategory = &pNsErrorSmPerfpolDedupeUnassignedAppCategory

	pNsErrorSmInvalidInitiatorLabel = cNsErrorSmInvalidInitiatorLabel
	NsErrorSmInvalidInitiatorLabel = &pNsErrorSmInvalidInitiatorLabel

	pNsErrorSmDuplicateInitiatorWithArgs = cNsErrorSmDuplicateInitiatorWithArgs
	NsErrorSmDuplicateInitiatorWithArgs = &pNsErrorSmDuplicateInitiatorWithArgs

	pNsErrorSmErrShelfForeign = cNsErrorSmErrShelfForeign
	NsErrorSmErrShelfForeign = &pNsErrorSmErrShelfForeign

	pNsErrorSmInvalidAgentType = cNsErrorSmInvalidAgentType
	NsErrorSmInvalidAgentType = &pNsErrorSmInvalidAgentType

	pNsErrorSmEinprogress = cNsErrorSmEinprogress
	NsErrorSmEinprogress = &pNsErrorSmEinprogress

	pNsErrorSmNotEnoughCache = cNsErrorSmNotEnoughCache
	NsErrorSmNotEnoughCache = &pNsErrorSmNotEnoughCache

	pNsErrorSmEexist = cNsErrorSmEexist
	NsErrorSmEexist = &pNsErrorSmEexist

	pNsErrorSmMissingArrayNetconfig = cNsErrorSmMissingArrayNetconfig
	NsErrorSmMissingArrayNetconfig = &pNsErrorSmMissingArrayNetconfig

	pNsErrorSmInvalidInitiatorAlias = cNsErrorSmInvalidInitiatorAlias
	NsErrorSmInvalidInitiatorAlias = &pNsErrorSmInvalidInitiatorAlias

	pNsErrorSmProtpolAppSyncOracle = cNsErrorSmProtpolAppSyncOracle
	NsErrorSmProtpolAppSyncOracle = &pNsErrorSmProtpolAppSyncOracle

	pNsErrorSmNoSupport = cNsErrorSmNoSupport
	NsErrorSmNoSupport = &pNsErrorSmNoSupport

	pNsErrorSmDataIpMissingSubnet = cNsErrorSmDataIpMissingSubnet
	NsErrorSmDataIpMissingSubnet = &pNsErrorSmDataIpMissingSubnet

	pNsErrorSmErrShelfHddsInAfs = cNsErrorSmErrShelfHddsInAfs
	NsErrorSmErrShelfHddsInAfs = &pNsErrorSmErrShelfHddsInAfs

	pNsErrorSmStartRowBeyondTotalRows = cNsErrorSmStartRowBeyondTotalRows
	NsErrorSmStartRowBeyondTotalRows = &pNsErrorSmStartRowBeyondTotalRows

	pNsErrorSmExtraneousArrayNetconfig = cNsErrorSmExtraneousArrayNetconfig
	NsErrorSmExtraneousArrayNetconfig = &pNsErrorSmExtraneousArrayNetconfig

	pNsErrorSmPoolCachePinNotsupp = cNsErrorSmPoolCachePinNotsupp
	NsErrorSmPoolCachePinNotsupp = &pNsErrorSmPoolCachePinNotsupp

	pNsErrorSmUsageUnavailable = cNsErrorSmUsageUnavailable
	NsErrorSmUsageUnavailable = &pNsErrorSmUsageUnavailable

	pNsErrorSmReplAgentTypeUnsup = cNsErrorSmReplAgentTypeUnsup
	NsErrorSmReplAgentTypeUnsup = &pNsErrorSmReplAgentTypeUnsup

	pNsErrorSmReplFanoutMaximumPartnersExceeded = cNsErrorSmReplFanoutMaximumPartnersExceeded
	NsErrorSmReplFanoutMaximumPartnersExceeded = &pNsErrorSmReplFanoutMaximumPartnersExceeded

	pNsErrorSmAsupPingfromCtrlraError = cNsErrorSmAsupPingfromCtrlraError
	NsErrorSmAsupPingfromCtrlraError = &pNsErrorSmAsupPingfromCtrlraError

	pNsErrorSmErrShelfInvalidLoc = cNsErrorSmErrShelfInvalidLoc
	NsErrorSmErrShelfInvalidLoc = &pNsErrorSmErrShelfInvalidLoc

	pNsErrorSmErrShelfSsdDegraded = cNsErrorSmErrShelfSsdDegraded
	NsErrorSmErrShelfSsdDegraded = &pNsErrorSmErrShelfSsdDegraded

	pNsErrorSmSyncReplConfigure = cNsErrorSmSyncReplConfigure
	NsErrorSmSyncReplConfigure = &pNsErrorSmSyncReplConfigure

	pNsErrorSmAsupHeartbeatError = cNsErrorSmAsupHeartbeatError
	NsErrorSmAsupHeartbeatError = &pNsErrorSmAsupHeartbeatError

	pNsErrorSmLimitHretSnapRetentionPoolWarn = cNsErrorSmLimitHretSnapRetentionPoolWarn
	NsErrorSmLimitHretSnapRetentionPoolWarn = &pNsErrorSmLimitHretSnapRetentionPoolWarn

	pNsErrorSmErrAuth = cNsErrorSmErrAuth
	NsErrorSmErrAuth = &pNsErrorSmErrAuth

	pNsErrorSmPartnerSubnetDoesNotExist = cNsErrorSmPartnerSubnetDoesNotExist
	NsErrorSmPartnerSubnetDoesNotExist = &pNsErrorSmPartnerSubnetDoesNotExist

	pNsErrorSmErrShelfLocOrder = cNsErrorSmErrShelfLocOrder
	NsErrorSmErrShelfLocOrder = &pNsErrorSmErrShelfLocOrder

	pNsErrorSmFolderEnospace = cNsErrorSmFolderEnospace
	NsErrorSmFolderEnospace = &pNsErrorSmFolderEnospace

	pNsErrorSmErrPoolHasGroupPartners = cNsErrorSmErrPoolHasGroupPartners
	NsErrorSmErrPoolHasGroupPartners = &pNsErrorSmErrPoolHasGroupPartners

	pNsErrorSmProtpolNotSpecified = cNsErrorSmProtpolNotSpecified
	NsErrorSmProtpolNotSpecified = &pNsErrorSmProtpolNotSpecified

	pNsErrorSmUnexpectedQueryParam = cNsErrorSmUnexpectedQueryParam
	NsErrorSmUnexpectedQueryParam = &pNsErrorSmUnexpectedQueryParam

	pNsErrorSmVolmvAbortEalready = cNsErrorSmVolmvAbortEalready
	NsErrorSmVolmvAbortEalready = &pNsErrorSmVolmvAbortEalready

	pNsErrorSmFcIntfAlreadyInState = cNsErrorSmFcIntfAlreadyInState
	NsErrorSmFcIntfAlreadyInState = &pNsErrorSmFcIntfAlreadyInState

	pNsErrorSmLimitHretSnapPool = cNsErrorSmLimitHretSnapPool
	NsErrorSmLimitHretSnapPool = &pNsErrorSmLimitHretSnapPool

	pNsErrorSmIpUpdateNoForce = cNsErrorSmIpUpdateNoForce
	NsErrorSmIpUpdateNoForce = &pNsErrorSmIpUpdateNoForce

	pNsErrorSmSrepAgentTypeMismatchDownstreamVol = cNsErrorSmSrepAgentTypeMismatchDownstreamVol
	NsErrorSmSrepAgentTypeMismatchDownstreamVol = &pNsErrorSmSrepAgentTypeMismatchDownstreamVol

	pNsErrorSmVssValidationTimedout = cNsErrorSmVssValidationTimedout
	NsErrorSmVssValidationTimedout = &pNsErrorSmVssValidationTimedout

	pNsErrorSmConfigSyncInprogress = cNsErrorSmConfigSyncInprogress
	NsErrorSmConfigSyncInprogress = &pNsErrorSmConfigSyncInprogress

	pNsErrorSmAsyncJobId = cNsErrorSmAsyncJobId
	NsErrorSmAsyncJobId = &pNsErrorSmAsyncJobId

	pNsErrorSmEagain = cNsErrorSmEagain
	NsErrorSmEagain = &pNsErrorSmEagain

	pNsErrorSmPerfpolVolMoveAppCategory = cNsErrorSmPerfpolVolMoveAppCategory
	NsErrorSmPerfpolVolMoveAppCategory = &pNsErrorSmPerfpolVolMoveAppCategory

	pNsErrorSmLimitHretSnapRetentionPoolMax = cNsErrorSmLimitHretSnapRetentionPoolMax
	NsErrorSmLimitHretSnapRetentionPoolMax = &pNsErrorSmLimitHretSnapRetentionPoolMax

	pNsErrorSmVolHasConnections = cNsErrorSmVolHasConnections
	NsErrorSmVolHasConnections = &pNsErrorSmVolHasConnections

	pNsErrorSmNoCommonLun = cNsErrorSmNoCommonLun
	NsErrorSmNoCommonLun = &pNsErrorSmNoCommonLun

	pNsErrorSmErrShelfSasLanesDegraded = cNsErrorSmErrShelfSasLanesDegraded
	NsErrorSmErrShelfSasLanesDegraded = &pNsErrorSmErrShelfSasLanesDegraded

	pNsErrorSmVolAppCategoryMoveInvalid = cNsErrorSmVolAppCategoryMoveInvalid
	NsErrorSmVolAppCategoryMoveInvalid = &pNsErrorSmVolAppCategoryMoveInvalid

	pNsErrorSmExtTrigSchedAlreadyPresent = cNsErrorSmExtTrigSchedAlreadyPresent
	NsErrorSmExtTrigSchedAlreadyPresent = &pNsErrorSmExtTrigSchedAlreadyPresent

	pNsErrorSmNoDataIpSpecified = cNsErrorSmNoDataIpSpecified
	NsErrorSmNoDataIpSpecified = &pNsErrorSmNoDataIpSpecified

	pNsErrorSmInvalidVolAssoc = cNsErrorSmInvalidVolAssoc
	NsErrorSmInvalidVolAssoc = &pNsErrorSmInvalidVolAssoc

	pNsErrorSmReplObjectBusy = cNsErrorSmReplObjectBusy
	NsErrorSmReplObjectBusy = &pNsErrorSmReplObjectBusy

	pNsErrorSmVolcollOwner = cNsErrorSmVolcollOwner
	NsErrorSmVolcollOwner = &pNsErrorSmVolcollOwner

	pNsErrorSmReservedUsername = cNsErrorSmReservedUsername
	NsErrorSmReservedUsername = &pNsErrorSmReservedUsername

	pNsErrorSmVvolsnapOnline = cNsErrorSmVvolsnapOnline
	NsErrorSmVvolsnapOnline = &pNsErrorSmVvolsnapOnline

	pNsErrorSmFolderVolmvEinprog = cNsErrorSmFolderVolmvEinprog
	NsErrorSmFolderVolmvEinprog = &pNsErrorSmFolderVolmvEinprog

	pNsErrorSmUnexpectedObjectSetQuery = cNsErrorSmUnexpectedObjectSetQuery
	NsErrorSmUnexpectedObjectSetQuery = &pNsErrorSmUnexpectedObjectSetQuery

	pNsErrorSmProtpolInvalidValue = cNsErrorSmProtpolInvalidValue
	NsErrorSmProtpolInvalidValue = &pNsErrorSmProtpolInvalidValue

	pNsErrorSmFolderIncompatibleAgentType = cNsErrorSmFolderIncompatibleAgentType
	NsErrorSmFolderIncompatibleAgentType = &pNsErrorSmFolderIncompatibleAgentType

	pNsErrorSmProtpolVmwareInvalidVcenterHostname = cNsErrorSmProtpolVmwareInvalidVcenterHostname
	NsErrorSmProtpolVmwareInvalidVcenterHostname = &pNsErrorSmProtpolVmwareInvalidVcenterHostname

	pNsErrorSmErrVolCollMultipleSchedules = cNsErrorSmErrVolCollMultipleSchedules
	NsErrorSmErrVolCollMultipleSchedules = &pNsErrorSmErrVolCollMultipleSchedules

	pNsErrorSmReplPartnerNameMismatch = cNsErrorSmReplPartnerNameMismatch
	NsErrorSmReplPartnerNameMismatch = &pNsErrorSmReplPartnerNameMismatch

	pNsErrorSmInvalidFolder = cNsErrorSmInvalidFolder
	NsErrorSmInvalidFolder = &pNsErrorSmInvalidFolder

	pNsErrorSmSrvUpdatePrecheckArray = cNsErrorSmSrvUpdatePrecheckArray
	NsErrorSmSrvUpdatePrecheckArray = &pNsErrorSmSrvUpdatePrecheckArray

	pNsErrorSmGatewayNotInSubnets = cNsErrorSmGatewayNotInSubnets
	NsErrorSmGatewayNotInSubnets = &pNsErrorSmGatewayNotInSubnets

	pNsErrorSmDeprecatedPerfpol = cNsErrorSmDeprecatedPerfpol
	NsErrorSmDeprecatedPerfpol = &pNsErrorSmDeprecatedPerfpol

	pNsErrorSmTakeoverSplitBrain = cNsErrorSmTakeoverSplitBrain
	NsErrorSmTakeoverSplitBrain = &pNsErrorSmTakeoverSplitBrain

	pNsErrorSmPeIgroupProtocolMismatched = cNsErrorSmPeIgroupProtocolMismatched
	NsErrorSmPeIgroupProtocolMismatched = &pNsErrorSmPeIgroupProtocolMismatched

	pNsErrorSmOnlyVvolFolderAppsrvr = cNsErrorSmOnlyVvolFolderAppsrvr
	NsErrorSmOnlyVvolFolderAppsrvr = &pNsErrorSmOnlyVvolFolderAppsrvr

	pNsErrorSmVersionMismatch = cNsErrorSmVersionMismatch
	NsErrorSmVersionMismatch = &pNsErrorSmVersionMismatch

	pNsErrorSmPoolLastArray = cNsErrorSmPoolLastArray
	NsErrorSmPoolLastArray = &pNsErrorSmPoolLastArray

	pNsErrorSmEaccess = cNsErrorSmEaccess
	NsErrorSmEaccess = &pNsErrorSmEaccess

	pNsErrorSmInvalidSubnetLabel = cNsErrorSmInvalidSubnetLabel
	NsErrorSmInvalidSubnetLabel = &pNsErrorSmInvalidSubnetLabel

	pNsErrorSmInvalidArg = cNsErrorSmInvalidArg
	NsErrorSmInvalidArg = &pNsErrorSmInvalidArg

	pNsErrorSmDedupeVolfamAppcat = cNsErrorSmDedupeVolfamAppcat
	NsErrorSmDedupeVolfamAppcat = &pNsErrorSmDedupeVolfamAppcat

	pNsErrorSmSrvUnreach = cNsErrorSmSrvUnreach
	NsErrorSmSrvUnreach = &pNsErrorSmSrvUnreach

	pNsErrorSmPoolPartnerPauseUnsup = cNsErrorSmPoolPartnerPauseUnsup
	NsErrorSmPoolPartnerPauseUnsup = &pNsErrorSmPoolPartnerPauseUnsup

	pNsErrorSmNetconfigCreateDraftOnly = cNsErrorSmNetconfigCreateDraftOnly
	NsErrorSmNetconfigCreateDraftOnly = &pNsErrorSmNetconfigCreateDraftOnly

	pNsErrorSmErrArrayNotFound = cNsErrorSmErrArrayNotFound
	NsErrorSmErrArrayNotFound = &pNsErrorSmErrArrayNotFound

	pNsErrorSmFolsetInUse = cNsErrorSmFolsetInUse
	NsErrorSmFolsetInUse = &pNsErrorSmFolsetInUse

	pNsErrorSmErrShelfSesDeviceNotReady = cNsErrorSmErrShelfSesDeviceNotReady
	NsErrorSmErrShelfSesDeviceNotReady = &pNsErrorSmErrShelfSesDeviceNotReady

	pNsErrorSmInvalidIp = cNsErrorSmInvalidIp
	NsErrorSmInvalidIp = &pNsErrorSmInvalidIp

	pNsErrorSmEalready = cNsErrorSmEalready
	NsErrorSmEalready = &pNsErrorSmEalready

	pNsErrorSmInvalidNonIscsiDataSubnet = cNsErrorSmInvalidNonIscsiDataSubnet
	NsErrorSmInvalidNonIscsiDataSubnet = &pNsErrorSmInvalidNonIscsiDataSubnet

	pNsErrorSmReplHandoverUnsupPtype = cNsErrorSmReplHandoverUnsupPtype
	NsErrorSmReplHandoverUnsupPtype = &pNsErrorSmReplHandoverUnsupPtype

	pNsErrorSmArrayNotAssigned = cNsErrorSmArrayNotAssigned
	NsErrorSmArrayNotAssigned = &pNsErrorSmArrayNotAssigned

	pNsErrorSmEpartial = cNsErrorSmEpartial
	NsErrorSmEpartial = &pNsErrorSmEpartial

	pNsErrorSmErrProtpolMissingName = cNsErrorSmErrProtpolMissingName
	NsErrorSmErrProtpolMissingName = &pNsErrorSmErrProtpolMissingName

	pNsErrorSmVfVolCachePinned = cNsErrorSmVfVolCachePinned
	NsErrorSmVfVolCachePinned = &pNsErrorSmVfVolCachePinned

	pNsErrorSmAtLeastOneIscsiCluster = cNsErrorSmAtLeastOneIscsiCluster
	NsErrorSmAtLeastOneIscsiCluster = &pNsErrorSmAtLeastOneIscsiCluster

	pNsErrorSmErrTooMany = cNsErrorSmErrTooMany
	NsErrorSmErrTooMany = &pNsErrorSmErrTooMany

	pNsErrorSmErrShelfExprMfgVerInval = cNsErrorSmErrShelfExprMfgVerInval
	NsErrorSmErrShelfExprMfgVerInval = &pNsErrorSmErrShelfExprMfgVerInval

	pNsErrorSmMgmtIpInvalid = cNsErrorSmMgmtIpInvalid
	NsErrorSmMgmtIpInvalid = &pNsErrorSmMgmtIpInvalid

	pNsErrorSmErrShelfWrongCtrlrSide = cNsErrorSmErrShelfWrongCtrlrSide
	NsErrorSmErrShelfWrongCtrlrSide = &pNsErrorSmErrShelfWrongCtrlrSide

}
