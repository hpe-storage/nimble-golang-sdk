// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsADReportStatusReturn - Status of the Active Directory domain.
// Export NsADReportStatusReturnFields for advance operations like search filter etc.
var NsADReportStatusReturnFields *NsADReportStatusReturn

func init() {

	NsADReportStatusReturnFields = &NsADReportStatusReturn{}
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
