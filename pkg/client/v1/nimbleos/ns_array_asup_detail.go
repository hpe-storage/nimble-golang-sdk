// Copyright 2020 Hewlett Packard Enterprise Development LP

package nimbleos

import (
	"encoding/json"
)

// NsArrayAsupDetail - Detailed array asup information.
// Export NsArrayAsupDetailFields for advance operations like search filter etc.
var NsArrayAsupDetailFields *NsArrayAsupDetail

func init() {

	NsArrayAsupDetailFields = &NsArrayAsupDetail{
		ArrayName: "array_name",
	}
}

type NsArrayAsupDetail struct {
	// ArrayName - Unique name of array.
	ArrayName string `json:"array_name,omitempty"`
	// AsupValidate - This indicates whether the ASUP setting is enabled or disabled.
	AsupValidate *bool `json:"asup_validate,omitempty"`
	// NameResolution - Indicates whether DNS resolution succeeded.
	NameResolution *bool `json:"name_resolution,omitempty"`
	// PingfromMgmtip - Indicates whether connection from Management IP address to support server succeed.
	PingfromMgmtip *bool `json:"pingfrom_mgmtip,omitempty"`
	// PingfromCtrlra - Indicates whether connection from Controller-A to support server succeed.
	PingfromCtrlra *bool `json:"pingfrom_ctrlra,omitempty"`
	// PingfromCtrlrb - Indicates whether connection from Controller-B to support server succeed.
	PingfromCtrlrb *bool `json:"pingfrom_ctrlrb,omitempty"`
	// Heartbeat - Indicates whether heartbeat to support server succeed.
	Heartbeat *bool `json:"heartbeat,omitempty"`
	// Messages - A list of error messages.
	Messages []*NsErrorWithArguments `json:"messages,omitempty"`
}

// sdk internal struct
type nsArrayAsupDetail struct {
	ArrayName      *string                 `json:"array_name,omitempty"`
	AsupValidate   *bool                   `json:"asup_validate,omitempty"`
	NameResolution *bool                   `json:"name_resolution,omitempty"`
	PingfromMgmtip *bool                   `json:"pingfrom_mgmtip,omitempty"`
	PingfromCtrlra *bool                   `json:"pingfrom_ctrlra,omitempty"`
	PingfromCtrlrb *bool                   `json:"pingfrom_ctrlrb,omitempty"`
	Heartbeat      *bool                   `json:"heartbeat,omitempty"`
	Messages       []*NsErrorWithArguments `json:"messages,omitempty"`
}

// EncodeNsArrayAsupDetail - Transform NsArrayAsupDetail to nsArrayAsupDetail type
func EncodeNsArrayAsupDetail(request interface{}) (*nsArrayAsupDetail, error) {
	reqNsArrayAsupDetail := request.(*NsArrayAsupDetail)
	byte, err := json.Marshal(reqNsArrayAsupDetail)
	if err != nil {
		return nil, err
	}
	respNsArrayAsupDetailPtr := &nsArrayAsupDetail{}
	err = json.Unmarshal(byte, respNsArrayAsupDetailPtr)
	return respNsArrayAsupDetailPtr, err
}

// DecodeNsArrayAsupDetail Transform nsArrayAsupDetail to NsArrayAsupDetail type
func DecodeNsArrayAsupDetail(request interface{}) (*NsArrayAsupDetail, error) {
	reqNsArrayAsupDetail := request.(*nsArrayAsupDetail)
	byte, err := json.Marshal(reqNsArrayAsupDetail)
	if err != nil {
		return nil, err
	}
	respNsArrayAsupDetail := &NsArrayAsupDetail{}
	err = json.Unmarshal(byte, respNsArrayAsupDetail)
	return respNsArrayAsupDetail, err
}
