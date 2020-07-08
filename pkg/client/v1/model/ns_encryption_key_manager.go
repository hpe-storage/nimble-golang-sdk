// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsEncryptionKeyManager Enum.

type NsEncryptionKeyManager string

const (
	cNsEncryptionKeyManagerExternal NsEncryptionKeyManager = "external"
	cNsEncryptionKeyManagerLocal    NsEncryptionKeyManager = "local"
)

var pNsEncryptionKeyManagerExternal NsEncryptionKeyManager
var pNsEncryptionKeyManagerLocal NsEncryptionKeyManager

// Export
var NsEncryptionKeyManagerExternal *NsEncryptionKeyManager
var NsEncryptionKeyManagerLocal *NsEncryptionKeyManager

func init() {
	pNsEncryptionKeyManagerExternal = cNsEncryptionKeyManagerExternal
	NsEncryptionKeyManagerExternal = &pNsEncryptionKeyManagerExternal

	pNsEncryptionKeyManagerLocal = cNsEncryptionKeyManagerLocal
	NsEncryptionKeyManagerLocal = &pNsEncryptionKeyManagerLocal

}
