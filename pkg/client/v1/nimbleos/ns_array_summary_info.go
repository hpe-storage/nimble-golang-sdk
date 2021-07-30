// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsArraySummaryInfoFields provides field names to use in filter parameters, for example.
var NsArraySummaryInfoFields *NsArraySummaryInfoFieldHandles

func init() {
	NsArraySummaryInfoFields = &NsArraySummaryInfoFieldHandles{
		Name:           "name",
		Version:        "version",
		Serial:         "serial",
		Model:          "model",
		ZconfIpaddrs:   "zconf_ipaddrs",
		Status:         "status",
		CountOfFcPorts: "count_of_fc_ports",
		AllFlash:       "all_flash",
	}
}

// NsArraySummaryInfo - Array summary information, including version, model, and IP configurations.
type NsArraySummaryInfo struct {
	// Name - Unique name of array.
	Name *string `json:"name,omitempty"`
	// Version - Version of array.
	Version *string `json:"version,omitempty"`
	// Serial - Array serial number.
	Serial *string `json:"serial,omitempty"`
	// Model - Array hardware model.
	Model *string `json:"model,omitempty"`
	// ZconfIpaddrs - Zero Conf IP of array, including Nic, local and remote IP addresses.
	ZconfIpaddrs []*NsZeroConfIPAddr `json:"zconf_ipaddrs,omitempty"`
	// Status - Status of array.
	Status *NsArrayStatus `json:"status,omitempty"`
	// CountOfFcPorts - Count of Fibre Channel ports per controller.
	CountOfFcPorts *int64 `json:"count_of_fc_ports,omitempty"`
	// AllFlash - Whether it is an all-flash array.
	AllFlash *bool `json:"all_flash,omitempty"`
}

// NsArraySummaryInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsArraySummaryInfoFieldHandles struct {
	Name           string
	Version        string
	Serial         string
	Model          string
	ZconfIpaddrs   string
	Status         string
	CountOfFcPorts string
	AllFlash       string
}
