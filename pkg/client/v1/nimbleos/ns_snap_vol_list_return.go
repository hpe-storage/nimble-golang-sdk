// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsSnapVolListReturn - Object returned after creating snapshot collection.

// Export NsSnapVolListReturnFields provides field names to use in filter parameters, for example.
var NsSnapVolListReturnFields *NsSnapVolListReturnStringFields

func init() {
	fieldSnapIds := "snap_ids"

	NsSnapVolListReturnFields = &NsSnapVolListReturnStringFields{
		SnapIds: &fieldSnapIds,
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
