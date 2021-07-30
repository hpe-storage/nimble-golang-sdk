// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcLinkInfo - Fibre Channel link information.

// Export NsFcLinkInfoFields provides field names to use in filter parameters, for example.
var NsFcLinkInfoFields *NsFcLinkInfoStringFields

func init() {
	fieldLinkSpeed := "link_speed"
	fieldMaxLinkSpeed := "max_link_speed"
	fieldLinkStatus := "link_status"
	fieldOperationalStatus := "operational_status"

	NsFcLinkInfoFields = &NsFcLinkInfoStringFields{
		LinkSpeed:         &fieldLinkSpeed,
		MaxLinkSpeed:      &fieldMaxLinkSpeed,
		LinkStatus:        &fieldLinkStatus,
		OperationalStatus: &fieldOperationalStatus,
	}
}

type NsFcLinkInfo struct {
	// LinkSpeed - Speed of the Fibre Channel link.
	LinkSpeed *NsPlatFcLinkSpeed `json:"link_speed,omitempty"`
	// MaxLinkSpeed - Maximum speed of the Fibre Channel link.
	MaxLinkSpeed *NsPlatFcLinkSpeed `json:"max_link_speed,omitempty"`
	// LinkStatus - Status of the Fibre Channel link.
	LinkStatus *NsPlatFcLinkStatus `json:"link_status,omitempty"`
	// OperationalStatus - Operational status of the Fibre Channel link.
	OperationalStatus *NsPlatFcOperationalStatus `json:"operational_status,omitempty"`
}

// Struct for NsFcLinkInfoFields
type NsFcLinkInfoStringFields struct {
	LinkSpeed         *string
	MaxLinkSpeed      *string
	LinkStatus        *string
	OperationalStatus *string
}
