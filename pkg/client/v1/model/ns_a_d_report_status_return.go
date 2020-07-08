// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsADReportStatusReturn - Status of the Active Directory domain.
// Export NsADReportStatusReturnFields for advance operations like search filter etc.
var NsADReportStatusReturnFields *NsADReportStatusReturn

func init() {

	NsADReportStatusReturnFields = &NsADReportStatusReturn{}
}

type NsADReportStatusReturn struct {
	// Joined - Joined the Active Directory group.
	Joined *bool `json:"joined,omitempty"`
	// Enabled - Active Directory group is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// LocalServiceStatus - Status of the local service.
	LocalServiceStatus *bool `json:"local_service_status,omitempty"`
	// RemoteServiceStatus - Status of the remote service.
	RemoteServiceStatus *bool `json:"remote_service_status,omitempty"`
	// TrustValid - Trust is valid.
	TrustValid *bool `json:"trust_valid,omitempty"`
}

// sdk internal struct
type nsADReportStatusReturn struct {
	// Joined - Joined the Active Directory group.
	Joined *bool `json:"joined,omitempty"`
	// Enabled - Active Directory group is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// LocalServiceStatus - Status of the local service.
	LocalServiceStatus *bool `json:"local_service_status,omitempty"`
	// RemoteServiceStatus - Status of the remote service.
	RemoteServiceStatus *bool `json:"remote_service_status,omitempty"`
	// TrustValid - Trust is valid.
	TrustValid *bool `json:"trust_valid,omitempty"`
}

// EncodeNsADReportStatusReturn - Transform NsADReportStatusReturn to nsADReportStatusReturn type
func EncodeNsADReportStatusReturn(request interface{}) (*nsADReportStatusReturn, error) {
	reqNsADReportStatusReturn := request.(*NsADReportStatusReturn)
	byte, err := json.Marshal(reqNsADReportStatusReturn)
	objPtr := &nsADReportStatusReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsADReportStatusReturn Transform nsADReportStatusReturn to NsADReportStatusReturn type
func DecodeNsADReportStatusReturn(request interface{}) (*NsADReportStatusReturn, error) {
	reqNsADReportStatusReturn := request.(*nsADReportStatusReturn)
	byte, err := json.Marshal(reqNsADReportStatusReturn)
	obj := &NsADReportStatusReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
