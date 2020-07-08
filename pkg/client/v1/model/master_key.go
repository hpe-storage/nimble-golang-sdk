// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// MasterKey - Manage the master key. Data encryption keys for volumes are encrypted by using a master key that must be initialized before encrypted volumes can be created. The master key in turn is protected by a passphrase that is set when the master key is created. The passphrase may have to be entered to enable the master key when it is not available, for example, after an array reboot.
// Export MasterKeyFields for advance operations like search filter etc.
var MasterKeyFields *MasterKey

func init() {

	MasterKeyFields = &MasterKey{
		ID:            "id",
		Name:          "name",
		Passphrase:    "passphrase",
		NewPassphrase: "new_passphrase",
	}
}

type MasterKey struct {
	// ID - Identifier of the master key.
	ID string `json:"id,omitempty"`
	// Name - Name of the master key. The only allowed value is "default".
	Name string `json:"name,omitempty"`
	// Passphrase - Passphrase used to protect the master key, required during creation, enabling/disabling the key and change the passphrase to a new value.
	Passphrase string `json:"passphrase,omitempty"`
	// NewPassphrase - When changing the passphrase, this attribute specifies the new value of the passphrase.
	NewPassphrase string `json:"new_passphrase,omitempty"`
	// Active - Whether the master key is active or not.
	Active *bool `json:"active,omitempty"`
	// PurgeAge - Default minimum age (in hours) of inactive encryption keys to be purged. '0' indicates to purge keys immediately.
	PurgeAge int64 `json:"purge_age,omitempty"`
}

// sdk internal struct
type masterKey struct {
	// ID - Identifier of the master key.
	ID *string `json:"id,omitempty"`
	// Name - Name of the master key. The only allowed value is "default".
	Name *string `json:"name,omitempty"`
	// Passphrase - Passphrase used to protect the master key, required during creation, enabling/disabling the key and change the passphrase to a new value.
	Passphrase *string `json:"passphrase,omitempty"`
	// NewPassphrase - When changing the passphrase, this attribute specifies the new value of the passphrase.
	NewPassphrase *string `json:"new_passphrase,omitempty"`
	// Active - Whether the master key is active or not.
	Active *bool `json:"active,omitempty"`
	// PurgeAge - Default minimum age (in hours) of inactive encryption keys to be purged. '0' indicates to purge keys immediately.
	PurgeAge *int64 `json:"purge_age,omitempty"`
}

// EncodeMasterKey - Transform MasterKey to masterKey type
func EncodeMasterKey(request interface{}) (*masterKey, error) {
	reqMasterKey := request.(*MasterKey)
	byte, err := json.Marshal(reqMasterKey)
	objPtr := &masterKey{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeMasterKey Transform masterKey to MasterKey type
func DecodeMasterKey(request interface{}) (*MasterKey, error) {
	reqMasterKey := request.(*masterKey)
	byte, err := json.Marshal(reqMasterKey)
	obj := &MasterKey{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
