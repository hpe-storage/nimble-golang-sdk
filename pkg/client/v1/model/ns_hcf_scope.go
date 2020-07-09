// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsHcfScope Enum.

type NsHcfScope string

const (
	cNsHcfScopeController NsHcfScope = "controller"
	cNsHcfScopeArray      NsHcfScope = "array"
	cNsHcfScopeGroup      NsHcfScope = "group"
)

var pNsHcfScopeController NsHcfScope
var pNsHcfScopeArray NsHcfScope
var pNsHcfScopeGroup NsHcfScope

// NsHcfScopeController enum exports
var NsHcfScopeController *NsHcfScope

// NsHcfScopeArray enum exports
var NsHcfScopeArray *NsHcfScope

// NsHcfScopeGroup enum exports
var NsHcfScopeGroup *NsHcfScope

func init() {
	pNsHcfScopeController = cNsHcfScopeController
	NsHcfScopeController = &pNsHcfScopeController

	pNsHcfScopeArray = cNsHcfScopeArray
	NsHcfScopeArray = &pNsHcfScopeArray

	pNsHcfScopeGroup = cNsHcfScopeGroup
	NsHcfScopeGroup = &pNsHcfScopeGroup

}
