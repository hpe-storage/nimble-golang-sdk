/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsResponse - May contain anything that any REST API response can contain.
// Export NsResponseFields for advance operations like search filter etc.
var NsResponseFields *NsResponse

func init(){
		
	NsResponseFields= &NsResponse{
		
	}
}

type NsResponse struct {
   
    // Response data.
    
	Data *NsObject `json:"data,omitempty"`
   
    // A list of error messages.
    
   	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}
