// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfIdentifyStatusReturn - Status of the shelf identifier.

// Export NsShelfIdentifyStatusReturnFields provides field names to use in filter parameters, for example.
var NsShelfIdentifyStatusReturnFields *NsShelfIdentifyStatusReturnStringFields

func init() {
	fieldEnabled := "enabled"

	NsShelfIdentifyStatusReturnFields = &NsShelfIdentifyStatusReturnStringFields{
		Enabled: &fieldEnabled,
	}
}

type NsShelfIdentifyStatusReturn struct {
	// Enabled - Shelf identifier is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// Struct for NsShelfIdentifyStatusReturnFields
type NsShelfIdentifyStatusReturnStringFields struct {
	Enabled *string
}
