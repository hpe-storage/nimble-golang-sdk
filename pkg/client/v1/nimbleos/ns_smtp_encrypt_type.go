// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsSmtpEncryptType Enum.

type NsSmtpEncryptType string

const (
 cNsSmtpEncryptTypeStarttls NsSmtpEncryptType = "starttls"
 cNsSmtpEncryptTypeNone NsSmtpEncryptType = "none"
 cNsSmtpEncryptTypeSsl NsSmtpEncryptType = "ssl"
)

var pNsSmtpEncryptTypeStarttls NsSmtpEncryptType
var pNsSmtpEncryptTypeNone NsSmtpEncryptType
var pNsSmtpEncryptTypeSsl NsSmtpEncryptType

// NsSmtpEncryptTypeStarttls enum exports
var NsSmtpEncryptTypeStarttls *NsSmtpEncryptType

// NsSmtpEncryptTypeNone enum exports
var NsSmtpEncryptTypeNone *NsSmtpEncryptType

// NsSmtpEncryptTypeSsl enum exports
var NsSmtpEncryptTypeSsl *NsSmtpEncryptType

func init() {
 pNsSmtpEncryptTypeStarttls = cNsSmtpEncryptTypeStarttls
 NsSmtpEncryptTypeStarttls = &pNsSmtpEncryptTypeStarttls

 pNsSmtpEncryptTypeNone = cNsSmtpEncryptTypeNone
 NsSmtpEncryptTypeNone = &pNsSmtpEncryptTypeNone

 pNsSmtpEncryptTypeSsl = cNsSmtpEncryptTypeSsl
 NsSmtpEncryptTypeSsl = &pNsSmtpEncryptTypeSsl

}

