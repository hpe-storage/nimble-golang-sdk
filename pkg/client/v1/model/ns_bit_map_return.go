// Copyright 2020 Hewlett Packard Enterprise Development LP

package model


// NsBitMapReturn - Return bitmap under certain request.
// Export NsBitMapReturnFields for advance operations like search filter etc.
var NsBitMapReturnFields *NsBitMapReturn

func init(){
	Bitmapfield:= "bitmap"
		
	NsBitMapReturnFields= &NsBitMapReturn{
	Bitmap: &Bitmapfield,
		
	}
}

type NsBitMapReturn struct {
	// Bitmap - Returned bitmap.
 	Bitmap *string `json:"bitmap,omitempty"`
}
