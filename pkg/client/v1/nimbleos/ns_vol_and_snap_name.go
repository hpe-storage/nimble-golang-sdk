// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsVolAndSnapName - Snapshot name and the belonging volume name.
// Export NsVolAndSnapNameFields for advance operations like search filter etc.
var NsVolAndSnapNameFields *NsVolAndSnapName

func init() {

	NsVolAndSnapNameFields = &NsVolAndSnapName{
		VolName:  "vol_name",
		SnapName: "snap_name",
	}
}

type NsVolAndSnapName struct {
	// VolName - The name of the volume that the snapshot belongs to.
	VolName string `json:"vol_name,omitempty"`
	// SnapName - Snapshot name.
	SnapName string `json:"snap_name,omitempty"`
}

// sdk internal struct
type nsVolAndSnapName struct {
	// VolName - The name of the volume that the snapshot belongs to.
	VolName *string `json:"vol_name,omitempty"`
	// SnapName - Snapshot name.
	SnapName *string `json:"snap_name,omitempty"`
}

// EncodeNsVolAndSnapName - Transform NsVolAndSnapName to nsVolAndSnapName type
func EncodeNsVolAndSnapName(request interface{}) (*nsVolAndSnapName, error) {
	reqNsVolAndSnapName := request.(*NsVolAndSnapName)
	byte, err := json.Marshal(reqNsVolAndSnapName)
	objPtr := &nsVolAndSnapName{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsVolAndSnapName Transform nsVolAndSnapName to NsVolAndSnapName type
func DecodeNsVolAndSnapName(request interface{}) (*NsVolAndSnapName, error) {
	reqNsVolAndSnapName := request.(*nsVolAndSnapName)
	byte, err := json.Marshal(reqNsVolAndSnapName)
	obj := &NsVolAndSnapName{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
