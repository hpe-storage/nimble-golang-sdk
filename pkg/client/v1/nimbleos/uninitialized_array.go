// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// UninitializedArray - Lists discovered arrays that are not members of any group and are in the same subnet.
// Export UninitializedArrayFields for advance operations like search filter etc.
var UninitializedArrayFields *UninitializedArray

func init() {

	UninitializedArrayFields = &UninitializedArray{
		ID:        "id",
		Serial:    "serial",
		ArrayName: "array_name",
		Model:     "model",
		ModelStr:  "model_str",
		Version:   "version",
		IpAddress: "ip_address",
	}
}

type UninitializedArray struct {
	// ID - Identifier for the interface.
	ID string `json:"id,omitempty"`
	// Serial - Serial Number of the uninitialized array.
	Serial string `json:"serial,omitempty"`
	// ArrayName - Name of the uninitialized array.
	ArrayName string `json:"array_name,omitempty"`
	// Model - Model of the uninitialized array.
	Model string `json:"model,omitempty"`
	// ModelStr - Model description of the uninitialized array.
	ModelStr string `json:"model_str,omitempty"`
	// Version - Version of the uninitialized array.
	Version string `json:"version,omitempty"`
	// IpAddress - Link local zero conf address of the uninitialized array.
	IpAddress string `json:"ip_address,omitempty"`
	// ZconfIpaddrs - List of link local zero conf address of the uninitialized array.
	ZconfIpaddrs []*NsIPAddressObject `json:"zconf_ipaddrs,omitempty"`
	// CountOfFcPorts - Number of Fibre Channel ports of the uninitialized array.
	CountOfFcPorts int64 `json:"count_of_fc_ports,omitempty"`
	// AllFlash - True if it is an All-Flash array, False otherwise.
	AllFlash *bool `json:"all_flash,omitempty"`
}

// sdk internal struct
type uninitializedArray struct {
	ID             *string              `json:"id,omitempty"`
	Serial         *string              `json:"serial,omitempty"`
	ArrayName      *string              `json:"array_name,omitempty"`
	Model          *string              `json:"model,omitempty"`
	ModelStr       *string              `json:"model_str,omitempty"`
	Version        *string              `json:"version,omitempty"`
	IpAddress      *string              `json:"ip_address,omitempty"`
	ZconfIpaddrs   []*NsIPAddressObject `json:"zconf_ipaddrs,omitempty"`
	CountOfFcPorts *int64               `json:"count_of_fc_ports,omitempty"`
	AllFlash       *bool                `json:"all_flash,omitempty"`
}

// EncodeUninitializedArray - Transform UninitializedArray to uninitializedArray type
func EncodeUninitializedArray(request interface{}) (*uninitializedArray, error) {
	reqUninitializedArray := request.(*UninitializedArray)
	byte, err := json.Marshal(reqUninitializedArray)
	if err != nil {
		return nil, err
	}
	respUninitializedArrayPtr := &uninitializedArray{}
	err = json.Unmarshal(byte, respUninitializedArrayPtr)
	return respUninitializedArrayPtr, err
}

// DecodeUninitializedArray Transform uninitializedArray to UninitializedArray type
func DecodeUninitializedArray(request interface{}) (*UninitializedArray, error) {
	reqUninitializedArray := request.(*uninitializedArray)
	byte, err := json.Marshal(reqUninitializedArray)
	if err != nil {
		return nil, err
	}
	respUninitializedArray := &UninitializedArray{}
	err = json.Unmarshal(byte, respUninitializedArray)
	return respUninitializedArray, err
}
