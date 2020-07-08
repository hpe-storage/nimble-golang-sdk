// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsSnapshotFromVolumes - Snapshot as presented in volumes object set.
// Export NsSnapshotFromVolumesFields for advance operations like search filter etc.
var NsSnapshotFromVolumesFields *NsSnapshotFromVolumes

func init() {

	NsSnapshotFromVolumesFields = &NsSnapshotFromVolumes{
		ID:       "id",
		SnapId:   "snap_id",
		Name:     "name",
		SnapName: "snap_name",
	}
}

type NsSnapshotFromVolumes struct {
	// ID - Snapshot id.
	ID string `json:"id,omitempty"`
	// SnapId - Snapshot id.
	SnapId string `json:"snap_id,omitempty"`
	// Name - Snapshot name.
	Name string `json:"name,omitempty"`
	// SnapName - Snapshot name.
	SnapName string `json:"snap_name,omitempty"`
}

// sdk internal struct
type nsSnapshotFromVolumes struct {
	// ID - Snapshot id.
	ID *string `json:"id,omitempty"`
	// SnapId - Snapshot id.
	SnapId *string `json:"snap_id,omitempty"`
	// Name - Snapshot name.
	Name *string `json:"name,omitempty"`
	// SnapName - Snapshot name.
	SnapName *string `json:"snap_name,omitempty"`
}

// EncodeNsSnapshotFromVolumes - Transform NsSnapshotFromVolumes to nsSnapshotFromVolumes type
func EncodeNsSnapshotFromVolumes(request interface{}) (*nsSnapshotFromVolumes, error) {
	reqNsSnapshotFromVolumes := request.(*NsSnapshotFromVolumes)
	byte, err := json.Marshal(reqNsSnapshotFromVolumes)
	objPtr := &nsSnapshotFromVolumes{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSnapshotFromVolumes Transform nsSnapshotFromVolumes to NsSnapshotFromVolumes type
func DecodeNsSnapshotFromVolumes(request interface{}) (*NsSnapshotFromVolumes, error) {
	reqNsSnapshotFromVolumes := request.(*nsSnapshotFromVolumes)
	byte, err := json.Marshal(reqNsSnapshotFromVolumes)
	obj := &NsSnapshotFromVolumes{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
