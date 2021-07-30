// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsLdapReportStatusReturn - A report of the current LDAP status for the group.

// Export NsLdapReportStatusReturnFields provides field names to use in filter parameters, for example.
var NsLdapReportStatusReturnFields *NsLdapReportStatusReturnStringFields

func init() {
	fieldEnabled := "enabled"
	fieldLocalServiceStatusGood := "local_service_status_good"
	fieldRemoteServiceStatusGood := "remote_service_status_good"
	fieldRemoteServiceStatus := "remote_service_status"

	NsLdapReportStatusReturnFields = &NsLdapReportStatusReturnStringFields{
		Enabled:                 &fieldEnabled,
		LocalServiceStatusGood:  &fieldLocalServiceStatusGood,
		RemoteServiceStatusGood: &fieldRemoteServiceStatusGood,
		RemoteServiceStatus:     &fieldRemoteServiceStatus,
	}
}

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

// Struct for NsLdapReportStatusReturnFields
type NsLdapReportStatusReturnStringFields struct {
	Enabled                 *string
	LocalServiceStatusGood  *string
	RemoteServiceStatusGood *string
	RemoteServiceStatus     *string
}
