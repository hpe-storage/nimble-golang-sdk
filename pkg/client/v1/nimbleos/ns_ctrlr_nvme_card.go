// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCtrlrNvmeCard - NVMe accelerator card.

// Export NsCtrlrNvmeCardFields provides field names to use in filter parameters, for example.
var NsCtrlrNvmeCardFields *NsCtrlrNvmeCardStringFields

func init() {
	fieldSerialNumber := "serial_number"
	fieldSize := "size"
	fieldState := "state"

	NsCtrlrNvmeCardFields = &NsCtrlrNvmeCardStringFields{
		SerialNumber: &fieldSerialNumber,
		Size:         &fieldSize,
		State:        &fieldState,
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
