/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArrayAsupDetail - Detailed array asup information.
// Export NsArrayAsupDetailFields for advance operations like search filter etc.
var NsArrayAsupDetailFields *NsArrayAsupDetail

func init(){
	ArrayNamefield:= "array_name"
		
	NsArrayAsupDetailFields= &NsArrayAsupDetail{
		ArrayName: &ArrayNamefield,
		
	}
}

type NsArrayAsupDetail struct {
   
    // Unique name of array.
    
 	ArrayName *string `json:"array_name,omitempty"`
   
    // This indicates whether the ASUP setting is enabled or disabled.
    
 	AsupValIDate *bool `json:"asup_validate,omitempty"`
   
    // Indicates whether DNS resolution succeeded.
    
 	NameResolution *bool `json:"name_resolution,omitempty"`
   
    // Indicates whether connection from Management IP address to support server succeed.
    
 	PingfromMgmtip *bool `json:"pingfrom_mgmtip,omitempty"`
   
    // Indicates whether connection from Controller-A to support server succeed.
    
 	PingfromCtrlra *bool `json:"pingfrom_ctrlra,omitempty"`
   
    // Indicates whether connection from Controller-B to support server succeed.
    
 	PingfromCtrlrb *bool `json:"pingfrom_ctrlrb,omitempty"`
   
    // Indicates whether heartbeat to support server succeed.
    
 	Heartbeat *bool `json:"heartbeat,omitempty"`
   
    // A list of error messages.
    
   	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}
