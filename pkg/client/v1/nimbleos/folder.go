// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export FolderFields provides field names to use in filter parameters, for example.
var FolderFields *FolderFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldFqn := "fqn"
	fieldFullName := "full_name"
	fieldSearchName := "search_name"
	fieldDescription := "description"
	fieldPoolName := "pool_name"
	fieldPoolId := "pool_id"
	fieldLimitBytesSpecified := "limit_bytes_specified"
	fieldLimitBytes := "limit_bytes"
	fieldLimitSizeBytes := "limit_size_bytes"
	fieldProvisionedLimitSizeBytes := "provisioned_limit_size_bytes"
	fieldOverdraftLimitPct := "overdraft_limit_pct"
	fieldCapacityBytes := "capacity_bytes"
	fieldFreeSpaceBytes := "free_space_bytes"
	fieldProvisionedBytes := "provisioned_bytes"
	fieldUsageBytes := "usage_bytes"
	fieldVolumeMappedBytes := "volume_mapped_bytes"
	fieldUsageValid := "usage_valid"
	fieldAgentType := "agent_type"
	fieldInheritedVolPerfpolId := "inherited_vol_perfpol_id"
	fieldInheritedVolPerfpolName := "inherited_vol_perfpol_name"
	fieldUnusedReserveBytes := "unused_reserve_bytes"
	fieldUnusedSnapReserveBytes := "unused_snap_reserve_bytes"
	fieldCompressedVolUsageBytes := "compressed_vol_usage_bytes"
	fieldCompressedSnapUsageBytes := "compressed_snap_usage_bytes"
	fieldUncompressedVolUsageBytes := "uncompressed_vol_usage_bytes"
	fieldUncompressedSnapUsageBytes := "uncompressed_snap_usage_bytes"
	fieldVolCompressionRatio := "vol_compression_ratio"
	fieldSnapCompressionRatio := "snap_compression_ratio"
	fieldCompressionRatio := "compression_ratio"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"
	fieldNumSnaps := "num_snaps"
	fieldNumSnapcolls := "num_snapcolls"
	fieldAppUuid := "app_uuid"
	fieldVolumeList := "volume_list"
	fieldAppserverId := "appserver_id"
	fieldAppserverName := "appserver_name"
	fieldFolsetId := "folset_id"
	fieldFolsetName := "folset_name"
	fieldLimitIops := "limit_iops"
	fieldLimitMbps := "limit_mbps"
	fieldAccessProtocol := "access_protocol"
	fieldTenantId := "tenant_id"

	FolderFields = &FolderFieldHandles{
		ID:                         &fieldID,
		Name:                       &fieldName,
		Fqn:                        &fieldFqn,
		FullName:                   &fieldFullName,
		SearchName:                 &fieldSearchName,
		Description:                &fieldDescription,
		PoolName:                   &fieldPoolName,
		PoolId:                     &fieldPoolId,
		LimitBytesSpecified:        &fieldLimitBytesSpecified,
		LimitBytes:                 &fieldLimitBytes,
		LimitSizeBytes:             &fieldLimitSizeBytes,
		ProvisionedLimitSizeBytes:  &fieldProvisionedLimitSizeBytes,
		OverdraftLimitPct:          &fieldOverdraftLimitPct,
		CapacityBytes:              &fieldCapacityBytes,
		FreeSpaceBytes:             &fieldFreeSpaceBytes,
		ProvisionedBytes:           &fieldProvisionedBytes,
		UsageBytes:                 &fieldUsageBytes,
		VolumeMappedBytes:          &fieldVolumeMappedBytes,
		UsageValid:                 &fieldUsageValid,
		AgentType:                  &fieldAgentType,
		InheritedVolPerfpolId:      &fieldInheritedVolPerfpolId,
		InheritedVolPerfpolName:    &fieldInheritedVolPerfpolName,
		UnusedReserveBytes:         &fieldUnusedReserveBytes,
		UnusedSnapReserveBytes:     &fieldUnusedSnapReserveBytes,
		CompressedVolUsageBytes:    &fieldCompressedVolUsageBytes,
		CompressedSnapUsageBytes:   &fieldCompressedSnapUsageBytes,
		UncompressedVolUsageBytes:  &fieldUncompressedVolUsageBytes,
		UncompressedSnapUsageBytes: &fieldUncompressedSnapUsageBytes,
		VolCompressionRatio:        &fieldVolCompressionRatio,
		SnapCompressionRatio:       &fieldSnapCompressionRatio,
		CompressionRatio:           &fieldCompressionRatio,
		CreationTime:               &fieldCreationTime,
		LastModified:               &fieldLastModified,
		NumSnaps:                   &fieldNumSnaps,
		NumSnapcolls:               &fieldNumSnapcolls,
		AppUuid:                    &fieldAppUuid,
		VolumeList:                 &fieldVolumeList,
		AppserverId:                &fieldAppserverId,
		AppserverName:              &fieldAppserverName,
		FolsetId:                   &fieldFolsetId,
		FolsetName:                 &fieldFolsetName,
		LimitIops:                  &fieldLimitIops,
		LimitMbps:                  &fieldLimitMbps,
		AccessProtocol:             &fieldAccessProtocol,
		TenantId:                   &fieldTenantId,
	}
}

// Folder - Folders are a way to group volumes, as well as a way to apply space constraints to them.
type Folder struct {
	// ID - Identifier for the folder.
	ID *string `json:"id,omitempty"`
	// Name - Name of folder.
	Name *string `json:"name,omitempty"`
	// Fqn - Fully qualified name of folder in the pool.
	Fqn *string `json:"fqn,omitempty"`
	// FullName - Fully qualified name of folder in the group.
	FullName *string `json:"full_name,omitempty"`
	// SearchName - Name of folder used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// Description - Text description of folder.
	Description *string `json:"description,omitempty"`
	// PoolName - Name of the pool where the folder resides.
	PoolName *string `json:"pool_name,omitempty"`
	// PoolId - ID of the pool where the folder resides.
	PoolId *string `json:"pool_id,omitempty"`
	// LimitBytesSpecified - Indicates whether the folder has a limit.
	LimitBytesSpecified *bool `json:"limit_bytes_specified,omitempty"`
	// LimitBytes - Folder limit size in bytes. By default, a folder (except SMIS and VVol types) does not have a limit. If limit_bytes is not specified when a folder is created, or if limit_bytes is set to the largest possible 64-bit signed integer (9223372036854775807), then the folder has no limit. Otherwise, a limit smaller than the capacity of the pool can be set. On output, if the folder has a limit, the limit_bytes_specified attribute will be true and limit_bytes will be the limit. If the folder does not have a limit, the limit_bytes_specified attribute will be false and limit_bytes will be interpreted based on the value of the usage_valid attribute. If the usage_valid attribute is true, limits_byte will be the capacity of the pool. Otherwise, limits_bytes is not meaningful and can be null. SMIS and VVol folders require a size limit. This attribute is superseded by limit_size_bytes.
	LimitBytes *int64 `json:"limit_bytes,omitempty"`
	// LimitSizeBytes - Folder size limit in bytes. If limit_size_bytes is not specified when a folder is created, or if limit_size_bytes is set to -1, then the folder has no limit. Otherwise, a limit smaller than the capacity of the pool can be set. Folders with an agent_type of 'smis' or 'vvol' must have a size limit.
	LimitSizeBytes *int64 `json:"limit_size_bytes,omitempty"`
	// ProvisionedLimitSizeBytes - Limit on the provisioned size of volumes in a folder. If provisioned_limit_size_bytes is not specified when a folder is created, or if provisioned_limit_size_bytes is set to -1, then the folder has no provisioned size limit.
	ProvisionedLimitSizeBytes *int64 `json:"provisioned_limit_size_bytes,omitempty"`
	// OverdraftLimitPct - Amount of space to consider as overdraft range for this folder as a percentage of folder used limit. Valid values are from 0% - 200%. This is the limit above the folder usage limit beyond which enforcement action(volume offline/non-writable) is issued.
	OverdraftLimitPct *int64 `json:"overdraft_limit_pct,omitempty"`
	// CapacityBytes - Capacity of the folder in bytes. If the folder's size has a usage limit, capacity_bytes will be the folder's usage limit. If the folder's size does not have a usage limit, capacity_bytes will be the pool's capacity. This field is meaningful only when the usage_valid attribute is true.
	CapacityBytes *int64 `json:"capacity_bytes,omitempty"`
	// FreeSpaceBytes - Free space in the folder in bytes. If the folder has a usage limit, free_space_bytes will be the folder's free space (the folder's usage limit minus the folder's space usage). If the folder does not have a usage limit, free_space_bytes will be the pool's free space. This field is meaningful only when the usage_valid attribute is true.
	FreeSpaceBytes *int64 `json:"free_space_bytes,omitempty"`
	// ProvisionedBytes - Sum of provisioned size of volumes in the folder.
	ProvisionedBytes *int64 `json:"provisioned_bytes,omitempty"`
	// UsageBytes - Sum of mapped usage and snapshot uncompressed usage of volumes in the folder.
	UsageBytes *int64 `json:"usage_bytes,omitempty"`
	// VolumeMappedBytes - Sum of mapped usage of volumes in the folder.
	VolumeMappedBytes *int64 `json:"volume_mapped_bytes,omitempty"`
	// UsageValid - Indicate whether the space usage attributes of folder are valid.
	UsageValid *bool `json:"usage_valid,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsFolderAgentType `json:"agent_type,omitempty"`
	// InheritedVolPerfpolId - Identifier of the default performance policy for a newly created volume.
	InheritedVolPerfpolId *string `json:"inherited_vol_perfpol_id,omitempty"`
	// InheritedVolPerfpolName - Name of the default performance policy for a newly created volume.
	InheritedVolPerfpolName *string `json:"inherited_vol_perfpol_name,omitempty"`
	// UnusedReserveBytes - Unused reserve of volumes in the folder in bytes. This field is meaningful only when the usage_valid attribute is true.
	UnusedReserveBytes *int64 `json:"unused_reserve_bytes,omitempty"`
	// UnusedSnapReserveBytes - Unused reserve of snapshots of volumes in the folder in bytes. This field is meaningful only when the usage_valid attribute is true.
	UnusedSnapReserveBytes *int64 `json:"unused_snap_reserve_bytes,omitempty"`
	// CompressedVolUsageBytes - Compressed usage of volumes in the folder. This field is meaningful only when the usage_valid attribute is true.
	CompressedVolUsageBytes *int64 `json:"compressed_vol_usage_bytes,omitempty"`
	// CompressedSnapUsageBytes - Compressed usage of snapshots in the folder. This field is meaningful only when the usage_valid attribute is true.
	CompressedSnapUsageBytes *int64 `json:"compressed_snap_usage_bytes,omitempty"`
	// UncompressedVolUsageBytes - Uncompressed usage of volumes in the folder. This field is meaningful only when the usage_valid attribute is true.
	UncompressedVolUsageBytes *int64 `json:"uncompressed_vol_usage_bytes,omitempty"`
	// UncompressedSnapUsageBytes - Uncompressed usage of snapshots in the folder. This field is meaningful only when the usage_valid attribute is true.
	UncompressedSnapUsageBytes *int64 `json:"uncompressed_snap_usage_bytes,omitempty"`
	// VolCompressionRatio - Compression ratio of volumes in the folder. This field is meaningful only when the usage_valid attribute is true.
	VolCompressionRatio *float64 `json:"vol_compression_ratio,omitempty"`
	// SnapCompressionRatio - Compression ratio of snapshots in the folder. This field is meaningful only when the usage_valid attribute is true.
	SnapCompressionRatio *float64 `json:"snap_compression_ratio,omitempty"`
	// CompressionRatio - Compression savings for the folder expressed as ratio. This field is meaningful only when the usage_valid attribute is true.
	CompressionRatio *float64 `json:"compression_ratio,omitempty"`
	// CreationTime - Time when this folder was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this folder was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// NumSnaps - Number of snapshots inside the folder. This attribute is deprecated and has no meaningful value.
	NumSnaps *int64 `json:"num_snaps,omitempty"`
	// NumSnapcolls - Number of snapshot collections inside the folder. This attribute is deprecated and has no meaningful value.
	NumSnapcolls *int64 `json:"num_snapcolls,omitempty"`
	// AppUuid - Application identifier of the folder.
	AppUuid *string `json:"app_uuid,omitempty"`
	// VolumeList - List of volumes contained by the folder.
	VolumeList []*NsVolumeSummary `json:"volume_list,omitempty"`
	// AppserverId - Identifier of the application server associated with the folder.
	AppserverId *string `json:"appserver_id,omitempty"`
	// AppserverName - Name of the application server associated with the folder.
	AppserverName *string `json:"appserver_name,omitempty"`
	// FolsetId - Identifier of the folder set associated with the folder. Only VVol folder can be associated with the folder set. The folder and the containing folder set must be associated with the same application server.
	FolsetId *string `json:"folset_id,omitempty"`
	// FolsetName - Name of the folder set associated with the folder. Only VVol folder can be associated with the folder set. The folder and the containing folder set must be associated with the same application server.
	FolsetName *string `json:"folset_name,omitempty"`
	// LimitIops - IOPS limit for this folder. If limit_iops is not specified when a folder is created, or if limit_iops is set to -1, then the folder has no IOPS limit. IOPS limit should be in range [256, 4294967294] or -1 for unlimited.
	LimitIops *int64 `json:"limit_iops,omitempty"`
	// LimitMbps - Throughput limit for this folder in MB/s. If limit_mbps is not specified when a folder is created, or if limit_mbps is set to -1, then the folder has no throughput limit. MBPS limit should be in range [1, 4294967294] or -1 for unlimited.
	LimitMbps *int64 `json:"limit_mbps,omitempty"`
	// AccessProtocol - Access protocol of the folder. This attribute is used by the VASA Provider to determine the access protocol of the bind request. If not specified in the creation request, it will be the access protocol supported by the group. If the group supports multiple protocols, the default will be Fibre Channel. This field is meaningful only to VVol folder.
	AccessProtocol *NsAccessProtocol `json:"access_protocol,omitempty"`
	// TenantId - Tenant ID of the folder. This is used to determine what tenant context the folder belongs to.
	TenantId *string `json:"tenant_id,omitempty"`
}

// FolderFieldHandles provides a string representation for each AccessControlRecord field.
type FolderFieldHandles struct {
	ID                         *string
	Name                       *string
	Fqn                        *string
	FullName                   *string
	SearchName                 *string
	Description                *string
	PoolName                   *string
	PoolId                     *string
	LimitBytesSpecified        *string
	LimitBytes                 *string
	LimitSizeBytes             *string
	ProvisionedLimitSizeBytes  *string
	OverdraftLimitPct          *string
	CapacityBytes              *string
	FreeSpaceBytes             *string
	ProvisionedBytes           *string
	UsageBytes                 *string
	VolumeMappedBytes          *string
	UsageValid                 *string
	AgentType                  *string
	InheritedVolPerfpolId      *string
	InheritedVolPerfpolName    *string
	UnusedReserveBytes         *string
	UnusedSnapReserveBytes     *string
	CompressedVolUsageBytes    *string
	CompressedSnapUsageBytes   *string
	UncompressedVolUsageBytes  *string
	UncompressedSnapUsageBytes *string
	VolCompressionRatio        *string
	SnapCompressionRatio       *string
	CompressionRatio           *string
	CreationTime               *string
	LastModified               *string
	NumSnaps                   *string
	NumSnapcolls               *string
	AppUuid                    *string
	VolumeList                 *string
	AppserverId                *string
	AppserverName              *string
	FolsetId                   *string
	FolsetName                 *string
	LimitIops                  *string
	LimitMbps                  *string
	AccessProtocol             *string
	TenantId                   *string
}
