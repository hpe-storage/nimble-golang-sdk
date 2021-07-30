// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// FibreChannelPort - Fibre Channel ports provide data access. This API provides the list of all Fibre Channel ports configured on the arrays.

// Export FibreChannelPortFields provides field names to use in filter parameters, for example.
var FibreChannelPortFields *FibreChannelPortStringFields

func init() {
	fieldID := "id"
	fieldArrayNameOrSerial := "array_name_or_serial"
	fieldControllerName := "controller_name"
	fieldFcPortName := "fc_port_name"
	fieldBusLocation := "bus_location"
	fieldPort := "port"
	fieldSlot := "slot"
	fieldOrientation := "orientation"
	fieldLinkInfo := "link_info"
	fieldRxPower := "rx_power"
	fieldTxPower := "tx_power"

	FibreChannelPortFields = &FibreChannelPortStringFields{
		ID:                &fieldID,
		ArrayNameOrSerial: &fieldArrayNameOrSerial,
		ControllerName:    &fieldControllerName,
		FcPortName:        &fieldFcPortName,
		BusLocation:       &fieldBusLocation,
		Port:              &fieldPort,
		Slot:              &fieldSlot,
		Orientation:       &fieldOrientation,
		LinkInfo:          &fieldLinkInfo,
		RxPower:           &fieldRxPower,
		TxPower:           &fieldTxPower,
	}
}

type FibreChannelPort struct {
	// ID - Identifier for the Fibre Channel port.
	ID *string `json:"id,omitempty"`
	// ArrayNameOrSerial - Name or serial number of the array.
	ArrayNameOrSerial *string `json:"array_name_or_serial,omitempty"`
	// ControllerName - Name (A or B) of the controller to which the port belongs.
	ControllerName *string `json:"controller_name,omitempty"`
	// FcPortName - Name of the Fibre Channel port.
	FcPortName *string `json:"fc_port_name,omitempty"`
	// BusLocation - PCI bus location of the HBA (Host Bus Adapter) for this Fibre Channel port.
	BusLocation *string `json:"bus_location,omitempty"`
	// Port - HBA (Host Bus Adapter) port number for this Fibre Channel port.
	Port *int64 `json:"port,omitempty"`
	// Slot - HBA (Host Bus Adapter) slot number for this Fibre Channel port.
	Slot *int64 `json:"slot,omitempty"`
	// Orientation - Orientation of FC ports on a HBA. An orientation of 'right_to_left' indicates that ports are ordered as 3,2,1,0 on the slot. Possible values: 'left_to_right', 'right_to_left'.
	Orientation *NsPlatOrientation `json:"orientation,omitempty"`
	// LinkInfo - Information about the Fibre Channel link associated with this port.
	LinkInfo *NsFcLinkInfo `json:"link_info,omitempty"`
	// RxPower - SFP RX power in uW.
	RxPower *int64 `json:"rx_power,omitempty"`
	// TxPower - SFP TX power in uW.
	TxPower *int64 `json:"tx_power,omitempty"`
}

// Struct for FibreChannelPortFields
type FibreChannelPortStringFields struct {
	ID                *string
	ArrayNameOrSerial *string
	ControllerName    *string
	FcPortName        *string
	BusLocation       *string
	Port              *string
	Slot              *string
	Orientation       *string
	LinkInfo          *string
	RxPower           *string
	TxPower           *string
}
