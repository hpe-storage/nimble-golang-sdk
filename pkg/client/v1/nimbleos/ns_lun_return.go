// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsLunReturn - Return LU number.

// Export NsLunReturnFields provides field names to use in filter parameters, for example.
var NsLunReturnFields *NsLunReturnStringFields

func init() {
	fieldLun := "lun"
	fieldLuNumber := "lu_number"

	NsLunReturnFields = &NsLunReturnStringFields{
		Lun:      &fieldLun,
		LuNumber: &fieldLuNumber,
	}
}

type NsLunReturn struct {
	// Lun - LU number in hexadecimal.
	Lun *int64 `json:"lun,omitempty"`
	// LuNumber - LU number in decimal.
	LuNumber *int64 `json:"lu_number,omitempty"`
}

// Struct for NsLunReturnFields
type NsLunReturnStringFields struct {
	Lun      *string
	LuNumber *string
}
