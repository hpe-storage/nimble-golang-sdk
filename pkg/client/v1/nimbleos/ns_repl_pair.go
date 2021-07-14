// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsReplPair - Replicated objects (vol/snap/snapcoll), their UIDs are the same.
// Export NsReplPairFields for advance operations like search filter etc.
var NsReplPairFields *NsReplPair

func init(){
 SrcNamefield:= "src_name"
 DstNamefield:= "dst_name"

 NsReplPairFields= &NsReplPair{
  SrcName: &SrcNamefield,
  DstName: &DstNamefield,
 }
}


type NsReplPair struct {
 // SrcName - Name of the replicated obj on the source group.
  SrcName *string `json:"src_name,omitempty"`
 // DstName - Name of the replicated obj on the destination group.
  DstName *string `json:"dst_name,omitempty"`
}
