// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsCertData - Detailed certificate information for an installed certificate.
// Export NsCertDataFields for advance operations like search filter etc.
var NsCertDataFields *NsCertDataStringFields

func init() {
	Descriptionfield := "description"
	Subjectfield := "subject"
	Dnslistfield := "dnslist"
	Iplistfield := "iplist"
	Trustedfield := "trusted"
	PemTextfield := "pem_text"

	NsCertDataFields = &NsCertDataStringFields{
		Description: &Descriptionfield,
		Subject:     &Subjectfield,
		Dnslist:     &Dnslistfield,
		Iplist:      &Iplistfield,
		Trusted:     &Trustedfield,
		PemText:     &PemTextfield,
	}
}

type NsCertData struct {
	// Description - Complete certificate description.
	Description *string `json:"description,omitempty"`
	// Subject - Certificate subject name.
	Subject *string `json:"subject,omitempty"`
	// Dnslist - Comma-separated list of DNS names from the Subject Alternate Name.
	Dnslist *string `json:"dnslist,omitempty"`
	// Iplist - Comma-separated list of IP addresses from the Subject Alternate Name.
	Iplist *string `json:"iplist,omitempty"`
	// Trusted - Certificate is an imported, trusted certificate.
	Trusted *bool `json:"trusted,omitempty"`
	// PemText - PEM text of the actual certificate.
	PemText *string `json:"pem_text,omitempty"`
}

// Struct for NsCertDataFields
type NsCertDataStringFields struct {
	Description *string
	Subject     *string
	Dnslist     *string
	Iplist      *string
	Trusted     *string
	PemText     *string
}
