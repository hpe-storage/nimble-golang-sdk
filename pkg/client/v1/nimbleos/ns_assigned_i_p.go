// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsAssignedIP - IP address assignment details to network interface.
// Export NsAssignedIPFields for advance operations like search filter etc.
var NsAssignedIPFields *NsAssignedIP

func init() {

	NsAssignedIPFields = &NsAssignedIP{
		Ip: "ip",
	}
}

type NsAssignedIP struct {
	// VlanId - VLAN id of network interface.
	VlanId int64 `json:"vlan_id,omitempty"`
	// Ip - Assigned IP address to network interface.
	Ip string `json:"ip,omitempty"`
}

// sdk internal struct
type nsAssignedIP struct {
	VlanId *int64  `json:"vlan_id,omitempty"`
	Ip     *string `json:"ip,omitempty"`
}

// EncodeNsAssignedIP - Transform NsAssignedIP to nsAssignedIP type
func EncodeNsAssignedIP(request interface{}) (*nsAssignedIP, error) {
	reqNsAssignedIP := request.(*NsAssignedIP)
	byte, err := json.Marshal(reqNsAssignedIP)
	if err != nil {
		return nil, err
	}
	respNsAssignedIPPtr := &nsAssignedIP{}
	err = json.Unmarshal(byte, respNsAssignedIPPtr)
	return respNsAssignedIPPtr, err
}

// DecodeNsAssignedIP Transform nsAssignedIP to NsAssignedIP type
func DecodeNsAssignedIP(request interface{}) (*NsAssignedIP, error) {
	reqNsAssignedIP := request.(*nsAssignedIP)
	byte, err := json.Marshal(reqNsAssignedIP)
	if err != nil {
		return nil, err
	}
	respNsAssignedIP := &NsAssignedIP{}
	err = json.Unmarshal(byte, respNsAssignedIP)
	return respNsAssignedIP, err
}
