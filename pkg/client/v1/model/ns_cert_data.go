/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsCertData - Detailed certificate information for an installed certificate.
// Export NsCertDataFields for advance operations like search filter etc.
var NsCertDataFields *NsCertData

func init(){
	Descriptionfield:= "description"
	Subjectfield:= "subject"
	Dnslistfield:= "dnslist"
	Iplistfield:= "iplist"
	PemTextfield:= "pem_text"
		
	NsCertDataFields= &NsCertData{
		Description: &Descriptionfield,
		Subject: &Subjectfield,
		Dnslist: &Dnslistfield,
		Iplist: &Iplistfield,
		PemText: &PemTextfield,
		
	}
}

type NsCertData struct {
   
    // Complete certificate description.
    
 	Description *string `json:"description,omitempty"`
   
    // Certificate subject name.
    
 	Subject *string `json:"subject,omitempty"`
   
    // Comma-separated list of DNS names from the Subject Alternate Name.
    
 	Dnslist *string `json:"dnslist,omitempty"`
   
    // Comma-separated list of IP addresses from the Subject Alternate Name.
    
 	Iplist *string `json:"iplist,omitempty"`
   
    // Certificate is an imported, trusted certificate.
    
 	Trusted *bool `json:"trusted,omitempty"`
   
    // PEM text of the actual certificate.
    
 	PemText *string `json:"pem_text,omitempty"`
}
