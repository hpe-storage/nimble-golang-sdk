// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCtrlrNvmeCard - NVMe accelerator card.
// Export NsCtrlrNvmeCardFields for advance operations like search filter etc.
var NsCtrlrNvmeCardFields *NsCtrlrNvmeCardStringFields

func init() {
	SerialNumberfield := "serial_number"
	Sizefield := "size"
	Statefield := "state"

	NsCtrlrNvmeCardFields = &NsCtrlrNvmeCardStringFields{
		SerialNumber: &SerialNumberfield,
		Size:         &Sizefield,
		State:        &Statefield,
	}
}

type NsCtrlrNvmeCard struct {
	// SerialNumber - Serial number.
	SerialNumber *string `json:"serial_number,omitempty"`
	// Size - NVMe card cache size in bytes.
	Size *int64 `json:"size,omitempty"`
	// State - Online state.
	State *NsCtrlrNvmeCardState `json:"state,omitempty"`
}

// Struct for NsCtrlrNvmeCardFields
type NsCtrlrNvmeCardStringFields struct {
	SerialNumber *string
	Size         *string
	State        *string
}
