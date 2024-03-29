// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// SnapshotCollectionFields provides field names to use in filter parameters, for example.
var SnapshotCollectionFields *SnapshotCollectionFieldHandles

func init() {
	SnapshotCollectionFields = &SnapshotCollectionFieldHandles{
		ID:                      "id",
		Name:                    "name",
		Description:             "description",
		VolcollName:             "volcoll_name",
		VolcollId:               "volcoll_id",
		OriginName:              "origin_name",
		IsReplica:               "is_replica",
		SrepOwnerName:           "srep_owner_name",
		SrepOwnerId:             "srep_owner_id",
		PeerSnapcollId:          "peer_snapcoll_id",
		NumSnaps:                "num_snaps",
		IsComplete:              "is_complete",
		IsManual:                "is_manual",
		IsExternalTrigger:       "is_external_trigger",
		IsUnmanaged:             "is_unmanaged",
		IsManuallyManaged:       "is_manually_managed",
		ReplStatus:              "repl_status",
		ReplStartTime:           "repl_start_time",
		ReplCompleteTime:        "repl_complete_time",
		ReplBytesTransferred:    "repl_bytes_transferred",
		CreationTime:            "creation_time",
		LastModified:            "last_modified",
		OnlineStatus:            "online_status",
		VolSnapAttrList:         "vol_snap_attr_list",
		SnapshotsList:           "snapshots_list",
		Replicate:               "replicate",
		ReplicateTo:             "replicate_to",
		StartOnline:             "start_online",
		AllowWrites:             "allow_writes",
		DisableAppsync:          "disable_appsync",
		SnapVerify:              "snap_verify",
		SkipDbConsistencyCheck:  "skip_db_consistency_check",
		SchedId:                 "sched_id",
		SchedName:               "sched_name",
		InvokeOnUpstreamPartner: "invoke_on_upstream_partner",
		AgentType:               "agent_type",
		ExpiryAfter:             "expiry_after",
		Metadata:                "metadata",
		Force:                   "force",
	}
}

// SnapshotCollection - Snapshot collections are collections of scheduled snapshots that are taken from volumes sharing a volume collection. Snapshot collections are replicated in the order that the collections were taken.
type SnapshotCollection struct {
	// ID - Identifier for snapshot collection.
	ID *string `json:"id,omitempty"`
	// Name - Name of snapshot collection.
	Name *string `json:"name,omitempty"`
	// Description - Text description of snapshot collection.
	Description *string `json:"description,omitempty"`
	// VolcollName - Volume collection name.
	VolcollName *string `json:"volcoll_name,omitempty"`
	// VolcollId - Parent volume collection ID.
	VolcollId *string `json:"volcoll_id,omitempty"`
	// OriginName - Origination group name/ID.
	OriginName *string `json:"origin_name,omitempty"`
	// IsReplica - Indicates if snapshot collection was created as a replica.
	IsReplica *bool `json:"is_replica,omitempty"`
	// SrepOwnerName - Name of the partner where the snapshots in this snapshot collection reside.
	SrepOwnerName *string `json:"srep_owner_name,omitempty"`
	// SrepOwnerId - ID of the partner where snapshots for this snapshot collection reside which were created by synchronous replication. Field will be null if no peer snapshot_collection was created by synchronous replication.
	SrepOwnerId *string `json:"srep_owner_id,omitempty"`
	// PeerSnapcollId - ID of the peer snapshot collection created by synchronous replication. Field will be null if no peer snapshot_collection was created by synchronous replication.
	PeerSnapcollId *string `json:"peer_snapcoll_id,omitempty"`
	// NumSnaps - Current number of live, non-hidden snaps in this collection.
	NumSnaps *int64 `json:"num_snaps,omitempty"`
	// IsComplete - Is complete.
	IsComplete *bool `json:"is_complete,omitempty"`
	// IsManual - Is manual.
	IsManual *bool `json:"is_manual,omitempty"`
	// IsExternalTrigger - Is externally triggered.
	IsExternalTrigger *bool `json:"is_external_trigger,omitempty"`
	// IsUnmanaged - Indicates whether a snapshot collection is unmanaged. This is based on the state of individual snapshots.
	IsUnmanaged *bool `json:"is_unmanaged,omitempty"`
	// IsManuallyManaged - Indicates whether a snapshot collection is managed.
	IsManuallyManaged *bool `json:"is_manually_managed,omitempty"`
	// ReplStatus - Replication status.
	ReplStatus *NsSnapReplStatus `json:"repl_status,omitempty"`
	// ReplStartTime - Replication start time.
	ReplStartTime *int64 `json:"repl_start_time,omitempty"`
	// ReplCompleteTime - Replication complete time.
	ReplCompleteTime *int64 `json:"repl_complete_time,omitempty"`
	// ReplBytesTransferred - Number of bytes transferred after replication completes.
	ReplBytesTransferred *int64 `json:"repl_bytes_transferred,omitempty"`
	// CreationTime - Time when this snapshot collection was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this snapshot collection was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// OnlineStatus - Online status of snapcoll. This is based on the online status of the individual snapshots.
	OnlineStatus *NsOnlineStatus `json:"online_status,omitempty"`
	// VolSnapAttrList - List of snapshot attributes for snapshots being created as part of snapshot collection creation.
	VolSnapAttrList []*NsVolumeSnapshotAttribute `json:"vol_snap_attr_list,omitempty"`
	// SnapshotsList - List of snapshots in the snapshot collection.
	SnapshotsList []*NsSnapshotFromSnapshotCollections `json:"snapshots_list,omitempty"`
	// Replicate - True if this snapshot collection has been marked for replication. This attribute cannot be updated for synchronous replication.
	Replicate *bool `json:"replicate,omitempty"`
	// ReplicateTo - Specifies the partner name that the snapshots in this snapshot collection are replicated to.
	ReplicateTo *string `json:"replicate_to,omitempty"`
	// StartOnline - Start with snapshot set online.
	StartOnline *bool `json:"start_online,omitempty"`
	// AllowWrites - Allow applications to write to created snapshot(s). Mandatory and must be set to 'true' for VSS application synchronized snapshots.
	AllowWrites *bool `json:"allow_writes,omitempty"`
	// DisableAppsync - Do not perform application synchronization for this snapshot, create a crash-consistent snapshot instead.
	DisableAppsync *bool `json:"disable_appsync,omitempty"`
	// SnapVerify - Run verification tool on this snapshot. This option can only be used with a volume collection that has application synchronization.
	SnapVerify *bool `json:"snap_verify,omitempty"`
	// SkipDbConsistencyCheck - Skip consistency check for database files on this snapshot. This option only applies to volume collections with application synchronization set to VSS, application ID set to MS Exchange 2010 or later with Database Availability Group (DAG), snap_verify option set to true, and disable_appsync option set to false.
	SkipDbConsistencyCheck *bool `json:"skip_db_consistency_check,omitempty"`
	// SchedId - ID of protection schedule of snapshot collection.
	SchedId *string `json:"sched_id,omitempty"`
	// SchedName - Name of protection schedule of snapshot collection.
	SchedName *string `json:"sched_name,omitempty"`
	// InvokeOnUpstreamPartner - Invoke snapshot request on upstream partner. This operation is not supported for synchronous replication volume vollections.
	InvokeOnUpstreamPartner *bool `json:"invoke_on_upstream_partner,omitempty"`
	// AgentType - External management agent type for snapshots being created as part of snapshot collection.
	AgentType *NsAgentType `json:"agent_type,omitempty"`
	// ExpiryAfter - Number of seconds after which this snapcoll is considered expired by the snapshot TTL. A value of 0 indicates that the snapshot never expires, 1 indicates that the snapshot uses a group-level configured TTL value and any other value indicates the number of seconds.
	ExpiryAfter *int64 `json:"expiry_after,omitempty"`
	// Metadata - Key-value pairs that augment a snapshot collection's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// Force - Forcibly delete the specified snapshot collection even if it is the last replicated snapshot. Doing so could lead to full re-seeding at the next replication.
	Force *bool `json:"force,omitempty"`
}

// SnapshotCollectionFieldHandles provides a string representation for each SnapshotCollection field.
type SnapshotCollectionFieldHandles struct {
	ID                      string
	Name                    string
	Description             string
	VolcollName             string
	VolcollId               string
	OriginName              string
	IsReplica               string
	SrepOwnerName           string
	SrepOwnerId             string
	PeerSnapcollId          string
	NumSnaps                string
	IsComplete              string
	IsManual                string
	IsExternalTrigger       string
	IsUnmanaged             string
	IsManuallyManaged       string
	ReplStatus              string
	ReplStartTime           string
	ReplCompleteTime        string
	ReplBytesTransferred    string
	CreationTime            string
	LastModified            string
	OnlineStatus            string
	VolSnapAttrList         string
	SnapshotsList           string
	Replicate               string
	ReplicateTo             string
	StartOnline             string
	AllowWrites             string
	DisableAppsync          string
	SnapVerify              string
	SkipDbConsistencyCheck  string
	SchedId                 string
	SchedName               string
	InvokeOnUpstreamPartner string
	AgentType               string
	ExpiryAfter             string
	Metadata                string
	Force                   string
}
