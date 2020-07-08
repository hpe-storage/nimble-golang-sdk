// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsVolumeSummaryWithAppCategory - Select fields containing volume info.
// Export NsVolumeSummaryWithAppCategoryFields for advance operations like search filter etc.
var NsVolumeSummaryWithAppCategoryFields *NsVolumeSummaryWithAppCategory

func init() {

	NsVolumeSummaryWithAppCategoryFields = &NsVolumeSummaryWithAppCategory{
		ID:          "id",
		Name:        "name",
		AppCategory: "app_category",
		FullName:    "full_name",
	}
}

type NsVolumeSummaryWithAppCategory struct {
	// ID - ID of volume.
	ID string `json:"id,omitempty"`
	// Name - Name of volume.
	Name string `json:"name,omitempty"`
	// AppCategory - Application category that the volume belongs to.
	AppCategory string `json:"app_category,omitempty"`
	// FullName - Fully qualified name of volume.
	FullName string `json:"full_name,omitempty"`
	// Lun - LUN of volume. Secondary LUN if this is Virtual Volume.
	Lun int64 `json:"lun,omitempty"`
}

// sdk internal struct
type nsVolumeSummaryWithAppCategory struct {
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

// EncodeNsVolumeSummaryWithAppCategory - Transform NsVolumeSummaryWithAppCategory to nsVolumeSummaryWithAppCategory type
func EncodeNsVolumeSummaryWithAppCategory(request interface{}) (*nsVolumeSummaryWithAppCategory, error) {
	reqNsVolumeSummaryWithAppCategory := request.(*NsVolumeSummaryWithAppCategory)
	byte, err := json.Marshal(reqNsVolumeSummaryWithAppCategory)
	objPtr := &nsVolumeSummaryWithAppCategory{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsVolumeSummaryWithAppCategory Transform nsVolumeSummaryWithAppCategory to NsVolumeSummaryWithAppCategory type
func DecodeNsVolumeSummaryWithAppCategory(request interface{}) (*NsVolumeSummaryWithAppCategory, error) {
	reqNsVolumeSummaryWithAppCategory := request.(*nsVolumeSummaryWithAppCategory)
	byte, err := json.Marshal(reqNsVolumeSummaryWithAppCategory)
	obj := &NsVolumeSummaryWithAppCategory{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
