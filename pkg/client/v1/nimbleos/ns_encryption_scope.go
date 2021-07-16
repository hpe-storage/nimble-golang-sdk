// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsEncryptionScope Enum.

type NsEncryptionScope string

const (
	cNsEncryptionScopeVolume NsEncryptionScope = "volume"
	cNsEncryptionScopePool   NsEncryptionScope = "pool"
	cNsEncryptionScopeNone   NsEncryptionScope = "none"
	cNsEncryptionScopeGroup  NsEncryptionScope = "group"
)

var pNsEncryptionScopeVolume NsEncryptionScope
var pNsEncryptionScopePool NsEncryptionScope
var pNsEncryptionScopeNone NsEncryptionScope
var pNsEncryptionScopeGroup NsEncryptionScope

// NsEncryptionScopeVolume enum exports
var NsEncryptionScopeVolume *NsEncryptionScope

// NsEncryptionScopePool enum exports
var NsEncryptionScopePool *NsEncryptionScope

// NsEncryptionScopeNone enum exports
var NsEncryptionScopeNone *NsEncryptionScope

// NsEncryptionScopeGroup enum exports
var NsEncryptionScopeGroup *NsEncryptionScope

func init() {
	pNsEncryptionScopeVolume = cNsEncryptionScopeVolume
	NsEncryptionScopeVolume = &pNsEncryptionScopeVolume

	pNsEncryptionScopePool = cNsEncryptionScopePool
	NsEncryptionScopePool = &pNsEncryptionScopePool

	pNsEncryptionScopeNone = cNsEncryptionScopeNone
	NsEncryptionScopeNone = &pNsEncryptionScopeNone

	pNsEncryptionScopeGroup = cNsEncryptionScopeGroup
	NsEncryptionScopeGroup = &pNsEncryptionScopeGroup

}
