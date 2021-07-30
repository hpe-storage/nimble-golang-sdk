// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// TrustedOauthIssuer - Oauth Credential Issuers that this device will trust.
// Export TrustedOauthIssuerFields for advance operations like search filter etc.
var TrustedOauthIssuerFields *TrustedOauthIssuerStringFields

func init() {
	IDfield := "id"
	Namefield := "name"
	JwksUrlfield := "jwks_url"
	KeySetfield := "key_set"

	TrustedOauthIssuerFields = &TrustedOauthIssuerStringFields{
		ID:      &IDfield,
		Name:    &Namefield,
		JwksUrl: &JwksUrlfield,
		KeySet:  &KeySetfield,
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

// Struct for TrustedOauthIssuerFields
type TrustedOauthIssuerStringFields struct {
	ID      *string
	Name    *string
	JwksUrl *string
	KeySet  *string
}
