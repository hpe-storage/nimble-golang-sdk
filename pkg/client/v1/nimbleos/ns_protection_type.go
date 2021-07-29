// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsProtectionType Enum.

type NsProtectionType string

const (
	cNsProtectionTypeUnprotected NsProtectionType = "unprotected"
	cNsProtectionTypeRemote      NsProtectionType = "remote"
	cNsProtectionTypeLocal       NsProtectionType = "local"
)

var pNsProtectionTypeUnprotected NsProtectionType
var pNsProtectionTypeRemote NsProtectionType
var pNsProtectionTypeLocal NsProtectionType

// NsProtectionTypeUnprotected enum exports
var NsProtectionTypeUnprotected *NsProtectionType

// NsProtectionTypeRemote enum exports
var NsProtectionTypeRemote *NsProtectionType

// NsProtectionTypeLocal enum exports
var NsProtectionTypeLocal *NsProtectionType

func init() {
	pNsProtectionTypeUnprotected = cNsProtectionTypeUnprotected
	NsProtectionTypeUnprotected = &pNsProtectionTypeUnprotected

	pNsProtectionTypeRemote = cNsProtectionTypeRemote
	NsProtectionTypeRemote = &pNsProtectionTypeRemote

	pNsProtectionTypeLocal = cNsProtectionTypeLocal
	NsProtectionTypeLocal = &pNsProtectionTypeLocal

}
