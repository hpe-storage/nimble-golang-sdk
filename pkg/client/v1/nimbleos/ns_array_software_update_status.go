// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArraySoftwareUpdateStatus - List of software update errors for a specific array.
// Export NsArraySoftwareUpdateStatusFields for advance operations like search filter etc.
var NsArraySoftwareUpdateStatusFields *NsArraySoftwareUpdateStatus

func init() {

	NsArraySoftwareUpdateStatusFields = &NsArraySoftwareUpdateStatus{
		ArrayName: "array_name",
		Error:     "error",
	}
}

type NsArraySoftwareUpdateStatus struct {
	// ArrayName - Name of the array.
	ArrayName string `json:"array_name,omitempty"`
	// Error - Top level software update error for the array.
	Error string `json:"error,omitempty"`
	// CtrlrErrorMask - Error mask for the controller.
	CtrlrErrorMask int64 `json:"ctrlr_error_mask,omitempty"`
	// CtrlrAErrs - List of software update errors for controller A. The key is always "error" and the value is the actual error code string.
	CtrlrAErrs []*NsArraySoftwareUpdateError `json:"ctrlr_a_errs,omitempty"`
	// CtrlrBErrs - List of software update errors for controller B. The key is always "error" and the value is the actual error code string.
	CtrlrBErrs []*NsArraySoftwareUpdateError `json:"ctrlr_b_errs,omitempty"`
}

// sdk internal struct
type nsArraySoftwareUpdateStatus struct {
	ArrayName      *string                       `json:"array_name,omitempty"`
	Error          *string                       `json:"error,omitempty"`
	CtrlrErrorMask *int64                        `json:"ctrlr_error_mask,omitempty"`
	CtrlrAErrs     []*NsArraySoftwareUpdateError `json:"ctrlr_a_errs,omitempty"`
	CtrlrBErrs     []*NsArraySoftwareUpdateError `json:"ctrlr_b_errs,omitempty"`
}

// EncodeNsArraySoftwareUpdateStatus - Transform NsArraySoftwareUpdateStatus to nsArraySoftwareUpdateStatus type
func EncodeNsArraySoftwareUpdateStatus(request interface{}) (*nsArraySoftwareUpdateStatus, error) {
	reqNsArraySoftwareUpdateStatus := request.(*NsArraySoftwareUpdateStatus)
	byte, err := json.Marshal(reqNsArraySoftwareUpdateStatus)
	if err != nil {
		return nil, err
	}
	respNsArraySoftwareUpdateStatusPtr := &nsArraySoftwareUpdateStatus{}
	err = json.Unmarshal(byte, respNsArraySoftwareUpdateStatusPtr)
	return respNsArraySoftwareUpdateStatusPtr, err
}

// DecodeNsArraySoftwareUpdateStatus Transform nsArraySoftwareUpdateStatus to NsArraySoftwareUpdateStatus type
func DecodeNsArraySoftwareUpdateStatus(request interface{}) (*NsArraySoftwareUpdateStatus, error) {
	reqNsArraySoftwareUpdateStatus := request.(*nsArraySoftwareUpdateStatus)
	byte, err := json.Marshal(reqNsArraySoftwareUpdateStatus)
	if err != nil {
		return nil, err
	}
	respNsArraySoftwareUpdateStatus := &NsArraySoftwareUpdateStatus{}
	err = json.Unmarshal(byte, respNsArraySoftwareUpdateStatus)
	return respNsArraySoftwareUpdateStatus, err
}
