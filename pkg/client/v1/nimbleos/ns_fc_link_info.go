// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcLinkInfo - Fibre Channel link information.
// Export NsFcLinkInfoFields for advance operations like search filter etc.
var NsFcLinkInfoFields *NsFcLinkInfoStringFields

func init() {
	LinkSpeedfield := "link_speed"
	MaxLinkSpeedfield := "max_link_speed"
	LinkStatusfield := "link_status"
	OperationalStatusfield := "operational_status"

	NsFcLinkInfoFields = &NsFcLinkInfoStringFields{
		LinkSpeed:         &LinkSpeedfield,
		MaxLinkSpeed:      &MaxLinkSpeedfield,
		LinkStatus:        &LinkStatusfield,
		OperationalStatus: &OperationalStatusfield,
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
