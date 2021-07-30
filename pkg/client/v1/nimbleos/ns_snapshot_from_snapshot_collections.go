// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSnapshotFromSnapshotCollectionsFields provides field names to use in filter parameters, for example.
var NsSnapshotFromSnapshotCollectionsFields *NsSnapshotFromSnapshotCollectionsFieldHandles

func init() {
	NsSnapshotFromSnapshotCollectionsFields = &NsSnapshotFromSnapshotCollectionsFieldHandles{
		ID:           "id",
		SnapcollId:   "snapcoll_id",
		Name:         "name",
		SnapcollName: "snapcoll_name",
		VolId:        "vol_id",
		VolName:      "vol_name",
		SnapId:       "snap_id",
		SnapName:     "snap_name",
		ScheduleId:   "schedule_id",
		ScheduleName: "schedule_name",
		ExpiryTime:   "expiry_time",
	}
}

// NsSnapshotFromSnapshotCollections - Snapshot as presented in snapshot collections object set.
type NsSnapshotFromSnapshotCollections struct {
	// ID - Snapshot ID.
	ID *string `json:"id,omitempty"`
	// SnapcollId - ID of the snapshot collection which owns this snapshot.
	SnapcollId *string `json:"snapcoll_id,omitempty"`
	// Name - Snapshot name.
	Name *string `json:"name,omitempty"`
	// SnapcollName - Name of the snapshot collection which owns this snapshot.
	SnapcollName *string `json:"snapcoll_name,omitempty"`
	// VolId - ID of the volume that is parent to this snapshot.
	VolId *string `json:"vol_id,omitempty"`
	// VolName - Name of the volume that is parent to this snapshot.
	VolName *string `json:"vol_name,omitempty"`
	// SnapId - Snapshot ID.
	SnapId *string `json:"snap_id,omitempty"`
	// SnapName - Snapshot name.
	SnapName *string `json:"snap_name,omitempty"`
	// ScheduleId - ID of protection schedule.
	ScheduleId *string `json:"schedule_id,omitempty"`
	// ScheduleName - Name of protection schedule.
	ScheduleName *string `json:"schedule_name,omitempty"`
	// ExpiryTime - Unix timestamp indicating that the snapshot is considered expired by Snapshot Time-to-live(TTL). A value of 0 indicates that snapshot never expires.
	ExpiryTime *int64 `json:"expiry_time,omitempty"`
}

// NsSnapshotFromSnapshotCollectionsFieldHandles provides a string representation for each AccessControlRecord field.
type NsSnapshotFromSnapshotCollectionsFieldHandles struct {
	ID           string
	SnapcollId   string
	Name         string
	SnapcollName string
	VolId        string
	VolName      string
	SnapId       string
	SnapName     string
	ScheduleId   string
	ScheduleName string
	ExpiryTime   string
}
