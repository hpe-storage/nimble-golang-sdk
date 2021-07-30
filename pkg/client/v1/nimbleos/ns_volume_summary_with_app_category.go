// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsVolumeSummaryWithAppCategoryFields provides field names to use in filter parameters, for example.
var NsVolumeSummaryWithAppCategoryFields *NsVolumeSummaryWithAppCategoryFieldHandles

func init() {
	NsVolumeSummaryWithAppCategoryFields = &NsVolumeSummaryWithAppCategoryFieldHandles{
		ID:          "id",
		Name:        "name",
		AppCategory: "app_category",
		FullName:    "full_name",
		Lun:         "lun",
	}
}

// NsVolumeSummaryWithAppCategory - Select fields containing volume info.
type NsVolumeSummaryWithAppCategory struct {
	// ID - ID of volume.
	ID *string `json:"id,omitempty"`
	// Name - Name of volume.
	Name *string `json:"name,omitempty"`
	// AppCategory - Application category that the volume belongs to.
	AppCategory *string `json:"app_category,omitempty"`
	// FullName - Fully qualified name of volume.
	FullName *string `json:"full_name,omitempty"`
	// Lun - LUN of volume. Secondary LUN if this is Virtual Volume.
	Lun *int64 `json:"lun,omitempty"`
}

// NsVolumeSummaryWithAppCategoryFieldHandles provides a string representation for each AccessControlRecord field.
type NsVolumeSummaryWithAppCategoryFieldHandles struct {
	ID          string
	Name        string
	AppCategory string
	FullName    string
	Lun         string
}
