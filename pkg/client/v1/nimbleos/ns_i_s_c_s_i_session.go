// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsISCSISession - ISCSI initiator session information.
// Export NsISCSISessionFields for advance operations like search filter etc.
var NsISCSISessionFields *NsISCSISession

func init() {

	NsISCSISessionFields = &NsISCSISession{
		ID:              "id",
		SessionId:       "session_id",
		InitiatorName:   "initiator_name",
		InitiatorIpAddr: "initiator_ip_addr",
		TargetIpAddr:    "target_ip_addr",
	}
}

type NsISCSISession struct {
	// ID - Unique identifier of the iSCSI session.
	ID string `json:"id,omitempty"`
	// SessionId - Unique identifier of the iSCSI session.
	SessionId string `json:"session_id,omitempty"`
	// InitiatorName - Name of the iSCSI initiator (IQN).
	InitiatorName string `json:"initiator_name,omitempty"`
	// NumConnections - Number of connections in iSCSI session.
	NumConnections int64 `json:"num_connections,omitempty"`
	// PrKey - Registered persistent reservation key for the I-T nexus on this LU (if there is no registration, its value will be zero).
	PrKey int64 `json:"pr_key,omitempty"`
	// InitiatorIpAddr - IP address of the iSCSI initiator.
	InitiatorIpAddr string `json:"initiator_ip_addr,omitempty"`
	// TargetIpAddr - Target IP address of the iSCSI initiator.
	TargetIpAddr string `json:"target_ip_addr,omitempty"`
	// HeaderDigestEnabled - Indicate whether header digest is enabled.
	HeaderDigestEnabled *bool `json:"header_digest_enabled,omitempty"`
	// DataDigestEnabled - Indicate whether data digest is enabled.
	DataDigestEnabled *bool `json:"data_digest_enabled,omitempty"`
}

// sdk internal struct
type nsISCSISession struct {
	ID                  *string `json:"id,omitempty"`
	SessionId           *string `json:"session_id,omitempty"`
	InitiatorName       *string `json:"initiator_name,omitempty"`
	NumConnections      *int64  `json:"num_connections,omitempty"`
	PrKey               *int64  `json:"pr_key,omitempty"`
	InitiatorIpAddr     *string `json:"initiator_ip_addr,omitempty"`
	TargetIpAddr        *string `json:"target_ip_addr,omitempty"`
	HeaderDigestEnabled *bool   `json:"header_digest_enabled,omitempty"`
	DataDigestEnabled   *bool   `json:"data_digest_enabled,omitempty"`
}

// EncodeNsISCSISession - Transform NsISCSISession to nsISCSISession type
func EncodeNsISCSISession(request interface{}) (*nsISCSISession, error) {
	reqNsISCSISession := request.(*NsISCSISession)
	byte, err := json.Marshal(reqNsISCSISession)
	if err != nil {
		return nil, err
	}
	respNsISCSISessionPtr := &nsISCSISession{}
	err = json.Unmarshal(byte, respNsISCSISessionPtr)
	return respNsISCSISessionPtr, err
}

// DecodeNsISCSISession Transform nsISCSISession to NsISCSISession type
func DecodeNsISCSISession(request interface{}) (*NsISCSISession, error) {
	reqNsISCSISession := request.(*nsISCSISession)
	byte, err := json.Marshal(reqNsISCSISession)
	if err != nil {
		return nil, err
	}
	respNsISCSISession := &NsISCSISession{}
	err = json.Unmarshal(byte, respNsISCSISession)
	return respNsISCSISession, err
}
