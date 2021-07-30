// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArraySoftwareUpdateStatus - List of software update errors for a specific array.
// Export NsArraySoftwareUpdateStatusFields for advance operations like search filter etc.
var NsArraySoftwareUpdateStatusFields *NsArraySoftwareUpdateStatusStringFields

func init() {
	ArrayNamefield := "array_name"
	Errorfield := "error"
	CtrlrErrorMaskfield := "ctrlr_error_mask"
	CtrlrAErrsfield := "ctrlr_a_errs"
	CtrlrBErrsfield := "ctrlr_b_errs"

	NsArraySoftwareUpdateStatusFields = &NsArraySoftwareUpdateStatusStringFields{
		ArrayName:      &ArrayNamefield,
		Error:          &Errorfield,
		CtrlrErrorMask: &CtrlrErrorMaskfield,
		CtrlrAErrs:     &CtrlrAErrsfield,
		CtrlrBErrs:     &CtrlrBErrsfield,
	}
}

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

// Struct for NsArraySoftwareUpdateStatusFields
type NsArraySoftwareUpdateStatusStringFields struct {
	ArrayName      *string
	Error          *string
	CtrlrErrorMask *string
	CtrlrAErrs     *string
	CtrlrBErrs     *string
}
