// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFcInterfaceInfo - Fibre Channel interface information.
// Export NsFcInterfaceInfoFields for advance operations like search filter etc.
var NsFcInterfaceInfoFields *NsFcInterfaceInfo

func init() {

	NsFcInterfaceInfoFields = &NsFcInterfaceInfo{
		Name:        "name",
		Wwnn:        "wwnn",
		Wwpn:        "wwpn",
		BusLocation: "bus_location",
	}
}

type NsFcInterfaceInfo struct {
	// Name - Name of Fibre Channel interface.
	Name string `json:"name,omitempty"`
	// Wwnn - WWNN (World Wide Node Name) for this Fibre Channel interface.
	Wwnn string `json:"wwnn,omitempty"`
	// Wwpn - WWPN (World Wide Port Name)for this Fibre Channel interface.
	Wwpn string `json:"wwpn,omitempty"`
	// Online - Identify whether the Fibre Channel interface is online.
	Online *bool `json:"online,omitempty"`
	// BusLocation - PCI bus location of the HBA.
	BusLocation string `json:"bus_location,omitempty"`
	// Slot - Information about the Fibre Channel slot.
	Slot int64 `json:"slot,omitempty"`
	// Port - Information about the Fibre Channel port on which the interface is hosted.
	Port int64 `json:"port,omitempty"`
}

// sdk internal struct
type nsFcInterfaceInfo struct {
	// Name - Name of Fibre Channel interface.
	Name *string `json:"name,omitempty"`
	// Wwnn - WWNN (World Wide Node Name) for this Fibre Channel interface.
	Wwnn *string `json:"wwnn,omitempty"`
	// Wwpn - WWPN (World Wide Port Name)for this Fibre Channel interface.
	Wwpn *string `json:"wwpn,omitempty"`
	// Online - Identify whether the Fibre Channel interface is online.
	Online *bool `json:"online,omitempty"`
	// BusLocation - PCI bus location of the HBA.
	BusLocation *string `json:"bus_location,omitempty"`
	// Slot - Information about the Fibre Channel slot.
	Slot *int64 `json:"slot,omitempty"`
	// Port - Information about the Fibre Channel port on which the interface is hosted.
	Port *int64 `json:"port,omitempty"`
}

// EncodeNsFcInterfaceInfo - Transform NsFcInterfaceInfo to nsFcInterfaceInfo type
func EncodeNsFcInterfaceInfo(request interface{}) (*nsFcInterfaceInfo, error) {
	reqNsFcInterfaceInfo := request.(*NsFcInterfaceInfo)
	byte, err := json.Marshal(reqNsFcInterfaceInfo)
	objPtr := &nsFcInterfaceInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsFcInterfaceInfo Transform nsFcInterfaceInfo to NsFcInterfaceInfo type
func DecodeNsFcInterfaceInfo(request interface{}) (*NsFcInterfaceInfo, error) {
	reqNsFcInterfaceInfo := request.(*nsFcInterfaceInfo)
	byte, err := json.Marshal(reqNsFcInterfaceInfo)
	obj := &NsFcInterfaceInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
