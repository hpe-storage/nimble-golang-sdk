// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// WitnessFields provides field names to use in filter parameters, for example.
var WitnessFields *WitnessFieldHandles

func init() {
	WitnessFields = &WitnessFieldHandles{
		ID:                     "id",
		Username:               "username",
		Password:               "password",
		Host:                   "host",
		Port:                   "port",
		SecureMode:             "secure_mode",
		AutoSwitchoverMessages: "auto_switchover_messages",
	}
}

// Witness - Manage witness host configuration.
type Witness struct {
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

// WitnessFieldHandles provides a string representation for each AccessControlRecord field.
type WitnessFieldHandles struct {
	ID                     string
	Username               string
	Password               string
	Host                   string
	Port                   string
	SecureMode             string
	AutoSwitchoverMessages string
}
