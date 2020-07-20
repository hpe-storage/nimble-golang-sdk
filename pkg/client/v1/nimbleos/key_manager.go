// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// KeyManager - Key Manager stores encryption keys for the array volumes / dedupe domains.
// Export KeyManagerFields for advance operations like search filter etc.
var KeyManagerFields *KeyManager

func init() {

	KeyManagerFields = &KeyManager{
		ID:          "id",
		Name:        "name",
		Description: "description",
		Hostname:    "hostname",
		Username:    "username",
		Password:    "password",
		Status:      "status",
		Vendor:      "vendor",
	}
}

type KeyManager struct {
	// ID - Identifier for External Key Manager.
	ID string `json:"id,omitempty"`
	// Name - Name of external key manager.
	Name string `json:"name,omitempty"`
	// Description - Description of external key manager.
	Description string `json:"description,omitempty"`
	// Hostname - Hostname or IP Address for the External Key Manager.
	Hostname string `json:"hostname,omitempty"`
	// Port - Port number for the External Key Manager.
	Port int64 `json:"port,omitempty"`
	// Protocol - KMIP protocol supported by External Key Manager.
	Protocol *NsKmipProtocol `json:"protocol,omitempty"`
	// Username - External Key Manager username. String up to 255 printable characters.
	Username string `json:"username,omitempty"`
	// Password - External Key Manager user password. String up to 255 printable characters.
	Password string `json:"password,omitempty"`
	// Active - Whether the given key manager is active or not.
	Active *bool `json:"active,omitempty"`
	// Status - Connection status of a given external key manager.
	Status string `json:"status,omitempty"`
	// Vendor - KMIP vendor name.
	Vendor string `json:"vendor,omitempty"`
}

// sdk internal struct
type keyManager struct {
	ID          *string         `json:"id,omitempty"`
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Hostname    *string         `json:"hostname,omitempty"`
	Port        *int64          `json:"port,omitempty"`
	Protocol    *NsKmipProtocol `json:"protocol,omitempty"`
	Username    *string         `json:"username,omitempty"`
	Password    *string         `json:"password,omitempty"`
	Active      *bool           `json:"active,omitempty"`
	Status      *string         `json:"status,omitempty"`
	Vendor      *string         `json:"vendor,omitempty"`
}

// EncodeKeyManager - Transform KeyManager to keyManager type
func EncodeKeyManager(request interface{}) (*keyManager, error) {
	reqKeyManager := request.(*KeyManager)
	byte, err := json.Marshal(reqKeyManager)
	if err != nil {
		return nil, err
	}
	respKeyManagerPtr := &keyManager{}
	err = json.Unmarshal(byte, respKeyManagerPtr)
	return respKeyManagerPtr, err
}

// DecodeKeyManager Transform keyManager to KeyManager type
func DecodeKeyManager(request interface{}) (*KeyManager, error) {
	reqKeyManager := request.(*keyManager)
	byte, err := json.Marshal(reqKeyManager)
	if err != nil {
		return nil, err
	}
	respKeyManager := &KeyManager{}
	err = json.Unmarshal(byte, respKeyManager)
	return respKeyManager, err
}
