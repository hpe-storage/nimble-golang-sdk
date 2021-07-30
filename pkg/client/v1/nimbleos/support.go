// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Support - View and alter support-based parameters.

// Export SupportFields provides field names to use in filter parameters, for example.
var SupportFields *SupportStringFields

func init() {
	fieldID := "id"
	fieldPasswordMode := "password_mode"
	fieldArrayCount := "array_count"
	fieldArrayList := "array_list"

	SupportFields = &SupportStringFields{
		ID:           &fieldID,
		PasswordMode: &fieldPasswordMode,
		ArrayCount:   &fieldArrayCount,
		ArrayList:    &fieldArrayList,
	}
}

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

// Struct for SupportFields
type SupportStringFields struct {
	ID           *string
	PasswordMode *string
	ArrayCount   *string
	ArrayList    *string
}
