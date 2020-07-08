// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsProtectionPolicyType Enum.

type NsProtectionPolicyType string

const (
	cNsProtectionPolicyTypeProtectionTemplate NsProtectionPolicyType = "protection_template"
	cNsProtectionPolicyTypeVolumeCollection   NsProtectionPolicyType = "volume_collection"
)

var pNsProtectionPolicyTypeProtectionTemplate NsProtectionPolicyType
var pNsProtectionPolicyTypeVolumeCollection NsProtectionPolicyType

// Export
var NsProtectionPolicyTypeProtectionTemplate *NsProtectionPolicyType
var NsProtectionPolicyTypeVolumeCollection *NsProtectionPolicyType

func init() {
	pNsProtectionPolicyTypeProtectionTemplate = cNsProtectionPolicyTypeProtectionTemplate
	NsProtectionPolicyTypeProtectionTemplate = &pNsProtectionPolicyTypeProtectionTemplate

	pNsProtectionPolicyTypeVolumeCollection = cNsProtectionPolicyTypeVolumeCollection
	NsProtectionPolicyTypeVolumeCollection = &pNsProtectionPolicyTypeVolumeCollection

}
