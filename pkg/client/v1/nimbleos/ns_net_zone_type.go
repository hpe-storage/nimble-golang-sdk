// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsNetZoneType Enum.

type NsNetZoneType string

const (
 cNsNetZoneTypeSingle NsNetZoneType = "single"
 cNsNetZoneTypeEvenodd NsNetZoneType = "evenodd"
 cNsNetZoneTypeBisect NsNetZoneType = "bisect"
 cNsNetZoneTypeNone NsNetZoneType = "none"
)

var pNsNetZoneTypeSingle NsNetZoneType
var pNsNetZoneTypeEvenodd NsNetZoneType
var pNsNetZoneTypeBisect NsNetZoneType
var pNsNetZoneTypeNone NsNetZoneType

// NsNetZoneTypeSingle enum exports
var NsNetZoneTypeSingle *NsNetZoneType

// NsNetZoneTypeEvenodd enum exports
var NsNetZoneTypeEvenodd *NsNetZoneType

// NsNetZoneTypeBisect enum exports
var NsNetZoneTypeBisect *NsNetZoneType

// NsNetZoneTypeNone enum exports
var NsNetZoneTypeNone *NsNetZoneType

func init() {
 pNsNetZoneTypeSingle = cNsNetZoneTypeSingle
 NsNetZoneTypeSingle = &pNsNetZoneTypeSingle

 pNsNetZoneTypeEvenodd = cNsNetZoneTypeEvenodd
 NsNetZoneTypeEvenodd = &pNsNetZoneTypeEvenodd

 pNsNetZoneTypeBisect = cNsNetZoneTypeBisect
 NsNetZoneTypeBisect = &pNsNetZoneTypeBisect

 pNsNetZoneTypeNone = cNsNetZoneTypeNone
 NsNetZoneTypeNone = &pNsNetZoneTypeNone

}

