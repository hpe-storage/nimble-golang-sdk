// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfIdentifyStatusReturnFields provides field names to use in filter parameters, for example.
var NsShelfIdentifyStatusReturnFields *NsShelfIdentifyStatusReturnFieldHandles

func init() {
	NsShelfIdentifyStatusReturnFields = &NsShelfIdentifyStatusReturnFieldHandles{
		Enabled: "enabled",
	}
}

// NsShelfIdentifyStatusReturn - Status of the shelf identifier.
type NsShelfIdentifyStatusReturn struct {
	// Enabled - Shelf identifier is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// NsShelfIdentifyStatusReturnFieldHandles provides a string representation for each NsShelfIdentifyStatusReturn field.
type NsShelfIdentifyStatusReturnFieldHandles struct {
	Enabled string
}
