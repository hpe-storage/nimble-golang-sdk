// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsBulkVolSettingReturnFields provides field names to use in filter parameters, for example.
var NsBulkVolSettingReturnFields *NsBulkVolSettingReturnFieldHandles

func init() {
	NsBulkVolSettingReturnFields = &NsBulkVolSettingReturnFieldHandles{
		ErrorCodes: "error_codes",
	}
}

// NsBulkVolSettingReturn - Return codes for setting an attribute to a list of items.
type NsBulkVolSettingReturn struct {
	// ErrorCodes - Error codes for every element in a list of items.
	ErrorCodes []*string `json:"error_codes,omitempty"`
}

// NsBulkVolSettingReturnFieldHandles provides a string representation for each NsBulkVolSettingReturn field.
type NsBulkVolSettingReturnFieldHandles struct {
	ErrorCodes string
}
