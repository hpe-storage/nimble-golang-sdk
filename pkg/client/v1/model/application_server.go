/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// ApplicationServer - An application server is an external agent that collaborates with an array to manage storage resources; for example, Volume Shadow Copy Service (VSS) or VMware.
// Export ApplicationServerFields for advance operations like search filter etc.
var ApplicationServerFields *ApplicationServer

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Hostnamefield:= "hostname"
	Usernamefield:= "username"
	Descriptionfield:= "description"
	Passwordfield:= "password"
		
	ApplicationServerFields= &ApplicationServer{
		ID: &IDfield,
		Name: &Namefield,
		Hostname: &Hostnamefield,
		Username: &Usernamefield,
		Description: &Descriptionfield,
		Password: &Passwordfield,
		
	}
}

type ApplicationServer struct {
   
    // Identifier for the application server.
    
 	ID *string `json:"id,omitempty"`
   
    // Name for the application server.
    
 	Name *string `json:"name,omitempty"`
   
    // Application server hostname.
    
 	Hostname *string `json:"hostname,omitempty"`
   
    // Application server port number.
    
   	Port *int64 `json:"port,omitempty"`
   
    // Application server username.
    
 	Username *string `json:"username,omitempty"`
   
    // Text description of application server.
    
 	Description *string `json:"description,omitempty"`
   
    // Application server password.
    
 	Password *string `json:"password,omitempty"`
   
    // Application server type ({invalid|vss|vmware|cisco|stack_vision|container_node}).
    
   	ServerType *NsAppServerType `json:"server_type,omitempty"`
   
    // Key-value pairs that augment an application server's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
   
    // Time when this application server was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this application server was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
}
