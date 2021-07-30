// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArrayNet - Array network config.

// Export NsArrayNetFields provides field names to use in filter parameters, for example.
var NsArrayNetFields *NsArrayNetStringFields

func init() {
	fieldName := "name"
	fieldMemberGid := "member_gid"
	fieldCtrlrASupportIp := "ctrlr_a_support_ip"
	fieldCtrlrBSupportIp := "ctrlr_b_support_ip"
	fieldNicList := "nic_list"

	NsArrayNetFields = &NsArrayNetStringFields{
		Name:            &fieldName,
		MemberGid:       &fieldMemberGid,
		CtrlrASupportIp: &fieldCtrlrASupportIp,
		CtrlrBSupportIp: &fieldCtrlrBSupportIp,
		NicList:         &fieldNicList,
	}
}

type NsArrayNet struct {
	// Name - Name of the array.
	Name *string `json:"name,omitempty"`
	// MemberGid - GID of member. This field cannot be updated.
	MemberGid *int64 `json:"member_gid,omitempty"`
	// CtrlrASupportIp - IP address of controller A.
	CtrlrASupportIp *string `json:"ctrlr_a_support_ip,omitempty"`
	// CtrlrBSupportIp - IP address of controller B.
	CtrlrBSupportIp *string `json:"ctrlr_b_support_ip,omitempty"`
	// NicList - List of NICs.
	NicList []*NsNIC `json:"nic_list,omitempty"`
}

// Struct for NsArrayNetFields
type NsArrayNetStringFields struct {
	Name            *string
	MemberGid       *string
	CtrlrASupportIp *string
	CtrlrBSupportIp *string
	NicList         *string
}
