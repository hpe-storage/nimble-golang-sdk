// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFibreChannelInterfaceFullName - Fully qualified name information for a Fibre Channel interface (array/controller/interface).

// Export NsFibreChannelInterfaceFullNameFields provides field names to use in filter parameters, for example.
var NsFibreChannelInterfaceFullNameFields *NsFibreChannelInterfaceFullNameStringFields

func init() {
	fieldArrayName := "array_name"
	fieldCtrlrName := "ctrlr_name"
	fieldIntfName := "intf_name"

	NsFibreChannelInterfaceFullNameFields = &NsFibreChannelInterfaceFullNameStringFields{
		ArrayName: &fieldArrayName,
		CtrlrName: &fieldCtrlrName,
		IntfName:  &fieldIntfName,
	}
}

type NsFibreChannelInterfaceFullName struct {
	// ArrayName - Array name.
	ArrayName *string `json:"array_name,omitempty"`
	// CtrlrName - Controller name.
	CtrlrName *string `json:"ctrlr_name,omitempty"`
	// IntfName - Fibre Channel interface name.
	IntfName *string `json:"intf_name,omitempty"`
}

// Struct for NsFibreChannelInterfaceFullNameFields
type NsFibreChannelInterfaceFullNameStringFields struct {
	ArrayName *string
	CtrlrName *string
	IntfName  *string
}
