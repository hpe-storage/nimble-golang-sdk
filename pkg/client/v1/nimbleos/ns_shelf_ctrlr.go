// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsShelfCtrlr - A shelf controller information.
// Export NsShelfCtrlrFields for advance operations like search filter etc.
var NsShelfCtrlrFields *NsShelfCtrlr

func init() {

	NsShelfCtrlrFields = &NsShelfCtrlr{
		ExpSasAddr:   "exp_sas_addr",
		CachedSerial: "cached_serial",
	}
}

type NsShelfCtrlr struct {
	// ExpSasAddr - Expander SAS address.
	ExpSasAddr string `json:"exp_sas_addr,omitempty"`
	// CtrlrSide - Position of this controller in the chassis.
	CtrlrSide *NsShelfCtrlrSide `json:"ctrlr_side,omitempty"`
	// EncLocId - Location ID based on SAS topology.
	EncLocId int64 `json:"enc_loc_id,omitempty"`
	// CachedSerial - Cached serial.
	CachedSerial string `json:"cached_serial,omitempty"`
	// CtrlrSensorLastRun - The time of last valid sensor reading, in epoch seconds.
	CtrlrSensorLastRun int64 `json:"ctrlr_sensor_last_run,omitempty"`
	// HwMshipFailure - SES device hardware mastership failure.
	HwMshipFailure *bool `json:"hw_mship_failure,omitempty"`
	// HwMasterState - SES device hardware mastership state.
	HwMasterState *NsShelfSesMasterHwState `json:"hw_master_state,omitempty"`
	// SwMasterState - SES device software mastership state.
	SwMasterState *NsShelfSesMasterSwState `json:"sw_master_state,omitempty"`
	// CtrlrSensors - The list of individual sensor reading in this controller.
	CtrlrSensors []*NsShelfSensor `json:"ctrlr_sensors,omitempty"`
	// FanOverallStatus - The overall status for the fans on this controller.
	FanOverallStatus *NsShelfSensorState `json:"fan_overall_status,omitempty"`
	// TempOverallStatus - The overall status for the temperature of this controller.
	TempOverallStatus *NsShelfSensorState `json:"temp_overall_status,omitempty"`
	// CtrlrHwModel - Controller hardware model.
	CtrlrHwModel *NsShelfCtrlrModel `json:"ctrlr_hw_model,omitempty"`
	// PortInfo - Port info for each SAS port.
	PortInfo []*NsShelfPortInfo `json:"port_info,omitempty"`
	// ExtraAttributes - Extra attributes as attribute value pairs.
	ExtraAttributes []*string `json:"extra_attributes,omitempty"`
	// CtrlrAttrsetList - List of ctrlr attribute set for each logical controller.
	CtrlrAttrsetList []*NsShelfCtrlrAttrSet `json:"ctrlr_attrset_list,omitempty"`
	// IDentifyStatus - Status of chassis identifier.
	IDentifyStatus *bool `json:"identify_status,omitempty"`
}

// sdk internal struct
type nsShelfCtrlr struct {
	ExpSasAddr         *string                  `json:"exp_sas_addr,omitempty"`
	CtrlrSide          *NsShelfCtrlrSide        `json:"ctrlr_side,omitempty"`
	EncLocId           *int64                   `json:"enc_loc_id,omitempty"`
	CachedSerial       *string                  `json:"cached_serial,omitempty"`
	CtrlrSensorLastRun *int64                   `json:"ctrlr_sensor_last_run,omitempty"`
	HwMshipFailure     *bool                    `json:"hw_mship_failure,omitempty"`
	HwMasterState      *NsShelfSesMasterHwState `json:"hw_master_state,omitempty"`
	SwMasterState      *NsShelfSesMasterSwState `json:"sw_master_state,omitempty"`
	CtrlrSensors       []*NsShelfSensor         `json:"ctrlr_sensors,omitempty"`
	FanOverallStatus   *NsShelfSensorState      `json:"fan_overall_status,omitempty"`
	TempOverallStatus  *NsShelfSensorState      `json:"temp_overall_status,omitempty"`
	CtrlrHwModel       *NsShelfCtrlrModel       `json:"ctrlr_hw_model,omitempty"`
	PortInfo           []*NsShelfPortInfo       `json:"port_info,omitempty"`
	ExtraAttributes    []*string                `json:"extra_attributes,omitempty"`
	CtrlrAttrsetList   []*NsShelfCtrlrAttrSet   `json:"ctrlr_attrset_list,omitempty"`
	IDentifyStatus     *bool                    `json:"identify_status,omitempty"`
}

// EncodeNsShelfCtrlr - Transform NsShelfCtrlr to nsShelfCtrlr type
func EncodeNsShelfCtrlr(request interface{}) (*nsShelfCtrlr, error) {
	reqNsShelfCtrlr := request.(*NsShelfCtrlr)
	byte, err := json.Marshal(reqNsShelfCtrlr)
	if err != nil {
		return nil, err
	}
	respNsShelfCtrlrPtr := &nsShelfCtrlr{}
	err = json.Unmarshal(byte, respNsShelfCtrlrPtr)
	return respNsShelfCtrlrPtr, err
}

// DecodeNsShelfCtrlr Transform nsShelfCtrlr to NsShelfCtrlr type
func DecodeNsShelfCtrlr(request interface{}) (*NsShelfCtrlr, error) {
	reqNsShelfCtrlr := request.(*nsShelfCtrlr)
	byte, err := json.Marshal(reqNsShelfCtrlr)
	if err != nil {
		return nil, err
	}
	respNsShelfCtrlr := &NsShelfCtrlr{}
	err = json.Unmarshal(byte, respNsShelfCtrlr)
	return respNsShelfCtrlr, err
}
