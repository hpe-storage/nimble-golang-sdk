// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export TrustedOauthIssuerFields provides field names to use in filter parameters, for example.
var TrustedOauthIssuerFields *TrustedOauthIssuerFieldHandles

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldJwksUrl := "jwks_url"
	fieldKeySet := "key_set"

	TrustedOauthIssuerFields = &TrustedOauthIssuerFieldHandles{
		ID:      &fieldID,
		Name:    &fieldName,
		JwksUrl: &fieldJwksUrl,
		KeySet:  &fieldKeySet,
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

// TrustedOauthIssuerFieldHandles provides a string representation for each AccessControlRecord field.
type TrustedOauthIssuerFieldHandles struct {
	ID      *string
	Name    *string
	JwksUrl *string
	KeySet  *string
}
