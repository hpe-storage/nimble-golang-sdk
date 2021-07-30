// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsSupportPasswordArrayFields provides field names to use in filter parameters, for example.
var NsSupportPasswordArrayFields *NsSupportPasswordArrayFieldHandles

func init() {
	fieldArrayName := "array_name"
	fieldArraySn := "array_sn"
	fieldModel := "model"
	fieldBlobCount := "blob_count"
	fieldBlobList := "blob_list"

	NsSupportPasswordArrayFields = &NsSupportPasswordArrayFieldHandles{
		ArrayName: &fieldArrayName,
		ArraySn:   &fieldArraySn,
		Model:     &fieldModel,
		BlobCount: &fieldBlobCount,
		BlobList:  &fieldBlobList,
	}
}

// NsSupportPasswordArray - Support password blobs for an array.
type NsSupportPasswordArray struct {
	// ArrayName - The name of the array.
	ArrayName *string `json:"array_name,omitempty"`
	// ArraySn - The serial number of the array.
	ArraySn *string `json:"array_sn,omitempty"`
	// Model - The model of the array.
	Model *string `json:"model,omitempty"`
	// BlobCount - The number of blobs stored for the array.
	BlobCount *int64 `json:"blob_count,omitempty"`
	// BlobList - The blobs stored for the array.
	BlobList []*NsSupportPasswordObject `json:"blob_list,omitempty"`
}

// NsSupportPasswordArrayFieldHandles provides a string representation for each AccessControlRecord field.
type NsSupportPasswordArrayFieldHandles struct {
	ArrayName *string
	ArraySn   *string
	Model     *string
	BlobCount *string
	BlobList  *string
}
