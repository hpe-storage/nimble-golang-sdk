/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Witness - Manage witness host configuration.
// Export WitnessFields for advance operations like search filter etc.
var WitnessFields *Witness

func init(){
	IDfield:= "id"
	Usernamefield:= "username"
	Passwordfield:= "password"
	Hostfield:= "host"
		
	WitnessFields= &Witness{
		ID: &IDfield,
		Username: &Usernamefield,
		Password: &Passwordfield,
		Host: &Hostfield,
		
	}
}

type Witness struct {
   
    // Identifier of the witness configuration.
    
 	ID *string `json:"id,omitempty"`
   
    // Username of witness. This has to be a non-root that can login to the witness host.
    
 	Username *string `json:"username,omitempty"`
   
    // Password of witness.
    
 	Password *string `json:"password,omitempty"`
   
    // Hostname or ip addresses of witness.
    
 	Host *string `json:"host,omitempty"`
   
    // Port of witness.
    
   	Port *int64 `json:"port,omitempty"`
   
    // To verify the witness host against CA cert and to apply possible security modes.
    
 	SecureMode *bool `json:"secure_mode,omitempty"`
   
    // List of validation messages for automatic switchover of Group Management. This will be empty when there are no conflicts found.
    
   	AutoSwitchoverMessages []*NsErrorWithArguments `json:"auto_switchover_messages,omitempty"`
}
