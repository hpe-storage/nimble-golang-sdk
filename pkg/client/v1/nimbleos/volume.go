// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Volume - Volumes are the basic storage units from which the total capacity is apportioned. The terms volume and LUN are used interchangeably.The number of volumes per array depends on storage allocation.
// Export VolumeFields for advance operations like search filter etc.
var VolumeFields *Volume

func init() {
	IDfield := "id"
	Namefield := "name"
	FullNamefield := "full_name"
	SearchNamefield := "search_name"
	Descriptionfield := "description"
	PerfpolicyNamefield := "perfpolicy_name"
	PerfpolicyIdfield := "perfpolicy_id"
	OwnedByGroupfield := "owned_by_group"
	OwnedByGroupIdfield := "owned_by_group_id"
	PoolNamefield := "pool_name"
	PoolIdfield := "pool_id"
	SerialNumberfield := "serial_number"
	SecondarySerialNumberfield := "secondary_serial_number"
	TargetNamefield := "target_name"
	ParentVolNamefield := "parent_vol_name"
	ParentVolIdfield := "parent_vol_id"
	BaseSnapNamefield := "base_snap_name"
	BaseSnapIdfield := "base_snap_id"
	VolcollNamefield := "volcoll_name"
	VolcollIdfield := "volcoll_id"
	DestPoolNamefield := "dest_pool_name"
	DestPoolIdfield := "dest_pool_id"
	AppUuidfield := "app_uuid"
	FolderIdfield := "folder_id"
	FolderNamefield := "folder_name"
	VpdT10field := "vpd_t10"
	VpdIeee0field := "vpd_ieee0"
	VpdIeee1field := "vpd_ieee1"
	AppCategoryfield := "app_category"
	PreFilterfield := "pre_filter"

	VolumeFields = &Volume{
		ID:                    &IDfield,
		Name:                  &Namefield,
		FullName:              &FullNamefield,
		SearchName:            &SearchNamefield,
		Description:           &Descriptionfield,
		PerfpolicyName:        &PerfpolicyNamefield,
		PerfpolicyId:          &PerfpolicyIdfield,
		OwnedByGroup:          &OwnedByGroupfield,
		OwnedByGroupId:        &OwnedByGroupIdfield,
		PoolName:              &PoolNamefield,
		PoolId:                &PoolIdfield,
		SerialNumber:          &SerialNumberfield,
		SecondarySerialNumber: &SecondarySerialNumberfield,
		TargetName:            &TargetNamefield,
		ParentVolName:         &ParentVolNamefield,
		ParentVolId:           &ParentVolIdfield,
		BaseSnapName:          &BaseSnapNamefield,
		BaseSnapId:            &BaseSnapIdfield,
		VolcollName:           &VolcollNamefield,
		VolcollId:             &VolcollIdfield,
		DestPoolName:          &DestPoolNamefield,
		DestPoolId:            &DestPoolIdfield,
		AppUuid:               &AppUuidfield,
		FolderId:              &FolderIdfield,
		FolderName:            &FolderNamefield,
		VpdT10:                &VpdT10field,
		VpdIeee0:              &VpdIeee0field,
		VpdIeee1:              &VpdIeee1field,
		AppCategory:           &AppCategoryfield,
		PreFilter:             &PreFilterfield,
	}
}

type Volume struct {
	// ID - Identifier for the volume.
	ID *string `json:"id,omitempty"`
	// Name - Name of the volume.
	Name *string `json:"name,omitempty"`
	// FullName - Fully qualified name of volume.
	FullName *string `json:"full_name,omitempty"`
	// SearchName - Name of volume used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// Size - Volume size in mebibytes. Size is required for creating a volume but not for cloning an existing volume.
	Size *int64 `json:"size,omitempty"`
	// Description - Text description of volume.
	Description *string `json:"description,omitempty"`
	// PerfpolicyName - Name of performance policy.
	PerfpolicyName *string `json:"perfpolicy_name,omitempty"`
	// PerfpolicyId - Identifier of the performance policy. After creating a volume, performance policy for the volume can only be changed to another performance policy with same block size.
	PerfpolicyId *string `json:"perfpolicy_id,omitempty"`
	// Reserve - Amount of space to reserve for this volume as a percentage of volume size.
	Reserve *int64 `json:"reserve,omitempty"`
	// WarnLevel - This attribute is deprecated. Alert threshold for the volume's mapped usage, expressed as a percentage of the volume's size. When the volume's mapped usage exceeds warn_level, the array issues an alert. If this option is not specified, array default volume warn level setting is used to decide the warning level for this volume.
	WarnLevel *int64 `json:"warn_level,omitempty"`
	// Limit - Limit on the volume's mapped usage, expressed as a percentage of the volume's size. When the volume's mapped usage exceeds limit, the volume will be offlined or made non-writable. If this option is not specified, array default volume limit setting is used to decide the limit for this volume.
	Limit *int64 `json:"limit,omitempty"`
	// SnapReserve - Amount of space to reserve for snapshots of this volume as a percentage of volume size.
	SnapReserve *int64 `json:"snap_reserve,omitempty"`
	// SnapWarnLevel - Threshold for available space as a percentage of volume size below which an alert is raised.
	SnapWarnLevel *int64 `json:"snap_warn_level,omitempty"`
	// SnapLimit - This attribute is deprecated. The array does not limit a volume's snapshot space usage. The attribute is ignored on input and returns max int64 on output.
	SnapLimit *int64 `json:"snap_limit,omitempty"`
	// SnapLimitPercent - This attribute is deprecated. The array does not limit a volume's snapshot space usage. The attribute is ignored on input and returns -1 on output.
	SnapLimitPercent *int64 `json:"snap_limit_percent,omitempty"`
	// NumSnaps - Number of live, non-hidden snapshots for this volume.
	NumSnaps *int64 `json:"num_snaps,omitempty"`
	// ProjectedNumSnaps - Deprecated. Projected number of snapshots (including scheduled and manual) for this volume.
	ProjectedNumSnaps *int64 `json:"projected_num_snaps,omitempty"`
	// Online - Online state of volume, available for host initiators to establish connections.
	Online *bool `json:"online,omitempty"`
	// OwnedByGroup - Name of group that currently owns the volume.
	OwnedByGroup *string `json:"owned_by_group,omitempty"`
	// OwnedByGroupId - ID of group that currently owns the volume.
	OwnedByGroupId *string `json:"owned_by_group_id,omitempty"`
	// MultiInitiator - For iSCSI Volume Target, this flag indicates whether the volume and its snapshots can be accessed from multiple initiators at the same time. The default is false. For iSCSI Group Target or FC access protocol, the attribute cannot be modified and always reads as false.
	MultiInitiator *bool `json:"multi_initiator,omitempty"`
	// IscsiTargetScope - This indicates whether volume is exported under iSCSI Group Target or iSCSI Volume Target. This attribute is only meaningful to iSCSI system. On FC system, all volumes are exported under the FC Group Target. In create operation, the volume's target type will be set by this attribute. If not specified, it will be set as the group-setting. In clone operation, the clone's target type will inherit from the parent' setting.
	IscsiTargetScope *NsTargetScope `json:"iscsi_target_scope,omitempty"`
	// PoolName - Name of the pool where the volume resides. Volume data will be distributed across arrays over which specified pool is defined. If pool option is not specified, volume is assigned to the default pool.
	PoolName *string `json:"pool_name,omitempty"`
	// PoolId - Identifier associated with the pool in the storage pool table.
	PoolId *string `json:"pool_id,omitempty"`
	// ReadOnly - Volume is read-only.
	ReadOnly *bool `json:"read_only,omitempty"`
	// SerialNumber - Identifier associated with the volume for the SCSI protocol.
	SerialNumber *string `json:"serial_number,omitempty"`
	// SecondarySerialNumber - Secondary identifier associated with the volume for the SCSI protocol.
	SecondarySerialNumber *string `json:"secondary_serial_number,omitempty"`
	// TargetName - The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target volume.
	TargetName *string `json:"target_name,omitempty"`
	// BlockSize - Size in bytes of blocks in the volume.
	BlockSize *int64 `json:"block_size,omitempty"`
	// OfflineReason - Volume offline reason.
	OfflineReason *NsOfflineReason `json:"offline_reason,omitempty"`
	// Clone - Whether this volume is a clone. Use this attribute in combination with name and base_snap_id to create a clone by setting clone = true.
	Clone *bool `json:"clone,omitempty"`
	// ParentVolName - Name of parent volume.
	ParentVolName *string `json:"parent_vol_name,omitempty"`
	// ParentVolId - Parent volume ID.
	ParentVolId *string `json:"parent_vol_id,omitempty"`
	// BaseSnapName - Name of base snapshot.
	BaseSnapName *string `json:"base_snap_name,omitempty"`
	// BaseSnapId - Base snapshot ID. This attribute is required together with name and clone when cloning a volume with the create operation.
	BaseSnapId *string `json:"base_snap_id,omitempty"`
	// ReplicationRole - Replication role that this volume performs.
	ReplicationRole *NsVolumeReplicationRole `json:"replication_role,omitempty"`
	// VolcollName - Name of volume collection of which this volume is a member.
	VolcollName *string `json:"volcoll_name,omitempty"`
	// VolcollId - ID of volume collection of which this volume is a member. Use this attribute in update operation to associate or dissociate volumes with or from volume collections. When associating, set this attribute to the ID of the volume collection. When dissociating, set this attribute to empty string.
	VolcollId *string `json:"volcoll_id,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// Force - Forcibly offline, reduce size or change read-only status a volume.
	Force *bool `json:"force,omitempty"`
	// CreationTime - Time when this volume was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this volume was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ProtectionType - Specifies if volume is protected with schedules. If protected, indicate whether replication is setup.
	ProtectionType *NsProtectionType `json:"protection_type,omitempty"`
	// LastSnap - Last snapshot for this volume.
	LastSnap *NsSnapSummary `json:"last_snap,omitempty"`
	// LastReplicatedSnap - Last replicated snapshot for this volume.
	LastReplicatedSnap *NsSnapSummary `json:"last_replicated_snap,omitempty"`
	// DestPoolName - Name of the destination pool where the volume is moving to.
	DestPoolName *string `json:"dest_pool_name,omitempty"`
	// DestPoolId - ID of the destination pool where the volume is moving to.
	DestPoolId *string `json:"dest_pool_id,omitempty"`
	// MoveStartTime - The Start time when this volume was moved.
	MoveStartTime *int64 `json:"move_start_time,omitempty"`
	// MoveAborting - This indicates whether the move of the volume is aborting or not.
	MoveAborting *bool `json:"move_aborting,omitempty"`
	// MoveBytesMigrated - The bytes of volume which have been moved.
	MoveBytesMigrated *int64 `json:"move_bytes_migrated,omitempty"`
	// MoveBytesRemaining - The bytes of volume which have not been moved.
	MoveBytesRemaining *int64 `json:"move_bytes_remaining,omitempty"`
	// MoveEstComplTime - The estimated time of completion of a move.
	MoveEstComplTime *int64 `json:"move_est_compl_time,omitempty"`
	// UsageValid - This indicates whether usage information of volume and snapshots are valid or not.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// SpaceUsageLevel - Indicates space usage level based on warning level.
	SpaceUsageLevel *NsSpaceUsageLevel `json:"space_usage_level,omitempty"`
	// TotalUsageBytes - Sum of volume mapped usage and uncompressed backup data(including pending deletes) in bytes of this volume.
	TotalUsageBytes *int64 `json:"total_usage_bytes,omitempty"`
	// VolUsageCompressedBytes - Compressed data in bytes for this volume.
	VolUsageCompressedBytes *int64 `json:"vol_usage_compressed_bytes,omitempty"`
	// VolUsageUncompressedBytes - Uncompressed data in bytes for this volume.
	VolUsageUncompressedBytes *int64 `json:"vol_usage_uncompressed_bytes,omitempty"`
	// VolUsageMappedBytes - Mapped data in bytes for this volume.
	VolUsageMappedBytes *int64 `json:"vol_usage_mapped_bytes,omitempty"`
	// SnapUsageCompressedBytes - Sum of compressed backup data in bytes stored in snapshots of this volume.
	SnapUsageCompressedBytes *int64 `json:"snap_usage_compressed_bytes,omitempty"`
	// SnapUsageUncompressedBytes - Sum of uncompressed unique backup data in bytes stored in snapshots of this volume.
	SnapUsageUncompressedBytes *int64 `json:"snap_usage_uncompressed_bytes,omitempty"`
	// SnapUsagePopulatedBytes - Sum of backup data in bytes stored in snapshots of this volume without accounting for the sharing of data between snapshots.
	SnapUsagePopulatedBytes *int64 `json:"snap_usage_populated_bytes,omitempty"`
	// CachePinned - If set to true, all the contents of this volume are kept in flash cache. This provides for consistent performance guarantees for all types of workloads. The amount of flash needed to pin the volume is equal to the limit for the volume.
	CachePinned *bool `json:"cache_pinned,omitempty"`
	// PinnedCacheSize - The amount of flash pinned on the volume.
	PinnedCacheSize *int64 `json:"pinned_cache_size,omitempty"`
	// CacheNeededForPin - The amount of flash needed to pin the volume.
	CacheNeededForPin *int64 `json:"cache_needed_for_pin,omitempty"`
	// UpstreamCachePinned - This indicates whether the upstream volume is cache pinned or not.
	UpstreamCachePinned *bool `json:"upstream_cache_pinned,omitempty"`
	// CachePolicy - Cache policy applied to the volume.
	CachePolicy *NsCachePolicy `json:"cache_policy,omitempty"`
	// ThinlyProvisioned - Set volume's provisioning level to thin.  Also advertises volume as thinly provisioned to initiators supporting thin provisioning. For such volumes, soft limit notification is set to initiators when the volume space usage crosses its volume_warn_level. Default is yes. The volume's space is provisioned immediately, but for advertising status, this change takes effect only for new connections to the volume. Initiators must disconnect and reconnect for the new setting to be take effect at the initiator level consistently.
	ThinlyProvisioned *bool `json:"thinly_provisioned,omitempty"`
	// VolState - Status of the volume.
	VolState *NsVolStatus `json:"vol_state,omitempty"`
	// OnlineSnaps - The list of online snapshots of this volume.
	OnlineSnaps []*NsSnapshotFromVolumes `json:"online_snaps,omitempty"`
	// NumConnections - Number of connections of this volume.
	NumConnections *int64 `json:"num_connections,omitempty"`
	// NumIscsiConnections - Number of iscsi connections of this volume.
	NumIscsiConnections *int64 `json:"num_iscsi_connections,omitempty"`
	// NumFcConnections - Number of Fibre Channel connections of this volume.
	NumFcConnections *int64 `json:"num_fc_connections,omitempty"`
	// AccessControlRecords - List of access control records that apply to this volume.
	AccessControlRecords []*NsAccessControlRecord `json:"access_control_records,omitempty"`
	// InheritAcl - In a volume clone operation, if both the parent and the clone have no external management agent (their agent_type property is "none"), then inherit_acl controls whether the clone will inherit a copy of the parent's access control list. If either the parent or the clone have an external management agent, then the clone will not inherit the parent's access control list.
	InheritAcl *bool `json:"inherit_acl,omitempty"`
	// EncryptionCipher - The encryption cipher of the volume.
	EncryptionCipher *NsEncryptionCipher `json:"encryption_cipher,omitempty"`
	// AppUuid - Application identifier of volume.
	AppUuid *string `json:"app_uuid,omitempty"`
	// FolderId - ID of the folder holding this volume.
	FolderId *string `json:"folder_id,omitempty"`
	// FolderName - Name of the folder holding this volume. It can be empty.
	FolderName *string `json:"folder_name,omitempty"`
	// Metadata - Key-value pairs that augment an volume's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// IscsiSessions - List of iSCSI sessions connected to this volume.
	IscsiSessions []*NsISCSISession `json:"iscsi_sessions,omitempty"`
	// FcSessions - List of Fibre Channel sessions connected to this volume.
	FcSessions []*NsFCSession `json:"fc_sessions,omitempty"`
	// CachingEnabled - Indicate caching the volume is enabled.
	CachingEnabled *bool `json:"caching_enabled,omitempty"`
	// PreviouslyDeduped - Indicate whether dedupe has ever been enabled on this volume.
	PreviouslyDeduped *bool `json:"previously_deduped,omitempty"`
	// DedupeEnabled - Indicate whether dedupe is enabled.
	DedupeEnabled *bool `json:"dedupe_enabled,omitempty"`
	// VpdT10 - The volume's T10 Vendor ID-based identifier.
	VpdT10 *string `json:"vpd_t10,omitempty"`
	// VpdIeee0 - The first 64 bits of the volume's EUI-64 identifier, encoded as a hexadecimal string.
	VpdIeee0 *string `json:"vpd_ieee0,omitempty"`
	// VpdIeee1 - The last 64 bits of the volume's EUI-64 identifier, encoded as a hexadecimal string.
	VpdIeee1 *string `json:"vpd_ieee1,omitempty"`
	// AppCategory - Application category that the volume belongs to.
	AppCategory *string `json:"app_category,omitempty"`
	// LimitIops - IOPS limit for this volume. If limit_iops is not specified when a volume is created, or if limit_iops is set to -1, then the volume has no IOPS limit. If limit_iops is not specified while creating a clone, IOPS limit of parent volume will be used as limit. IOPS limit should be in range [256, 4294967294] or -1 for unlimited. If both limit_iops and limit_mbps are specified, limit_mbps must not be hit before limit_iops. In other words, IOPS and MBPS limits should honor limit_iops <= ((limit_mbps MB/s * 2^20 B/MB) / block_size B).
	LimitIops *int64 `json:"limit_iops,omitempty"`
	// LimitMbps - Throughput limit for this volume in MB/s. If limit_mbps is not specified when a volume is created, or if limit_mbps is set to -1, then the volume has no MBPS limit. MBPS limit should be in range [1, 4294967294] or -1 for unlimited. If both limit_iops and limit_mbps are specified, limit_mbps must not be hit before limit_iops. In other words, IOPS and MBPS limits should honor limit_iops <= ((limit_mbps MB/s * 2^20 B/MB) / block_size B).
	LimitMbps *int64 `json:"limit_mbps,omitempty"`
	// NeedsContentRepl - Indicates whether the volume needs content based replication.
	NeedsContentRepl *bool `json:"needs_content_repl,omitempty"`
	// ContentReplErrorsFound - Indicates whether the last content based replication had errors.
	ContentReplErrorsFound *bool `json:"content_repl_errors_found,omitempty"`
	// LastContentSnapBrCgUid - The branch cg uid of the content based snapshot that was last replicated.
	LastContentSnapBrCgUid *int64 `json:"last_content_snap_br_cg_uid,omitempty"`
	// LastContentSnapBrGid - The branch gid of the content based snapshot that was last replicated.
	LastContentSnapBrGid *int64 `json:"last_content_snap_br_gid,omitempty"`
	// LastContentSnapId - The ID of the content based snapshot that was last replicated.
	LastContentSnapId *int64 `json:"last_content_snap_id,omitempty"`
	// CksumLastVerified - Last checksum verification time.
	CksumLastVerified *int64 `json:"cksum_last_verified,omitempty"`
	// PreFilter - Pre-filtering criteria.
	PreFilter *string `json:"pre_filter,omitempty"`
	// AvgStatsLast5mins - Average statistics in last 5 minutes.
	AvgStatsLast5mins *NsAverageStats `json:"avg_stats_last_5mins,omitempty"`
	// SrepLastSync - Time when synchronously replicated volume was last synchronized.
	SrepLastSync *int64 `json:"srep_last_sync,omitempty"`
	// SrepResyncPercent - Percentage of resync progress for synchronously replicated volume.
	SrepResyncPercent *int64 `json:"srep_resync_percent,omitempty"`
}
