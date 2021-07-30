// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfPortInfo - A shelf sas port information.
// Export NsShelfPortInfoFields for advance operations like search filter etc.
var NsShelfPortInfoFields *NsShelfPortInfoStringFields

func init() {
	PortIdxfield := "port_idx"
	PortNamefield := "port_name"
	PortTypefield := "port_type"
	PortStatusfield := "port_status"
	PortErrorsfield := "port_errors"
	RemoteSasAddrfield := "remote_sas_addr"
	RemoteSasPhyIdfield := "remote_sas_phy_id"
	RemoteSasDomainfield := "remote_sas_domain"
	RemoteLocIdfield := "remote_loc_id"
	RemotePortIdfield := "remote_port_id"

	NsShelfPortInfoFields = &NsShelfPortInfoStringFields{
		PortIdx:         &PortIdxfield,
		PortName:        &PortNamefield,
		PortType:        &PortTypefield,
		PortStatus:      &PortStatusfield,
		PortErrors:      &PortErrorsfield,
		RemoteSasAddr:   &RemoteSasAddrfield,
		RemoteSasPhyId:  &RemoteSasPhyIdfield,
		RemoteSasDomain: &RemoteSasDomainfield,
		RemoteLocId:     &RemoteLocIdfield,
		RemotePortId:    &RemotePortIdfield,
	}
}

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

// Struct for NsShelfPortInfoFields
type NsShelfPortInfoStringFields struct {
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
