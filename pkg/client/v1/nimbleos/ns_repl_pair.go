// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReplPair - Replicated objects (vol/snap/snapcoll), their UIDs are the same.

// Export NsReplPairFields provides field names to use in filter parameters, for example.
var NsReplPairFields *NsReplPairStringFields

func init() {
	fieldSrcName := "src_name"
	fieldDstName := "dst_name"

	NsReplPairFields = &NsReplPairStringFields{
		SrcName: &fieldSrcName,
		DstName: &fieldDstName,
	}
}

type NsReplPair struct {
	// SrcName - Name of the replicated obj on the source group.
	SrcName *string `json:"src_name,omitempty"`
	// DstName - Name of the replicated obj on the destination group.
	DstName *string `json:"dst_name,omitempty"`
}

// Struct for NsReplPairFields
type NsReplPairStringFields struct {
	SrcName *string
	DstName *string
}
