// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsProtectionPolicyType Enum.

type NsProtectionPolicyType string

const (
	cNsProtectionPolicyTypeProtectionTemplate NsProtectionPolicyType = "protection_template"
	cNsProtectionPolicyTypeVolumeCollection   NsProtectionPolicyType = "volume_collection"
)

var pNsProtectionPolicyTypeProtectionTemplate NsProtectionPolicyType
var pNsProtectionPolicyTypeVolumeCollection NsProtectionPolicyType

// NsProtectionPolicyTypeProtectionTemplate enum exports
var NsProtectionPolicyTypeProtectionTemplate *NsProtectionPolicyType

// NsProtectionPolicyTypeVolumeCollection enum exports
var NsProtectionPolicyTypeVolumeCollection *NsProtectionPolicyType

func init() {
	pNsProtectionPolicyTypeProtectionTemplate = cNsProtectionPolicyTypeProtectionTemplate
	NsProtectionPolicyTypeProtectionTemplate = &pNsProtectionPolicyTypeProtectionTemplate

	pNsProtectionPolicyTypeVolumeCollection = cNsProtectionPolicyTypeVolumeCollection
	NsProtectionPolicyTypeVolumeCollection = &pNsProtectionPolicyTypeVolumeCollection

}
