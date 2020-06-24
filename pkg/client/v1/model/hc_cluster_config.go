/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// HcClusterConfig - Configuration information for virtual appliance that provides highly available storage and compute.
// Export HcClusterConfigFields for advance operations like search filter etc.
var HcClusterConfigFields *HcClusterConfig

func init(){
	IDfield:= "id"
	UniqueIDfield:= "unique_id"
	Namefield:= "name"
	Descriptionfield:= "description"
	Usernamefield:= "username"
	Passwordfield:= "password"
		
	HcClusterConfigFields= &HcClusterConfig{
		ID: &IDfield,
		UniqueID: &UniqueIDfield,
		Name: &Namefield,
		Description: &Descriptionfield,
		Username: &Usernamefield,
		Password: &Passwordfield,
		
	}
}

type HcClusterConfig struct {
   
    // Identifier for the hc cluster config.
    
 	ID *string `json:"id,omitempty"`
   
    // Unique identifier for the HC component.
    
 	UniqueID *string `json:"unique_id,omitempty"`
   
    // Name for the HC component.
    
 	Name *string `json:"name,omitempty"`
   
    // Text description of HC component.
    
 	Description *string `json:"description,omitempty"`
   
    // HC component username.
    
 	Username *string `json:"username,omitempty"`
   
    // HC component password.
    
 	Password *string `json:"password,omitempty"`
   
    // HCI config type ({invalid|node|block|cluster}).
    
   	Type *NsHCIConfigType `json:"type,omitempty"`
   
    // Key-value pairs that augment a HC cluster config's attributes.
    
   	Metadata []*NsKeyValue `json:"metadata,omitempty"`
   
    // Time when this HC cluster configuration was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this HC cluster configuration was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
}
