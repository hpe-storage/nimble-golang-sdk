// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsNotificationType Enum.

type NsNotificationType string

const (
	cNsNotificationTypeAlerts    NsNotificationType = "alerts"
	cNsNotificationTypeAuditlogs NsNotificationType = "auditlogs"
)

var pNsNotificationTypeAlerts NsNotificationType
var pNsNotificationTypeAuditlogs NsNotificationType

// Export
var NsNotificationTypeAlerts *NsNotificationType
var NsNotificationTypeAuditlogs *NsNotificationType

func init() {
	pNsNotificationTypeAlerts = cNsNotificationTypeAlerts
	NsNotificationTypeAlerts = &pNsNotificationTypeAlerts

	pNsNotificationTypeAuditlogs = cNsNotificationTypeAuditlogs
	NsNotificationTypeAuditlogs = &pNsNotificationTypeAuditlogs

}
