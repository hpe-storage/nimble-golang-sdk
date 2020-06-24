/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// SnapshotCollection - Snapshot collections are collections of scheduled snapshots that are taken from volumes sharing a volume collection. Snapshot collections are replicated in the order that the collections were taken.
// Export SnapshotCollectionFields for advance operations like search filter etc.
var SnapshotCollectionFields *SnapshotCollection

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Descriptionfield:= "description"
	VolcollNamefield:= "volcoll_name"
	VolcollIDfield:= "volcoll_id"
	OriginNamefield:= "origin_name"
	SrepOwnerNamefield:= "srep_owner_name"
	SrepOwnerIDfield:= "srep_owner_id"
	PeerSnapcollIDfield:= "peer_snapcoll_id"
	ReplicateTofield:= "replicate_to"
	SchedIDfield:= "sched_id"
	SchedNamefield:= "sched_name"
		
	SnapshotCollectionFields= &SnapshotCollection{
		ID: &IDfield,
		Name: &Namefield,
		Description: &Descriptionfield,
		VolcollName: &VolcollNamefield,
		VolcollID: &VolcollIDfield,
		OriginName: &OriginNamefield,
		SrepOwnerName: &SrepOwnerNamefield,
		SrepOwnerID: &SrepOwnerIDfield,
		PeerSnapcollID: &PeerSnapcollIDfield,
		ReplicateTo: &ReplicateTofield,
		SchedID: &SchedIDfield,
		SchedName: &SchedNamefield,
		
	}
}

type SnapshotCollection struct {
   
    // Identifier for snapshot collection.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of snapshot collection.
    
 	Name *string `json:"name,omitempty"`
   
    // Text description of snapshot collection.
    
 	Description *string `json:"description,omitempty"`
   
    // Volume collection name.
    
 	VolcollName *string `json:"volcoll_name,omitempty"`
   
    // Parent volume collection ID.
    
 	VolcollID *string `json:"volcoll_id,omitempty"`
   
    // Origination group name/ID.
    
 	OriginName *string `json:"origin_name,omitempty"`
   
    // Indicates if snapshot collection was created as a replica.
    
 	IsReplica *bool `json:"is_replica,omitempty"`
   
    // Name of the partner where the snapshots in this snapshot collection reside.
    
 	SrepOwnerName *string `json:"srep_owner_name,omitempty"`
   
    // ID of the partner where snapshots for this snapshot collection reside which were created by synchronous replication. Field will be null if no peer snapshot_collection was created by synchronous replication.
    
 	SrepOwnerID *string `json:"srep_owner_id,omitempty"`
   
    // ID of the peer snapshot collection created by synchronous replication. Field will be null if no peer snapshot_collection was created by synchronous replication.
    
 	PeerSnapcollID *string `json:"peer_snapcoll_id,omitempty"`
   
    // Is complete.
    
 	IsComplete *bool `json:"is_complete,omitempty"`
   
    // Is manual.
    
 	IsManual *bool `json:"is_manual,omitempty"`
   
    // Is externally triggered.
    
 	IsExternalTrigger *bool `json:"is_external_trigger,omitempty"`
   
    // Indicates whether a snapshot collection is unmanaged. This is based on the state of individual snapshots.
    
 	IsUnmanaged *bool `json:"is_unmanaged,omitempty"`
   
    // Indicates whether a snapshot collection is managed.
    
 	IsManuallyManaged *bool `json:"is_manually_managed,omitempty"`
   
    // Replication status.
    
   	ReplStatus *NsSnapReplStatus `json:"repl_status,omitempty"`
   
    // Replication start time.
    
   	ReplStartTime *int64 `json:"repl_start_time,omitempty"`
   
    // Replication complete time.
    
   	ReplCompleteTime *int64 `json:"repl_complete_time,omitempty"`
   
    // Number of bytes transferred after replication completes.
    
   	ReplBytesTransferred *int64 `json:"repl_bytes_transferred,omitempty"`
   
    // Time when this snapshot collection was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this snapshot collection was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Online status of snapcoll. This is based on the online status of the individual snapshots.
    
   	OnlineStatus *NsOnlineStatus `json:"online_status,omitempty"`
   
    // List of snapshot attributes for snapshots being created as part of snapshot collection creation.
    
   	VolSnapAttrList []*NsVolumeSnapshotAttribute `json:"vol_snap_attr_list,omitempty"`
   
    // List of snapshots in the snapshot collection.
    
   	SnapshotsList []*NsSnapshotFromSnapshotCollections `json:"snapshots_list,omitempty"`
   
    // True if this snapshot collection has been marked for replication. This attribute cannot be updated for synchronous replication.
    
 	Replicate *bool `json:"replicate,omitempty"`
   
    // Specifies the partner name that the snapshots in this snapshot collection are replicated to.
    
 	ReplicateTo *string `json:"replicate_to,omitempty"`
   
    // Start with snapshot set online.
    
 	StartOnline *bool `json:"start_online,omitempty"`
   
    // Allow applications to write to created snapshot(s). Mandatory and must be set to 'true' for VSS application synchronized snapshots.
    
 	AllowWrites *bool `json:"allow_writes,omitempty"`
   
    // Do not perform application synchronization for this snapshot, create a crash-consistent snapshot instead.
    
 	DisableAppsync *bool `json:"disable_appsync,omitempty"`
   
    // Run verification tool on this snapshot. This option can only be used with a volume collection that has application synchronization.
    
 	SnapVerify *bool `json:"snap_verify,omitempty"`
   
    // Skip consistency check for database files on this snapshot. This option only applies to volume collections with application synchronization set to VSS, application ID set to MS Exchange 2010 or later with Database Availability Group (DAG), snap_verify option set to true, and disable_appsync option set to false.
    
 	SkipDbConsistencyCheck *bool `json:"skip_db_consistency_check,omitempty"`
   
    // ID of protection schedule of snapshot collection.
    
 	SchedID *string `json:"sched_id,omitempty"`
   
    // Name of protection schedule of snapshot collection.
    
 	SchedName *string `json:"sched_name,omitempty"`
   
    // Invoke snapshot request on upstream partner. This operation is not supported for synchronous replication volume vollections.
    
 	InvokeOnUpstreamPartner *bool `json:"invoke_on_upstream_partner,omitempty"`
   
    // External management agent type for snapshots being created as part of snapshot collection.
    
   	AgentType *NsAgentType `json:"agent_type,omitempty"`
   
    // Number of seconds after which this snapcoll is considered expired by the snapshot TTL. A value of 0 indicates that the snapshot never expires, 1 indicates that the snapshot uses a group-level configured TTL value and any other value indicates the number of seconds.
    
  	ExpiryAfter  *int64 `json:"expiry_after,omitempty"`
   
    // Key-value pairs that augment a snapshot collection's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
   
    // Forcibly delete the specified snapshot collection even if it is the last replicated snapshot. Doing so could lead to full re-seeding at the next replication.
    
 	Force *bool `json:"force,omitempty"`
}
