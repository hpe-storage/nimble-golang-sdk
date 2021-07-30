// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArgMap - Just a string to string map.

// Export NsArgMapFields provides field names to use in filter parameters, for example.
var NsArgMapFields *NsArgMapStringFields

func init() {

	NsArgMapFields = &NsArgMapStringFields{}
}

type NsArgMap struct {
}

// Struct for NsArgMapFields
type NsArgMapStringFields struct {
}
