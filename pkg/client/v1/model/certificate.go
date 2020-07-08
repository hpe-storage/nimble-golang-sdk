// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// Certificate - Manage certificates used by SSL/TLS.
// Export CertificateFields for advance operations like search filter etc.
var CertificateFields *Certificate

func init() {

	CertificateFields = &Certificate{
		ID:       "id",
		Name:     "name",
		Subject:  "subject",
		Dnslist:  "dnslist",
		Iplist:   "iplist",
		Input:    "input",
		Password: "password",
	}
}

type Certificate struct {
	// ID - Dummy identifier (copy of name).
	ID string `json:"id,omitempty"`
	// Name - Name of certficiate.
	Name string `json:"name,omitempty"`
	// Subject - Subject for the certificate.
	Subject string `json:"subject,omitempty"`
	// Dnslist - DNS list for subject alternate name.
	Dnslist string `json:"dnslist,omitempty"`
	// Iplist - List of IP addresses for subject alternate name.
	Iplist string `json:"iplist,omitempty"`
	// CertList - Certificate information for certificates in file.
	CertList []*NsCertData `json:"cert_list,omitempty"`
	// Csr - Certificate Signing Request information.
	Csr *NsCertData `json:"csr,omitempty"`
	// ValidDays - Number of days certificate is valid.
	ValidDays int64 `json:"valid_days,omitempty"`
	// Input - Certificate input for import operation. The format of the input is a base-64 encoded string for either PKCS-12 input or PEM encoded certificate(s).  PEM certificate(s) may also be specified as-is, with newlines replaced by "\n".
	Input string `json:"input,omitempty"`
	// Password - Password used to protect PKCS-12 certificate bundle.
	Password string `json:"password,omitempty"`
	// Https - For use command, Use specified type for HTTPS access.
	Https *bool `json:"https,omitempty"`
	// Apis - For use command, use specified type for REST API access.
	Apis *bool `json:"apis,omitempty"`
	// Force - Force regeneration or import when certificate already exists.
	Force *bool `json:"force,omitempty"`
	// Check - Check existing group certificate parameters against inputs.
	Check *bool `json:"check,omitempty"`
	// Pkcs12 - Input is a PKCS-12 bundle if true.
	Pkcs12 *bool `json:"pkcs12,omitempty"`
	// Trusted - Certificate is an imported trusted certificate.
	Trusted *bool `json:"trusted,omitempty"`
	// ReloadHttps - HTTPS certificate changed, so needs to be reloaded.
	ReloadHttps *bool `json:"reload_https,omitempty"`
	// ReloadApis - API certificate changed, so needs to be reloaded.
	ReloadApis *bool `json:"reload_apis,omitempty"`
}

// sdk internal struct
type certificate struct {
	// ID - Dummy identifier (copy of name).
	ID *string `json:"id,omitempty"`
	// Name - Name of certficiate.
	Name *string `json:"name,omitempty"`
	// Subject - Subject for the certificate.
	Subject *string `json:"subject,omitempty"`
	// Dnslist - DNS list for subject alternate name.
	Dnslist *string `json:"dnslist,omitempty"`
	// Iplist - List of IP addresses for subject alternate name.
	Iplist *string `json:"iplist,omitempty"`
	// CertList - Certificate information for certificates in file.
	CertList []*NsCertData `json:"cert_list,omitempty"`
	// Csr - Certificate Signing Request information.
	Csr *NsCertData `json:"csr,omitempty"`
	// ValidDays - Number of days certificate is valid.
	ValidDays *int64 `json:"valid_days,omitempty"`
	// Input - Certificate input for import operation. The format of the input is a base-64 encoded string for either PKCS-12 input or PEM encoded certificate(s).  PEM certificate(s) may also be specified as-is, with newlines replaced by "\n".
	Input *string `json:"input,omitempty"`
	// Password - Password used to protect PKCS-12 certificate bundle.
	Password *string `json:"password,omitempty"`
	// Https - For use command, Use specified type for HTTPS access.
	Https *bool `json:"https,omitempty"`
	// Apis - For use command, use specified type for REST API access.
	Apis *bool `json:"apis,omitempty"`
	// Force - Force regeneration or import when certificate already exists.
	Force *bool `json:"force,omitempty"`
	// Check - Check existing group certificate parameters against inputs.
	Check *bool `json:"check,omitempty"`
	// Pkcs12 - Input is a PKCS-12 bundle if true.
	Pkcs12 *bool `json:"pkcs12,omitempty"`
	// Trusted - Certificate is an imported trusted certificate.
	Trusted *bool `json:"trusted,omitempty"`
	// ReloadHttps - HTTPS certificate changed, so needs to be reloaded.
	ReloadHttps *bool `json:"reload_https,omitempty"`
	// ReloadApis - API certificate changed, so needs to be reloaded.
	ReloadApis *bool `json:"reload_apis,omitempty"`
}

// EncodeCertificate - Transform Certificate to certificate type
func EncodeCertificate(request interface{}) (*certificate, error) {
	reqCertificate := request.(*Certificate)
	byte, err := json.Marshal(reqCertificate)
	objPtr := &certificate{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeCertificate Transform certificate to Certificate type
func DecodeCertificate(request interface{}) (*Certificate, error) {
	reqCertificate := request.(*certificate)
	byte, err := json.Marshal(reqCertificate)
	obj := &Certificate{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
