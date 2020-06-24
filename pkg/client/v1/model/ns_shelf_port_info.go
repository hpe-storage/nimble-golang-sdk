/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsShelfPortInfo - A shelf sas port information.
// Export NsShelfPortInfoFields for advance operations like search filter etc.
var NsShelfPortInfoFields *NsShelfPortInfo

func init(){
	PortNamefield:= "port_name"
	PortErrorsfield:= "port_errors"
	RemoteSasAddrfield:= "remote_sas_addr"
	RemoteSasPhyIDfield:= "remote_sas_phy_id"
		
	NsShelfPortInfoFields= &NsShelfPortInfo{
		PortName: &PortNamefield,
		PortErrors: &PortErrorsfield,
		RemoteSasAddr: &RemoteSasAddrfield,
		RemoteSasPhyID: &RemoteSasPhyIDfield,
		
	}
}

type NsShelfPortInfo struct {
   
    // Index of the port, starting from 0.
    
   	PortIDx *int64 `json:"port_idx,omitempty"`
   
    // Name of the port.
    
 	PortName *string `json:"port_name,omitempty"`
   
    // Type of the sas port (e.g. upstream/downstream).
    
   	PortType *NsShelfPortType `json:"port_type,omitempty"`
   
    // Status of the port.
    
   	PortStatus *NsShelfPortStatus `json:"port_status,omitempty"`
   
    // Comma separated list of integers to indicate error conditions.
    
 	PortErrors *string `json:"port_errors,omitempty"`
   
    // SAS address for the connected.
    
 	RemoteSasAddr *string `json:"remote_sas_addr,omitempty"`
   
    // Comma separated list of phy ids that this port connects to.
    
 	RemoteSasPhyID *string `json:"remote_sas_phy_id,omitempty"`
   
    // The sas domain (A or B side) it connects to.
    
   	RemoteSasDomain *NsShelfCtrlrSIDe `json:"remote_sas_domain,omitempty"`
   
    // The location ID of the controller that connects to this port.
    
  	RemoteLocID  *int64 `json:"remote_loc_id,omitempty"`
   
    // The pord_id of the remote SAS port that connects to this port.
    
  	RemotePortID  *int64 `json:"remote_port_id,omitempty"`
}
