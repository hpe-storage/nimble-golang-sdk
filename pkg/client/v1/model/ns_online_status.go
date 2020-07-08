// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsOnlineStatus Enum.

type NsOnlineStatus string

const (
	cNsOnlineStatusOffline NsOnlineStatus = "offline"
	cNsOnlineStatusOnline  NsOnlineStatus = "online"
	cNsOnlineStatusPartial NsOnlineStatus = "partial"
)

var pNsOnlineStatusOffline NsOnlineStatus
var pNsOnlineStatusOnline NsOnlineStatus
var pNsOnlineStatusPartial NsOnlineStatus

// Export
var NsOnlineStatusOffline *NsOnlineStatus
var NsOnlineStatusOnline *NsOnlineStatus
var NsOnlineStatusPartial *NsOnlineStatus

func init() {
	pNsOnlineStatusOffline = cNsOnlineStatusOffline
	NsOnlineStatusOffline = &pNsOnlineStatusOffline

	pNsOnlineStatusOnline = cNsOnlineStatusOnline
	NsOnlineStatusOnline = &pNsOnlineStatusOnline

	pNsOnlineStatusPartial = cNsOnlineStatusPartial
	NsOnlineStatusPartial = &pNsOnlineStatusPartial

}
