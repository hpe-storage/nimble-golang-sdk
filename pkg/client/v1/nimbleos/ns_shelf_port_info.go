// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// Export NsShelfPortInfoFields provides field names to use in filter parameters, for example.
var NsShelfPortInfoFields *NsShelfPortInfoFieldHandles

func init() {
	fieldPortIdx := "port_idx"
	fieldPortName := "port_name"
	fieldPortType := "port_type"
	fieldPortStatus := "port_status"
	fieldPortErrors := "port_errors"
	fieldRemoteSasAddr := "remote_sas_addr"
	fieldRemoteSasPhyId := "remote_sas_phy_id"
	fieldRemoteSasDomain := "remote_sas_domain"
	fieldRemoteLocId := "remote_loc_id"
	fieldRemotePortId := "remote_port_id"

	NsShelfPortInfoFields = &NsShelfPortInfoFieldHandles{
		PortIdx:         &fieldPortIdx,
		PortName:        &fieldPortName,
		PortType:        &fieldPortType,
		PortStatus:      &fieldPortStatus,
		PortErrors:      &fieldPortErrors,
		RemoteSasAddr:   &fieldRemoteSasAddr,
		RemoteSasPhyId:  &fieldRemoteSasPhyId,
		RemoteSasDomain: &fieldRemoteSasDomain,
		RemoteLocId:     &fieldRemoteLocId,
		RemotePortId:    &fieldRemotePortId,
	}
}

// NsShelfPortInfo - A shelf sas port information.
type NsShelfPortInfo struct {
	// PortIdx - Index of the port, starting from 0.
	PortIdx *int64 `json:"port_idx,omitempty"`
	// PortName - Name of the port.
	PortName *string `json:"port_name,omitempty"`
	// PortType - Type of the sas port (e.g. upstream/downstream).
	PortType *NsShelfPortType `json:"port_type,omitempty"`
	// PortStatus - Status of the port.
	PortStatus *NsShelfPortStatus `json:"port_status,omitempty"`
	// PortErrors - Comma separated list of integers to indicate error conditions.
	PortErrors *string `json:"port_errors,omitempty"`
	// RemoteSasAddr - SAS address for the connected.
	RemoteSasAddr *string `json:"remote_sas_addr,omitempty"`
	// RemoteSasPhyId - Comma separated list of phy ids that this port connects to.
	RemoteSasPhyId *string `json:"remote_sas_phy_id,omitempty"`
	// RemoteSasDomain - The sas domain (A or B side) it connects to.
	RemoteSasDomain *NsShelfCtrlrSide `json:"remote_sas_domain,omitempty"`
	// RemoteLocId - The location ID of the controller that connects to this port.
	RemoteLocId *int64 `json:"remote_loc_id,omitempty"`
	// RemotePortId - The pord_id of the remote SAS port that connects to this port.
	RemotePortId *int64 `json:"remote_port_id,omitempty"`
}

// NsShelfPortInfoFieldHandles provides a string representation for each AccessControlRecord field.
type NsShelfPortInfoFieldHandles struct {
	PortIdx         *string
	PortName        *string
	PortType        *string
	PortStatus      *string
	PortErrors      *string
	RemoteSasAddr   *string
	RemoteSasPhyId  *string
	RemoteSasDomain *string
	RemoteLocId     *string
	RemotePortId    *string
}
