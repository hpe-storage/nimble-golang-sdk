// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsChksumReturn - Return computed checksum.
// Export NsChksumReturnFields for advance operations like search filter etc.
var NsChksumReturnFields *NsChksumReturnStringFields

func init() {
	Cksumfield := "cksum"

	NsChksumReturnFields = &NsChksumReturnStringFields{
		Cksum: &Cksumfield,
	}
}

type NsChksumReturn struct {
	// Cksum - Computed checksum.
	Cksum *string `json:"cksum,omitempty"`
}

// Struct for NsChksumReturnFields
type NsChksumReturnStringFields struct {
	Cksum *string
}
