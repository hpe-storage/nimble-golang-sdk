// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsFolderCreateAttr - Attributes for folder creation.
// Export NsFolderCreateAttrFields for advance operations like search filter etc.
var NsFolderCreateAttrFields *NsFolderCreateAttr

func init() {

	NsFolderCreateAttrFields = &NsFolderCreateAttr{
		Name:   "name",
		PoolId: "pool_id",
	}
}

type NsFolderCreateAttr struct {
	// Name - Name of folder.
	Name string `json:"name,omitempty"`
	// PoolId - ID of pool to create the folder in.
	PoolId string `json:"pool_id,omitempty"`
}

// sdk internal struct
type nsFolderCreateAttr struct {
	// Name - Name of folder.
	Name *string `json:"name,omitempty"`
	// PoolId - ID of pool to create the folder in.
	PoolId *string `json:"pool_id,omitempty"`
}

// EncodeNsFolderCreateAttr - Transform NsFolderCreateAttr to nsFolderCreateAttr type
func EncodeNsFolderCreateAttr(request interface{}) (*nsFolderCreateAttr, error) {
	reqNsFolderCreateAttr := request.(*NsFolderCreateAttr)
	byte, err := json.Marshal(reqNsFolderCreateAttr)
	objPtr := &nsFolderCreateAttr{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsFolderCreateAttr Transform nsFolderCreateAttr to NsFolderCreateAttr type
func DecodeNsFolderCreateAttr(request interface{}) (*NsFolderCreateAttr, error) {
	reqNsFolderCreateAttr := request.(*nsFolderCreateAttr)
	byte, err := json.Marshal(reqNsFolderCreateAttr)
	obj := &NsFolderCreateAttr{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
