// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsReplPair - Replicated objects (vol/snap/snapcoll), their UIDs are the same.
// Export NsReplPairFields for advance operations like search filter etc.
var NsReplPairFields *NsReplPair

func init() {

	NsReplPairFields = &NsReplPair{
		SrcName: "src_name",
		DstName: "dst_name",
	}
}

type NsReplPair struct {
	// SrcName - Name of the replicated obj on the source group.
	SrcName string `json:"src_name,omitempty"`
	// DstName - Name of the replicated obj on the destination group.
	DstName string `json:"dst_name,omitempty"`
}

// sdk internal struct
type nsReplPair struct {
	// SrcName - Name of the replicated obj on the source group.
	SrcName *string `json:"src_name,omitempty"`
	// DstName - Name of the replicated obj on the destination group.
	DstName *string `json:"dst_name,omitempty"`
}

// EncodeNsReplPair - Transform NsReplPair to nsReplPair type
func EncodeNsReplPair(request interface{}) (*nsReplPair, error) {
	reqNsReplPair := request.(*NsReplPair)
	byte, err := json.Marshal(reqNsReplPair)
	objPtr := &nsReplPair{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsReplPair Transform nsReplPair to NsReplPair type
func DecodeNsReplPair(request interface{}) (*NsReplPair, error) {
	reqNsReplPair := request.(*nsReplPair)
	byte, err := json.Marshal(reqNsReplPair)
	obj := &NsReplPair{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
