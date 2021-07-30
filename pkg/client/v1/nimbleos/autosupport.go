// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Autosupport - Get status of autosupport.
// Export AutosupportFields for advance operations like search filter etc.
var AutosupportFields *AutosupportStringFields

func init() {
	IDfield := "id"
	ArrayListfield := "array_list"
	ArrayCountfield := "array_count"
	GroupIdfield := "group_id"
	GroupNamefield := "group_name"

	AutosupportFields = &AutosupportStringFields{
		ID:         &IDfield,
		ArrayList:  &ArrayListfield,
		ArrayCount: &ArrayCountfield,
		GroupId:    &GroupIdfield,
		GroupName:  &GroupNamefield,
	}
}

type Autosupport struct {
	// ID - Identifier of the autosupport.
	ID *string `json:"id,omitempty"`
	// ArrayList - List of arrays in the group with autosupport information.
	ArrayList []*NsArrayAsupDetail `json:"array_list,omitempty"`
	// ArrayCount - Number of arrays in the group.
	ArrayCount *int64 `json:"array_count,omitempty"`
	// GroupId - Identifier for the group.
	GroupId *string `json:"group_id,omitempty"`
	// GroupName - Name of the group.
	GroupName *string `json:"group_name,omitempty"`
}

// Struct for AutosupportFields
type AutosupportStringFields struct {
	ID         *string
	ArrayList  *string
	ArrayCount *string
	GroupId    *string
	GroupName  *string
}
