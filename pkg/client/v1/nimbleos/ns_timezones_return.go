// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsTimezonesReturn - Group timezone list attribute.
// Export NsTimezonesReturnFields for advance operations like search filter etc.
var NsTimezonesReturnFields *NsTimezonesReturnStringFields

func init() {
	Timezonesfield := "timezones"

	NsTimezonesReturnFields = &NsTimezonesReturnStringFields{
		Timezones: &Timezonesfield,
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
