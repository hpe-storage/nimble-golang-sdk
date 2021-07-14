// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// NsArgMap - Just a string to string map.
// Export NsArgMapFields for advance operations like search filter etc.
var NsArgMapFields *NsArgMap

func init(){

 NsArgMapFields= &NsArgMap{
 }
}


type NsArgMap struct {
}
