// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// ApplicationCategoryFields provides field names to use in filter parameters, for example.
var ApplicationCategoryFields *ApplicationCategoryFieldHandles

func init() {
	ApplicationCategoryFields = &ApplicationCategoryFieldHandles{
		ID:            "id",
		Name:          "name",
		DedupeEnabled: "dedupe_enabled",
		CreationTime:  "creation_time",
		LastModified:  "last_modified",
	}
}

// ApplicationCategory - Provides the list of application categories that are available, to classify volumes depending on the applications that use them.
type ApplicationCategory struct {
	// ID - Identifier for the application category.
	ID *string `json:"id,omitempty"`
	// Name - Name of application category.
	Name *string `json:"name,omitempty"`
	// DedupeEnabled - Specifies if dedupe is enabled for performance policies associated with this application category.
	DedupeEnabled *bool `json:"dedupe_enabled,omitempty"`
	// CreationTime - Time when this application category was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this application category was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
}

// ApplicationCategoryFieldHandles provides a string representation for each ApplicationCategory field.
type ApplicationCategoryFieldHandles struct {
	ID            string
	Name          string
	DedupeEnabled string
	CreationTime  string
	LastModified  string
}
