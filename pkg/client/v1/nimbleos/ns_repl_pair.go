// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsReplPairFields provides field names to use in filter parameters, for example.
var NsReplPairFields *NsReplPairFieldHandles

func init() {
	fieldSrcName := "src_name"
	fieldDstName := "dst_name"

	NsReplPairFields = &NsReplPairFieldHandles{
		SrcName: &fieldSrcName,
		DstName: &fieldDstName,
	}
}

// NsReplPair - Replicated objects (vol/snap/snapcoll), their UIDs are the same.
type NsReplPair struct {
	// SrcName - Name of the replicated obj on the source group.
	SrcName *string `json:"src_name,omitempty"`
	// DstName - Name of the replicated obj on the destination group.
	DstName *string `json:"dst_name,omitempty"`
}

// NsReplPairFieldHandles provides a string representation for each AccessControlRecord field.
type NsReplPairFieldHandles struct {
	SrcName *string
	DstName *string
}
