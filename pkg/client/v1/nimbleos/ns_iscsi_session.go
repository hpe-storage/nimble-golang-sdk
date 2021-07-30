// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsISCSISession - ISCSI initiator session information.

// Export NsISCSISessionFields provides field names to use in filter parameters, for example.
var NsISCSISessionFields *NsISCSISessionStringFields

func init() {
	fieldID := "id"
	fieldSessionId := "session_id"
	fieldInitiatorName := "initiator_name"
	fieldNumConnections := "num_connections"
	fieldPrKey := "pr_key"
	fieldInitiatorIpAddr := "initiator_ip_addr"
	fieldTargetIpAddr := "target_ip_addr"
	fieldHeaderDigestEnabled := "header_digest_enabled"
	fieldDataDigestEnabled := "data_digest_enabled"

	NsISCSISessionFields = &NsISCSISessionStringFields{
		ID:                  &fieldID,
		SessionId:           &fieldSessionId,
		InitiatorName:       &fieldInitiatorName,
		NumConnections:      &fieldNumConnections,
		PrKey:               &fieldPrKey,
		InitiatorIpAddr:     &fieldInitiatorIpAddr,
		TargetIpAddr:        &fieldTargetIpAddr,
		HeaderDigestEnabled: &fieldHeaderDigestEnabled,
		DataDigestEnabled:   &fieldDataDigestEnabled,
	}
}

type NsISCSISession struct {
	// ID - Unique identifier of the iSCSI session.
	ID *string `json:"id,omitempty"`
	// SessionId - Unique identifier of the iSCSI session.
	SessionId *string `json:"session_id,omitempty"`
	// InitiatorName - Name of the iSCSI initiator (IQN).
	InitiatorName *string `json:"initiator_name,omitempty"`
	// NumConnections - Number of connections in iSCSI session.
	NumConnections *int64 `json:"num_connections,omitempty"`
	// PrKey - Registered persistent reservation key for the I-T nexus on this LU (if there is no registration, its value will be zero).
	PrKey *int64 `json:"pr_key,omitempty"`
	// InitiatorIpAddr - IP address of the iSCSI initiator.
	InitiatorIpAddr *string `json:"initiator_ip_addr,omitempty"`
	// TargetIpAddr - Target IP address of the iSCSI initiator.
	TargetIpAddr *string `json:"target_ip_addr,omitempty"`
	// HeaderDigestEnabled - Indicate whether header digest is enabled.
	HeaderDigestEnabled *bool `json:"header_digest_enabled,omitempty"`
	// DataDigestEnabled - Indicate whether data digest is enabled.
	DataDigestEnabled *bool `json:"data_digest_enabled,omitempty"`
}

// Struct for NsISCSISessionFields
type NsISCSISessionStringFields struct {
	ID                  *string
	SessionId           *string
	InitiatorName       *string
	NumConnections      *string
	PrKey               *string
	InitiatorIpAddr     *string
	TargetIpAddr        *string
	HeaderDigestEnabled *string
	DataDigestEnabled   *string
}
