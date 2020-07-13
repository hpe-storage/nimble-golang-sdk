// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsCtrlrNvmeCard - NVMe accelerator card.
// Export NsCtrlrNvmeCardFields for advance operations like search filter etc.
var NsCtrlrNvmeCardFields *NsCtrlrNvmeCard

func init() {

	NsCtrlrNvmeCardFields = &NsCtrlrNvmeCard{
		SerialNumber: "serial_number",
	}
}

type NsCtrlrNvmeCard struct {
	// SerialNumber - Serial number.
	SerialNumber string `json:"serial_number,omitempty"`
	// Size - NVMe card cache size in bytes.
	Size int64 `json:"size,omitempty"`
	// State - Online state.
	State *NsCtrlrNvmeCardState `json:"state,omitempty"`
}

// sdk internal struct
type nsCtrlrNvmeCard struct {
	// SerialNumber - Serial number.
	SerialNumber *string `json:"serial_number,omitempty"`
	// Size - NVMe card cache size in bytes.
	Size *int64 `json:"size,omitempty"`
	// State - Online state.
	State *NsCtrlrNvmeCardState `json:"state,omitempty"`
}

// EncodeNsCtrlrNvmeCard - Transform NsCtrlrNvmeCard to nsCtrlrNvmeCard type
func EncodeNsCtrlrNvmeCard(request interface{}) (*nsCtrlrNvmeCard, error) {
	reqNsCtrlrNvmeCard := request.(*NsCtrlrNvmeCard)
	byte, err := json.Marshal(reqNsCtrlrNvmeCard)
	objPtr := &nsCtrlrNvmeCard{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsCtrlrNvmeCard Transform nsCtrlrNvmeCard to NsCtrlrNvmeCard type
func DecodeNsCtrlrNvmeCard(request interface{}) (*NsCtrlrNvmeCard, error) {
	reqNsCtrlrNvmeCard := request.(*nsCtrlrNvmeCard)
	byte, err := json.Marshal(reqNsCtrlrNvmeCard)
	obj := &NsCtrlrNvmeCard{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
