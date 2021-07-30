// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export AutosupportFields provides field names to use in filter parameters, for example.
var AutosupportFields *AutosupportFieldHandles

func init() {
	fieldID := "id"
	fieldArrayList := "array_list"
	fieldArrayCount := "array_count"
	fieldGroupId := "group_id"
	fieldGroupName := "group_name"

	AutosupportFields = &AutosupportFieldHandles{
		ID:         &fieldID,
		ArrayList:  &fieldArrayList,
		ArrayCount: &fieldArrayCount,
		GroupId:    &fieldGroupId,
		GroupName:  &fieldGroupName,
	}
}

// Autosupport - Get status of autosupport.
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

// AutosupportFieldHandles provides a string representation for each AccessControlRecord field.
type AutosupportFieldHandles struct {
	ID         *string
	ArrayList  *string
	ArrayCount *string
	GroupId    *string
	GroupName  *string
}
