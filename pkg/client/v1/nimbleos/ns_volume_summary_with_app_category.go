// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

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
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	AppCategory *string `json:"app_category,omitempty"`
	FullName    *string `json:"full_name,omitempty"`
	Lun         *int64  `json:"lun,omitempty"`
}

// EncodeNsVolumeSummaryWithAppCategory - Transform NsVolumeSummaryWithAppCategory to nsVolumeSummaryWithAppCategory type
func EncodeNsVolumeSummaryWithAppCategory(request interface{}) (*nsVolumeSummaryWithAppCategory, error) {
	reqNsVolumeSummaryWithAppCategory := request.(*NsVolumeSummaryWithAppCategory)
	byte, err := json.Marshal(reqNsVolumeSummaryWithAppCategory)
	if err != nil {
		return nil, err
	}
	respNsVolumeSummaryWithAppCategoryPtr := &nsVolumeSummaryWithAppCategory{}
	err = json.Unmarshal(byte, respNsVolumeSummaryWithAppCategoryPtr)
	return respNsVolumeSummaryWithAppCategoryPtr, err
}

// DecodeNsVolumeSummaryWithAppCategory Transform nsVolumeSummaryWithAppCategory to NsVolumeSummaryWithAppCategory type
func DecodeNsVolumeSummaryWithAppCategory(request interface{}) (*NsVolumeSummaryWithAppCategory, error) {
	reqNsVolumeSummaryWithAppCategory := request.(*nsVolumeSummaryWithAppCategory)
	byte, err := json.Marshal(reqNsVolumeSummaryWithAppCategory)
	if err != nil {
		return nil, err
	}
	respNsVolumeSummaryWithAppCategory := &NsVolumeSummaryWithAppCategory{}
	err = json.Unmarshal(byte, respNsVolumeSummaryWithAppCategory)
	return respNsVolumeSummaryWithAppCategory, err
}
