// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcInterfaceInfo - Fibre Channel interface information.
// Export NsFcInterfaceInfoFields for advance operations like search filter etc.
var NsFcInterfaceInfoFields *NsFcInterfaceInfoStringFields

func init() {
	Namefield := "name"
	Wwnnfield := "wwnn"
	Wwpnfield := "wwpn"
	Onlinefield := "online"
	BusLocationfield := "bus_location"
	Slotfield := "slot"
	Portfield := "port"

	NsFcInterfaceInfoFields = &NsFcInterfaceInfoStringFields{
		Name:        &Namefield,
		Wwnn:        &Wwnnfield,
		Wwpn:        &Wwpnfield,
		Online:      &Onlinefield,
		BusLocation: &BusLocationfield,
		Slot:        &Slotfield,
		Port:        &Portfield,
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

// Struct for NsFcInterfaceInfoFields
type NsFcInterfaceInfoStringFields struct {
	Name        *string
	Wwnn        *string
	Wwpn        *string
	Online      *string
	BusLocation *string
	Slot        *string
	Port        *string
}
