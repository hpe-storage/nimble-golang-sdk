// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// FolderSet - Folder set represents a set of folder each on separate pools that represent a group-scope datastore spanning the entire group.

// Export FolderSetFields provides field names to use in filter parameters, for example.
var FolderSetFields *FolderSetStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldFullName := "full_name"
	fieldSearchName := "search_name"
	fieldDescription := "description"
	fieldCreationTime := "creation_time"
	fieldLastModified := "last_modified"
	fieldAppUuid := "app_uuid"
	fieldFolderList := "folder_list"
	fieldAppserverId := "appserver_id"
	fieldAppserverName := "appserver_name"

	FolderSetFields = &FolderSetStringFields{
		ID:            &fieldID,
		Name:          &fieldName,
		FullName:      &fieldFullName,
		SearchName:    &fieldSearchName,
		Description:   &fieldDescription,
		CreationTime:  &fieldCreationTime,
		LastModified:  &fieldLastModified,
		AppUuid:       &fieldAppUuid,
		FolderList:    &fieldFolderList,
		AppserverId:   &fieldAppserverId,
		AppserverName: &fieldAppserverName,
	}
}

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

// Struct for FolderSetFields
type FolderSetStringFields struct {
	ID            *string
	Name          *string
	FullName      *string
	SearchName    *string
	Description   *string
	CreationTime  *string
	LastModified  *string
	AppUuid       *string
	FolderList    *string
	AppserverId   *string
	AppserverName *string
}
