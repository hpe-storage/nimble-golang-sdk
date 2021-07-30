// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapVolListReturn - Object returned after creating snapshot collection.
// Export NsSnapVolListReturnFields for advance operations like search filter etc.
var NsSnapVolListReturnFields *NsSnapVolListReturnStringFields

func init() {
	SnapIdsfield := "snap_ids"

	NsSnapVolListReturnFields = &NsSnapVolListReturnStringFields{
		SnapIds: &SnapIdsfield,
	}
}

type NsSnapVolListReturn struct {
	// SnapIds - A list of snapshot ids.
	SnapIds []*NsObjectIDKV `json:"snap_ids,omitempty"`
}

// Struct for NsSnapVolListReturnFields
type NsSnapVolListReturnStringFields struct {
	SnapIds *string
}
