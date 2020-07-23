// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcInterfaceInfo - Fibre Channel interface information.
// Export NsFcInterfaceInfoFields for advance operations like search filter etc.
var NsFcInterfaceInfoFields *NsFcInterfaceInfo

func init() {
	Namefield := "name"
	Wwnnfield := "wwnn"
	Wwpnfield := "wwpn"
	BusLocationfield := "bus_location"

	NsFcInterfaceInfoFields = &NsFcInterfaceInfo{
		Name:        &Namefield,
		Wwnn:        &Wwnnfield,
		Wwpn:        &Wwpnfield,
		BusLocation: &BusLocationfield,
	}
}

type NsFcInterfaceInfo struct {
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
