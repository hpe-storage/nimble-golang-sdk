// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeSummaryWithAppCategory - Select fields containing volume info.

// Export NsVolumeSummaryWithAppCategoryFields provides field names to use in filter parameters, for example.
var NsVolumeSummaryWithAppCategoryFields *NsVolumeSummaryWithAppCategoryStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldAppCategory := "app_category"
	fieldFullName := "full_name"
	fieldLun := "lun"

	NsVolumeSummaryWithAppCategoryFields = &NsVolumeSummaryWithAppCategoryStringFields{
		ID:          &fieldID,
		Name:        &fieldName,
		AppCategory: &fieldAppCategory,
		FullName:    &fieldFullName,
		Lun:         &fieldLun,
	}
}

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

// Struct for NsVolumeSummaryWithAppCategoryFields
type NsVolumeSummaryWithAppCategoryStringFields struct {
	ID          *string
	Name        *string
	AppCategory *string
	FullName    *string
	Lun         *string
}
