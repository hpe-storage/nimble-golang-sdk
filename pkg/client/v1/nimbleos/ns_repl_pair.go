// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsReplPairFields provides field names to use in filter parameters, for example.
var NsReplPairFields *NsReplPairFieldHandles

func init() {
	NsReplPairFields = &NsReplPairFieldHandles{
		SrcName: "src_name",
		DstName: "dst_name",
	}
}

// NsReplPair - Replicated objects (vol/snap/snapcoll), their UIDs are the same.
type NsReplPair struct {
	// SrcName - Name of the replicated obj on the source group.
	SrcName *string `json:"src_name,omitempty"`
	// DstName - Name of the replicated obj on the destination group.
	DstName *string `json:"dst_name,omitempty"`
}

// NsReplPairFieldHandles provides a string representation for each NsReplPair field.
type NsReplPairFieldHandles struct {
	SrcName string
	DstName string
}
