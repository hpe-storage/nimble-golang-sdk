// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsArraySummaryInfo - Array summary information, including version, model, and IP configurations.
// Export NsArraySummaryInfoFields for advance operations like search filter etc.
var NsArraySummaryInfoFields *NsArraySummaryInfo

func init() {

	NsArraySummaryInfoFields = &NsArraySummaryInfo{
		Name:    "name",
		Version: "version",
		Serial:  "serial",
		Model:   "model",
	}
}

type NsArraySummaryInfo struct {
	// Name - Unique name of array.
	Name string `json:"name,omitempty"`
	// Version - Version of array.
	Version string `json:"version,omitempty"`
	// Serial - Array serial number.
	Serial string `json:"serial,omitempty"`
	// Model - Array hardware model.
	Model string `json:"model,omitempty"`
	// ZconfIpaddrs - Zero Conf IP of array, including Nic, local and remote IP addresses.
	ZconfIpaddrs []*NsZeroConfIPAddr `json:"zconf_ipaddrs,omitempty"`
	// Status - Status of array.
	Status *NsArrayStatus `json:"status,omitempty"`
	// CountOfFcPorts - Count of Fibre Channel ports per controller.
	CountOfFcPorts int64 `json:"count_of_fc_ports,omitempty"`
	// AllFlash - Whether it is an all-flash array.
	AllFlash *bool `json:"all_flash,omitempty"`
}

// sdk internal struct
type nsArraySummaryInfo struct {
	// Name - Unique name of array.
	Name *string `json:"name,omitempty"`
	// Version - Version of array.
	Version *string `json:"version,omitempty"`
	// Serial - Array serial number.
	Serial *string `json:"serial,omitempty"`
	// Model - Array hardware model.
	Model *string `json:"model,omitempty"`
	// ZconfIpaddrs - Zero Conf IP of array, including Nic, local and remote IP addresses.
	ZconfIpaddrs []*NsZeroConfIPAddr `json:"zconf_ipaddrs,omitempty"`
	// Status - Status of array.
	Status *NsArrayStatus `json:"status,omitempty"`
	// CountOfFcPorts - Count of Fibre Channel ports per controller.
	CountOfFcPorts *int64 `json:"count_of_fc_ports,omitempty"`
	// AllFlash - Whether it is an all-flash array.
	AllFlash *bool `json:"all_flash,omitempty"`
}

// EncodeNsArraySummaryInfo - Transform NsArraySummaryInfo to nsArraySummaryInfo type
func EncodeNsArraySummaryInfo(request interface{}) (*nsArraySummaryInfo, error) {
	reqNsArraySummaryInfo := request.(*NsArraySummaryInfo)
	byte, err := json.Marshal(reqNsArraySummaryInfo)
	objPtr := &nsArraySummaryInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsArraySummaryInfo Transform nsArraySummaryInfo to NsArraySummaryInfo type
func DecodeNsArraySummaryInfo(request interface{}) (*NsArraySummaryInfo, error) {
	reqNsArraySummaryInfo := request.(*nsArraySummaryInfo)
	byte, err := json.Marshal(reqNsArraySummaryInfo)
	obj := &NsArraySummaryInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
