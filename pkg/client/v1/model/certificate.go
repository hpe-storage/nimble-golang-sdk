/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// Certificate - Manage certificates used by SSL/TLS.
// Export CertificateFields for advance operations like search filter etc.
var CertificateFields *Certificate

func init(){
	IDfield:= "id"
	Namefield:= "name"
	Subjectfield:= "subject"
	Dnslistfield:= "dnslist"
	Iplistfield:= "iplist"
	Inputfield:= "input"
	Passwordfield:= "password"
		
	CertificateFields= &Certificate{
		ID: &IDfield,
		Name: &Namefield,
		Subject: &Subjectfield,
		Dnslist: &Dnslistfield,
		Iplist: &Iplistfield,
		Input: &Inputfield,
		Password: &Passwordfield,
		
	}
}

type Certificate struct {
   
    // Dummy identifier (copy of name).
    
 	ID *string `json:"id,omitempty"`
   
    // Name of certficiate.
    
 	Name *string `json:"name,omitempty"`
   
    // Subject for the certificate.
    
 	Subject *string `json:"subject,omitempty"`
   
    // DNS list for subject alternate name.
    
 	Dnslist *string `json:"dnslist,omitempty"`
   
    // List of IP addresses for subject alternate name.
    
 	Iplist *string `json:"iplist,omitempty"`
   
    // Certificate information for certificates in file.
    
   	CertList []*NsCertData `json:"cert_list,omitempty"`
   
    // Certificate Signing Request information.
    
	Csr *NsCertData `json:"csr,omitempty"`
   
    // Number of days certificate is valid.
    
   	ValIDDays *int64 `json:"valid_days,omitempty"`
   
    // Certificate input for import operation. The format of the input is a base-64 encoded string for either PKCS-12 input or PEM encoded certificate(s).  PEM certificate(s) may also be specified as-is, with newlines replaced by "\n".
    
 	Input *string `json:"input,omitempty"`
   
    // Password used to protect PKCS-12 certificate bundle.
    
 	Password *string `json:"password,omitempty"`
   
    // For use command, Use specified type for HTTPS access.
    
 	Https *bool `json:"https,omitempty"`
   
    // For use command, use specified type for REST API access.
    
 	Apis *bool `json:"apis,omitempty"`
   
    // Force regeneration or import when certificate already exists.
    
 	Force *bool `json:"force,omitempty"`
   
    // Check existing group certificate parameters against inputs.
    
 	Check *bool `json:"check,omitempty"`
   
    // Input is a PKCS-12 bundle if true.
    
 	Pkcs12 *bool `json:"pkcs12,omitempty"`
   
    // Certificate is an imported trusted certificate.
    
 	Trusted *bool `json:"trusted,omitempty"`
   
    // HTTPS certificate changed, so needs to be reloaded.
    
 	ReloadHttps *bool `json:"reload_https,omitempty"`
   
    // API certificate changed, so needs to be reloaded.
    
 	ReloadApis *bool `json:"reload_apis,omitempty"`
}
