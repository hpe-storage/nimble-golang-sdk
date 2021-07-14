// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsEncryptionMode Enum.

type NsEncryptionMode string

const (
 cNsEncryptionModeAvailable NsEncryptionMode = "available"
 cNsEncryptionModeNone NsEncryptionMode = "none"
 cNsEncryptionModeSecure NsEncryptionMode = "secure"
)

var pNsEncryptionModeAvailable NsEncryptionMode
var pNsEncryptionModeNone NsEncryptionMode
var pNsEncryptionModeSecure NsEncryptionMode

// NsEncryptionModeAvailable enum exports
var NsEncryptionModeAvailable *NsEncryptionMode

// NsEncryptionModeNone enum exports
var NsEncryptionModeNone *NsEncryptionMode

// NsEncryptionModeSecure enum exports
var NsEncryptionModeSecure *NsEncryptionMode

func init() {
 pNsEncryptionModeAvailable = cNsEncryptionModeAvailable
 NsEncryptionModeAvailable = &pNsEncryptionModeAvailable

 pNsEncryptionModeNone = cNsEncryptionModeNone
 NsEncryptionModeNone = &pNsEncryptionModeNone

 pNsEncryptionModeSecure = cNsEncryptionModeSecure
 NsEncryptionModeSecure = &pNsEncryptionModeSecure

}

