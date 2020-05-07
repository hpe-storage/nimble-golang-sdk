// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

// GetAccessControlRecordObjectSet returns a reference to the access control record object set
func (client *GroupMgmtClient) GetAccessControlRecordObjectSet() *AccessControlRecordObjectSet {
	return &AccessControlRecordObjectSet{
		Client: client,
	}
}

// GetPerformancePolicyObjectSet returns a reference to the performance policy object set
func (client *GroupMgmtClient) GetPerformancePolicyObjectSet() *PerformancePolicyObjectSet {
	return &PerformancePolicyObjectSet{
		Client: client,
	}
}

// GetTokenObjectSet returns a reference to the token object set
func (client *GroupMgmtClient) GetTokenObjectSet() *TokenObjectSet {
	return &TokenObjectSet{
		Client: client,
	}
}

// GetVolumeObjectSet returns a reference to the volume object set
func (client *GroupMgmtClient) GetVolumeObjectSet() *VolumeObjectSet {
	return &VolumeObjectSet{
		Client: client,
	}
}
