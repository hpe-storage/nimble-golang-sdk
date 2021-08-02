// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// UninitializedArrayFields provides field names to use in filter parameters, for example.
var UninitializedArrayFields *UninitializedArrayFieldHandles

func init() {
	UninitializedArrayFields = &UninitializedArrayFieldHandles{
		ID:                 "id",
		Serial:             "serial",
		ArrayName:          "array_name",
		Model:              "model",
		ModelStr:           "model_str",
		Version:            "version",
		IpAddress:          "ip_address",
		ZconfIpaddrs:       "zconf_ipaddrs",
		CountOfFcPorts:     "count_of_fc_ports",
		AllFlash:           "all_flash",
		DedupeConfigurable: "dedupe_configurable",
	}
}

// UninitializedArray - Lists discovered arrays that are not members of any group and are in the same subnet.
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

// UninitializedArrayFieldHandles provides a string representation for each AccessControlRecord field.
type UninitializedArrayFieldHandles struct {
	ID                 string
	Serial             string
	ArrayName          string
	Model              string
	ModelStr           string
	Version            string
	IpAddress          string
	ZconfIpaddrs       string
	CountOfFcPorts     string
	AllFlash           string
	DedupeConfigurable string
}
