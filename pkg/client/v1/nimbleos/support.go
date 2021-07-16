// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Support - View and alter support-based parameters.
// Export SupportFields for advance operations like search filter etc.
var SupportFields *Support

func init() {
	IDfield := "id"

	SupportFields = &Support{
		ID: &IDfield,
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
