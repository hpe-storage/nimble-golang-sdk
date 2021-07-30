// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsChksumReturnFields provides field names to use in filter parameters, for example.
var NsChksumReturnFields *NsChksumReturnFieldHandles

func init() {
	fieldCksum := "cksum"

	NsChksumReturnFields = &NsChksumReturnFieldHandles{
		Cksum: &fieldCksum,
	}
}

// NsChksumReturn - Return computed checksum.
type NsChksumReturn struct {
	// Cksum - Computed checksum.
	Cksum *string `json:"cksum,omitempty"`
}

// NsChksumReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsChksumReturnFieldHandles struct {
	Cksum *string
}
