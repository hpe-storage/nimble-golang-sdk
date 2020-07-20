// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsObjectOwnerPair - Objects and their owners.
// Export NsObjectOwnerPairFields for advance operations like search filter etc.
var NsObjectOwnerPairFields *NsObjectOwnerPair

func init() {

	NsObjectOwnerPairFields = &NsObjectOwnerPair{
		ObjName:  "obj_name",
		SrcOwner: "src_owner",
		DstOwner: "dst_owner",
	}
}

type NsObjectOwnerPair struct {
	// ObjName - Object name. Same on source and destination.
	ObjName string `json:"obj_name,omitempty"`
	// SrcOwner - Name of the owner on the source group.
	SrcOwner string `json:"src_owner,omitempty"`
	// DstOwner - Name of the owner on the destination group.
	DstOwner string `json:"dst_owner,omitempty"`
}

// sdk internal struct
type nsObjectOwnerPair struct {
	ObjName  *string `json:"obj_name,omitempty"`
	SrcOwner *string `json:"src_owner,omitempty"`
	DstOwner *string `json:"dst_owner,omitempty"`
}

// EncodeNsObjectOwnerPair - Transform NsObjectOwnerPair to nsObjectOwnerPair type
func EncodeNsObjectOwnerPair(request interface{}) (*nsObjectOwnerPair, error) {
	reqNsObjectOwnerPair := request.(*NsObjectOwnerPair)
	byte, err := json.Marshal(reqNsObjectOwnerPair)
	if err != nil {
		return nil, err
	}
	respNsObjectOwnerPairPtr := &nsObjectOwnerPair{}
	err = json.Unmarshal(byte, respNsObjectOwnerPairPtr)
	return respNsObjectOwnerPairPtr, err
}

// DecodeNsObjectOwnerPair Transform nsObjectOwnerPair to NsObjectOwnerPair type
func DecodeNsObjectOwnerPair(request interface{}) (*NsObjectOwnerPair, error) {
	reqNsObjectOwnerPair := request.(*nsObjectOwnerPair)
	byte, err := json.Marshal(reqNsObjectOwnerPair)
	if err != nil {
		return nil, err
	}
	respNsObjectOwnerPair := &NsObjectOwnerPair{}
	err = json.Unmarshal(byte, respNsObjectOwnerPair)
	return respNsObjectOwnerPair, err
}
