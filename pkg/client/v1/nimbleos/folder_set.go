// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export FolderSetFields provides field names to use in filter parameters, for example.
var FolderSetFields *FolderSetFieldHandles

func init() {
	FolderSetFields = &FolderSetFieldHandles{
		ID:            "id",
		Name:          "name",
		FullName:      "full_name",
		SearchName:    "search_name",
		Description:   "description",
		CreationTime:  "creation_time",
		LastModified:  "last_modified",
		AppUuid:       "app_uuid",
		FolderList:    "folder_list",
		AppserverId:   "appserver_id",
		AppserverName: "appserver_name",
	}
}

// FolderSet - Folder set represents a set of folder each on separate pools that represent a group-scope datastore spanning the entire group.
type FolderSet struct {
	// ID - Identifier of the folder set.
	ID *string `json:"id,omitempty"`
	// Name - Name of folder set.
	Name *string `json:"name,omitempty"`
	// FullName - Fully qualified name of folder set in the group.
	FullName *string `json:"full_name,omitempty"`
	// SearchName - Name of folder set used for object search.
	SearchName *string `json:"search_name,omitempty"`
	// Description - Text description of folder set.
	Description *string `json:"description,omitempty"`
	// CreationTime - Time when this folder set was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this folder set was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
	// AppUuid - Application identifier of the folder set. If this attribute is not provided at creation, the system will generate one.
	AppUuid *string `json:"app_uuid,omitempty"`
	// FolderList - List of folders contained by the folder set. The folders must be of agent type VVol and share the same application server with the folder set.
	FolderList []*NsFolderSummary `json:"folder_list,omitempty"`
	// AppserverId - Identifier of the application server associated with the folder set and member folders.
	AppserverId *string `json:"appserver_id,omitempty"`
	// AppserverName - Name of the application server associated with the folder set and member folders.
	AppserverName *string `json:"appserver_name,omitempty"`
}

// FolderSetFieldHandles provides a string representation for each AccessControlRecord field.
type FolderSetFieldHandles struct {
	ID            string
	Name          string
	FullName      string
	SearchName    string
	Description   string
	CreationTime  string
	LastModified  string
	AppUuid       string
	FolderList    string
	AppserverId   string
	AppserverName string
}
