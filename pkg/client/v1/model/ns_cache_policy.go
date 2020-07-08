// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsCachePolicy Enum.

type NsCachePolicy string

const (
	cNsCachePolicyNormal                NsCachePolicy = "normal"
	cNsCachePolicyNoWrite               NsCachePolicy = "no_write"
	cNsCachePolicyAggressiveReadNoWrite NsCachePolicy = "aggressive_read_no_write"
	cNsCachePolicyDisabled              NsCachePolicy = "disabled"
	cNsCachePolicyAggressive            NsCachePolicy = "aggressive"
)

var pNsCachePolicyNormal NsCachePolicy
var pNsCachePolicyNoWrite NsCachePolicy
var pNsCachePolicyAggressiveReadNoWrite NsCachePolicy
var pNsCachePolicyDisabled NsCachePolicy
var pNsCachePolicyAggressive NsCachePolicy

// Export
var NsCachePolicyNormal *NsCachePolicy
var NsCachePolicyNoWrite *NsCachePolicy
var NsCachePolicyAggressiveReadNoWrite *NsCachePolicy
var NsCachePolicyDisabled *NsCachePolicy
var NsCachePolicyAggressive *NsCachePolicy

func init() {
	pNsCachePolicyNormal = cNsCachePolicyNormal
	NsCachePolicyNormal = &pNsCachePolicyNormal

	pNsCachePolicyNoWrite = cNsCachePolicyNoWrite
	NsCachePolicyNoWrite = &pNsCachePolicyNoWrite

	pNsCachePolicyAggressiveReadNoWrite = cNsCachePolicyAggressiveReadNoWrite
	NsCachePolicyAggressiveReadNoWrite = &pNsCachePolicyAggressiveReadNoWrite

	pNsCachePolicyDisabled = cNsCachePolicyDisabled
	NsCachePolicyDisabled = &pNsCachePolicyDisabled

	pNsCachePolicyAggressive = cNsCachePolicyAggressive
	NsCachePolicyAggressive = &pNsCachePolicyAggressive

}
