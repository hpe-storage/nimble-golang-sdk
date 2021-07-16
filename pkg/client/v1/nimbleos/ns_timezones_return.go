// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsTimezonesReturn - Group timezone list attribute.
// Export NsTimezonesReturnFields for advance operations like search filter etc.
var NsTimezonesReturnFields *NsTimezonesReturn

func init() {

	NsTimezonesReturnFields = &NsTimezonesReturn{}
}

type NsTimezonesReturn struct {
	// Timezones - Group timezone list.
	Timezones []*string `json:"timezones,omitempty"`
}
