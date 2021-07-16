// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsChksumReturn - Return computed checksum.
// Export NsChksumReturnFields for advance operations like search filter etc.
var NsChksumReturnFields *NsChksumReturn

func init() {
	Cksumfield := "cksum"

	NsChksumReturnFields = &NsChksumReturn{
		Cksum: &Cksumfield,
	}
}

type NsChksumReturn struct {
	// Cksum - Computed checksum.
	Cksum *string `json:"cksum,omitempty"`
}
