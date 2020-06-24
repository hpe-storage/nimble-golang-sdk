/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Volume - Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on storage allocation.
// Export VolumeFields for advance operations like search filter etc.
var VolumeFields *Volume

func init(){
	IDfield:= "id"
	Namefield:= "name"
	FullNamefield:= "full_name"
	SearchNamefield:= "search_name"
	Descriptionfield:= "description"
	PerfpolicyNamefield:= "perfpolicy_name"
	PerfpolicyIDfield:= "perfpolicy_id"
	OwnedByGroupfield:= "owned_by_group"
	OwnedByGroupIDfield:= "owned_by_group_id"
	PoolNamefield:= "pool_name"
	PoolIDfield:= "pool_id"
	SerialNumberfield:= "serial_number"
	SecondarySerialNumberfield:= "secondary_serial_number"
	TargetNamefield:= "target_name"
	ParentVolNamefield:= "parent_vol_name"
	ParentVolIDfield:= "parent_vol_id"
	BaseSnapNamefield:= "base_snap_name"
	BaseSnapIDfield:= "base_snap_id"
	VolcollNamefield:= "volcoll_name"
	VolcollIDfield:= "volcoll_id"
	DestPoolNamefield:= "dest_pool_name"
	DestPoolIDfield:= "dest_pool_id"
	AppUuIDfield:= "app_uuid"
	FolderIDfield:= "folder_id"
	FolderNamefield:= "folder_name"
	VpdT10field:= "vpd_t10"
	VpdIeee0field:= "vpd_ieee0"
	VpdIeee1field:= "vpd_ieee1"
	AppCategoryfield:= "app_category"
	PreFilterfield:= "pre_filter"
		
	VolumeFields= &Volume{
		ID: &IDfield,
		Name: &Namefield,
		FullName: &FullNamefield,
		SearchName: &SearchNamefield,
		Description: &Descriptionfield,
		PerfpolicyName: &PerfpolicyNamefield,
		PerfpolicyID: &PerfpolicyIDfield,
		OwnedByGroup: &OwnedByGroupfield,
		OwnedByGroupID: &OwnedByGroupIDfield,
		PoolName: &PoolNamefield,
		PoolID: &PoolIDfield,
		SerialNumber: &SerialNumberfield,
		SecondarySerialNumber: &SecondarySerialNumberfield,
		TargetName: &TargetNamefield,
		ParentVolName: &ParentVolNamefield,
		ParentVolID: &ParentVolIDfield,
		BaseSnapName: &BaseSnapNamefield,
		BaseSnapID: &BaseSnapIDfield,
		VolcollName: &VolcollNamefield,
		VolcollID: &VolcollIDfield,
		DestPoolName: &DestPoolNamefield,
		DestPoolID: &DestPoolIDfield,
		AppUuID: &AppUuIDfield,
		FolderID: &FolderIDfield,
		FolderName: &FolderNamefield,
		VpdT10: &VpdT10field,
		VpdIeee0: &VpdIeee0field,
		VpdIeee1: &VpdIeee1field,
		AppCategory: &AppCategoryfield,
		PreFilter: &PreFilterfield,
		
	}
}

type Volume struct {
   
    // Identifier for the volume.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of the volume.
    
 	Name *string `json:"name,omitempty"`
   
    // Fully qualified name of volume.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // Name of volume used for object search.
    
 	SearchName *string `json:"search_name,omitempty"`
   
    // Volume size in mebibytes. Size is required for creating a volume but not for cloning an existing volume.
    
   	Size *int64 `json:"size,omitempty"`
   
    // Text description of volume.
    
 	Description *string `json:"description,omitempty"`
   
    // Name of performance policy.
    
 	PerfpolicyName *string `json:"perfpolicy_name,omitempty"`
   
    // Identifier of the performance policy. After creating a volume, performance policy for the volume can only be changed to another performance policy with same block size.
    
 	PerfpolicyID *string `json:"perfpolicy_id,omitempty"`
   
    // Amount of space to reserve for this volume as a percentage of volume size.
    
   	Reserve *int64 `json:"reserve,omitempty"`
   
    // This attribute is deprecated. Alert threshold for the volume's mapped usage, expressed as a percentage of the volume's size. When the volume's mapped usage exceeds warn_level, the array issues an alert. If this option is not specified, array default volume warn level setting is used to decide the warning level for this volume.
    
   	WarnLevel *int64 `json:"warn_level,omitempty"`
   
    // Limit on the volume's mapped usage, expressed as a percentage of the volume's size. When the volume's mapped usage exceeds limit, the volume will be offlined or made non-writable. If this option is not specified, array default volume limit setting is used to decide the limit for this volume.
    
   	Limit *int64 `json:"limit,omitempty"`
   
    // Amount of space to reserve for snapshots of this volume as a percentage of volume size.
    
   	SnapReserve *int64 `json:"snap_reserve,omitempty"`
   
    // Threshold for available space as a percentage of volume size below which an alert is raised.
    
   	SnapWarnLevel *int64 `json:"snap_warn_level,omitempty"`
   
    // This attribute is deprecated. The array does not limit a volume's snapshot space usage. The attribute is ignored on input and returns max int64 on output.
    
   	SnapLimit *int64 `json:"snap_limit,omitempty"`
   
    // This attribute is deprecated. The array does not limit a volume's snapshot space usage. The attribute is ignored on input and returns -1 on output.
    
  	SnapLimitPercent  *int64 `json:"snap_limit_percent,omitempty"`
   
    // Number of live, non-hidden snapshots for this volume.
    
   	NumSnaps *int64 `json:"num_snaps,omitempty"`
   
    // Deprecated. Projected number of snapshots (including scheduled and manual) for this volume.
    
   	ProjectedNumSnaps *int64 `json:"projected_num_snaps,omitempty"`
   
    // Online state of volume, available for host initiators to establish connections.
    
 	Online *bool `json:"online,omitempty"`
   
    // Name of group that currently owns the volume.
    
 	OwnedByGroup *string `json:"owned_by_group,omitempty"`
   
    // ID of group that currently owns the volume.
    
 	OwnedByGroupID *string `json:"owned_by_group_id,omitempty"`
   
    // For iSCSI Volume Target, this flag indicates whether the volume and its snapshots can be accessed from multiple initiators at the same time. The default is false. For iSCSI Group Target or FC access protocol, the attribute cannot be modified and always reads as false.
    
 	MultiInitiator *bool `json:"multi_initiator,omitempty"`
   
    // This indicates whether volume is exported under iSCSI Group Target or iSCSI Volume Target. This attribute is only meaningful to iSCSI system. On FC system, all volumes are exported under the FC Group Target. In create operation, the volume's target type will be set by this attribute. If not specified, it will be set as the group-setting. In clone operation, the clone's target type will inherit from the parent' setting.
    
   	IscsiTargetScope *NsTargetScope `json:"iscsi_target_scope,omitempty"`
   
    // Name of the pool where the volume resides. Volume data will be distributed across arrays over which specified pool is defined. If pool option is not specified, volume is assigned to the default pool.
    
 	PoolName *string `json:"pool_name,omitempty"`
   
    // Identifier associated with the pool in the storage pool table.
    
 	PoolID *string `json:"pool_id,omitempty"`
   
    // Volume is read-only.
    
 	ReadOnly *bool `json:"read_only,omitempty"`
   
    // Identifier associated with the volume for the SCSI protocol.
    
 	SerialNumber *string `json:"serial_number,omitempty"`
   
    // Secondary identifier associated with the volume for the SCSI protocol.
    
 	SecondarySerialNumber *string `json:"secondary_serial_number,omitempty"`
   
    // The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target volume.
    
 	TargetName *string `json:"target_name,omitempty"`
   
    // Size in bytes of blocks in the volume.
    
   	BlockSize *int64 `json:"block_size,omitempty"`
   
    // Volume offline reason.
    
   	OfflineReason *NsOfflineReason `json:"offline_reason,omitempty"`
   
    // Whether this volume is a clone. Use this attribute in combination with name and base_snap_id to create a clone by setting clone = true.
    
 	Clone *bool `json:"clone,omitempty"`
   
    // Name of parent volume.
    
 	ParentVolName *string `json:"parent_vol_name,omitempty"`
   
    // Parent volume ID.
    
 	ParentVolID *string `json:"parent_vol_id,omitempty"`
   
    // Name of base snapshot.
    
 	BaseSnapName *string `json:"base_snap_name,omitempty"`
   
    // Base snapshot ID. This attribute is required together with name and clone when cloning a volume with the create operation.
    
 	BaseSnapID *string `json:"base_snap_id,omitempty"`
   
    // Replication role that this volume performs.
    
   	ReplicationRole *NsVolumeReplicationRole `json:"replication_role,omitempty"`
   
    // Name of volume collection of which this volume is a member.
    
 	VolcollName *string `json:"volcoll_name,omitempty"`
   
    // ID of volume collection of which this volume is a member. Use this attribute in update operation to associate or dissociate volumes with or from volume collections. When associating, set this attribute to the ID of the volume collection. When dissociating, set this attribute to empty string.
    
 	VolcollID *string `json:"volcoll_id,omitempty"`
   
    // External management agent type.
    
   	AgentType *NsAgentType `json:"agent_type,omitempty"`
   
    // Forcibly offline, reduce size or change read-only status a volume.
    
 	Force *bool `json:"force,omitempty"`
   
    // Time when this volume was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this volume was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Specifies if volume is protected with schedules. If protected, indicate whether replication is setup.
    
   	ProtectionType *NsProtectionType `json:"protection_type,omitempty"`
   
    // Last snapshot for this volume.
    
	LastSnap *NsSnapSummary `json:"last_snap,omitempty"`
   
    // Last replicated snapshot for this volume.
    
	LastReplicatedSnap *NsSnapSummary `json:"last_replicated_snap,omitempty"`
   
    // Name of the destination pool where the volume is moving to.
    
 	DestPoolName *string `json:"dest_pool_name,omitempty"`
   
    // ID of the destination pool where the volume is moving to.
    
 	DestPoolID *string `json:"dest_pool_id,omitempty"`
   
    // The Start time when this volume was moved.
    
   	MoveStartTime *int64 `json:"move_start_time,omitempty"`
   
    // This indicates whether the move of the volume is aborting or not.
    
 	MoveAborting *bool `json:"move_aborting,omitempty"`
   
    // The bytes of volume which have been moved.
    
   	MoveBytesMigrated *int64 `json:"move_bytes_migrated,omitempty"`
   
    // The bytes of volume which have not been moved.
    
   	MoveBytesRemaining *int64 `json:"move_bytes_remaining,omitempty"`
   
    // The estimated time of completion of a move.
    
   	MoveEstComplTime *int64 `json:"move_est_compl_time,omitempty"`
   
    // This indicates whether usage information of volume and snapshots are valid or not.
    
 	UsageValID *bool `json:"usage_valid,omitempty"`
   
    // Indicates space usage level based on warning level.
    
   	SpaceUsageLevel *NsSpaceUsageLevel `json:"space_usage_level,omitempty"`
   
    // Sum of volume mapped usage and uncompressed backup data(including pending deletes) in bytes of this volume.
    
   	TotalUsageBytes *int64 `json:"total_usage_bytes,omitempty"`
   
    // Compressed data in bytes for this volume.
    
   	VolUsageCompressedBytes *int64 `json:"vol_usage_compressed_bytes,omitempty"`
   
    // Uncompressed data in bytes for this volume.
    
   	VolUsageUncompressedBytes *int64 `json:"vol_usage_uncompressed_bytes,omitempty"`
   
    // Mapped data in bytes for this volume.
    
   	VolUsageMappedBytes *int64 `json:"vol_usage_mapped_bytes,omitempty"`
   
    // Sum of compressed backup data in bytes stored in snapshots of this volume.
    
   	SnapUsageCompressedBytes *int64 `json:"snap_usage_compressed_bytes,omitempty"`
   
    // Sum of uncompressed unique backup data in bytes stored in snapshots of this volume.
    
   	SnapUsageUncompressedBytes *int64 `json:"snap_usage_uncompressed_bytes,omitempty"`
   
    // Sum of backup data in bytes stored in snapshots of this volume without accounting for the sharing of data between snapshots.
    
   	SnapUsagePopulatedBytes *int64 `json:"snap_usage_populated_bytes,omitempty"`
   
    // If set to true, all the contents of this volume are kept in flash cache. This provides for consistent performance guarantees for all types of workloads. The amount of flash needed to pin the volume is equal to the limit for the volume.
    
 	CachePinned *bool `json:"cache_pinned,omitempty"`
   
    // The amount of flash pinned on the volume.
    
   	PinnedCacheSize *int64 `json:"pinned_cache_size,omitempty"`
   
    // The amount of flash needed to pin the volume.
    
   	CacheNeededForPin *int64 `json:"cache_needed_for_pin,omitempty"`
   
    // This indicates whether the upstream volume is cache pinned or not.
    
 	UpstreamCachePinned *bool `json:"upstream_cache_pinned,omitempty"`
   
    // Cache policy applied to the volume.
    
   	CachePolicy *NsCachePolicy `json:"cache_policy,omitempty"`
   
    // Set volume's provisioning level to thin.  Also advertises volume as thinly provisioned to initiators supporting thin provisioning. For such volumes, soft limit notification is set to initiators when the volume space usage crosses its volume_warn_level. Default is yes. The volume's space is provisioned immediately, but for advertising status, this change takes effect only for new connections to the volume. Initiators must disconnect and reconnect for the new setting to be take effect at the initiator level consistently.
    
 	ThinlyProvisioned *bool `json:"thinly_provisioned,omitempty"`
   
    // Status of the volume.
    
   	VolState *NsVolStatus `json:"vol_state,omitempty"`
   
    // The list of online snapshots of this volume.
    
   	OnlineSnaps []*NsSnapshotFromVolumes `json:"online_snaps,omitempty"`
   
    // Number of connections of this volume.
    
   	NumConnections *int64 `json:"num_connections,omitempty"`
   
    // Number of iscsi connections of this volume.
    
   	NumIscsiConnections *int64 `json:"num_iscsi_connections,omitempty"`
   
    // Number of Fibre Channel connections of this volume.
    
   	NumFcConnections *int64 `json:"num_fc_connections,omitempty"`
   
    // List of access control records that apply to this volume.
    
   	AccessControlRecords []*NsAccessControlRecord `json:"access_control_records,omitempty"`
   
    // In a volume clone operation, if both the parent and the clone have no external management agent (their agent_type property is "none"), then inherit_acl controls whether the clone will inherit a copy of the parent's access control list. If either the parent or the clone have an external management agent, then the clone will not inherit the parent's access control list.
    
 	InheritAcl *bool `json:"inherit_acl,omitempty"`
   
    // The encryption cipher of the volume.
    
   	EncryptionCipher *NsEncryptionCipher `json:"encryption_cipher,omitempty"`
   
    // Application identifier of volume.
    
 	AppUuID *string `json:"app_uuid,omitempty"`
   
    // ID of the folder holding this volume.
    
 	FolderID *string `json:"folder_id,omitempty"`
   
    // Name of the folder holding this volume. It can be empty.
    
 	FolderName *string `json:"folder_name,omitempty"`
   
    // Key-value pairs that augment an volume's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
   
    // List of iSCSI sessions connected to this volume.
    
   	IscsiSessions []*NsISCSISession `json:"iscsi_sessions,omitempty"`
   
    // List of Fibre Channel sessions connected to this volume.
    
   	FcSessions []*NsFCSession `json:"fc_sessions,omitempty"`
   
    // Indicate caching the volume is enabled.
    
 	CachingEnabled *bool `json:"caching_enabled,omitempty"`
   
    // Indicate whether dedupe has ever been enabled on this volume.
    
 	PreviouslyDeduped *bool `json:"previously_deduped,omitempty"`
   
    // Indicate whether dedupe is enabled.
    
 	DedupeEnabled *bool `json:"dedupe_enabled,omitempty"`
   
    // The volume's T10 Vendor ID-based identifier.
    
 	VpdT10 *string `json:"vpd_t10,omitempty"`
   
    // The first 64 bits of the volume's EUI-64 identifier, encoded as a hexadecimal string.
    
 	VpdIeee0 *string `json:"vpd_ieee0,omitempty"`
   
    // The last 64 bits of the volume's EUI-64 identifier, encoded as a hexadecimal string.
    
 	VpdIeee1 *string `json:"vpd_ieee1,omitempty"`
   
    // Application category that the volume belongs to.
    
 	AppCategory *string `json:"app_category,omitempty"`
   
    // IOPS limit for this volume. If limit_iops is not specified when a volume is created, or if limit_iops is set to -1, then the volume has no IOPS limit. If limit_iops is not specified while creating a clone, IOPS limit of parent volume will be used as limit. IOPS limit should be in range [256, 4294967294] or -1 for unlimited. If both limit_iops and limit_mbps are specified, limit_mbps must not be hit before limit_iops. In other words, IOPS and MBPS limits should honor limit_iops <= ((limit_mbps MB/s * 2^20 B/MB) / block_size B).
    
  	LimitIops  *int64 `json:"limit_iops,omitempty"`
   
    // Throughput limit for this volume in MB/s. If limit_mbps is not specified when a volume is created, or if limit_mbps is set to -1, then the volume has no MBPS limit. MBPS limit should be in range [1, 4294967294] or -1 for unlimited. If both limit_iops and limit_mbps are specified, limit_mbps must not be hit before limit_iops. In other words, IOPS and MBPS limits should honor limit_iops <= ((limit_mbps MB/s * 2^20 B/MB) / block_size B).
    
  	LimitMbps  *int64 `json:"limit_mbps,omitempty"`
   
    // Indicates whether the volume needs content based replication.
    
 	NeedsContentRepl *bool `json:"needs_content_repl,omitempty"`
   
    // Indicates whether the last content based replication had errors.
    
 	ContentReplErrorsFound *bool `json:"content_repl_errors_found,omitempty"`
   
    // The branch cg uid of the content based snapshot that was last replicated.
    
   	LastContentSnapBrCgUID *int64 `json:"last_content_snap_br_cg_uid,omitempty"`
   
    // The branch gid of the content based snapshot that was last replicated.
    
   	LastContentSnapBrGID *int64 `json:"last_content_snap_br_gid,omitempty"`
   
    // The ID of the content based snapshot that was last replicated.
    
   	LastContentSnapID *int64 `json:"last_content_snap_id,omitempty"`
   
    // Last checksum verification time.
    
   	CksumLastVerified *int64 `json:"cksum_last_verified,omitempty"`
   
    // Pre-filtering criteria.
    
 	PreFilter *string `json:"pre_filter,omitempty"`
   
    // Average statistics in last 5 minutes.
    
	AvgStatsLast5mins *NsAverageStats `json:"avg_stats_last5mins,omitempty"`
   
    // Time when synchronously replicated volume was last synchronized.
    
   	SrepLastSync *int64 `json:"srep_last_sync,omitempty"`
   
    // Percentage of resync progress for synchronously replicated volume.
    
   	SrepResyncPercent *int64 `json:"srep_resync_percent,omitempty"`
}
