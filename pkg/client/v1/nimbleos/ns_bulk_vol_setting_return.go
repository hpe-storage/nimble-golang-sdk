// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsBulkVolSettingReturn - Return codes for setting an attribute to a list of items.
// Export NsBulkVolSettingReturnFields for advance operations like search filter etc.
var NsBulkVolSettingReturnFields *NsBulkVolSettingReturn

func init() {

	NsBulkVolSettingReturnFields = &NsBulkVolSettingReturn{}
}

type NsBulkVolSettingReturn struct {
	// ErrorCodes - Error codes for every element in a list of items.
	ErrorCodes []*string `json:"error_codes,omitempty"`
}
