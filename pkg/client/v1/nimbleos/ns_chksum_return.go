// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsChksumReturn - Return computed checksum.

// Export NsChksumReturnFields provides field names to use in filter parameters, for example.
var NsChksumReturnFields *NsChksumReturnStringFields

func init() {
	fieldCksum := "cksum"

	NsChksumReturnFields = &NsChksumReturnStringFields{
		Cksum: &fieldCksum,
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
