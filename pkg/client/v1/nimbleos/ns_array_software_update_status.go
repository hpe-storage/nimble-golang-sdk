// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArraySoftwareUpdateStatusFields provides field names to use in filter parameters, for example.
var NsArraySoftwareUpdateStatusFields *NsArraySoftwareUpdateStatusFieldHandles

func init() {
	NsArraySoftwareUpdateStatusFields = &NsArraySoftwareUpdateStatusFieldHandles{
		ArrayName:      "array_name",
		Error:          "error",
		CtrlrErrorMask: "ctrlr_error_mask",
		CtrlrAErrs:     "ctrlr_a_errs",
		CtrlrBErrs:     "ctrlr_b_errs",
	}
}

// NsArraySoftwareUpdateStatus - List of software update errors for a specific array.
type NsArraySoftwareUpdateStatus struct {
	// ArrayName - Name of the array.
	ArrayName *string `json:"array_name,omitempty"`
	// Error - Top level software update error for the array.
	Error *string `json:"error,omitempty"`
	// CtrlrErrorMask - Error mask for the controller.
	CtrlrErrorMask *int64 `json:"ctrlr_error_mask,omitempty"`
	// CtrlrAErrs - List of software update errors for controller A. The key is always "error" and the value is the actual error code string.
	CtrlrAErrs []*NsArraySoftwareUpdateError `json:"ctrlr_a_errs,omitempty"`
	// CtrlrBErrs - List of software update errors for controller B. The key is always "error" and the value is the actual error code string.
	CtrlrBErrs []*NsArraySoftwareUpdateError `json:"ctrlr_b_errs,omitempty"`
}

// NsArraySoftwareUpdateStatusFieldHandles provides a string representation for each AccessControlRecord field.
type NsArraySoftwareUpdateStatusFieldHandles struct {
	ArrayName      string
	Error          string
	CtrlrErrorMask string
	CtrlrAErrs     string
	CtrlrBErrs     string
}
