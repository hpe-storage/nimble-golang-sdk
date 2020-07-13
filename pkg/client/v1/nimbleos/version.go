// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// Version - Show the API version.
// Export VersionFields for advance operations like search filter etc.
var VersionFields *Version

func init() {

	VersionFields = &Version{
		Name:            "name",
		SoftwareVersion: "software_version",
	}
}

type Version struct {
	// Name - API version number.
	Name string `json:"name,omitempty"`
	// SoftwareVersion - Software version of the group.
	SoftwareVersion string `json:"software_version,omitempty"`
}

// sdk internal struct
type version struct {
	// Name - API version number.
	Name *string `json:"name,omitempty"`
	// SoftwareVersion - Software version of the group.
	SoftwareVersion *string `json:"software_version,omitempty"`
}

// EncodeVersion - Transform Version to version type
func EncodeVersion(request interface{}) (*version, error) {
	reqVersion := request.(*Version)
	byte, err := json.Marshal(reqVersion)
	objPtr := &version{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeVersion Transform version to Version type
func DecodeVersion(request interface{}) (*Version, error) {
	reqVersion := request.(*version)
	byte, err := json.Marshal(reqVersion)
	obj := &Version{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
