// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// ControllerFields provides field names to use in filter parameters, for example.
var ControllerFields *ControllerFieldHandles

func init() {
	ControllerFields = &ControllerFieldHandles{
		ID:                 "id",
		Name:               "name",
		ArrayName:          "array_name",
		ArrayId:            "array_id",
		PartialResponseOk:  "partial_response_ok",
		Serial:             "serial",
		Model:              "model",
		Hostname:           "hostname",
		SupportAddress:     "support_address",
		SupportNetmask:     "support_netmask",
		SupportNic:         "support_nic",
		PowerStatus:        "power_status",
		FanStatus:          "fan_status",
		TemperatureStatus:  "temperature_status",
		PowerSupplies:      "power_supplies",
		Fans:               "fans",
		TemperatureSensors: "temperature_sensors",
		PartitionStatus:    "partition_status",
		CtrlrSide:          "ctrlr_side",
		State:              "state",
		NvmeCardsEnabled:   "nvme_cards_enabled",
		NvmeCards:          "nvme_cards",
		AsupTime:           "asup_time",
	}
}

// Controller - Controller is a redundant collection of hardware capable of running the array software.
type Controller struct {
	// ID - Identifier of the controller.
	ID *string `json:"id,omitempty"`
	// Name - Name of the controller.
	Name *string `json:"name,omitempty"`
	// ArrayName - Name of the array containing this controller.
	ArrayName *string `json:"array_name,omitempty"`
	// ArrayId - Rest ID of the array containing this controller.
	ArrayId *string `json:"array_id,omitempty"`
	// PartialResponseOk - Indicate that it is ok to provide partially available response.
	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
	// Serial - Serial number for this controller.
	Serial *string `json:"serial,omitempty"`
	// Model - Model of this controller.
	Model *string `json:"model,omitempty"`
	// Hostname - Host name for the controller.
	Hostname *string `json:"hostname,omitempty"`
	// SupportAddress - IP address used for support.
	SupportAddress *string `json:"support_address,omitempty"`
	// SupportNetmask - IP netmask used for support.
	SupportNetmask *string `json:"support_netmask,omitempty"`
	// SupportNic - Network card used for support.
	SupportNic *string `json:"support_nic,omitempty"`
	// PowerStatus - Overall power supply status for the controller.
	PowerStatus *NsPowerSupplyStatus `json:"power_status,omitempty"`
	// FanStatus - Overall fan status for the controller.
	FanStatus *NsFanStatus `json:"fan_status,omitempty"`
	// TemperatureStatus - Overall temperature status for the controller.
	TemperatureStatus *NsTemperatureStatus `json:"temperature_status,omitempty"`
	// PowerSupplies - Status for each power supply in the controller.
	PowerSupplies []*NsCtrlrHwSensorInfo `json:"power_supplies,omitempty"`
	// Fans - Status for each fan in the controller.
	Fans []*NsCtrlrHwSensorInfo `json:"fans,omitempty"`
	// TemperatureSensors - Status for temperature sensor in the controller.
	TemperatureSensors []*NsCtrlrHwSensorInfo `json:"temperature_sensors,omitempty"`
	// PartitionStatus - Status of the system's raid partitions.
	PartitionStatus []*NsCtrlrRaidInfo `json:"partition_status,omitempty"`
	// CtrlrSide - Identifies which controller this is on its array.
	CtrlrSide *NsControllerSide `json:"ctrlr_side,omitempty"`
	// State - Indicates whether this controller is active or not.
	State *NsControllerState `json:"state,omitempty"`
	// NvmeCardsEnabled - Indicates if the NVMe accelerator card is enabled.
	NvmeCardsEnabled *int64 `json:"nvme_cards_enabled,omitempty"`
	// NvmeCards - List of NVMe accelerator cards.
	NvmeCards []*NsCtrlrNvmeCard `json:"nvme_cards,omitempty"`
	// AsupTime - Time of the last autosupport by the controller.
	AsupTime *int64 `json:"asup_time,omitempty"`
}

// ControllerFieldHandles provides a string representation for each AccessControlRecord field.
type ControllerFieldHandles struct {
	ID                 string
	Name               string
	ArrayName          string
	ArrayId            string
	PartialResponseOk  string
	Serial             string
	Model              string
	Hostname           string
	SupportAddress     string
	SupportNetmask     string
	SupportNic         string
	PowerStatus        string
	FanStatus          string
	TemperatureStatus  string
	PowerSupplies      string
	Fans               string
	TemperatureSensors string
	PartitionStatus    string
	CtrlrSide          string
	State              string
	NvmeCardsEnabled   string
	NvmeCards          string
	AsupTime           string
}
