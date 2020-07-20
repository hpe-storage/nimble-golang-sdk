// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// SoftwareVersion - Show the software version.
// Export SoftwareVersionFields for advance operations like search filter etc.
var SoftwareVersionFields *SoftwareVersion

func init() {

	SoftwareVersionFields = &SoftwareVersion{
		Version:         "version",
		Signature:       "signature",
		Name:            "name",
		Status:          "status",
		BlacklistReason: "blacklist_reason",
		ReleaseStatus:   "release_status",
	}
}

type SoftwareVersion struct {
	// Version - Software version, used as identifier in URL.
	Version string `json:"version,omitempty"`
	// Signature - Keyed hash of download package.
	Signature string `json:"signature,omitempty"`
	// Name - Name of version.
	Name string `json:"name,omitempty"`
	// Status - Status of version.
	Status string `json:"status,omitempty"`
	// TotalBytes - Size of version.
	TotalBytes int64 `json:"total_bytes,omitempty"`
	// DownloadedBytes - Number of bytes downloaded for the version.
	DownloadedBytes int64 `json:"downloaded_bytes,omitempty"`
	// BlacklistReason - Reason for blacklisting the version. Empty if version is not blacklisted.
	BlacklistReason string `json:"blacklist_reason,omitempty"`
	// ReleaseDate - Date when software version was released.
	ReleaseDate int64 `json:"release_date,omitempty"`
	// IsManuallyDownloaded - Whether or not the version was downloaded manually.
	IsManuallyDownloaded *bool `json:"is_manually_downloaded,omitempty"`
	// ReleaseStatus - Release status of software version.
	ReleaseStatus string `json:"release_status,omitempty"`
	// NoPartialResponse - Indicate that it is not ok to provide partially available response.
	NoPartialResponse *bool `json:"no_partial_response,omitempty"`
}

// sdk internal struct
type softwareVersion struct {
	Version              *string `json:"version,omitempty"`
	Signature            *string `json:"signature,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Status               *string `json:"status,omitempty"`
	TotalBytes           *int64  `json:"total_bytes,omitempty"`
	DownloadedBytes      *int64  `json:"downloaded_bytes,omitempty"`
	BlacklistReason      *string `json:"blacklist_reason,omitempty"`
	ReleaseDate          *int64  `json:"release_date,omitempty"`
	IsManuallyDownloaded *bool   `json:"is_manually_downloaded,omitempty"`
	ReleaseStatus        *string `json:"release_status,omitempty"`
	NoPartialResponse    *bool   `json:"no_partial_response,omitempty"`
}

// EncodeSoftwareVersion - Transform SoftwareVersion to softwareVersion type
func EncodeSoftwareVersion(request interface{}) (*softwareVersion, error) {
	reqSoftwareVersion := request.(*SoftwareVersion)
	byte, err := json.Marshal(reqSoftwareVersion)
	if err != nil {
		return nil, err
	}
	respSoftwareVersionPtr := &softwareVersion{}
	err = json.Unmarshal(byte, respSoftwareVersionPtr)
	return respSoftwareVersionPtr, err
}

// DecodeSoftwareVersion Transform softwareVersion to SoftwareVersion type
func DecodeSoftwareVersion(request interface{}) (*SoftwareVersion, error) {
	reqSoftwareVersion := request.(*softwareVersion)
	byte, err := json.Marshal(reqSoftwareVersion)
	if err != nil {
		return nil, err
	}
	respSoftwareVersion := &SoftwareVersion{}
	err = json.Unmarshal(byte, respSoftwareVersion)
	return respSoftwareVersion, err
}
