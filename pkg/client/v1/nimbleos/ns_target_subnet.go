// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsTargetSubnet - List of subnet labels.
// Export NsTargetSubnetFields for advance operations like search filter etc.
var NsTargetSubnetFields *NsTargetSubnet

func init() {

	NsTargetSubnetFields = &NsTargetSubnet{
		ID:    "id",
		Label: "label",
	}
}

type NsTargetSubnet struct {
	// ID - Subnet ID.
	ID string `json:"id,omitempty"`
	// Label - Subnet label.
	Label string `json:"label,omitempty"`
}

// sdk internal struct
type nsTargetSubnet struct {
	// ID - Subnet ID.
	ID *string `json:"id,omitempty"`
	// Label - Subnet label.
	Label *string `json:"label,omitempty"`
}

// EncodeNsTargetSubnet - Transform NsTargetSubnet to nsTargetSubnet type
func EncodeNsTargetSubnet(request interface{}) (*nsTargetSubnet, error) {
	reqNsTargetSubnet := request.(*NsTargetSubnet)
	byte, err := json.Marshal(reqNsTargetSubnet)
	objPtr := &nsTargetSubnet{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsTargetSubnet Transform nsTargetSubnet to NsTargetSubnet type
func DecodeNsTargetSubnet(request interface{}) (*NsTargetSubnet, error) {
	reqNsTargetSubnet := request.(*nsTargetSubnet)
	byte, err := json.Marshal(reqNsTargetSubnet)
	obj := &NsTargetSubnet{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
