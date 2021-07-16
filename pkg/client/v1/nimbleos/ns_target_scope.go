// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsTargetScope Enum.

type NsTargetScope string

const (
	cNsTargetScopeVolume NsTargetScope = "volume"
	cNsTargetScopeGroup  NsTargetScope = "group"
)

var pNsTargetScopeVolume NsTargetScope
var pNsTargetScopeGroup NsTargetScope

// NsTargetScopeVolume enum exports
var NsTargetScopeVolume *NsTargetScope

// NsTargetScopeGroup enum exports
var NsTargetScopeGroup *NsTargetScope

func init() {
	pNsTargetScopeVolume = cNsTargetScopeVolume
	NsTargetScopeVolume = &pNsTargetScopeVolume

	pNsTargetScopeGroup = cNsTargetScopeGroup
	NsTargetScopeGroup = &pNsTargetScopeGroup

}
