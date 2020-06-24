/**
 * Copyright 2017 Hewlett Packard Enterprise Development LP
 */

package model



// NsFCSession - Fibre Channel initiator session information.
// Export NsFCSessionFields for advance operations like search filter etc.
var NsFCSessionFields *NsFCSession

func init(){
	IDfield:= "id"
	SessionIDfield:= "session_id"
	InitiatorAliasfield:= "initiator_alias"
	InitiatorWwpnfield:= "initiator_wwpn"
	InitiatorWwnnfield:= "initiator_wwnn"
	InitiatorSwitchNamefield:= "initiator_switch_name"
	InitiatorSwitchPortfield:= "initiator_switch_port"
	InitiatorSymbolicPortnamefield:= "initiator_symbolic_portname"
	InitiatorSymbolicNodenamefield:= "initiator_symbolic_nodename"
	TargetPortArrayNamefield:= "target_port_array_name"
	TargetPortInterfaceNamefield:= "target_port_interface_name"
	TargetWwnnfield:= "target_wwnn"
	TargetWwpnfield:= "target_wwpn"
		
	NsFCSessionFields= &NsFCSession{
		ID: &IDfield,
		SessionID: &SessionIDfield,
		InitiatorAlias: &InitiatorAliasfield,
		InitiatorWwpn: &InitiatorWwpnfield,
		InitiatorWwnn: &InitiatorWwnnfield,
		InitiatorSwitchName: &InitiatorSwitchNamefield,
		InitiatorSwitchPort: &InitiatorSwitchPortfield,
		InitiatorSymbolicPortname: &InitiatorSymbolicPortnamefield,
		InitiatorSymbolicNodename: &InitiatorSymbolicNodenamefield,
		TargetPortArrayName: &TargetPortArrayNamefield,
		TargetPortInterfaceName: &TargetPortInterfaceNamefield,
		TargetWwnn: &TargetWwnnfield,
		TargetWwpn: &TargetWwpnfield,
		
	}
}

type NsFCSession struct {
   
    // Unique identifier of the Fibre Channel session.
    
 	ID *string `json:"id,omitempty"`
   
    // Unique identifier of the Fibre Channel session.
    
 	SessionID *string `json:"session_id,omitempty"`
   
    // ALUA (Asymmetric Logical Unit Access) value of the Fibre Channel session.
    
   	Alua *NsALUA `json:"alua,omitempty"`
   
    // Registered persistent reservation key for this session on associated logical unit (if there is no registration, its value will be zero).
    
  	PrKey  *int64 `json:"pr_key,omitempty"`
   
    // Alias of the Fibre Channel initiator.
    
 	InitiatorAlias *string `json:"initiator_alias,omitempty"`
   
    // WWPN (World Wide Port Name) of the Fibre Channel initiator.
    
 	InitiatorWwpn *string `json:"initiator_wwpn,omitempty"`
   
    // WWNN (World Wide Node Name) of the Fibre Channel initiator.
    
 	InitiatorWwnn *string `json:"initiator_wwnn,omitempty"`
   
    // Name of the switch used by the Fibre Channel initiator.
    
 	InitiatorSwitchName *string `json:"initiator_switch_name,omitempty"`
   
    // Port on the switch used by the Fibre Channel initiator.
    
 	InitiatorSwitchPort *string `json:"initiator_switch_port,omitempty"`
   
    // Symbolic port name associated with the Fibre Channel initiator.
    
 	InitiatorSymbolicPortname *string `json:"initiator_symbolic_portname,omitempty"`
   
    // Symbolic node name associated with the Fibre Channel initiator.
    
 	InitiatorSymbolicNodename *string `json:"initiator_symbolic_nodename,omitempty"`
   
    // FCID assigned to the Fibre Channel initiator.
    
   	InitiatorFcID *int64 `json:"initiator_fcid,omitempty"`
   
    // Name of the array hosting the Fibre Channel target port.
    
 	TargetPortArrayName *string `json:"target_port_array_name,omitempty"`
   
    // Identifier of the controller hosting the Fibre Channel target port.
    
   	TargetPortCtrlrID *int64 `json:"target_port_ctrlr_id,omitempty"`
   
    // Name of the interface hosted on the Fibre Channel target port.
    
 	TargetPortInterfaceName *string `json:"target_port_interface_name,omitempty"`
   
    // WWNN (World Wide Node Name) of the Fibre Channel target port.
    
 	TargetWwnn *string `json:"target_wwnn,omitempty"`
   
    // WWPN (World Wide Port Name) of the Fibre Channel target port.
    
 	TargetWwpn *string `json:"target_wwpn,omitempty"`
   
    // FCID assigned to the Fibre Channel target port.
    
   	TargetFcID *int64 `json:"target_fcid,omitempty"`
}
