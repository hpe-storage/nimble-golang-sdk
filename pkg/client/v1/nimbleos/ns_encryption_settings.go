// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsEncryptionSettings - Group encryption settings.
// Export NsEncryptionSettingsFields for advance operations like search filter etc.
var NsEncryptionSettingsFields *NsEncryptionSettingsStringFields

func init() {
	MasterKeySetfield := "master_key_set"
	Modefield := "mode"
	Scopefield := "scope"
	Cipherfield := "cipher"
	EncryptionActivefield := "encryption_active"
	EncryptionKeyManagerfield := "encryption_key_manager"

	NsEncryptionSettingsFields = &NsEncryptionSettingsStringFields{
		MasterKeySet:         &MasterKeySetfield,
		Mode:                 &Modefield,
		Scope:                &Scopefield,
		Cipher:               &Cipherfield,
		EncryptionActive:     &EncryptionActivefield,
		EncryptionKeyManager: &EncryptionKeyManagerfield,
	}
}

type NsEncryptionSettings struct {
	// MasterKeySet - Is the master key set (output only).
	MasterKeySet *bool `json:"master_key_set,omitempty"`
	// Mode - Mode of encryption.
	Mode *NsEncryptionMode `json:"mode,omitempty"`
	// Scope - Encryption scope.
	Scope *NsEncryptionScope `json:"scope,omitempty"`
	// Cipher - Type of encryption cipher used.
	Cipher *NsEncryptionCipher `json:"cipher,omitempty"`
	// EncryptionActive - Is encryption active (output only).
	EncryptionActive *bool `json:"encryption_active,omitempty"`
	// EncryptionKeyManager - Is the master key on local or external key manager (output only).
	EncryptionKeyManager *NsEncryptionKeyManager `json:"encryption_key_manager,omitempty"`
}

// Struct for NsEncryptionSettingsFields
type NsEncryptionSettingsStringFields struct {
	MasterKeySet         *string
	Mode                 *string
	Scope                *string
	Cipher               *string
	EncryptionActive     *string
	EncryptionKeyManager *string
}
