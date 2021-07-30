// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Version - Show the API version.

// Export VersionFields provides field names to use in filter parameters, for example.
var VersionFields *VersionStringFields

func init() {
	fieldName := "name"
	fieldSoftwareVersion := "software_version"

	VersionFields = &VersionStringFields{
		Name:            &fieldName,
		SoftwareVersion: &fieldSoftwareVersion,
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
