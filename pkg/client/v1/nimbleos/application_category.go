// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// ApplicationCategory - Provides the list of application categories that are available, to classify volumes depending on the applications that use them.
// Export ApplicationCategoryFields for advance operations like search filter etc.
var ApplicationCategoryFields *ApplicationCategory

func init() {

	ApplicationCategoryFields = &ApplicationCategory{
		ID:   "id",
		Name: "name",
	}
}

type ApplicationCategory struct {
	// ID - Identifier for the application category.
	ID string `json:"id,omitempty"`
	// Name - Name of application category.
	Name string `json:"name,omitempty"`
	// DedupeEnabled - Specifies if dedupe is enabled for performance policies associated with this application category.
	DedupeEnabled *bool `json:"dedupe_enabled,omitempty"`
	// CreationTime - Time when this application category was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this application category was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
}

// sdk internal struct
type applicationCategory struct {
	ID            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	DedupeEnabled *bool   `json:"dedupe_enabled,omitempty"`
	CreationTime  *int64  `json:"creation_time,omitempty"`
	LastModified  *int64  `json:"last_modified,omitempty"`
}

// EncodeApplicationCategory - Transform ApplicationCategory to applicationCategory type
func EncodeApplicationCategory(request interface{}) (*applicationCategory, error) {
	reqApplicationCategory := request.(*ApplicationCategory)
	byte, err := json.Marshal(reqApplicationCategory)
	if err != nil {
		return nil, err
	}
	respApplicationCategoryPtr := &applicationCategory{}
	err = json.Unmarshal(byte, respApplicationCategoryPtr)
	return respApplicationCategoryPtr, err
}

// DecodeApplicationCategory Transform applicationCategory to ApplicationCategory type
func DecodeApplicationCategory(request interface{}) (*ApplicationCategory, error) {
	reqApplicationCategory := request.(*applicationCategory)
	byte, err := json.Marshal(reqApplicationCategory)
	if err != nil {
		return nil, err
	}
	respApplicationCategory := &ApplicationCategory{}
	err = json.Unmarshal(byte, respApplicationCategory)
	return respApplicationCategory, err
}
