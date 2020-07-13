// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Snapshot - Snapshots are point-in-time copies of a volume. Snapshots are managed the same way you manage volumes. In reality, snapshots are volumes: they can be accessed by initiators, are subject to the same controls, can be modified, and have the same restrictions as volumes. Snapshots can be cloned and replicated. The initial snapshot uses no space: it shares the original data with the source volume. Each successive snapshot captures the changes that have occurred on the volume. The changed blocks are compressed.
// Export SnapshotFields for advance operations like search filter etc.
var SnapshotFields *Snapshot

func init() {

	SnapshotFields = &Snapshot{
		ID:                 "id",
		Name:               "name",
		Description:        "description",
		VolName:            "vol_name",
		PoolName:           "pool_name",
		VolId:              "vol_id",
		SnapCollectionName: "snap_collection_name",
		SnapCollectionId:   "snap_collection_id",
		OriginName:         "origin_name",
		SerialNumber:       "serial_number",
		TargetName:         "target_name",
		ScheduleName:       "schedule_name",
		ScheduleId:         "schedule_id",
		AppUuid:            "app_uuid",
		VpdT10:             "vpd_t10",
		VpdIeee0:           "vpd_ieee0",
		VpdIeee1:           "vpd_ieee1",
	}
}

type Snapshot struct {
	// ID - Identifier for the snapshot.
	ID string `json:"id,omitempty"`
	// Name - Name of snapshot.
	Name string `json:"name,omitempty"`
	// Description - Text description of snapshot.
	Description string `json:"description,omitempty"`
	// Size - Size of volume at time of snapshot (in bytes).
	Size int64 `json:"size,omitempty"`
	// VolName - Name of the parent volume in which the snapshot belongs to.
	VolName string `json:"vol_name,omitempty"`
	// PoolName - Name of the pool in which the parent volume belongs to.
	PoolName string `json:"pool_name,omitempty"`
	// VolId - Parent volume ID.
	VolId string `json:"vol_id,omitempty"`
	// SnapCollectionName - Name of snapshot collection.
	SnapCollectionName string `json:"snap_collection_name,omitempty"`
	// SnapCollectionId - Identifier of snapshot collection.
	SnapCollectionId string `json:"snap_collection_id,omitempty"`
	// Online - Online state for a snapshot means it could be mounted for data restore.
	Online *bool `json:"online,omitempty"`
	// Writable - Allow snapshot to be writable. Mandatory and must be set to 'true' for VSS application synchronized snapshots.
	Writable *bool `json:"writable,omitempty"`
	// OfflineReason - Snapshot offline reason - possible entries: one of 'user', 'recovery', 'replica', 'over_volume_limit', 'over_snapshot_limit', 'over_volume_reserve', 'nvram_loss_recovery', 'pool_free_space_exhausted' .
	OfflineReason *NsOfflineReason `json:"offline_reason,omitempty"`
	// ExpiryTime - Unix timestamp indicating that the snapshot is considered expired by Snapshot Time-to-live(TTL). A value of 0 indicates that snapshot never expires.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// ExpiryAfter - Number of seconds after which this snapshot is considered expired by snapshot TTL. A value of 0 indicates that snapshot never expires, 1 indicates that snapshot uses group-level configured TTL value and any other value indicates number of seconds.
	ExpiryAfter int64 `json:"expiry_after,omitempty"`
	// OriginName - Origination group name.
	OriginName string `json:"origin_name,omitempty"`
	// IsReplica - Snapshot is a replica from upstream replication partner.
	IsReplica *bool `json:"is_replica,omitempty"`
	// IsUnmanaged - Indicates whether the snapshot is unmanaged. The snapshot will not be deleted automatically unless the unmanaged cleanup feature is enabled.
	IsUnmanaged *bool `json:"is_unmanaged,omitempty"`
	// IsManuallyManaged - Is snapshot manually managed, i.e., snapshot is manually or third party created or created by system at the time of volume restore or resize.
	IsManuallyManaged *bool `json:"is_manually_managed,omitempty"`
	// ReplicationStatus - Replication status.
	ReplicationStatus *NsSnapReplStatus `json:"replication_status,omitempty"`
	// AccessControlRecords - List of access control records that apply to this snapshot.
	AccessControlRecords []*NsAccessControlRecord `json:"access_control_records,omitempty"`
	// SerialNumber - Identifier for the SCSI protocol.
	SerialNumber string `json:"serial_number,omitempty"`
	// TargetName - The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target snapshot.
	TargetName string `json:"target_name,omitempty"`
	// CreationTime - Time when this snapshot was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this snapshort was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// ScheduleName - Name of protection schedule.
	ScheduleName string `json:"schedule_name,omitempty"`
	// ScheduleId - Identifier of protection schedule.
	ScheduleId string `json:"schedule_id,omitempty"`
	// AppUuid - Application identifier of snapshots.
	AppUuid string `json:"app_uuid,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// NewDataValid - Indicate the usage infomation is valid.
	NewDataValid *bool `json:"new_data_valid,omitempty"`
	// NewDataCompressedBytes - The bytes of compressed new data.
	NewDataCompressedBytes int64 `json:"new_data_compressed_bytes,omitempty"`
	// NewDataUncompressedBytes - The bytes of uncompressed new data.
	NewDataUncompressedBytes int64 `json:"new_data_uncompressed_bytes,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// VpdT10 - The snapshot's T10 Vendor ID-based identifier.
	VpdT10 string `json:"vpd_t10,omitempty"`
	// VpdIeee0 - The first 64 bits of the snapshots's EUI-64 identifier, encoded as a hexadecimal string.
	VpdIeee0 string `json:"vpd_ieee0,omitempty"`
	// VpdIeee1 - The last 64 bits of the snapshots's EUI-64 identifier, encoded as a hexadecimal string.
	VpdIeee1 string `json:"vpd_ieee1,omitempty"`
	// Force - Forcibly delete the specified snapshot even if it is the last replicated collection. Doing so could lead to full re-seeding at the next replication.
	Force *bool `json:"force,omitempty"`
}

// sdk internal struct
type snapshot struct {
	// ID - Identifier for the snapshot.
	ID *string `json:"id,omitempty"`
	// Name - Name of snapshot.
	Name *string `json:"name,omitempty"`
	// Description - Text description of snapshot.
	Description *string `json:"description,omitempty"`
	// Size - Size of volume at time of snapshot (in bytes).
	Size *int64 `json:"size,omitempty"`
	// VolName - Name of the parent volume in which the snapshot belongs to.
	VolName *string `json:"vol_name,omitempty"`
	// PoolName - Name of the pool in which the parent volume belongs to.
	PoolName *string `json:"pool_name,omitempty"`
	// VolId - Parent volume ID.
	VolId *string `json:"vol_id,omitempty"`
	// SnapCollectionName - Name of snapshot collection.
	SnapCollectionName *string `json:"snap_collection_name,omitempty"`
	// SnapCollectionId - Identifier of snapshot collection.
	SnapCollectionId *string `json:"snap_collection_id,omitempty"`
	// Online - Online state for a snapshot means it could be mounted for data restore.
	Online *bool `json:"online,omitempty"`
	// Writable - Allow snapshot to be writable. Mandatory and must be set to 'true' for VSS application synchronized snapshots.
	Writable *bool `json:"writable,omitempty"`
	// OfflineReason - Snapshot offline reason - possible entries: one of 'user', 'recovery', 'replica', 'over_volume_limit', 'over_snapshot_limit', 'over_volume_reserve', 'nvram_loss_recovery', 'pool_free_space_exhausted' .
	OfflineReason *NsOfflineReason `json:"offline_reason,omitempty"`
	// ExpiryTime - Unix timestamp indicating that the snapshot is considered expired by Snapshot Time-to-live(TTL). A value of 0 indicates that snapshot never expires.
	ExpiryTime *int64 `json:"expiry_time,omitempty"`
	// ExpiryAfter - Number of seconds after which this snapshot is considered expired by snapshot TTL. A value of 0 indicates that snapshot never expires, 1 indicates that snapshot uses group-level configured TTL value and any other value indicates number of seconds.
	ExpiryAfter *int64 `json:"expiry_after,omitempty"`
	// OriginName - Origination group name.
	OriginName *string `json:"origin_name,omitempty"`
	// IsReplica - Snapshot is a replica from upstream replication partner.
	IsReplica *bool `json:"is_replica,omitempty"`
	// IsUnmanaged - Indicates whether the snapshot is unmanaged. The snapshot will not be deleted automatically unless the unmanaged cleanup feature is enabled.
	IsUnmanaged *bool `json:"is_unmanaged,omitempty"`
	// IsManuallyManaged - Is snapshot manually managed, i.e., snapshot is manually or third party created or created by system at the time of volume restore or resize.
	IsManuallyManaged *bool `json:"is_manually_managed,omitempty"`
	// ReplicationStatus - Replication status.
	ReplicationStatus *NsSnapReplStatus `json:"replication_status,omitempty"`
	// AccessControlRecords - List of access control records that apply to this snapshot.
	AccessControlRecords []*NsAccessControlRecord `json:"access_control_records,omitempty"`
	// SerialNumber - Identifier for the SCSI protocol.
	SerialNumber *string `json:"serial_number,omitempty"`
	// TargetName - The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target snapshot.
	TargetName *string `json:"target_name,omitempty"`
	// CreationTime - Time when this snapshot was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this snapshort was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// ScheduleName - Name of protection schedule.
	ScheduleName *string `json:"schedule_name,omitempty"`
	// ScheduleId - Identifier of protection schedule.
	ScheduleId *string `json:"schedule_id,omitempty"`
	// AppUuid - Application identifier of snapshots.
	AppUuid *string `json:"app_uuid,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// NewDataValid - Indicate the usage infomation is valid.
	NewDataValid *bool `json:"new_data_valid,omitempty"`
	// NewDataCompressedBytes - The bytes of compressed new data.
	NewDataCompressedBytes *int64 `json:"new_data_compressed_bytes,omitempty"`
	// NewDataUncompressedBytes - The bytes of uncompressed new data.
	NewDataUncompressedBytes *int64 `json:"new_data_uncompressed_bytes,omitempty"`
	// AgentType - External management agent type.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// VpdT10 - The snapshot's T10 Vendor ID-based identifier.
	VpdT10 *string `json:"vpd_t10,omitempty"`
	// VpdIeee0 - The first 64 bits of the snapshots's EUI-64 identifier, encoded as a hexadecimal string.
	VpdIeee0 *string `json:"vpd_ieee0,omitempty"`
	// VpdIeee1 - The last 64 bits of the snapshots's EUI-64 identifier, encoded as a hexadecimal string.
	VpdIeee1 *string `json:"vpd_ieee1,omitempty"`
	// Force - Forcibly delete the specified snapshot even if it is the last replicated collection. Doing so could lead to full re-seeding at the next replication.
	Force *bool `json:"force,omitempty"`
}

// EncodeSnapshot - Transform Snapshot to snapshot type
func EncodeSnapshot(request interface{}) (*snapshot, error) {
	reqSnapshot := request.(*Snapshot)
	byte, err := json.Marshal(reqSnapshot)
	objPtr := &snapshot{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeSnapshot Transform snapshot to Snapshot type
func DecodeSnapshot(request interface{}) (*Snapshot, error) {
	reqSnapshot := request.(*snapshot)
	byte, err := json.Marshal(reqSnapshot)
	obj := &Snapshot{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
