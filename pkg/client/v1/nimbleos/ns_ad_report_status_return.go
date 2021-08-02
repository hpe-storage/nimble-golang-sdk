// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsADReportStatusReturnFields provides field names to use in filter parameters, for example.
var NsADReportStatusReturnFields *NsADReportStatusReturnFieldHandles

func init() {
	NsADReportStatusReturnFields = &NsADReportStatusReturnFieldHandles{
		Joined:              "joined",
		Enabled:             "enabled",
		LocalServiceStatus:  "local_service_status",
		RemoteServiceStatus: "remote_service_status",
		TrustValid:          "trust_valid",
	}
}

// NsADReportStatusReturn - Status of the Active Directory domain.
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

// NsADReportStatusReturnFieldHandles provides a string representation for each NsADReportStatusReturn field.
type NsADReportStatusReturnFieldHandles struct {
	Joined              string
	Enabled             string
	LocalServiceStatus  string
	RemoteServiceStatus string
	TrustValid          string
}
