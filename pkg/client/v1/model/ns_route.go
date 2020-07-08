// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsRoute - Route configuration.
// Export NsRouteFields for advance operations like search filter etc.
var NsRouteFields *NsRoute

func init() {

	NsRouteFields = &NsRoute{
		TgtNetwork: "tgt_network",
		TgtNetmask: "tgt_netmask",
		Gateway:    "gateway",
	}
}

type NsRoute struct {
	// TgtNetwork - Target network address.
	TgtNetwork string `json:"tgt_network,omitempty"`
	// TgtNetmask - Target network mask.
	TgtNetmask string `json:"tgt_netmask,omitempty"`
	// Gateway - Gateway IP address.
	Gateway string `json:"gateway,omitempty"`
}

// sdk internal struct
type nsRoute struct {
	// TgtNetwork - Target network address.
	TgtNetwork *string `json:"tgt_network,omitempty"`
	// TgtNetmask - Target network mask.
	TgtNetmask *string `json:"tgt_netmask,omitempty"`
	// Gateway - Gateway IP address.
	Gateway *string `json:"gateway,omitempty"`
}

// EncodeNsRoute - Transform NsRoute to nsRoute type
func EncodeNsRoute(request interface{}) (*nsRoute, error) {
	reqNsRoute := request.(*NsRoute)
	byte, err := json.Marshal(reqNsRoute)
	objPtr := &nsRoute{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsRoute Transform nsRoute to NsRoute type
func DecodeNsRoute(request interface{}) (*NsRoute, error) {
	reqNsRoute := request.(*nsRoute)
	byte, err := json.Marshal(reqNsRoute)
	obj := &NsRoute{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
