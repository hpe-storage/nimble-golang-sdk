/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

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
   
    // Returned bitmap.
    
 	Bitmap *string `json:"bitmap,omitempty"`
}
