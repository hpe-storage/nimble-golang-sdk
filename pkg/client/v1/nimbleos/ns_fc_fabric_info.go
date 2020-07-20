// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsFcFabricInfo - Fibre Channel fabric information.
// Export NsFcFabricInfoFields for advance operations like search filter etc.
var NsFcFabricInfoFields *NsFcFabricInfo

func init() {

	NsFcFabricInfoFields = &NsFcFabricInfo{
		FabricSwitchName: "fabric_switch_name",
		FabricSwitchPort: "fabric_switch_port",
		FabricSwitchWwnn: "fabric_switch_wwnn",
		FabricSwitchWwpn: "fabric_switch_wwpn",
		FabricWwn:        "fabric_wwn",
		FcId:             "fc_id",
	}
}

type NsFcFabricInfo struct {
	// FabricSwitchName - Name of the Fibre Channel switch.
	FabricSwitchName string `json:"fabric_switch_name,omitempty"`
	// FabricSwitchPort - Port on the Fibre Channel switch to which connection is established.
	FabricSwitchPort string `json:"fabric_switch_port,omitempty"`
	// FabricSwitchWwnn - WWNN (World Wide Node Name) for the connected port on the fabric switch.
	FabricSwitchWwnn string `json:"fabric_switch_wwnn,omitempty"`
	// FabricSwitchWwpn - WWPN (World Wide Port Name) for the connected port on the fabric switch.
	FabricSwitchWwpn string `json:"fabric_switch_wwpn,omitempty"`
	// FabricWwn - WWNN(World Wide Node Name) for the Fabric Switch.
	FabricWwn string `json:"fabric_wwn,omitempty"`
	// FcId - FCID assigned to the Fabric Channel fabric port.
	FcId string `json:"fc_id,omitempty"`
	// LoggedIn - Login information for interface. True if interface has logged in to the Fibre Channel fabric, else false.
	LoggedIn *bool `json:"logged_in,omitempty"`
}

// sdk internal struct
type nsFcFabricInfo struct {
	FabricSwitchName *string `json:"fabric_switch_name,omitempty"`
	FabricSwitchPort *string `json:"fabric_switch_port,omitempty"`
	FabricSwitchWwnn *string `json:"fabric_switch_wwnn,omitempty"`
	FabricSwitchWwpn *string `json:"fabric_switch_wwpn,omitempty"`
	FabricWwn        *string `json:"fabric_wwn,omitempty"`
	FcId             *string `json:"fc_id,omitempty"`
	LoggedIn         *bool   `json:"logged_in,omitempty"`
}

// EncodeNsFcFabricInfo - Transform NsFcFabricInfo to nsFcFabricInfo type
func EncodeNsFcFabricInfo(request interface{}) (*nsFcFabricInfo, error) {
	reqNsFcFabricInfo := request.(*NsFcFabricInfo)
	byte, err := json.Marshal(reqNsFcFabricInfo)
	if err != nil {
		return nil, err
	}
	respNsFcFabricInfoPtr := &nsFcFabricInfo{}
	err = json.Unmarshal(byte, respNsFcFabricInfoPtr)
	return respNsFcFabricInfoPtr, err
}

// DecodeNsFcFabricInfo Transform nsFcFabricInfo to NsFcFabricInfo type
func DecodeNsFcFabricInfo(request interface{}) (*NsFcFabricInfo, error) {
	reqNsFcFabricInfo := request.(*nsFcFabricInfo)
	byte, err := json.Marshal(reqNsFcFabricInfo)
	if err != nil {
		return nil, err
	}
	respNsFcFabricInfo := &NsFcFabricInfo{}
	err = json.Unmarshal(byte, respNsFcFabricInfo)
	return respNsFcFabricInfo, err
}
