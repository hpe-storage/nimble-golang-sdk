// Copyright 2020 Hewlett Packard Enterprise Development LP
package nimbleos

// Golang package for NsShelfSwType Enum.

type NsShelfSwType string

const (
	cNsShelfSwTypeDiskShelf                NsShelfSwType = "Disk Shelf"
	cNsShelfSwTypeUnknownShelfSoftwareType NsShelfSwType = "unknown shelf software type"
	cNsShelfSwTypeAllFlashShelf            NsShelfSwType = "All Flash Shelf"
	cNsShelfSwTypeHeadShelf                NsShelfSwType = "Head Shelf"
)

var pNsShelfSwTypeDiskShelf NsShelfSwType
var pNsShelfSwTypeUnknownShelfSoftwareType NsShelfSwType
var pNsShelfSwTypeAllFlashShelf NsShelfSwType
var pNsShelfSwTypeHeadShelf NsShelfSwType

// NsShelfSwTypeDiskShelf enum exports
var NsShelfSwTypeDiskShelf *NsShelfSwType

// NsShelfSwTypeUnknownShelfSoftwareType enum exports
var NsShelfSwTypeUnknownShelfSoftwareType *NsShelfSwType

// NsShelfSwTypeAllFlashShelf enum exports
var NsShelfSwTypeAllFlashShelf *NsShelfSwType

// NsShelfSwTypeHeadShelf enum exports
var NsShelfSwTypeHeadShelf *NsShelfSwType

func init() {
	pNsShelfSwTypeDiskShelf = cNsShelfSwTypeDiskShelf
	NsShelfSwTypeDiskShelf = &pNsShelfSwTypeDiskShelf

	pNsShelfSwTypeUnknownShelfSoftwareType = cNsShelfSwTypeUnknownShelfSoftwareType
	NsShelfSwTypeUnknownShelfSoftwareType = &pNsShelfSwTypeUnknownShelfSoftwareType

	pNsShelfSwTypeAllFlashShelf = cNsShelfSwTypeAllFlashShelf
	NsShelfSwTypeAllFlashShelf = &pNsShelfSwTypeAllFlashShelf

	pNsShelfSwTypeHeadShelf = cNsShelfSwTypeHeadShelf
	NsShelfSwTypeHeadShelf = &pNsShelfSwTypeHeadShelf

}
