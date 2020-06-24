/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Snapshot - Snapshots are point-in-time copies of a volume. Snapshots are managed the same way you manage volumes. In reality, snapshots are volumes: they can be accessed by initiators, are subject to the same controls, can be modified, and have the same restrictions as volumes. Snapshots can be cloned and replicated. The initial snapshot uses no space: it shares the original data with the source volume. Each successive snapshot captures the changes that have occurred on the volume. The changed blocks are compressed.
// Export SnapshotFields for advance operations like search filter etc.
var SnapshotFields *Snapshot

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Descriptionfield:= "description"
	VolNamefield:= "vol_name"
	PoolNamefield:= "pool_name"
	VolIDfield:= "vol_id"
	SnapCollectionNamefield:= "snap_collection_name"
	SnapCollectionIDfield:= "snap_collection_id"
	OriginNamefield:= "origin_name"
	SerialNumberfield:= "serial_number"
	TargetNamefield:= "target_name"
	ScheduleNamefield:= "schedule_name"
	ScheduleIDfield:= "schedule_id"
	AppUuIDfield:= "app_uuid"
	VpdT10field:= "vpd_t10"
	VpdIeee0field:= "vpd_ieee0"
	VpdIeee1field:= "vpd_ieee1"
		
	SnapshotFields= &Snapshot{
		ID: &IDfield,
		Name: &Namefield,
		Description: &Descriptionfield,
		VolName: &VolNamefield,
		PoolName: &PoolNamefield,
		VolID: &VolIDfield,
		SnapCollectionName: &SnapCollectionNamefield,
		SnapCollectionID: &SnapCollectionIDfield,
		OriginName: &OriginNamefield,
		SerialNumber: &SerialNumberfield,
		TargetName: &TargetNamefield,
		ScheduleName: &ScheduleNamefield,
		ScheduleID: &ScheduleIDfield,
		AppUuID: &AppUuIDfield,
		VpdT10: &VpdT10field,
		VpdIeee0: &VpdIeee0field,
		VpdIeee1: &VpdIeee1field,
		
	}
}

type Snapshot struct {
   
    // Identifier for the snapshot.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of snapshot.
    
 	Name *string `json:"name,omitempty"`
   
    // Text description of snapshot.
    
 	Description *string `json:"description,omitempty"`
   
    // Size of volume at time of snapshot (in bytes).
    
   	Size *int64 `json:"size,omitempty"`
   
    // Name of the parent volume in which the snapshot belongs to.
    
 	VolName *string `json:"vol_name,omitempty"`
   
    // Name of the pool in which the parent volume belongs to.
    
 	PoolName *string `json:"pool_name,omitempty"`
   
    // Parent volume ID.
    
 	VolID *string `json:"vol_id,omitempty"`
   
    // Name of snapshot collection.
    
 	SnapCollectionName *string `json:"snap_collection_name,omitempty"`
   
    // Identifier of snapshot collection.
    
 	SnapCollectionID *string `json:"snap_collection_id,omitempty"`
   
    // Online state for a snapshot means it could be mounted for data restore.
    
 	Online *bool `json:"online,omitempty"`
   
    // Allow snapshot to be writable. Mandatory and must be set to 'true' for VSS application synchronized snapshots.
    
 	Writable *bool `json:"writable,omitempty"`
   
    // Snapshot offline reason - possible entries: one of 'user', 'recovery', 'replica', 'over_volume_limit', 'over_snapshot_limit', 'over_volume_reserve', 'nvram_loss_recovery', 'pool_free_space_exhausted' .
    
   	OfflineReason *NsOfflineReason `json:"offline_reason,omitempty"`
   
    // Unix timestamp indicating that the snapshot is considered expired by Snapshot Time-to-live(TTL). A value of 0 indicates that snapshot never expires.
    
   	ExpiryTime *int64 `json:"expiry_time,omitempty"`
   
    // Number of seconds after which this snapshot is considered expired by snapshot TTL. A value of 0 indicates that snapshot never expires, 1 indicates that snapshot uses group-level configured TTL value and any other value indicates number of seconds.
    
  	ExpiryAfter  *int64 `json:"expiry_after,omitempty"`
   
    // Origination group name.
    
 	OriginName *string `json:"origin_name,omitempty"`
   
    // Snapshot is a replica from upstream replication partner.
    
 	IsReplica *bool `json:"is_replica,omitempty"`
   
    // Indicates whether the snapshot is unmanaged. The snapshot will not be deleted automatically unless the unmanaged cleanup feature is enabled.
    
 	IsUnmanaged *bool `json:"is_unmanaged,omitempty"`
   
    // Is snapshot manually managed, i.e., snapshot is manually or third party created or created by system at the time of volume restore or resize.
    
 	IsManuallyManaged *bool `json:"is_manually_managed,omitempty"`
   
    // Replication status.
    
   	ReplicationStatus *NsSnapReplStatus `json:"replication_status,omitempty"`
   
    // List of access control records that apply to this snapshot.
    
   	AccessControlRecords []*NsAccessControlRecord `json:"access_control_records,omitempty"`
   
    // Identifier for the SCSI protocol.
    
 	SerialNumber *string `json:"serial_number,omitempty"`
   
    // The iSCSI Qualified Name (IQN) or the Fibre Channel World Wide Node Name (WWNN) of the target snapshot.
    
 	TargetName *string `json:"target_name,omitempty"`
   
    // Time when this snapshot was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this snapshort was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Name of protection schedule.
    
 	ScheduleName *string `json:"schedule_name,omitempty"`
   
    // Identifier of protection schedule.
    
 	ScheduleID *string `json:"schedule_id,omitempty"`
   
    // Application identifier of snapshots.
    
 	AppUuID *string `json:"app_uuid,omitempty"`
   
    // Key-value pairs that augment a snapshot's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
   
    // Indicate the usage infomation is valid.
    
 	NewDataValID *bool `json:"new_data_valid,omitempty"`
   
    // The bytes of compressed new data.
    
   	NewDataCompressedBytes *int64 `json:"new_data_compressed_bytes,omitempty"`
   
    // The bytes of uncompressed new data.
    
   	NewDataUncompressedBytes *int64 `json:"new_data_uncompressed_bytes,omitempty"`
   
    // External management agent type.
    
   	AgentType *NsAgentType `json:"agent_type,omitempty"`
   
    // The snapshot's T10 Vendor ID-based identifier.
    
 	VpdT10 *string `json:"vpd_t10,omitempty"`
   
    // The first 64 bits of the snapshots's EUI-64 identifier, encoded as a hexadecimal string.
    
 	VpdIeee0 *string `json:"vpd_ieee0,omitempty"`
   
    // The last 64 bits of the snapshots's EUI-64 identifier, encoded as a hexadecimal string.
    
 	VpdIeee1 *string `json:"vpd_ieee1,omitempty"`
   
    // Forcibly delete the specified snapshot even if it is the last replicated collection. Doing so could lead to full re-seeding at the next replication.
    
 	Force *bool `json:"force,omitempty"`
}
