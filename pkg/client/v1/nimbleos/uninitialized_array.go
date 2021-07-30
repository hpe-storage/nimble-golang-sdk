// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// UninitializedArray - Lists discovered arrays that are not members of any group and are in the same subnet.
// Export UninitializedArrayFields for advance operations like search filter etc.
var UninitializedArrayFields *UninitializedArrayStringFields

func init() {
	IDfield := "id"
	Serialfield := "serial"
	ArrayNamefield := "array_name"
	Modelfield := "model"
	ModelStrfield := "model_str"
	Versionfield := "version"
	IpAddressfield := "ip_address"
	ZconfIpaddrsfield := "zconf_ipaddrs"
	CountOfFcPortsfield := "count_of_fc_ports"
	AllFlashfield := "all_flash"
	DedupeConfigurablefield := "dedupe_configurable"

	UninitializedArrayFields = &UninitializedArrayStringFields{
		ID:                 &IDfield,
		Serial:             &Serialfield,
		ArrayName:          &ArrayNamefield,
		Model:              &Modelfield,
		ModelStr:           &ModelStrfield,
		Version:            &Versionfield,
		IpAddress:          &IpAddressfield,
		ZconfIpaddrs:       &ZconfIpaddrsfield,
		CountOfFcPorts:     &CountOfFcPortsfield,
		AllFlash:           &AllFlashfield,
		DedupeConfigurable: &DedupeConfigurablefield,
	}
}

type UninitializedArray struct {
	// ID - Identifier for the interface.
	ID *string `json:"id,omitempty"`
	// Serial - Serial Number of the uninitialized array.
	Serial *string `json:"serial,omitempty"`
	// ArrayName - Name of the uninitialized array.
	ArrayName *string `json:"array_name,omitempty"`
	// Model - Model of the uninitialized array.
	Model *string `json:"model,omitempty"`
	// ModelStr - Model description of the uninitialized array.
	ModelStr *string `json:"model_str,omitempty"`
	// Version - Version of the uninitialized array.
	Version *string `json:"version,omitempty"`
	// IpAddress - Link local zero conf address of the uninitialized array.
	IpAddress *string `json:"ip_address,omitempty"`
	// ZconfIpaddrs - List of link local zero conf address of the uninitialized array.
	ZconfIpaddrs []*NsIPAddressObject `json:"zconf_ipaddrs,omitempty"`
	// CountOfFcPorts - Number of Fibre Channel ports of the uninitialized array.
	CountOfFcPorts *int64 `json:"count_of_fc_ports,omitempty"`
	// AllFlash - True if it is an All-Flash array, False otherwise.
	AllFlash *bool `json:"all_flash,omitempty"`
	// DedupeConfigurable - True if it is a hybrid array that is capable of updating data deduplication setting, False otherwise.
	DedupeConfigurable *bool `json:"dedupe_configurable,omitempty"`
}

// Struct for UninitializedArrayFields
type UninitializedArrayStringFields struct {
	ID                 *string
	Serial             *string
	ArrayName          *string
	Model              *string
	ModelStr           *string
	Version            *string
	IpAddress          *string
	ZconfIpaddrs       *string
	CountOfFcPorts     *string
	AllFlash           *string
	DedupeConfigurable *string
}
