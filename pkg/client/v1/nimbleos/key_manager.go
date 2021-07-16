// Copyright 2021 Hewlett Packard Enterprise Development LP

package nimbleos

// KeyManager - Key Manager stores encryption keys for the array volumes / dedupe domains.
// Export KeyManagerFields for advance operations like search filter etc.
var KeyManagerFields *KeyManager

func init() {
	IDfield := "id"
	Namefield := "name"
	Descriptionfield := "description"
	Hostnamefield := "hostname"
	Usernamefield := "username"
	Passwordfield := "password"
	Statusfield := "status"
	Vendorfield := "vendor"

	KeyManagerFields = &KeyManager{
		ID:          &IDfield,
		Name:        &Namefield,
		Description: &Descriptionfield,
		Hostname:    &Hostnamefield,
		Username:    &Usernamefield,
		Password:    &Passwordfield,
		Status:      &Statusfield,
		Vendor:      &Vendorfield,
	}
}

type KeyManager struct {
	// ID - Identifier for External Key Manager.
	ID *string `json:"id,omitempty"`
	// Name - Name of external key manager.
	Name *string `json:"name,omitempty"`
	// Description - Description of external key manager.
	Description *string `json:"description,omitempty"`
	// Hostname - Hostname or IP Address for the External Key Manager.
	Hostname *string `json:"hostname,omitempty"`
	// Port - Port number for the External Key Manager.
	Port *int64 `json:"port,omitempty"`
	// Protocol - KMIP protocol supported by External Key Manager.
	Protocol *NsKmipProtocol `json:"protocol,omitempty"`
	// Username - External Key Manager username. String up to 255 printable characters.
	Username *string `json:"username,omitempty"`
	// Password - External Key Manager user password. String up to 255 printable characters.
	Password *string `json:"password,omitempty"`
	// Active - Whether the given key manager is active or not.
	Active *bool `json:"active,omitempty"`
	// Status - Connection status of a given external key manager.
	Status *string `json:"status,omitempty"`
	// Vendor - KMIP vendor name.
	Vendor *string `json:"vendor,omitempty"`
}
