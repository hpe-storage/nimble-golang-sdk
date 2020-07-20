// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFcPortInfo - Fibre Channel port information.
// Export NsFcPortInfoFields for advance operations like search filter etc.
var NsFcPortInfoFields *NsFcPortInfo

func init() {

	NsFcPortInfoFields = &NsFcPortInfo{
		Name:        "name",
		BusLocation: "bus_location",
	}
}

type NsFcPortInfo struct {
	// Name - Name of Fibre Channel port.
	Name string `json:"name,omitempty"`
	// BusLocation - PCI bus location of the HBA for this Fibre Channel port.
	BusLocation string `json:"bus_location,omitempty"`
	// Port - HBA port number for this Fibre Channel port.
	Port int64 `json:"port,omitempty"`
	// Slot - HBA slot number for this Fibre Channel port.
	Slot int64 `json:"slot,omitempty"`
}

// sdk internal struct
type nsFcPortInfo struct {
	Name        *string `json:"name,omitempty"`
	BusLocation *string `json:"bus_location,omitempty"`
	Port        *int64  `json:"port,omitempty"`
	Slot        *int64  `json:"slot,omitempty"`
}

// EncodeNsFcPortInfo - Transform NsFcPortInfo to nsFcPortInfo type
func EncodeNsFcPortInfo(request interface{}) (*nsFcPortInfo, error) {
	reqNsFcPortInfo := request.(*NsFcPortInfo)
	byte, err := json.Marshal(reqNsFcPortInfo)
	if err != nil {
		return nil, err
	}
	respNsFcPortInfoPtr := &nsFcPortInfo{}
	err = json.Unmarshal(byte, respNsFcPortInfoPtr)
	return respNsFcPortInfoPtr, err
}

// DecodeNsFcPortInfo Transform nsFcPortInfo to NsFcPortInfo type
func DecodeNsFcPortInfo(request interface{}) (*NsFcPortInfo, error) {
	reqNsFcPortInfo := request.(*nsFcPortInfo)
	byte, err := json.Marshal(reqNsFcPortInfo)
	if err != nil {
		return nil, err
	}
	respNsFcPortInfo := &NsFcPortInfo{}
	err = json.Unmarshal(byte, respNsFcPortInfo)
	return respNsFcPortInfo, err
}
