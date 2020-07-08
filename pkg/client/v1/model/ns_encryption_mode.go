// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsEncryptionMode Enum.

type NsEncryptionMode string

const (
	cNsEncryptionModeAvailable NsEncryptionMode = "available"
	cNsEncryptionModeNone      NsEncryptionMode = "none"
	cNsEncryptionModeSecure    NsEncryptionMode = "secure"
)

var pNsEncryptionModeAvailable NsEncryptionMode
var pNsEncryptionModeNone NsEncryptionMode
var pNsEncryptionModeSecure NsEncryptionMode

// Export
var NsEncryptionModeAvailable *NsEncryptionMode
var NsEncryptionModeNone *NsEncryptionMode
var NsEncryptionModeSecure *NsEncryptionMode

func init() {
	pNsEncryptionModeAvailable = cNsEncryptionModeAvailable
	NsEncryptionModeAvailable = &pNsEncryptionModeAvailable

	pNsEncryptionModeNone = cNsEncryptionModeNone
	NsEncryptionModeNone = &pNsEncryptionModeNone

	pNsEncryptionModeSecure = cNsEncryptionModeSecure
	NsEncryptionModeSecure = &pNsEncryptionModeSecure

}
