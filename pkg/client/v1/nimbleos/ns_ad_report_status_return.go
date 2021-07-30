// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsADReportStatusReturn - Status of the Active Directory domain.

// Export NsADReportStatusReturnFields provides field names to use in filter parameters, for example.
var NsADReportStatusReturnFields *NsADReportStatusReturnStringFields

func init() {
	fieldJoined := "joined"
	fieldEnabled := "enabled"
	fieldLocalServiceStatus := "local_service_status"
	fieldRemoteServiceStatus := "remote_service_status"
	fieldTrustValid := "trust_valid"

	NsADReportStatusReturnFields = &NsADReportStatusReturnStringFields{
		Joined:              &fieldJoined,
		Enabled:             &fieldEnabled,
		LocalServiceStatus:  &fieldLocalServiceStatus,
		RemoteServiceStatus: &fieldRemoteServiceStatus,
		TrustValid:          &fieldTrustValid,
	}
}

type NsADReportStatusReturn struct {
	// Joined - Joined the Active Directory group.
	Joined *bool `json:"joined,omitempty"`
	// Enabled - Active Directory group is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// LocalServiceStatus - Status of the local service.
	LocalServiceStatus *bool `json:"local_service_status,omitempty"`
	// RemoteServiceStatus - Status of the remote service.
	RemoteServiceStatus *bool `json:"remote_service_status,omitempty"`
	// TrustValid - Trust is valid.
	TrustValid *bool `json:"trust_valid,omitempty"`
}

// Struct for NsADReportStatusReturnFields
type NsADReportStatusReturnStringFields struct {
	Joined              *string
	Enabled             *string
	LocalServiceStatus  *string
	RemoteServiceStatus *string
	TrustValid          *string
}
