// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsEncryptionSettings - Group encryption settings.
// Export NsEncryptionSettingsFields for advance operations like search filter etc.
var NsEncryptionSettingsFields *NsEncryptionSettings

func init() {

	NsEncryptionSettingsFields = &NsEncryptionSettings{}
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

// sdk internal struct
type nsEncryptionSettings struct {
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

// EncodeNsEncryptionSettings - Transform NsEncryptionSettings to nsEncryptionSettings type
func EncodeNsEncryptionSettings(request interface{}) (*nsEncryptionSettings, error) {
	reqNsEncryptionSettings := request.(*NsEncryptionSettings)
	byte, err := json.Marshal(reqNsEncryptionSettings)
	objPtr := &nsEncryptionSettings{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsEncryptionSettings Transform nsEncryptionSettings to NsEncryptionSettings type
func DecodeNsEncryptionSettings(request interface{}) (*NsEncryptionSettings, error) {
	reqNsEncryptionSettings := request.(*nsEncryptionSettings)
	byte, err := json.Marshal(reqNsEncryptionSettings)
	obj := &NsEncryptionSettings{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
