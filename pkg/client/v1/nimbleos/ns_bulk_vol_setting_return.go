// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsBulkVolSettingReturn - Return codes for setting an attribute to a list of items.

// Export NsBulkVolSettingReturnFields provides field names to use in filter parameters, for example.
var NsBulkVolSettingReturnFields *NsBulkVolSettingReturnStringFields

func init() {
	fieldErrorCodes := "error_codes"

	NsBulkVolSettingReturnFields = &NsBulkVolSettingReturnStringFields{
		ErrorCodes: &fieldErrorCodes,
	}
}

type NsBulkVolSettingReturn struct {
	// ErrorCodes - Error codes for every element in a list of items.
	ErrorCodes []*string `json:"error_codes,omitempty"`
}

// Struct for NsBulkVolSettingReturnFields
type NsBulkVolSettingReturnStringFields struct {
	ErrorCodes *string
}
