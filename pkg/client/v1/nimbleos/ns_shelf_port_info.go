// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsShelfPortInfo - A shelf sas port information.
// Export NsShelfPortInfoFields for advance operations like search filter etc.
var NsShelfPortInfoFields *NsShelfPortInfo

func init() {
	PortNamefield := "port_name"
	PortErrorsfield := "port_errors"
	RemoteSasAddrfield := "remote_sas_addr"
	RemoteSasPhyIdfield := "remote_sas_phy_id"

	NsShelfPortInfoFields = &NsShelfPortInfo{
		PortName:       &PortNamefield,
		PortErrors:     &PortErrorsfield,
		RemoteSasAddr:  &RemoteSasAddrfield,
		RemoteSasPhyId: &RemoteSasPhyIdfield,
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
