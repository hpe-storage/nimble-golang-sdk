// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsEncryptionCipher Enum.

type NsEncryptionCipher string

const (
	cNsEncryptionCipherAes256Xts NsEncryptionCipher = "aes_256_xts"
	cNsEncryptionCipherNone      NsEncryptionCipher = "none"
)

var pNsEncryptionCipherAes256Xts NsEncryptionCipher
var pNsEncryptionCipherNone NsEncryptionCipher

// NsEncryptionCipherAes256Xts enum exports
var NsEncryptionCipherAes256Xts *NsEncryptionCipher

// NsEncryptionCipherNone enum exports
var NsEncryptionCipherNone *NsEncryptionCipher

func init() {
	pNsEncryptionCipherAes256Xts = cNsEncryptionCipherAes256Xts
	NsEncryptionCipherAes256Xts = &pNsEncryptionCipherAes256Xts

	pNsEncryptionCipherNone = cNsEncryptionCipherNone
	NsEncryptionCipherNone = &pNsEncryptionCipherNone

}
