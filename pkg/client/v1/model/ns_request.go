/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsRequest - May contain anything that any REST API request can contain.
// Export NsRequestFields for advance operations like search filter etc.
var NsRequestFields *NsRequest

func init(){
	Pathfield:= "path"
		
	NsRequestFields= &NsRequest{
		Path: &Pathfield,
		
	}
}

type NsRequest struct {
   
    // Request data.
    
	Data *NsObject `json:"data,omitempty"`
   
    // HTTP method.
    
   	Method *int64 `json:"method,omitempty"`
   
    // Path which identifies the target resource.
    
 	Path *string `json:"path,omitempty"`
}
