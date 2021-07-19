// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfIdentifyStatusReturn - Status of the shelf identifier.
// Export NsShelfIdentifyStatusReturnFields for advance operations like search filter etc.
var NsShelfIdentifyStatusReturnFields *NsShelfIdentifyStatusReturn

func init() {

	NsShelfIdentifyStatusReturnFields = &NsShelfIdentifyStatusReturn{}
}

type NsShelfIdentifyStatusReturn struct {
	// Enabled - Shelf identifier is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}
