/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsGenericResp - Response from generic app server.
// Export NsGenericRespFields for advance operations like search filter etc.
var NsGenericRespFields *NsGenericResp

func init(){
	GenericErrorfield:= "generic_error"
	GenericErrorMessagefield:= "generic_error_message"
	ConnMessagefield:= "conn_message"
		
	NsGenericRespFields= &NsGenericResp{
		GenericError: &GenericErrorfield,
		GenericErrorMessage: &GenericErrorMessagefield,
		ConnMessage: &ConnMessagefield,
		
	}
}

type NsGenericResp struct {
   
    // Error code from generic app server.
    
 	GenericError *string `json:"generic_error,omitempty"`
   
    // Detailed error message from generic app server.
    
 	GenericErrorMessage *string `json:"generic_error_message,omitempty"`
   
    // Is the connection status OK.
    
 	ConnStatusOk *bool `json:"conn_status_ok,omitempty"`
   
    // Detailed connection message.
    
 	ConnMessage *string `json:"conn_message,omitempty"`
}
