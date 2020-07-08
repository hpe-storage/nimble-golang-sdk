// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsNetConfigName Enum.

type NsNetConfigName string

const (
	cNsNetConfigNameBackup NsNetConfigName = "backup"
	cNsNetConfigNameDraft  NsNetConfigName = "draft"
	cNsNetConfigNameActive NsNetConfigName = "active"
)

var pNsNetConfigNameBackup NsNetConfigName
var pNsNetConfigNameDraft NsNetConfigName
var pNsNetConfigNameActive NsNetConfigName

// Export
var NsNetConfigNameBackup *NsNetConfigName
var NsNetConfigNameDraft *NsNetConfigName
var NsNetConfigNameActive *NsNetConfigName

func init() {
	pNsNetConfigNameBackup = cNsNetConfigNameBackup
	NsNetConfigNameBackup = &pNsNetConfigNameBackup

	pNsNetConfigNameDraft = cNsNetConfigNameDraft
	NsNetConfigNameDraft = &pNsNetConfigNameDraft

	pNsNetConfigNameActive = cNsNetConfigNameActive
	NsNetConfigNameActive = &pNsNetConfigNameActive

}
