// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsVolumeSummaryWithAppCategory - Select fields containing volume info.
// Export NsVolumeSummaryWithAppCategoryFields for advance operations like search filter etc.
var NsVolumeSummaryWithAppCategoryFields *NsVolumeSummaryWithAppCategory

func init() {
	IDfield := "id"
	Namefield := "name"
	AppCategoryfield := "app_category"
	FullNamefield := "full_name"

	NsVolumeSummaryWithAppCategoryFields = &NsVolumeSummaryWithAppCategory{
		ID:          &IDfield,
		Name:        &Namefield,
		AppCategory: &AppCategoryfield,
		FullName:    &FullNamefield,
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
