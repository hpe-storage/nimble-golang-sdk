// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsPartnerCfgSyncStatus Enum.

type NsPartnerCfgSyncStatus string

const (
	cNsPartnerCfgSyncStatusNA  NsPartnerCfgSyncStatus = "N/A"
	cNsPartnerCfgSyncStatusNo  NsPartnerCfgSyncStatus = "No"
	cNsPartnerCfgSyncStatusYes NsPartnerCfgSyncStatus = "Yes"
)

var pNsPartnerCfgSyncStatusNA NsPartnerCfgSyncStatus
var pNsPartnerCfgSyncStatusNo NsPartnerCfgSyncStatus
var pNsPartnerCfgSyncStatusYes NsPartnerCfgSyncStatus

// Export
var NsPartnerCfgSyncStatusNA *NsPartnerCfgSyncStatus
var NsPartnerCfgSyncStatusNo *NsPartnerCfgSyncStatus
var NsPartnerCfgSyncStatusYes *NsPartnerCfgSyncStatus

func init() {
	pNsPartnerCfgSyncStatusNA = cNsPartnerCfgSyncStatusNA
	NsPartnerCfgSyncStatusNA = &pNsPartnerCfgSyncStatusNA

	pNsPartnerCfgSyncStatusNo = cNsPartnerCfgSyncStatusNo
	NsPartnerCfgSyncStatusNo = &pNsPartnerCfgSyncStatusNo

	pNsPartnerCfgSyncStatusYes = cNsPartnerCfgSyncStatusYes
	NsPartnerCfgSyncStatusYes = &pNsPartnerCfgSyncStatusYes

}
