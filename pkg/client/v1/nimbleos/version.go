// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Version - Show the API version.
// Export VersionFields for advance operations like search filter etc.
var VersionFields *VersionStringFields

func init() {
	Namefield := "name"
	SoftwareVersionfield := "software_version"

	VersionFields = &VersionStringFields{
		Name:            &Namefield,
		SoftwareVersion: &SoftwareVersionfield,
	}
}

type Version struct {
	// Name - API version number.
	Name *string `json:"name,omitempty"`
	// SoftwareVersion - Software version of the group.
	SoftwareVersion *string `json:"software_version,omitempty"`
}

// Struct for VersionFields
type VersionStringFields struct {
	Name            *string
	SoftwareVersion *string
}
