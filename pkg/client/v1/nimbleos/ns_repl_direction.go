// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsReplDirection Enum.

type NsReplDirection string

const (
 cNsReplDirectionUpstream NsReplDirection = "upstream"
 cNsReplDirectionDownstream NsReplDirection = "downstream"
 cNsReplDirectionNone NsReplDirection = "none"
 cNsReplDirectionBiDirectional NsReplDirection = "bi_directional"
)

var pNsReplDirectionUpstream NsReplDirection
var pNsReplDirectionDownstream NsReplDirection
var pNsReplDirectionNone NsReplDirection
var pNsReplDirectionBiDirectional NsReplDirection

// NsReplDirectionUpstream enum exports
var NsReplDirectionUpstream *NsReplDirection

// NsReplDirectionDownstream enum exports
var NsReplDirectionDownstream *NsReplDirection

// NsReplDirectionNone enum exports
var NsReplDirectionNone *NsReplDirection

// NsReplDirectionBiDirectional enum exports
var NsReplDirectionBiDirectional *NsReplDirection

func init() {
 pNsReplDirectionUpstream = cNsReplDirectionUpstream
 NsReplDirectionUpstream = &pNsReplDirectionUpstream

 pNsReplDirectionDownstream = cNsReplDirectionDownstream
 NsReplDirectionDownstream = &pNsReplDirectionDownstream

 pNsReplDirectionNone = cNsReplDirectionNone
 NsReplDirectionNone = &pNsReplDirectionNone

 pNsReplDirectionBiDirectional = cNsReplDirectionBiDirectional
 NsReplDirectionBiDirectional = &pNsReplDirectionBiDirectional

}

