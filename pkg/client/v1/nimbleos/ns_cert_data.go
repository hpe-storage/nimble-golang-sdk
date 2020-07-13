// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsCertData - Detailed certificate information for an installed certificate.
// Export NsCertDataFields for advance operations like search filter etc.
var NsCertDataFields *NsCertData

func init() {

	NsCertDataFields = &NsCertData{
		Description: "description",
		Subject:     "subject",
		Dnslist:     "dnslist",
		Iplist:      "iplist",
		PemText:     "pem_text",
	}
}

type NsCertData struct {
	// Description - Complete certificate description.
	Description string `json:"description,omitempty"`
	// Subject - Certificate subject name.
	Subject string `json:"subject,omitempty"`
	// Dnslist - Comma-separated list of DNS names from the Subject Alternate Name.
	Dnslist string `json:"dnslist,omitempty"`
	// Iplist - Comma-separated list of IP addresses from the Subject Alternate Name.
	Iplist string `json:"iplist,omitempty"`
	// Trusted - Certificate is an imported, trusted certificate.
	Trusted *bool `json:"trusted,omitempty"`
	// PemText - PEM text of the actual certificate.
	PemText string `json:"pem_text,omitempty"`
}

// sdk internal struct
type nsCertData struct {
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

// EncodeNsCertData - Transform NsCertData to nsCertData type
func EncodeNsCertData(request interface{}) (*nsCertData, error) {
	reqNsCertData := request.(*NsCertData)
	byte, err := json.Marshal(reqNsCertData)
	objPtr := &nsCertData{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsCertData Transform nsCertData to NsCertData type
func DecodeNsCertData(request interface{}) (*NsCertData, error) {
	reqNsCertData := request.(*nsCertData)
	byte, err := json.Marshal(reqNsCertData)
	obj := &NsCertData{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
