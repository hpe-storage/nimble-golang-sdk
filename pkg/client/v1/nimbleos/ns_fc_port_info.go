// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcPortInfo - Fibre Channel port information.
// Export NsFcPortInfoFields for advance operations like search filter etc.
var NsFcPortInfoFields *NsFcPortInfoStringFields

func init() {
	Namefield := "name"
	BusLocationfield := "bus_location"
	Portfield := "port"
	Slotfield := "slot"

	NsFcPortInfoFields = &NsFcPortInfoStringFields{
		Name:        &Namefield,
		BusLocation: &BusLocationfield,
		Port:        &Portfield,
		Slot:        &Slotfield,
	}
}

type NsFcPortInfo struct {
	// Name - Name of Fibre Channel port.
	Name *string `json:"name,omitempty"`
	// BusLocation - PCI bus location of the HBA for this Fibre Channel port.
	BusLocation *string `json:"bus_location,omitempty"`
	// Port - HBA port number for this Fibre Channel port.
	Port *int64 `json:"port,omitempty"`
	// Slot - HBA slot number for this Fibre Channel port.
	Slot *int64 `json:"slot,omitempty"`
}

// Struct for NsFcPortInfoFields
type NsFcPortInfoStringFields struct {
	Name        *string
	BusLocation *string
	Port        *string
	Slot        *string
}
