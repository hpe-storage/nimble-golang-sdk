/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsShelfIDentifyStatusReturn - Status of the shelf identifier.
// Export NsShelfIDentifyStatusReturnFields for advance operations like search filter etc.
var NsShelfIDentifyStatusReturnFields *NsShelfIDentifyStatusReturn

func init(){
		
	NsShelfIDentifyStatusReturnFields= &NsShelfIDentifyStatusReturn{
		
	}
}

type NsShelfIDentifyStatusReturn struct {
   
    // Shelf identifier is enabled.
    
 	Enabled *bool `json:"enabled,omitempty"`
}
