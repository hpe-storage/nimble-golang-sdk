// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsFcLinkInfo - Fibre Channel link information.
// Export NsFcLinkInfoFields for advance operations like search filter etc.
var NsFcLinkInfoFields *NsFcLinkInfo

func init() {

	NsFcLinkInfoFields = &NsFcLinkInfo{}
}

type NsFcLinkInfo struct {
	// LinkSpeed - Speed of the Fibre Channel link.
	LinkSpeed *NsPlatFcLinkSpeed `json:"link_speed,omitempty"`
	// MaxLinkSpeed - Maximum speed of the Fibre Channel link.
	MaxLinkSpeed *NsPlatFcLinkSpeed `json:"max_link_speed,omitempty"`
	// LinkStatus - Status of the Fibre Channel link.
	LinkStatus *NsPlatFcLinkStatus `json:"link_status,omitempty"`
	// OperationalStatus - Operational status of the Fibre Channel link.
	OperationalStatus *NsPlatFcOperationalStatus `json:"operational_status,omitempty"`
}

// sdk internal struct
type nsFcLinkInfo struct {
	// LinkSpeed - Speed of the Fibre Channel link.
	LinkSpeed *NsPlatFcLinkSpeed `json:"link_speed,omitempty"`
	// MaxLinkSpeed - Maximum speed of the Fibre Channel link.
	MaxLinkSpeed *NsPlatFcLinkSpeed `json:"max_link_speed,omitempty"`
	// LinkStatus - Status of the Fibre Channel link.
	LinkStatus *NsPlatFcLinkStatus `json:"link_status,omitempty"`
	// OperationalStatus - Operational status of the Fibre Channel link.
	OperationalStatus *NsPlatFcOperationalStatus `json:"operational_status,omitempty"`
}

// EncodeNsFcLinkInfo - Transform NsFcLinkInfo to nsFcLinkInfo type
func EncodeNsFcLinkInfo(request interface{}) (*nsFcLinkInfo, error) {
	reqNsFcLinkInfo := request.(*NsFcLinkInfo)
	byte, err := json.Marshal(reqNsFcLinkInfo)
	objPtr := &nsFcLinkInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsFcLinkInfo Transform nsFcLinkInfo to NsFcLinkInfo type
func DecodeNsFcLinkInfo(request interface{}) (*NsFcLinkInfo, error) {
	reqNsFcLinkInfo := request.(*nsFcLinkInfo)
	byte, err := json.Marshal(reqNsFcLinkInfo)
	obj := &NsFcLinkInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
