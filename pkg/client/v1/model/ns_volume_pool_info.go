// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsVolumePoolInfo - Volume information along with the pool to which it belongs to.
// Export NsVolumePoolInfoFields for advance operations like search filter etc.
var NsVolumePoolInfoFields *NsVolumePoolInfo

func init() {

	NsVolumePoolInfoFields = &NsVolumePoolInfo{
		VolId:    "vol_id",
		VolName:  "vol_name",
		PoolId:   "pool_id",
		PoolName: "pool_name",
	}
}

type NsVolumePoolInfo struct {
	// VolId - ID of the volume.
	VolId string `json:"vol_id,omitempty"`
	// VolName - Name of the volume.
	VolName string `json:"vol_name,omitempty"`
	// PoolId - ID of the pool to which the volume belongs to.
	PoolId string `json:"pool_id,omitempty"`
	// PoolName - Name of the pool to which volume belongs to.
	PoolName string `json:"pool_name,omitempty"`
}

// sdk internal struct
type nsVolumePoolInfo struct {
	// VolId - ID of the volume.
	VolId *string `json:"vol_id,omitempty"`
	// VolName - Name of the volume.
	VolName *string `json:"vol_name,omitempty"`
	// PoolId - ID of the pool to which the volume belongs to.
	PoolId *string `json:"pool_id,omitempty"`
	// PoolName - Name of the pool to which volume belongs to.
	PoolName *string `json:"pool_name,omitempty"`
}

// EncodeNsVolumePoolInfo - Transform NsVolumePoolInfo to nsVolumePoolInfo type
func EncodeNsVolumePoolInfo(request interface{}) (*nsVolumePoolInfo, error) {
	reqNsVolumePoolInfo := request.(*NsVolumePoolInfo)
	byte, err := json.Marshal(reqNsVolumePoolInfo)
	objPtr := &nsVolumePoolInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsVolumePoolInfo Transform nsVolumePoolInfo to NsVolumePoolInfo type
func DecodeNsVolumePoolInfo(request interface{}) (*NsVolumePoolInfo, error) {
	reqNsVolumePoolInfo := request.(*nsVolumePoolInfo)
	byte, err := json.Marshal(reqNsVolumePoolInfo)
	obj := &NsVolumePoolInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
