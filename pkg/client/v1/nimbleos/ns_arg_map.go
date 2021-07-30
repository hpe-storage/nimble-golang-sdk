// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArgMap - Just a string to string map.
// Export NsArgMapFields for advance operations like search filter etc.
var NsArgMapFields *NsArgMapStringFields

func init() {

	NsArgMapFields = &NsArgMapStringFields{}
}

type NsArgMap struct {
}

// Struct for NsArgMapFields
type NsArgMapStringFields struct {
}
