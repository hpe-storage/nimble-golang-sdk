// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// TrustedOauthIssuerFields provides field names to use in filter parameters, for example.
var TrustedOauthIssuerFields *TrustedOauthIssuerFieldHandles

func init() {
	TrustedOauthIssuerFields = &TrustedOauthIssuerFieldHandles{
		ID:      "id",
		Name:    "name",
		JwksUrl: "jwks_url",
		KeySet:  "key_set",
	}
}

// TrustedOauthIssuer - Oauth Credential Issuers that this device will trust.
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

// TrustedOauthIssuerFieldHandles provides a string representation for each TrustedOauthIssuer field.
type TrustedOauthIssuerFieldHandles struct {
	ID      string
	Name    string
	JwksUrl string
	KeySet  string
}
