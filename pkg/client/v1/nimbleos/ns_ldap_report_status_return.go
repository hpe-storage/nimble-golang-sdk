// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsLdapReportStatusReturnFields provides field names to use in filter parameters, for example.
var NsLdapReportStatusReturnFields *NsLdapReportStatusReturnFieldHandles

func init() {
	NsLdapReportStatusReturnFields = &NsLdapReportStatusReturnFieldHandles{
		Enabled:                 "enabled",
		LocalServiceStatusGood:  "local_service_status_good",
		RemoteServiceStatusGood: "remote_service_status_good",
		RemoteServiceStatus:     "remote_service_status",
	}
}

// NsLdapReportStatusReturn - A report of the current LDAP status for the group.
type NsLdapReportStatusReturn struct {
	// Enabled - Whether the LDAP domain enabled on the group.
	Enabled *bool `json:"enabled,omitempty"`
	// LocalServiceStatusGood - Whether the (group's) local LDAP service running in a good state.
	LocalServiceStatusGood *bool `json:"local_service_status_good,omitempty"`
	// RemoteServiceStatusGood - Whether there appears to be an issue with the remote LDAP service. For example, misconfigured/expired TLS certificates.
	RemoteServiceStatusGood *bool `json:"remote_service_status_good,omitempty"`
	// RemoteServiceStatus - If the remote service status is not good then this provides a descriptive status message.
	RemoteServiceStatus *string `json:"remote_service_status,omitempty"`
}

// NsLdapReportStatusReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsLdapReportStatusReturnFieldHandles struct {
	Enabled                 string
	LocalServiceStatusGood  string
	RemoteServiceStatusGood string
	RemoteServiceStatus     string
}
