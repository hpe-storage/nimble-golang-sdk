// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// MasterKey - Manage the master key. Data encryption keys for volumes are encrypted by using a master key that must be initialized before encrypted volumes can be created. The master key in turn is protected by a passphrase that is set when the master key is created. The passphrase may have to be entered to enable the master key when it is not available, for example, after an array reboot.

// Export MasterKeyFields provides field names to use in filter parameters, for example.
var MasterKeyFields *MasterKeyStringFields

func init() {
	fieldID := "id"
	fieldName := "name"
	fieldPassphrase := "passphrase"
	fieldHalfkey := "halfkey"
	fieldNewPassphrase := "new_passphrase"
	fieldActive := "active"
	fieldPurgeAge := "purge_age"

	MasterKeyFields = &MasterKeyStringFields{
		ID:            &fieldID,
		Name:          &fieldName,
		Passphrase:    &fieldPassphrase,
		Halfkey:       &fieldHalfkey,
		NewPassphrase: &fieldNewPassphrase,
		Active:        &fieldActive,
		PurgeAge:      &fieldPurgeAge,
	}
}

type MasterKey struct {
	// ID - Identifier of the master key.
	ID *string `json:"id,omitempty"`
	// Name - Name of the master key. The only allowed value is "default".
	Name *string `json:"name,omitempty"`
	// Passphrase - Passphrase used to protect the master key, required during creation, enabling/disabling the key and change the passphrase to a new value.
	Passphrase *string `json:"passphrase,omitempty"`
	// Halfkey - When changing the passphrase, this authenticates the change operation, for support use only.
	Halfkey *string `json:"halfkey,omitempty"`
	// NewPassphrase - When changing the passphrase, this attribute specifies the new value of the passphrase.
	NewPassphrase *string `json:"new_passphrase,omitempty"`
	// Active - Whether the master key is active or not.
	Active *bool `json:"active,omitempty"`
	// PurgeAge - Default minimum age (in hours) of inactive encryption keys to be purged. '0' indicates to purge keys immediately.
	PurgeAge *int64 `json:"purge_age,omitempty"`
}

// Struct for MasterKeyFields
type MasterKeyStringFields struct {
	ID            *string
	Name          *string
	Passphrase    *string
	Halfkey       *string
	NewPassphrase *string
	Active        *string
	PurgeAge      *string
}
