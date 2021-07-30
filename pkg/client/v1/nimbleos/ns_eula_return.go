// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsEulaReturnFields provides field names to use in filter parameters, for example.
var NsEulaReturnFields *NsEulaReturnFieldHandles

func init() {
	NsEulaReturnFields = &NsEulaReturnFieldHandles{
		Eula: "eula",
	}
}

// NsEulaReturn - Return end-user license information.
type NsEulaReturn struct {
	// Eula - License information.
	Eula *string `json:"eula,omitempty"`
}

// NsEulaReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsEulaReturnFieldHandles struct {
	Eula string
}
