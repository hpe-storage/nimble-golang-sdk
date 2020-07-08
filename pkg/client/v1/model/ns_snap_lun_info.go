// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsSnapLunInfo - Snapshot LUN information as presented in the access_control_records object set. This information is only available for Fibre Channel.
// Export NsSnapLunInfoFields for advance operations like search filter etc.
var NsSnapLunInfoFields *NsSnapLunInfo

func init() {

	NsSnapLunInfoFields = &NsSnapLunInfo{
		ID:   "id",
		Name: "name",
	}
}

type NsSnapLunInfo struct {
	// ID - Snapshot ID.
	ID string `json:"id,omitempty"`
	// Name - Snapshot name.
	Name string `json:"name,omitempty"`
	// Lun - Snapshot LUN.
	Lun int64 `json:"lun,omitempty"`
}

// sdk internal struct
type nsSnapLunInfo struct {
	// ID - Snapshot ID.
	ID *string `json:"id,omitempty"`
	// Name - Snapshot name.
	Name *string `json:"name,omitempty"`
	// Lun - Snapshot LUN.
	Lun *int64 `json:"lun,omitempty"`
}

// EncodeNsSnapLunInfo - Transform NsSnapLunInfo to nsSnapLunInfo type
func EncodeNsSnapLunInfo(request interface{}) (*nsSnapLunInfo, error) {
	reqNsSnapLunInfo := request.(*NsSnapLunInfo)
	byte, err := json.Marshal(reqNsSnapLunInfo)
	objPtr := &nsSnapLunInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSnapLunInfo Transform nsSnapLunInfo to NsSnapLunInfo type
func DecodeNsSnapLunInfo(request interface{}) (*NsSnapLunInfo, error) {
	reqNsSnapLunInfo := request.(*nsSnapLunInfo)
	byte, err := json.Marshal(reqNsSnapLunInfo)
	obj := &NsSnapLunInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
