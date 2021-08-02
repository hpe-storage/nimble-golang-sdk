// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsTimezonesReturnFields provides field names to use in filter parameters, for example.
var NsTimezonesReturnFields *NsTimezonesReturnFieldHandles

func init() {
	NsTimezonesReturnFields = &NsTimezonesReturnFieldHandles{
		Timezones: "timezones",
	}
}

// NsTimezonesReturn - Group timezone list attribute.
type NsTimezonesReturn struct {
	// Timezones - Group timezone list.
	Timezones []*string `json:"timezones,omitempty"`
}

// NsTimezonesReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsTimezonesReturnFieldHandles struct {
	Timezones string
}
