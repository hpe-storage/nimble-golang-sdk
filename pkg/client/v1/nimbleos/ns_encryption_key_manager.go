// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsEncryptionKeyManager Enum.

type NsEncryptionKeyManager string

const (
	cNsEncryptionKeyManagerExternal NsEncryptionKeyManager = "external"
	cNsEncryptionKeyManagerLocal    NsEncryptionKeyManager = "local"
)

var pNsEncryptionKeyManagerExternal NsEncryptionKeyManager
var pNsEncryptionKeyManagerLocal NsEncryptionKeyManager

// NsEncryptionKeyManagerExternal enum exports
var NsEncryptionKeyManagerExternal *NsEncryptionKeyManager

// NsEncryptionKeyManagerLocal enum exports
var NsEncryptionKeyManagerLocal *NsEncryptionKeyManager

func init() {
	pNsEncryptionKeyManagerExternal = cNsEncryptionKeyManagerExternal
	NsEncryptionKeyManagerExternal = &pNsEncryptionKeyManagerExternal

	pNsEncryptionKeyManagerLocal = cNsEncryptionKeyManagerLocal
	NsEncryptionKeyManagerLocal = &pNsEncryptionKeyManagerLocal

}
