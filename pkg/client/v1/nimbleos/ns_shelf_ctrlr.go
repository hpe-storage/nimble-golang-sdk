// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfCtrlr - A shelf controller information.

// Export NsShelfCtrlrFields provides field names to use in filter parameters, for example.
var NsShelfCtrlrFields *NsShelfCtrlrStringFields

func init() {
	fieldExpSasAddr := "exp_sas_addr"
	fieldCtrlrSide := "ctrlr_side"
	fieldEncLocId := "enc_loc_id"
	fieldCachedSerial := "cached_serial"
	fieldCtrlrSensorLastRun := "ctrlr_sensor_last_run"
	fieldHwMshipFailure := "hw_mship_failure"
	fieldHwMasterState := "hw_master_state"
	fieldSwMasterState := "sw_master_state"
	fieldCtrlrSensors := "ctrlr_sensors"
	fieldFanOverallStatus := "fan_overall_status"
	fieldTempOverallStatus := "temp_overall_status"
	fieldPsuOverallStatus := "psu_overall_status"
	fieldCtrlrHwModel := "ctrlr_hw_model"
	fieldPortInfo := "port_info"
	fieldExtraAttributes := "extra_attributes"
	fieldCtrlrAttrsetList := "ctrlr_attrset_list"
	fieldIDentifyStatus := "identify_status"

	NsShelfCtrlrFields = &NsShelfCtrlrStringFields{
		ExpSasAddr:         &fieldExpSasAddr,
		CtrlrSide:          &fieldCtrlrSide,
		EncLocId:           &fieldEncLocId,
		CachedSerial:       &fieldCachedSerial,
		CtrlrSensorLastRun: &fieldCtrlrSensorLastRun,
		HwMshipFailure:     &fieldHwMshipFailure,
		HwMasterState:      &fieldHwMasterState,
		SwMasterState:      &fieldSwMasterState,
		CtrlrSensors:       &fieldCtrlrSensors,
		FanOverallStatus:   &fieldFanOverallStatus,
		TempOverallStatus:  &fieldTempOverallStatus,
		PsuOverallStatus:   &fieldPsuOverallStatus,
		CtrlrHwModel:       &fieldCtrlrHwModel,
		PortInfo:           &fieldPortInfo,
		ExtraAttributes:    &fieldExtraAttributes,
		CtrlrAttrsetList:   &fieldCtrlrAttrsetList,
		IDentifyStatus:     &fieldIDentifyStatus,
	}
}

type NsShelfCtrlr struct {
	// ExpSasAddr - Expander SAS address.
	ExpSasAddr *string `json:"exp_sas_addr,omitempty"`
	// CtrlrSide - Position of this controller in the chassis.
	CtrlrSide *NsShelfCtrlrSide `json:"ctrlr_side,omitempty"`
	// EncLocId - Location ID based on SAS topology.
	EncLocId *int64 `json:"enc_loc_id,omitempty"`
	// CachedSerial - Cached serial.
	CachedSerial *string `json:"cached_serial,omitempty"`
	// CtrlrSensorLastRun - The time of last valid sensor reading, in epoch seconds.
	CtrlrSensorLastRun *int64 `json:"ctrlr_sensor_last_run,omitempty"`
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
	// PsuOverallStatus - The overall status for the PSU on this controller.
	PsuOverallStatus *NsShelfSensorState `json:"psu_overall_status,omitempty"`
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

// Struct for NsShelfCtrlrFields
type NsShelfCtrlrStringFields struct {
	ExpSasAddr         *string
	CtrlrSide          *string
	EncLocId           *string
	CachedSerial       *string
	CtrlrSensorLastRun *string
	HwMshipFailure     *string
	HwMasterState      *string
	SwMasterState      *string
	CtrlrSensors       *string
	FanOverallStatus   *string
	TempOverallStatus  *string
	PsuOverallStatus   *string
	CtrlrHwModel       *string
	PortInfo           *string
	ExtraAttributes    *string
	CtrlrAttrsetList   *string
	IDentifyStatus     *string
}
