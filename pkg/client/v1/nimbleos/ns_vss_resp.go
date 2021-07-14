// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsVssResp - Response from VSS app server.
// Export NsVssRespFields for advance operations like search filter etc.
var NsVssRespFields *NsVssResp

func init(){
 VssErrorfield:= "vss_error"
 VssErrorMessagefield:= "vss_error_message"

 NsVssRespFields= &NsVssResp{
  VssError:         &VssErrorfield,
  VssErrorMessage:  &VssErrorMessagefield,
 }
}


type NsVssResp struct {
 // VssError - Error code from VSS app server.
  VssError *string `json:"vss_error,omitempty"`
 // VssErrorMessage - Detailed error message from VSS app server.
  VssErrorMessage *string `json:"vss_error_message,omitempty"`
}
