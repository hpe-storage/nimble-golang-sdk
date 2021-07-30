// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArraySummaryInfo - Array summary information, including version, model, and IP configurations.
// Export NsArraySummaryInfoFields for advance operations like search filter etc.
var NsArraySummaryInfoFields *NsArraySummaryInfoStringFields

func init() {
	Namefield := "name"
	Versionfield := "version"
	Serialfield := "serial"
	Modelfield := "model"
	ZconfIpaddrsfield := "zconf_ipaddrs"
	Statusfield := "status"
	CountOfFcPortsfield := "count_of_fc_ports"
	AllFlashfield := "all_flash"

	NsArraySummaryInfoFields = &NsArraySummaryInfoStringFields{
		Name:           &Namefield,
		Version:        &Versionfield,
		Serial:         &Serialfield,
		Model:          &Modelfield,
		ZconfIpaddrs:   &ZconfIpaddrsfield,
		Status:         &Statusfield,
		CountOfFcPorts: &CountOfFcPortsfield,
		AllFlash:       &AllFlashfield,
	}
}

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

// Struct for NsArraySummaryInfoFields
type NsArraySummaryInfoStringFields struct {
	Name           *string
	Version        *string
	Serial         *string
	Model          *string
	ZconfIpaddrs   *string
	Status         *string
	CountOfFcPorts *string
	AllFlash       *string
}
