// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsBitMapReturnFields provides field names to use in filter parameters, for example.
var NsBitMapReturnFields *NsBitMapReturnFieldHandles

func init() {
	NsBitMapReturnFields = &NsBitMapReturnFieldHandles{
		Bitmap: "bitmap",
	}
}

// NsBitMapReturn - Return bitmap under certain request.
type NsBitMapReturn struct {
	// Bitmap - Returned bitmap.
	Bitmap *string `json:"bitmap,omitempty"`
}

// NsBitMapReturnFieldHandles provides a string representation for each AccessControlRecord field.
type NsBitMapReturnFieldHandles struct {
	Bitmap string
}
