// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolumeBulkUpdateAttr - Volume object used in bulk update.
// Export NsVolumeBulkUpdateAttrFields for advance operations like search filter etc.
var NsVolumeBulkUpdateAttrFields *NsVolumeBulkUpdateAttr

func init() {

	NsVolumeBulkUpdateAttrFields = &NsVolumeBulkUpdateAttr{
		ID:       "id",
		FolderId: "folder_id",
	}
}

type NsVolumeBulkUpdateAttr struct {
	// ID - ID of volume.
	ID string `json:"id,omitempty"`
	// FolderId - ID of folder.
	FolderId string `json:"folder_id,omitempty"`
	// Online - Online state of the volume.
	Online *bool `json:"online,omitempty"`
}

// sdk internal struct
type nsVolumeBulkUpdateAttr struct {
	ID       *string `json:"id,omitempty"`
	FolderId *string `json:"folder_id,omitempty"`
	Online   *bool   `json:"online,omitempty"`
}

// EncodeNsVolumeBulkUpdateAttr - Transform NsVolumeBulkUpdateAttr to nsVolumeBulkUpdateAttr type
func EncodeNsVolumeBulkUpdateAttr(request interface{}) (*nsVolumeBulkUpdateAttr, error) {
	reqNsVolumeBulkUpdateAttr := request.(*NsVolumeBulkUpdateAttr)
	byte, err := json.Marshal(reqNsVolumeBulkUpdateAttr)
	if err != nil {
		return nil, err
	}
	respNsVolumeBulkUpdateAttrPtr := &nsVolumeBulkUpdateAttr{}
	err = json.Unmarshal(byte, respNsVolumeBulkUpdateAttrPtr)
	return respNsVolumeBulkUpdateAttrPtr, err
}

// DecodeNsVolumeBulkUpdateAttr Transform nsVolumeBulkUpdateAttr to NsVolumeBulkUpdateAttr type
func DecodeNsVolumeBulkUpdateAttr(request interface{}) (*NsVolumeBulkUpdateAttr, error) {
	reqNsVolumeBulkUpdateAttr := request.(*nsVolumeBulkUpdateAttr)
	byte, err := json.Marshal(reqNsVolumeBulkUpdateAttr)
	if err != nil {
		return nil, err
	}
	respNsVolumeBulkUpdateAttr := &NsVolumeBulkUpdateAttr{}
	err = json.Unmarshal(byte, respNsVolumeBulkUpdateAttr)
	return respNsVolumeBulkUpdateAttr, err
}
