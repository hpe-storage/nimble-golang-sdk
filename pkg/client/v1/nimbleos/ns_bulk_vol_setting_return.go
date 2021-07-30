// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsBulkVolSettingReturn - Return codes for setting an attribute to a list of items.
// Export NsBulkVolSettingReturnFields for advance operations like search filter etc.
var NsBulkVolSettingReturnFields *NsBulkVolSettingReturnStringFields

func init() {
	ErrorCodesfield := "error_codes"

	NsBulkVolSettingReturnFields = &NsBulkVolSettingReturnStringFields{
		ErrorCodes: &ErrorCodesfield,
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
