// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export VersionFields provides field names to use in filter parameters, for example.
var VersionFields *VersionFieldHandles

func init() {
	VersionFields = &VersionFieldHandles{
		Name:            "name",
		SoftwareVersion: "software_version",
	}
}

// Version - Show the API version.
type Version struct {
	// Name - API version number.
	Name *string `json:"name,omitempty"`
	// SoftwareVersion - Software version of the group.
	SoftwareVersion *string `json:"software_version,omitempty"`
}

// VersionFieldHandles provides a string representation for each AccessControlRecord field.
type VersionFieldHandles struct {
	Name            string
	SoftwareVersion string
}
