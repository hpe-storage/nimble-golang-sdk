/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Token - Manage user's session information.
// Export TokenFields for advance operations like search filter etc.
var TokenFields *Token

func init(){
	IDfield:= "id"
	SessionTokenfield:= "session_token"
	Usernamefield:= "username"
	Passwordfield:= "password"
	AppNamefield:= "app_name"
	SdkNamefield:= "sdk_name"
	SourceIpfield:= "source_ip"
	ServerUuIDfield:= "server_uuid"
		
	TokenFields= &Token{
		ID: &IDfield,
		SessionToken: &SessionTokenfield,
		Username: &Usernamefield,
		Password: &Passwordfield,
		AppName: &AppNamefield,
		SdkName: &SdkNamefield,
		SourceIp: &SourceIpfield,
		ServerUuID: &ServerUuIDfield,
		
	}
}

type Token struct {
   
    // Object identifier for the session token.
    
 	ID *string `json:"id,omitempty"`
   
    // Token used for authentication.
    
 	SessionToken *string `json:"session_token,omitempty"`
   
    // User name for the session.
    
 	Username *string `json:"username,omitempty"`
   
    // Password for the user. A password is required for creating a token.
    
 	Password *string `json:"password,omitempty"`
   
    // Application name.
    
 	AppName *string `json:"app_name,omitempty"`
   
    // SDK name.
    
 	SdkName *string `json:"sdk_name,omitempty"`
   
    // IP address from which the session originates.
    
 	SourceIp *string `json:"source_ip,omitempty"`
   
    // Time when this token was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this token was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Time when this token will expire.
    
   	ExpiryTime *int64 `json:"expiry_time,omitempty"`
   
    // Non mandatory 36 character uuid returned by the server. Currently only the witness REST server returns one.
    
 	ServerUuID *string `json:"server_uuid,omitempty"`
}
