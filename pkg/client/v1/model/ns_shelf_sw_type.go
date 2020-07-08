// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

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

// Export
var NsShelfSwTypeDiskShelf *NsShelfSwType
var NsShelfSwTypeUnknownShelfSoftwareType *NsShelfSwType
var NsShelfSwTypeAllFlashShelf *NsShelfSwType
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
