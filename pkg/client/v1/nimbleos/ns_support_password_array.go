// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsSupportPasswordArray - Support password blobs for an array.
// Export NsSupportPasswordArrayFields for advance operations like search filter etc.
var NsSupportPasswordArrayFields *NsSupportPasswordArray

func init(){
 ArrayNamefield:= "array_name"
 ArraySnfield:= "array_sn"
 Modelfield:= "model"

 NsSupportPasswordArrayFields= &NsSupportPasswordArray{
  ArrayName: &ArrayNamefield,
  ArraySn:   &ArraySnfield,
  Model:     &Modelfield,
 }
}


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
