// Copyright 2020-2021 Hewlett Packard Enterprise Development LP
package nimbleos

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

// NsOnlineStatusOffline enum exports
var NsOnlineStatusOffline *NsOnlineStatus

// NsOnlineStatusOnline enum exports
var NsOnlineStatusOnline *NsOnlineStatus

// NsOnlineStatusPartial enum exports
var NsOnlineStatusPartial *NsOnlineStatus

func init() {
	pNsOnlineStatusOffline = cNsOnlineStatusOffline
	NsOnlineStatusOffline = &pNsOnlineStatusOffline

	pNsOnlineStatusOnline = cNsOnlineStatusOnline
	NsOnlineStatusOnline = &pNsOnlineStatusOnline

	pNsOnlineStatusPartial = cNsOnlineStatusPartial
	NsOnlineStatusPartial = &pNsOnlineStatusPartial

}
