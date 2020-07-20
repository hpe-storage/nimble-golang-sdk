// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsSnapshotFromSnapshotCollections - Snapshot as presented in snapshot collections object set.
// Export NsSnapshotFromSnapshotCollectionsFields for advance operations like search filter etc.
var NsSnapshotFromSnapshotCollectionsFields *NsSnapshotFromSnapshotCollections

func init() {

	NsSnapshotFromSnapshotCollectionsFields = &NsSnapshotFromSnapshotCollections{
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
	}
}

type NsSnapshotFromSnapshotCollections struct {
	// ID - Snapshot ID.
	ID string `json:"id,omitempty"`
	// SnapcollId - ID of the snapshot collection which owns this snapshot.
	SnapcollId string `json:"snapcoll_id,omitempty"`
	// Name - Snapshot name.
	Name string `json:"name,omitempty"`
	// SnapcollName - Name of the snapshot collection which owns this snapshot.
	SnapcollName string `json:"snapcoll_name,omitempty"`
	// VolId - ID of the volume that is parent to this snapshot.
	VolId string `json:"vol_id,omitempty"`
	// VolName - Name of the volume that is parent to this snapshot.
	VolName string `json:"vol_name,omitempty"`
	// SnapId - Snapshot ID.
	SnapId string `json:"snap_id,omitempty"`
	// SnapName - Snapshot name.
	SnapName string `json:"snap_name,omitempty"`
	// ScheduleId - ID of protection schedule.
	ScheduleId string `json:"schedule_id,omitempty"`
	// ScheduleName - Name of protection schedule.
	ScheduleName string `json:"schedule_name,omitempty"`
	// ExpiryTime - Unix timestamp indicating that the snapshot is considered expired by Snapshot Time-to-live(TTL). A value of 0 indicates that snapshot never expires.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
}

// sdk internal struct
type nsSnapshotFromSnapshotCollections struct {
	ID           *string `json:"id,omitempty"`
	SnapcollId   *string `json:"snapcoll_id,omitempty"`
	Name         *string `json:"name,omitempty"`
	SnapcollName *string `json:"snapcoll_name,omitempty"`
	VolId        *string `json:"vol_id,omitempty"`
	VolName      *string `json:"vol_name,omitempty"`
	SnapId       *string `json:"snap_id,omitempty"`
	SnapName     *string `json:"snap_name,omitempty"`
	ScheduleId   *string `json:"schedule_id,omitempty"`
	ScheduleName *string `json:"schedule_name,omitempty"`
	ExpiryTime   *int64  `json:"expiry_time,omitempty"`
}

// EncodeNsSnapshotFromSnapshotCollections - Transform NsSnapshotFromSnapshotCollections to nsSnapshotFromSnapshotCollections type
func EncodeNsSnapshotFromSnapshotCollections(request interface{}) (*nsSnapshotFromSnapshotCollections, error) {
	reqNsSnapshotFromSnapshotCollections := request.(*NsSnapshotFromSnapshotCollections)
	byte, err := json.Marshal(reqNsSnapshotFromSnapshotCollections)
	if err != nil {
		return nil, err
	}
	respNsSnapshotFromSnapshotCollectionsPtr := &nsSnapshotFromSnapshotCollections{}
	err = json.Unmarshal(byte, respNsSnapshotFromSnapshotCollectionsPtr)
	return respNsSnapshotFromSnapshotCollectionsPtr, err
}

// DecodeNsSnapshotFromSnapshotCollections Transform nsSnapshotFromSnapshotCollections to NsSnapshotFromSnapshotCollections type
func DecodeNsSnapshotFromSnapshotCollections(request interface{}) (*NsSnapshotFromSnapshotCollections, error) {
	reqNsSnapshotFromSnapshotCollections := request.(*nsSnapshotFromSnapshotCollections)
	byte, err := json.Marshal(reqNsSnapshotFromSnapshotCollections)
	if err != nil {
		return nil, err
	}
	respNsSnapshotFromSnapshotCollections := &NsSnapshotFromSnapshotCollections{}
	err = json.Unmarshal(byte, respNsSnapshotFromSnapshotCollections)
	return respNsSnapshotFromSnapshotCollections, err
}
