// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// SoftwareVersion - Show the software version.
// Export SoftwareVersionFields for advance operations like search filter etc.
var SoftwareVersionFields *SoftwareVersion

func init() {
	Versionfield := "version"
	Signaturefield := "signature"
	Namefield := "name"
	Statusfield := "status"
	BlacklistReasonfield := "blacklist_reason"
	ReleaseStatusfield := "release_status"

	SoftwareVersionFields = &SoftwareVersion{
		Version:         &Versionfield,
		Signature:       &Signaturefield,
		Name:            &Namefield,
		Status:          &Statusfield,
		BlacklistReason: &BlacklistReasonfield,
		ReleaseStatus:   &ReleaseStatusfield,
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
