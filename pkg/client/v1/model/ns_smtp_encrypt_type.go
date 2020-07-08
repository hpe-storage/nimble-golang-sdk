// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsSmtpEncryptType Enum.

type NsSmtpEncryptType string

const (
	cNsSmtpEncryptTypeStarttls NsSmtpEncryptType = "starttls"
	cNsSmtpEncryptTypeNone     NsSmtpEncryptType = "none"
	cNsSmtpEncryptTypeSsl      NsSmtpEncryptType = "ssl"
)

var pNsSmtpEncryptTypeStarttls NsSmtpEncryptType
var pNsSmtpEncryptTypeNone NsSmtpEncryptType
var pNsSmtpEncryptTypeSsl NsSmtpEncryptType

// Export
var NsSmtpEncryptTypeStarttls *NsSmtpEncryptType
var NsSmtpEncryptTypeNone *NsSmtpEncryptType
var NsSmtpEncryptTypeSsl *NsSmtpEncryptType

func init() {
	pNsSmtpEncryptTypeStarttls = cNsSmtpEncryptTypeStarttls
	NsSmtpEncryptTypeStarttls = &pNsSmtpEncryptTypeStarttls

	pNsSmtpEncryptTypeNone = cNsSmtpEncryptTypeNone
	NsSmtpEncryptTypeNone = &pNsSmtpEncryptTypeNone

	pNsSmtpEncryptTypeSsl = cNsSmtpEncryptTypeSsl
	NsSmtpEncryptTypeSsl = &pNsSmtpEncryptTypeSsl

}
