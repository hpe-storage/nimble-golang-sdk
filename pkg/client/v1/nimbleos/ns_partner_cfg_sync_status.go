// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsPartnerCfgSyncStatus Enum.

type NsPartnerCfgSyncStatus string

const (
 cNsPartnerCfgSyncStatusNA NsPartnerCfgSyncStatus = "N/A"
 cNsPartnerCfgSyncStatusNo NsPartnerCfgSyncStatus = "No"
 cNsPartnerCfgSyncStatusYes NsPartnerCfgSyncStatus = "Yes"
)

var pNsPartnerCfgSyncStatusNA NsPartnerCfgSyncStatus
var pNsPartnerCfgSyncStatusNo NsPartnerCfgSyncStatus
var pNsPartnerCfgSyncStatusYes NsPartnerCfgSyncStatus

// NsPartnerCfgSyncStatusNA enum exports
var NsPartnerCfgSyncStatusNA *NsPartnerCfgSyncStatus

// NsPartnerCfgSyncStatusNo enum exports
var NsPartnerCfgSyncStatusNo *NsPartnerCfgSyncStatus

// NsPartnerCfgSyncStatusYes enum exports
var NsPartnerCfgSyncStatusYes *NsPartnerCfgSyncStatus

func init() {
 pNsPartnerCfgSyncStatusNA = cNsPartnerCfgSyncStatusNA
 NsPartnerCfgSyncStatusNA = &pNsPartnerCfgSyncStatusNA

 pNsPartnerCfgSyncStatusNo = cNsPartnerCfgSyncStatusNo
 NsPartnerCfgSyncStatusNo = &pNsPartnerCfgSyncStatusNo

 pNsPartnerCfgSyncStatusYes = cNsPartnerCfgSyncStatusYes
 NsPartnerCfgSyncStatusYes = &pNsPartnerCfgSyncStatusYes

}

