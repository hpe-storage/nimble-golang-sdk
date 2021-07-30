// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// ApplicationCategory - Provides the list of application categories that are available, to classify volumes depending on the applications that use them.
// Export ApplicationCategoryFields for advance operations like search filter etc.
var ApplicationCategoryFields *ApplicationCategoryStringFields

func init() {
	IDfield := "id"
	Namefield := "name"
	DedupeEnabledfield := "dedupe_enabled"
	CreationTimefield := "creation_time"
	LastModifiedfield := "last_modified"

	ApplicationCategoryFields = &ApplicationCategoryStringFields{
		ID:            &IDfield,
		Name:          &Namefield,
		DedupeEnabled: &DedupeEnabledfield,
		CreationTime:  &CreationTimefield,
		LastModified:  &LastModifiedfield,
	}
}

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

// Struct for ApplicationCategoryFields
type ApplicationCategoryStringFields struct {
	ID            *string
	Name          *string
	DedupeEnabled *string
	CreationTime  *string
	LastModified  *string
}
