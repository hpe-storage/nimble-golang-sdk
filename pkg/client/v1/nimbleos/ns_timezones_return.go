// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsTimezonesReturn - Group timezone list attribute.

// Export NsTimezonesReturnFields provides field names to use in filter parameters, for example.
var NsTimezonesReturnFields *NsTimezonesReturnStringFields

func init() {
	fieldTimezones := "timezones"

	NsTimezonesReturnFields = &NsTimezonesReturnStringFields{
		Timezones: &fieldTimezones,
	}
}

type NsTimezonesReturn struct {
	// Timezones - Group timezone list.
	Timezones []*string `json:"timezones,omitempty"`
}

// Struct for NsTimezonesReturnFields
type NsTimezonesReturnStringFields struct {
	Timezones *string
}
