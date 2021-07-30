// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export SoftwareVersionFields provides field names to use in filter parameters, for example.
var SoftwareVersionFields *SoftwareVersionFieldHandles

func init() {
	fieldVersion := "version"
	fieldSignature := "signature"
	fieldName := "name"
	fieldStatus := "status"
	fieldTotalBytes := "total_bytes"
	fieldDownloadedBytes := "downloaded_bytes"
	fieldBlacklistReason := "blacklist_reason"
	fieldReleaseDate := "release_date"
	fieldIsManuallyDownloaded := "is_manually_downloaded"
	fieldReleaseStatus := "release_status"
	fieldNoPartialResponse := "no_partial_response"

	SoftwareVersionFields = &SoftwareVersionFieldHandles{
		Version:              &fieldVersion,
		Signature:            &fieldSignature,
		Name:                 &fieldName,
		Status:               &fieldStatus,
		TotalBytes:           &fieldTotalBytes,
		DownloadedBytes:      &fieldDownloadedBytes,
		BlacklistReason:      &fieldBlacklistReason,
		ReleaseDate:          &fieldReleaseDate,
		IsManuallyDownloaded: &fieldIsManuallyDownloaded,
		ReleaseStatus:        &fieldReleaseStatus,
		NoPartialResponse:    &fieldNoPartialResponse,
	}
}

// SoftwareVersion - Show the software version.
type SoftwareVersion struct {
	// Version - Software version, used as identifier in URL.
	Version *string `json:"version,omitempty"`
	// Signature - Keyed hash of download package.
	Signature *string `json:"signature,omitempty"`
	// Name - Name of version.
	Name *string `json:"name,omitempty"`
	// Status - Status of version.
	Status *string `json:"status,omitempty"`
	// TotalBytes - Size of version.
	TotalBytes *int64 `json:"total_bytes,omitempty"`
	// DownloadedBytes - Number of bytes downloaded for the version.
	DownloadedBytes *int64 `json:"downloaded_bytes,omitempty"`
	// BlacklistReason - Reason for blacklisting the version. Empty if version is not blacklisted.
	BlacklistReason *string `json:"blacklist_reason,omitempty"`
	// ReleaseDate - Date when software version was released.
	ReleaseDate *int64 `json:"release_date,omitempty"`
	// IsManuallyDownloaded - Whether or not the version was downloaded manually.
	IsManuallyDownloaded *bool `json:"is_manually_downloaded,omitempty"`
	// ReleaseStatus - Release status of software version.
	ReleaseStatus *string `json:"release_status,omitempty"`
	// NoPartialResponse - Indicate that it is not ok to provide partially available response.
	NoPartialResponse *bool `json:"no_partial_response,omitempty"`
}

// SoftwareVersionFieldHandles provides a string representation for each AccessControlRecord field.
type SoftwareVersionFieldHandles struct {
	Version              *string
	Signature            *string
	Name                 *string
	Status               *string
	TotalBytes           *string
	DownloadedBytes      *string
	BlacklistReason      *string
	ReleaseDate          *string
	IsManuallyDownloaded *string
	ReleaseStatus        *string
	NoPartialResponse    *string
}
