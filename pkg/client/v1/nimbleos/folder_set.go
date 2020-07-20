// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// FolderSet - Folder set represents a set of folder each on separate pools that represent a group-scope datastore spanning the entire group.
// Export FolderSetFields for advance operations like search filter etc.
var FolderSetFields *FolderSet

func init() {

	FolderSetFields = &FolderSet{
		ID:            "id",
		Name:          "name",
		FullName:      "full_name",
		SearchName:    "search_name",
		Description:   "description",
		AppUuid:       "app_uuid",
		AppserverId:   "appserver_id",
		AppserverName: "appserver_name",
	}
}

type FolderSet struct {
	// ID - Identifier of the folder set.
	ID string `json:"id,omitempty"`
	// Name - Name of folder set.
	Name string `json:"name,omitempty"`
	// FullName - Fully qualified name of folder set in the group.
	FullName string `json:"full_name,omitempty"`
	// SearchName - Name of folder set used for object search.
	SearchName string `json:"search_name,omitempty"`
	// Description - Text description of folder set.
	Description string `json:"description,omitempty"`
	// CreationTime - Time when this folder set was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this folder set was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
	// AppUuid - Application identifier of the folder set. If this attribute is not provided at creation, the system will generate one.
	AppUuid string `json:"app_uuid,omitempty"`
	// FolderList - List of folders contained by the folder set. The folders must be of agent type VVol and share the same application server with the folder set.
	FolderList []*NsFolderSummary `json:"folder_list,omitempty"`
	// AppserverId - Identifier of the application server associated with the folder set and member folders.
	AppserverId string `json:"appserver_id,omitempty"`
	// AppserverName - Name of the application server associated with the folder set and member folders.
	AppserverName string `json:"appserver_name,omitempty"`
}

// sdk internal struct
type folderSet struct {
	ID            *string            `json:"id,omitempty"`
	Name          *string            `json:"name,omitempty"`
	FullName      *string            `json:"full_name,omitempty"`
	SearchName    *string            `json:"search_name,omitempty"`
	Description   *string            `json:"description,omitempty"`
	CreationTime  *int64             `json:"creation_time,omitempty"`
	LastModified  *int64             `json:"last_modified,omitempty"`
	AppUuid       *string            `json:"app_uuid,omitempty"`
	FolderList    []*NsFolderSummary `json:"folder_list,omitempty"`
	AppserverId   *string            `json:"appserver_id,omitempty"`
	AppserverName *string            `json:"appserver_name,omitempty"`
}

// EncodeFolderSet - Transform FolderSet to folderSet type
func EncodeFolderSet(request interface{}) (*folderSet, error) {
	reqFolderSet := request.(*FolderSet)
	byte, err := json.Marshal(reqFolderSet)
	if err != nil {
		return nil, err
	}
	respFolderSetPtr := &folderSet{}
	err = json.Unmarshal(byte, respFolderSetPtr)
	return respFolderSetPtr, err
}

// DecodeFolderSet Transform folderSet to FolderSet type
func DecodeFolderSet(request interface{}) (*FolderSet, error) {
	reqFolderSet := request.(*folderSet)
	byte, err := json.Marshal(reqFolderSet)
	if err != nil {
		return nil, err
	}
	respFolderSet := &FolderSet{}
	err = json.Unmarshal(byte, respFolderSet)
	return respFolderSet, err
}
