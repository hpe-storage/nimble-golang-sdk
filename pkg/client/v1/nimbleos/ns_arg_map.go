// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArgMapFields provides field names to use in filter parameters, for example.
var NsArgMapFields *NsArgMapFieldHandles

func init() {
	NsArgMapFields = &NsArgMapFieldHandles{}
}

// NsArgMap - Just a string to string map.
type NsArgMap struct {
}

// NsArgMapFieldHandles provides a string representation for each AccessControlRecord field.
type NsArgMapFieldHandles struct {
}
