// Copyright 2020-2021 Hewlett Packard Enterprise Development LP

package nimbleos

// NsArrayNet - Array network config.
// Export NsArrayNetFields for advance operations like search filter etc.
var NsArrayNetFields *NsArrayNet

func init() {
	Namefield := "name"
	CtrlrASupportIpfield := "ctrlr_a_support_ip"
	CtrlrBSupportIpfield := "ctrlr_b_support_ip"

	NsArrayNetFields = &NsArrayNet{
		Name:            &Namefield,
		CtrlrASupportIp: &CtrlrASupportIpfield,
		CtrlrBSupportIp: &CtrlrBSupportIpfield,
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
