// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsDomainType Enum.

type NsDomainType string

const (
	cNsDomainTypeAd    NsDomainType = "ad"
	cNsDomainTypeLdap  NsDomainType = "ldap"
	cNsDomainTypeLocal NsDomainType = "local"
)

var pNsDomainTypeAd NsDomainType
var pNsDomainTypeLdap NsDomainType
var pNsDomainTypeLocal NsDomainType

// NsDomainTypeAd enum exports
var NsDomainTypeAd *NsDomainType

// NsDomainTypeLdap enum exports
var NsDomainTypeLdap *NsDomainType

// NsDomainTypeLocal enum exports
var NsDomainTypeLocal *NsDomainType

func init() {
	pNsDomainTypeAd = cNsDomainTypeAd
	NsDomainTypeAd = &pNsDomainTypeAd

	pNsDomainTypeLdap = cNsDomainTypeLdap
	NsDomainTypeLdap = &pNsDomainTypeLdap

	pNsDomainTypeLocal = cNsDomainTypeLocal
	NsDomainTypeLocal = &pNsDomainTypeLocal

}
