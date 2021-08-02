// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsFcLinkInfoFields provides field names to use in filter parameters, for example.
var NsFcLinkInfoFields *NsFcLinkInfoFieldHandles

func init() {
	NsFcLinkInfoFields = &NsFcLinkInfoFieldHandles{
		LinkSpeed:         "link_speed",
		MaxLinkSpeed:      "max_link_speed",
		LinkStatus:        "link_status",
		OperationalStatus: "operational_status",
	}
}

// NsFcLinkInfo - Fibre Channel link information.
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

// NsFcLinkInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsFcLinkInfoFieldHandles struct {
	LinkSpeed         string
	MaxLinkSpeed      string
	LinkStatus        string
	OperationalStatus string
}
