// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsADReportStatusReturn - Status of the Active Directory domain.
// Export NsADReportStatusReturnFields for advance operations like search filter etc.
var NsADReportStatusReturnFields *NsADReportStatusReturnStringFields

func init() {
	Joinedfield := "joined"
	Enabledfield := "enabled"
	LocalServiceStatusfield := "local_service_status"
	RemoteServiceStatusfield := "remote_service_status"
	TrustValidfield := "trust_valid"

	NsADReportStatusReturnFields = &NsADReportStatusReturnStringFields{
		Joined:              &Joinedfield,
		Enabled:             &Enabledfield,
		LocalServiceStatus:  &LocalServiceStatusfield,
		RemoteServiceStatus: &RemoteServiceStatusfield,
		TrustValid:          &TrustValidfield,
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
