// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// Witness - Manage witness host configuration.
// Export WitnessFields for advance operations like search filter etc.
var WitnessFields *Witness

func init() {

	WitnessFields = &Witness{
		ID:       "id",
		Username: "username",
		Password: "password",
		Host:     "host",
	}
}

type Witness struct {
	// ID - Identifier of the witness configuration.
	ID string `json:"id,omitempty"`
	// Username - Username of witness. This has to be a non-root that can login to the witness host.
	Username string `json:"username,omitempty"`
	// Password - Password of witness.
	Password string `json:"password,omitempty"`
	// Host - Hostname or ip addresses of witness.
	Host string `json:"host,omitempty"`
	// Port - Port of witness.
	Port int64 `json:"port,omitempty"`
	// SecureMode - To verify the witness host against CA cert and to apply possible security modes.
	SecureMode *bool `json:"secure_mode,omitempty"`
	// AutoSwitchoverMessages - List of validation messages for automatic switchover of Group Management. This will be empty when there are no conflicts found.
	AutoSwitchoverMessages []*NsErrorWithArguments `json:"auto_switchover_messages,omitempty"`
}

// sdk internal struct
type witness struct {
	// ID - Identifier of the witness configuration.
	ID *string `json:"id,omitempty"`
	// Username - Username of witness. This has to be a non-root that can login to the witness host.
	Username *string `json:"username,omitempty"`
	// Password - Password of witness.
	Password *string `json:"password,omitempty"`
	// Host - Hostname or ip addresses of witness.
	Host *string `json:"host,omitempty"`
	// Port - Port of witness.
	Port *int64 `json:"port,omitempty"`
	// SecureMode - To verify the witness host against CA cert and to apply possible security modes.
	SecureMode *bool `json:"secure_mode,omitempty"`
	// AutoSwitchoverMessages - List of validation messages for automatic switchover of Group Management. This will be empty when there are no conflicts found.
	AutoSwitchoverMessages []*NsErrorWithArguments `json:"auto_switchover_messages,omitempty"`
}

// EncodeWitness - Transform Witness to witness type
func EncodeWitness(request interface{}) (*witness, error) {
	reqWitness := request.(*Witness)
	byte, err := json.Marshal(reqWitness)
	objPtr := &witness{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeWitness Transform witness to Witness type
func DecodeWitness(request interface{}) (*Witness, error) {
	reqWitness := request.(*witness)
	byte, err := json.Marshal(reqWitness)
	obj := &Witness{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
