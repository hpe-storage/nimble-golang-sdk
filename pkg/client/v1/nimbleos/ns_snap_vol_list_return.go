// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSnapVolListReturnFields provides field names to use in filter parameters, for example.
var NsSnapVolListReturnFields *NsSnapVolListReturnFieldHandles

func init() {
	fieldSnapIds := "snap_ids"

	NsSnapVolListReturnFields = &NsSnapVolListReturnFieldHandles{
		SnapIds: &fieldSnapIds,
	}
}

// NsSnapVolListReturn - Object returned after creating snapshot collection.
type NsSnapVolListReturn struct {
	// SnapIds - A list of snapshot ids.
	SnapIds []*NsObjectIDKV `json:"snap_ids,omitempty"`
}

// NsSnapVolListReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsSnapVolListReturnFieldHandles struct {
	SnapIds *string
}
