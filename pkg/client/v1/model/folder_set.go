/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// FolderSet - Folder set represents a set of folder each on separate pools that represent a group-scope datastore spanning the entire group.
// Export FolderSetFields for advance operations like search filter etc.
var FolderSetFields *FolderSet

func init(){
	IDfield:= "id"
	Namefield:= "name"
	FullNamefield:= "full_name"
	SearchNamefield:= "search_name"
	Descriptionfield:= "description"
	AppUuIDfield:= "app_uuid"
	AppserverIDfield:= "appserver_id"
	AppserverNamefield:= "appserver_name"
		
	FolderSetFields= &FolderSet{
		ID: &IDfield,
		Name: &Namefield,
		FullName: &FullNamefield,
		SearchName: &SearchNamefield,
		Description: &Descriptionfield,
		AppUuID: &AppUuIDfield,
		AppserverID: &AppserverIDfield,
		AppserverName: &AppserverNamefield,
		
	}
}

type FolderSet struct {
   
    // Identifier of the folder set.
    
 	ID *string `json:"id,omitempty"`
   
    // Name of folder set.
    
 	Name *string `json:"name,omitempty"`
   
    // Fully qualified name of folder set in the group.
    
 	FullName *string `json:"full_name,omitempty"`
   
    // Name of folder set used for object search.
    
 	SearchName *string `json:"search_name,omitempty"`
   
    // Text description of folder set.
    
 	Description *string `json:"description,omitempty"`
   
    // Time when this folder set was created.
    
   	CreationTime *int64 `json:"creation_time,omitempty"`
   
    // Time when this folder set was last modified.
    
   	LastModified *int64 `json:"last_modified,omitempty"`
   
    // Application identifier of the folder set. If this attribute is not provided at creation, the system will generate one.
    
 	AppUuID *string `json:"app_uuid,omitempty"`
   
    // List of folders contained by the folder set. The folders must be of agent type VVol and share the same application server with the folder set.
    
   	FolderList []*NsFolderSummary `json:"folder_list,omitempty"`
   
    // Identifier of the application server associated with the folder set and member folders.
    
 	AppserverID *string `json:"appserver_id,omitempty"`
   
    // Name of the application server associated with the folder set and member folders.
    
 	AppserverName *string `json:"appserver_name,omitempty"`
}
