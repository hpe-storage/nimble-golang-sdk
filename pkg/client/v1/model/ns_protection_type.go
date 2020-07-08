// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsProtectionTypeUnprotected *NsProtectionType
var NsProtectionTypeRemote *NsProtectionType
var NsProtectionTypeLocal *NsProtectionType

func init() {
	pNsProtectionTypeUnprotected = cNsProtectionTypeUnprotected
	NsProtectionTypeUnprotected = &pNsProtectionTypeUnprotected

	pNsProtectionTypeRemote = cNsProtectionTypeRemote
	NsProtectionTypeRemote = &pNsProtectionTypeRemote

	pNsProtectionTypeLocal = cNsProtectionTypeLocal
	NsProtectionTypeLocal = &pNsProtectionTypeLocal

}
