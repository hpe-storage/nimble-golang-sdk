// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsPinnedVolumeInfo - Select fields containing pinned volume info.
// Export NsPinnedVolumeInfoFields for advance operations like search filter etc.
var NsPinnedVolumeInfoFields *NsPinnedVolumeInfo

func init(){
 IDfield:= "id"
 Namefield:= "name"

 NsPinnedVolumeInfoFields= &NsPinnedVolumeInfo{
  ID:                     &IDfield,
  Name:                   &Namefield,
 }
}


type NsPinnedVolumeInfo struct {
 // ID - ID of volume.
  ID *string `json:"id,omitempty"`
 // Name - Name of volume.
  Name *string `json:"name,omitempty"`
 // SizePinnedCacheBytes - Total pinned cache size of the volume in bytes.
    SizePinnedCacheBytes *int64 `json:"size_pinned_cache_bytes,omitempty"`
}
