// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// SoftwareVersion - Show the software version.
// Export SoftwareVersionFields for advance operations like search filter etc.
var SoftwareVersionFields *SoftwareVersionStringFields

func init() {
	Versionfield := "version"
	Signaturefield := "signature"
	Namefield := "name"
	Statusfield := "status"
	TotalBytesfield := "total_bytes"
	DownloadedBytesfield := "downloaded_bytes"
	BlacklistReasonfield := "blacklist_reason"
	ReleaseDatefield := "release_date"
	IsManuallyDownloadedfield := "is_manually_downloaded"
	ReleaseStatusfield := "release_status"
	NoPartialResponsefield := "no_partial_response"

	SoftwareVersionFields = &SoftwareVersionStringFields{
		Version:              &Versionfield,
		Signature:            &Signaturefield,
		Name:                 &Namefield,
		Status:               &Statusfield,
		TotalBytes:           &TotalBytesfield,
		DownloadedBytes:      &DownloadedBytesfield,
		BlacklistReason:      &BlacklistReasonfield,
		ReleaseDate:          &ReleaseDatefield,
		IsManuallyDownloaded: &IsManuallyDownloadedfield,
		ReleaseStatus:        &ReleaseStatusfield,
		NoPartialResponse:    &NoPartialResponsefield,
	}
}

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

// Struct for SoftwareVersionFields
type SoftwareVersionStringFields struct {
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
