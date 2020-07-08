// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsTargetScope Enum.

type NsTargetScope string

const (
	cNsTargetScopeVolume NsTargetScope = "volume"
	cNsTargetScopeGroup  NsTargetScope = "group"
)

var pNsTargetScopeVolume NsTargetScope
var pNsTargetScopeGroup NsTargetScope

// Export
var NsTargetScopeVolume *NsTargetScope
var NsTargetScopeGroup *NsTargetScope

func init() {
	pNsTargetScopeVolume = cNsTargetScopeVolume
	NsTargetScopeVolume = &pNsTargetScopeVolume

	pNsTargetScopeGroup = cNsTargetScopeGroup
	NsTargetScopeGroup = &pNsTargetScopeGroup

}
