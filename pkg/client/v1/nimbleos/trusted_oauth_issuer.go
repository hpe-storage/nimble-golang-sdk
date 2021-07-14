// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos


// TrustedOauthIssuer - Oauth Credential Issuers that this device will trust.
// Export TrustedOauthIssuerFields for advance operations like search filter etc.
var TrustedOauthIssuerFields *TrustedOauthIssuer

func init(){
 IDfield:= "id"
 Namefield:= "name"
 JwksUrlfield:= "jwks_url"

 TrustedOauthIssuerFields= &TrustedOauthIssuer{
  ID:      &IDfield,
  Name:    &Namefield,
  JwksUrl: &JwksUrlfield,
 }
}


type TrustedOauthIssuer struct {
 // ID - Identifier for the trusted oauth issuer record.
  ID *string `json:"id,omitempty"`
 // Name - Issuer ID string.
  Name *string `json:"name,omitempty"`
 // JwksUrl - The URL from which the device will download the public key set for signature verification.
  JwksUrl *string `json:"jwks_url,omitempty"`
 // KeySet - List of public keys for verifying signatures.
    KeySet []*NsJWKey `json:"key_set,omitempty"`
}
