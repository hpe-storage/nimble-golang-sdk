// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// FibreChannelInterface - Represent information of specified Fibre Channel interfaces. Fibre Channel interfaces are hosted on Fibre Channel ports to provide data access.
// Export FibreChannelInterfaceFields for advance operations like search filter etc.
var FibreChannelInterfaceFields *FibreChannelInterface

func init() {
	IDfield := "id"
	ArrayNameOrSerialfield := "array_name_or_serial"
	ControllerNamefield := "controller_name"
	FcPortIdfield := "fc_port_id"
	Namefield := "name"
	Wwnnfield := "wwnn"
	Wwpnfield := "wwpn"
	Peerzonefield := "peerzone"
	FirmwareVersionfield := "firmware_version"
	FcPortNamefield := "fc_port_name"
	BusLocationfield := "bus_location"

	FibreChannelInterfaceFields = &FibreChannelInterface{
		ID:                &IDfield,
		ArrayNameOrSerial: &ArrayNameOrSerialfield,
		ControllerName:    &ControllerNamefield,
		FcPortId:          &FcPortIdfield,
		Name:              &Namefield,
		Wwnn:              &Wwnnfield,
		Wwpn:              &Wwpnfield,
		Peerzone:          &Peerzonefield,
		FirmwareVersion:   &FirmwareVersionfield,
		FcPortName:        &FcPortNamefield,
		BusLocation:       &BusLocationfield,
	}
}

type FibreChannelInterface struct {
	// ID - Identifier for the Fibre Channel interface.
	ID *string `json:"id,omitempty"`
	// ArrayNameOrSerial - Name or serial number of array where the interface is hosted.
	ArrayNameOrSerial *string `json:"array_name_or_serial,omitempty"`
	// PartialResponseOk - Indicate that it is ok to provide partially available response.
	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
	// ControllerName - Name (A or B) of the controller where the interface is hosted.
	ControllerName *string `json:"controller_name,omitempty"`
	// FcPortId - ID of the port with which the interface is associated.
	FcPortId *string `json:"fc_port_id,omitempty"`
	// Name - Name of Fibre Channel interface.
	Name *string `json:"name,omitempty"`
	// Wwnn - WWNN (World Wide Node Name) for this Fibre Channel interface.
	Wwnn *string `json:"wwnn,omitempty"`
	// Wwpn - WWPN (World Wide Port Name) for this Fibre Channel interface.
	Wwpn *string `json:"wwpn,omitempty"`
	// Peerzone - Active peer zone for this Fibre Channel interface.
	Peerzone *string `json:"peerzone,omitempty"`
	// Online - Identify whether the Fibre Channel interface is online.
	Online *bool `json:"online,omitempty"`
	// FirmwareVersion - Version of the Fibre Channel firmware.
	FirmwareVersion *string `json:"firmware_version,omitempty"`
	// LogicalPortNumber - Logical port number for the Fibre Channel port.
	LogicalPortNumber *int64 `json:"logical_port_number,omitempty"`
	// FcPortName - Name of Fibre Channel port.
	FcPortName *string `json:"fc_port_name,omitempty"`
	// BusLocation - PCI bus location of the HBA for this Fibre Channel port.
	BusLocation *string `json:"bus_location,omitempty"`
	// Slot - HBA slot number for this Fibre Channel port.
	Slot *int64 `json:"slot,omitempty"`
	// Orientation - Orientation of FC ports on a HBA. An orientation of 'right_to_left' indicates that ports are ordered as 3,2,1,0 on the slot. Possible values: 'left_to_right', 'right_to_left'.
	Orientation *NsPlatOrientation `json:"orientation,omitempty"`
	// Port - HBA port number for this Fibre Channel port.
	Port *int64 `json:"port,omitempty"`
	// LinkInfo - Information about the Fibre Channel link at which interface is operating.
	LinkInfo *NsFcLinkInfo `json:"link_info,omitempty"`
	// FabricInfo - Fibre Channel fabric information.
	FabricInfo *NsFcFabricInfo `json:"fabric_info,omitempty"`
}
