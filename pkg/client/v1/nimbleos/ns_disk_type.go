// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsDiskType Enum.

type NsDiskType string

const (
 cNsDiskTypeSsd NsDiskType = "ssd"
 cNsDiskTypeHdd NsDiskType = "hdd"
)

var pNsDiskTypeSsd NsDiskType
var pNsDiskTypeHdd NsDiskType

// NsDiskTypeSsd enum exports
var NsDiskTypeSsd *NsDiskType

// NsDiskTypeHdd enum exports
var NsDiskTypeHdd *NsDiskType

func init() {
 pNsDiskTypeSsd = cNsDiskTypeSsd
 NsDiskTypeSsd = &pNsDiskTypeSsd

 pNsDiskTypeHdd = cNsDiskTypeHdd
 NsDiskTypeHdd = &pNsDiskTypeHdd

}

