// Copyright 2021 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsJWKAlgorithm Enum.

type NsJWKAlgorithm string

const (
	cNsJWKAlgorithmRs256 NsJWKAlgorithm = "RS256"
)

var pNsJWKAlgorithmRs256 NsJWKAlgorithm

// NsJWKAlgorithmRs256 enum exports
var NsJWKAlgorithmRs256 *NsJWKAlgorithm

func init() {
	pNsJWKAlgorithmRs256 = cNsJWKAlgorithmRs256
	NsJWKAlgorithmRs256 = &pNsJWKAlgorithmRs256

}
