// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolFamMigStatus - Data migration status for a group of related volumes.
// Export NsVolFamMigStatusFields for advance operations like search filter etc.
var NsVolFamMigStatusFields *NsVolFamMigStatus

func init() {

	NsVolFamMigStatusFields = &NsVolFamMigStatus{
		RootVolId:      "root_vol_id",
		RootVolName:    "root_vol_name",
		SourcePoolId:   "source_pool_id",
		SourcePoolName: "source_pool_name",
		DestPoolId:     "dest_pool_id",
		DestPoolName:   "dest_pool_name",
	}
}

type NsVolFamMigStatus struct {
	// RootVolId - ID of the root volume in the group.
	RootVolId string `json:"root_vol_id,omitempty"`
	// RootVolName - Name of the root volume in the group.
	RootVolName string `json:"root_vol_name,omitempty"`
	// SourcePoolId - ID of the source pool, where the volumes originally locate.
	SourcePoolId string `json:"source_pool_id,omitempty"`
	// SourcePoolName - Name of the source pool, where the volumes originally locate.
	SourcePoolName string `json:"source_pool_name,omitempty"`
	// DestPoolId - ID of the destination pool, where the volumes are moved.
	DestPoolId string `json:"dest_pool_id,omitempty"`
	// DestPoolName - Name of the destination pool, where the volumes are moved.
	DestPoolName string `json:"dest_pool_name,omitempty"`
	// MoveBytesMigrated - The bytes of volumes which have been moved.
	MoveBytesMigrated int64 `json:"move_bytes_migrated,omitempty"`
	// MoveBytesRemaining - The bytes of volumes which have not been moved.
	MoveBytesRemaining int64 `json:"move_bytes_remaining,omitempty"`
	// MoveStartTime - The start time when the volumes was moved.
	MoveStartTime int64 `json:"move_start_time,omitempty"`
	// MoveEstComplTime - The estimated time of completion of a move.
	MoveEstComplTime int64 `json:"move_est_compl_time,omitempty"`
	// ArrayList - Data migration status for the arrays that store the volumes.
	ArrayList []*NsArrayMigStatus `json:"array_list,omitempty"`
}

// sdk internal struct
type nsVolFamMigStatus struct {
	RootVolId          *string             `json:"root_vol_id,omitempty"`
	RootVolName        *string             `json:"root_vol_name,omitempty"`
	SourcePoolId       *string             `json:"source_pool_id,omitempty"`
	SourcePoolName     *string             `json:"source_pool_name,omitempty"`
	DestPoolId         *string             `json:"dest_pool_id,omitempty"`
	DestPoolName       *string             `json:"dest_pool_name,omitempty"`
	MoveBytesMigrated  *int64              `json:"move_bytes_migrated,omitempty"`
	MoveBytesRemaining *int64              `json:"move_bytes_remaining,omitempty"`
	MoveStartTime      *int64              `json:"move_start_time,omitempty"`
	MoveEstComplTime   *int64              `json:"move_est_compl_time,omitempty"`
	ArrayList          []*NsArrayMigStatus `json:"array_list,omitempty"`
}

// EncodeNsVolFamMigStatus - Transform NsVolFamMigStatus to nsVolFamMigStatus type
func EncodeNsVolFamMigStatus(request interface{}) (*nsVolFamMigStatus, error) {
	reqNsVolFamMigStatus := request.(*NsVolFamMigStatus)
	byte, err := json.Marshal(reqNsVolFamMigStatus)
	if err != nil {
		return nil, err
	}
	respNsVolFamMigStatusPtr := &nsVolFamMigStatus{}
	err = json.Unmarshal(byte, respNsVolFamMigStatusPtr)
	return respNsVolFamMigStatusPtr, err
}

// DecodeNsVolFamMigStatus Transform nsVolFamMigStatus to NsVolFamMigStatus type
func DecodeNsVolFamMigStatus(request interface{}) (*NsVolFamMigStatus, error) {
	reqNsVolFamMigStatus := request.(*nsVolFamMigStatus)
	byte, err := json.Marshal(reqNsVolFamMigStatus)
	if err != nil {
		return nil, err
	}
	respNsVolFamMigStatus := &NsVolFamMigStatus{}
	err = json.Unmarshal(byte, respNsVolFamMigStatus)
	return respNsVolFamMigStatus, err
}
