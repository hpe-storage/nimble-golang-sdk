// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcInterfaceInfoFields provides field names to use in filter parameters, for example.
var NsFcInterfaceInfoFields *NsFcInterfaceInfoFieldHandles

func init() {
	NsFcInterfaceInfoFields = &NsFcInterfaceInfoFieldHandles{
		Name:        "name",
		Wwnn:        "wwnn",
		Wwpn:        "wwpn",
		Online:      "online",
		BusLocation: "bus_location",
		Slot:        "slot",
		Port:        "port",
	}
}

// NsFcInterfaceInfo - Fibre Channel interface information.
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

// NsFcInterfaceInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcInterfaceInfoFieldHandles struct {
	Name        string
	Wwnn        string
	Wwpn        string
	Online      string
	BusLocation string
	Slot        string
	Port        string
}
