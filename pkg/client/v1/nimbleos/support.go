// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export SupportFields provides field names to use in filter parameters, for example.
var SupportFields *SupportFieldHandles

func init() {
	fieldID := "id"
	fieldPasswordMode := "password_mode"
	fieldArrayCount := "array_count"
	fieldArrayList := "array_list"

	SupportFields = &SupportFieldHandles{
		ID:           &fieldID,
		PasswordMode: &fieldPasswordMode,
		ArrayCount:   &fieldArrayCount,
		ArrayList:    &fieldArrayList,
	}
}

// Support - View and alter support-based parameters.
type Support struct {
	// ID - Object identifier for the group.
	ID *string `json:"id,omitempty"`
	// PasswordMode - Mode for support password.
	PasswordMode *NsSupportPasswordMode `json:"password_mode,omitempty"`
	// ArrayCount - Count of arrays for support password.
	ArrayCount *int64 `json:"array_count,omitempty"`
	// ArrayList - Details of support passwords for arrays.
	ArrayList []*NsSupportPasswordArray `json:"array_list,omitempty"`
}

// SupportFieldHandles provides a string representation for each AccessControlRecord field.
type SupportFieldHandles struct {
	ID           *string
	PasswordMode *string
	ArrayCount   *string
	ArrayList    *string
}
