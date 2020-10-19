// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

// NsISCSISession - ISCSI initiator session information.
// Export NsISCSISessionFields for advance operations like search filter etc.
var NsISCSISessionFields *NsISCSISession

func init() {
	IDfield := "id"
	SessionIdfield := "session_id"
	InitiatorNamefield := "initiator_name"
	InitiatorIpAddrfield := "initiator_ip_addr"
	TargetIpAddrfield := "target_ip_addr"

	NsISCSISessionFields = &NsISCSISession{
		ID:              &IDfield,
		SessionId:       &SessionIdfield,
		InitiatorName:   &InitiatorNamefield,
		InitiatorIpAddr: &InitiatorIpAddrfield,
		TargetIpAddr:    &TargetIpAddrfield,
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
