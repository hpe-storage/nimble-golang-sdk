// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsEncryptionCipher Enum.

type NsEncryptionCipher string

const (
	cNsEncryptionCipherAes256Xts NsEncryptionCipher = "aes_256_xts"
	cNsEncryptionCipherNone      NsEncryptionCipher = "none"
)

var pNsEncryptionCipherAes256Xts NsEncryptionCipher
var pNsEncryptionCipherNone NsEncryptionCipher

// Export
var NsEncryptionCipherAes256Xts *NsEncryptionCipher
var NsEncryptionCipherNone *NsEncryptionCipher

func init() {
	pNsEncryptionCipherAes256Xts = cNsEncryptionCipherAes256Xts
	NsEncryptionCipherAes256Xts = &pNsEncryptionCipherAes256Xts

	pNsEncryptionCipherNone = cNsEncryptionCipherNone
	NsEncryptionCipherNone = &pNsEncryptionCipherNone

}
