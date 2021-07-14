// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsOTPType Enum.

type NsOTPType string

const (
 cNsOTPTypeGoogleAuth NsOTPType = "google_auth"
 cNsOTPTypeNone NsOTPType = "none"
)

var pNsOTPTypeGoogleAuth NsOTPType
var pNsOTPTypeNone NsOTPType

// NsOTPTypeGoogleAuth enum exports
var NsOTPTypeGoogleAuth *NsOTPType

// NsOTPTypeNone enum exports
var NsOTPTypeNone *NsOTPType

func init() {
 pNsOTPTypeGoogleAuth = cNsOTPTypeGoogleAuth
 NsOTPTypeGoogleAuth = &pNsOTPTypeGoogleAuth

 pNsOTPTypeNone = cNsOTPTypeNone
 NsOTPTypeNone = &pNsOTPTypeNone

}

