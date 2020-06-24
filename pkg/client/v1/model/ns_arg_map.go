/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsArgMap - Just a string to string map.
// Export NsArgMapFields for advance operations like search filter etc.
var NsArgMapFields *NsArgMap

func init(){
		
	NsArgMapFields= &NsArgMap{
		
	}
}

type NsArgMap struct {
}
