// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsLdapSchemaType Enum.

type NsLdapSchemaType string

const (
	cNsLdapSchemaTypeOpenldap NsLdapSchemaType = "OpenLDAP"
	cNsLdapSchemaTypeAd       NsLdapSchemaType = "AD"
)

var pNsLdapSchemaTypeOpenldap NsLdapSchemaType
var pNsLdapSchemaTypeAd NsLdapSchemaType

// NsLdapSchemaTypeOpenldap enum exports
var NsLdapSchemaTypeOpenldap *NsLdapSchemaType

// NsLdapSchemaTypeAd enum exports
var NsLdapSchemaTypeAd *NsLdapSchemaType

func init() {
	pNsLdapSchemaTypeOpenldap = cNsLdapSchemaTypeOpenldap
	NsLdapSchemaTypeOpenldap = &pNsLdapSchemaTypeOpenldap

	pNsLdapSchemaTypeAd = cNsLdapSchemaTypeAd
	NsLdapSchemaTypeAd = &pNsLdapSchemaTypeAd

}
