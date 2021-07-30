// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsLunReturnFields provides field names to use in filter parameters, for example.
var NsLunReturnFields *NsLunReturnFieldHandles

func init() {
	fieldLun := "lun"
	fieldLuNumber := "lu_number"

	NsLunReturnFields = &NsLunReturnFieldHandles{
		Lun:      &fieldLun,
		LuNumber: &fieldLuNumber,
	}
}

// NsLunReturn - Return LU number.
type NsLunReturn struct {
	// Lun - LU number in hexadecimal.
	Lun *int64 `json:"lun,omitempty"`
	// LuNumber - LU number in decimal.
	LuNumber *int64 `json:"lu_number,omitempty"`
}

// NsLunReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsLunReturnFieldHandles struct {
	Lun      *string
	LuNumber *string
}
