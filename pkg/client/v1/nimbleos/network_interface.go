// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NetworkInterface - Manage per array network interface configuration.
// Export NetworkInterfaceFields for advance operations like search filter etc.
var NetworkInterfaceFields *NetworkInterfaceStringFields

func init() {
	IDfield := "id"
	ArrayNameOrSerialfield := "array_name_or_serial"
	PartialResponseOkfield := "partial_response_ok"
	ArrayIdfield := "array_id"
	ControllerNamefield := "controller_name"
	ControllerIdfield := "controller_id"
	Namefield := "name"
	Macfield := "mac"
	IsPresentfield := "is_present"
	LinkSpeedfield := "link_speed"
	LinkStatusfield := "link_status"
	Mtufield := "mtu"
	Portfield := "port"
	Slotfield := "slot"
	MaxLinkSpeedfield := "max_link_speed"
	NicTypefield := "nic_type"
	IpListfield := "ip_list"

	NetworkInterfaceFields = &NetworkInterfaceStringFields{
		ID:                &IDfield,
		ArrayNameOrSerial: &ArrayNameOrSerialfield,
		PartialResponseOk: &PartialResponseOkfield,
		ArrayId:           &ArrayIdfield,
		ControllerName:    &ControllerNamefield,
		ControllerId:      &ControllerIdfield,
		Name:              &Namefield,
		Mac:               &Macfield,
		IsPresent:         &IsPresentfield,
		LinkSpeed:         &LinkSpeedfield,
		LinkStatus:        &LinkStatusfield,
		Mtu:               &Mtufield,
		Port:              &Portfield,
		Slot:              &Slotfield,
		MaxLinkSpeed:      &MaxLinkSpeedfield,
		NicType:           &NicTypefield,
		IpList:            &IpListfield,
	}
}

type NetworkInterface struct {
	// ID - Identifier for the interface.
	ID *string `json:"id,omitempty"`
	// ArrayNameOrSerial - Name or serial of the array where the interface is hosted.
	ArrayNameOrSerial *string `json:"array_name_or_serial,omitempty"`
	// PartialResponseOk - Indicate that it is ok to provide partially available response.
	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
	// ArrayId - Identifier for the array.
	ArrayId *string `json:"array_id,omitempty"`
	// ControllerName - Name (A or B) of the controller where the interface is hosted.
	ControllerName *string `json:"controller_name,omitempty"`
	// ControllerId - Identifier of the controller where the interface is hosted.
	ControllerId *string `json:"controller_id,omitempty"`
	// Name - Name of the interface.
	Name *string `json:"name,omitempty"`
	// Mac - MAC address of the interface.
	Mac *string `json:"mac,omitempty"`
	// IsPresent - Whether this interface is present on this controller.
	IsPresent *bool `json:"is_present,omitempty"`
	// LinkSpeed - Speed of the link.
	LinkSpeed *NsPlatLinkSpeed `json:"link_speed,omitempty"`
	// LinkStatus - Status of the link.
	LinkStatus *NsPlatLinkStatus `json:"link_status,omitempty"`
	// Mtu - MTU on the link.
	Mtu *int64 `json:"mtu,omitempty"`
	// Port - Port number for this interface.
	Port *int64 `json:"port,omitempty"`
	// Slot - Slot number for this interface.
	Slot *int64 `json:"slot,omitempty"`
	// MaxLinkSpeed - Maximum speed of the link.
	MaxLinkSpeed *NsPlatLinkSpeed `json:"max_link_speed,omitempty"`
	// NicType - Interface type.
	NicType *NsPlatNicType `json:"nic_type,omitempty"`
	// IpList - List of IP addresses assigned to this network interface.
	IpList []*NsAssignedIP `json:"ip_list,omitempty"`
}

// Struct for NetworkInterfaceFields
type NetworkInterfaceStringFields struct {
	ID                *string
	ArrayNameOrSerial *string
	PartialResponseOk *string
	ArrayId           *string
	ControllerName    *string
	ControllerId      *string
	Name              *string
	Mac               *string
	IsPresent         *string
	LinkSpeed         *string
	LinkStatus        *string
	Mtu               *string
	Port              *string
	Slot              *string
	MaxLinkSpeed      *string
	NicType           *string
	IpList            *string
}
