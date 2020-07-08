// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsShelfCtrlrAttrSet - A shelf logical controller attributes.
// Export NsShelfCtrlrAttrSetFields for advance operations like search filter etc.
var NsShelfCtrlrAttrSetFields *NsShelfCtrlrAttrSet

func init() {

	NsShelfCtrlrAttrSetFields = &NsShelfCtrlrAttrSet{
		SessionSerial: "session_serial",
		CachedSerial:  "cached_serial",
		DiskSerials:   "disk_serials",
		DiskTypes:     "disk_types",
	}
}

type NsShelfCtrlrAttrSet struct {
	// SessionSerial - Session serial.
	SessionSerial string `json:"session_serial,omitempty"`
	// CachedSerial - Cached serial.
	CachedSerial string `json:"cached_serial,omitempty"`
	// HwState - The hardware state for this logical controller.
	HwState *NsShelfHwState `json:"hw_state,omitempty"`
	// SwType - The software type of this logical controller.
	SwType *NsShelfSwType `json:"sw_type,omitempty"`
	// DiskSerials - Comma separated list of disk serials connected to this logical controller.
	DiskSerials string `json:"disk_serials,omitempty"`
	// DiskTypes - Comma separated list of disk types (H for HDD, S for SSD).
	DiskTypes string `json:"disk_types,omitempty"`
}

// sdk internal struct
type nsShelfCtrlrAttrSet struct {
	// SessionSerial - Session serial.
	SessionSerial *string `json:"session_serial,omitempty"`
	// CachedSerial - Cached serial.
	CachedSerial *string `json:"cached_serial,omitempty"`
	// HwState - The hardware state for this logical controller.
	HwState *NsShelfHwState `json:"hw_state,omitempty"`
	// SwType - The software type of this logical controller.
	SwType *NsShelfSwType `json:"sw_type,omitempty"`
	// DiskSerials - Comma separated list of disk serials connected to this logical controller.
	DiskSerials *string `json:"disk_serials,omitempty"`
	// DiskTypes - Comma separated list of disk types (H for HDD, S for SSD).
	DiskTypes *string `json:"disk_types,omitempty"`
}

// EncodeNsShelfCtrlrAttrSet - Transform NsShelfCtrlrAttrSet to nsShelfCtrlrAttrSet type
func EncodeNsShelfCtrlrAttrSet(request interface{}) (*nsShelfCtrlrAttrSet, error) {
	reqNsShelfCtrlrAttrSet := request.(*NsShelfCtrlrAttrSet)
	byte, err := json.Marshal(reqNsShelfCtrlrAttrSet)
	objPtr := &nsShelfCtrlrAttrSet{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsShelfCtrlrAttrSet Transform nsShelfCtrlrAttrSet to NsShelfCtrlrAttrSet type
func DecodeNsShelfCtrlrAttrSet(request interface{}) (*NsShelfCtrlrAttrSet, error) {
	reqNsShelfCtrlrAttrSet := request.(*nsShelfCtrlrAttrSet)
	byte, err := json.Marshal(reqNsShelfCtrlrAttrSet)
	obj := &NsShelfCtrlrAttrSet{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
