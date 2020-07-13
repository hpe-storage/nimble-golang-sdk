// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Controller - Controller is a redundant collection of hardware capable of running the array software.
// Export ControllerFields for advance operations like search filter etc.
var ControllerFields *Controller

func init() {

	ControllerFields = &Controller{
		ID:             "id",
		Name:           "name",
		ArrayName:      "array_name",
		ArrayId:        "array_id",
		Serial:         "serial",
		Hostname:       "hostname",
		SupportAddress: "support_address",
		SupportNetmask: "support_netmask",
		SupportNic:     "support_nic",
	}
}

type Controller struct {
	// ID - Identifier of the controller.
	ID string `json:"id,omitempty"`
	// Name - Name of the controller.
	Name string `json:"name,omitempty"`
	// ArrayName - Name of the array containing this controller.
	ArrayName string `json:"array_name,omitempty"`
	// ArrayId - Rest ID of the array containing this controller.
	ArrayId string `json:"array_id,omitempty"`
	// PartialResponseOk - Indicate that it is ok to provide partially available response.
	PartialResponseOk *bool `json:"partial_response_ok,omitempty"`
	// Serial - Serial number for this controller.
	Serial string `json:"serial,omitempty"`
	// Hostname - Host name for the controller.
	Hostname string `json:"hostname,omitempty"`
	// SupportAddress - IP address used for support.
	SupportAddress string `json:"support_address,omitempty"`
	// SupportNetmask - IP netmask used for support.
	SupportNetmask string `json:"support_netmask,omitempty"`
	// SupportNic - Network card used for support.
	SupportNic string `json:"support_nic,omitempty"`
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
	NvmeCardsEnabled int64 `json:"nvme_cards_enabled,omitempty"`
	// NvmeCards - List of NVMe accelerator cards.
	NvmeCards []*NsCtrlrNvmeCard `json:"nvme_cards,omitempty"`
	// AsupTime - Time of the last autosupport by the controller.
	AsupTime int64 `json:"asup_time,omitempty"`
}

// sdk internal struct
type controller struct {
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

// EncodeController - Transform Controller to controller type
func EncodeController(request interface{}) (*controller, error) {
	reqController := request.(*Controller)
	byte, err := json.Marshal(reqController)
	objPtr := &controller{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeController Transform controller to Controller type
func DecodeController(request interface{}) (*Controller, error) {
	reqController := request.(*controller)
	byte, err := json.Marshal(reqController)
	obj := &Controller{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
