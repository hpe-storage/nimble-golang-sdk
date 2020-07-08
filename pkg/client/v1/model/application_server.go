// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// ApplicationServer - An application server is an external agent that collaborates with an array to manage storage resources; for example, Volume Shadow Copy Service (VSS) or VMware.
// Export ApplicationServerFields for advance operations like search filter etc.
var ApplicationServerFields *ApplicationServer

func init() {

	ApplicationServerFields = &ApplicationServer{
		ID:          "id",
		Name:        "name",
		Hostname:    "hostname",
		Username:    "username",
		Description: "description",
		Password:    "password",
	}
}

type ApplicationServer struct {
	// ID - Identifier for the application server.
	ID string `json:"id,omitempty"`
	// Name - Name for the application server.
	Name string `json:"name,omitempty"`
	// Hostname - Application server hostname.
	Hostname string `json:"hostname,omitempty"`
	// Port - Application server port number.
	Port int64 `json:"port,omitempty"`
	// Username - Application server username.
	Username string `json:"username,omitempty"`
	// Description - Text description of application server.
	Description string `json:"description,omitempty"`
	// Password - Application server password.
	Password string `json:"password,omitempty"`
	// ServerType - Application server type ({invalid|vss|vmware|cisco|stack_vision|container_node}).
	ServerType *NsAppServerType `json:"server_type,omitempty"`
	// Metadata - Key-value pairs that augment an application server's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// CreationTime - Time when this application server was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this application server was last modified.
	LastModified int64 `json:"last_modified,omitempty"`
}

// sdk internal struct
type applicationServer struct {
	// ID - Identifier for the application server.
	ID *string `json:"id,omitempty"`
	// Name - Name for the application server.
	Name *string `json:"name,omitempty"`
	// Hostname - Application server hostname.
	Hostname *string `json:"hostname,omitempty"`
	// Port - Application server port number.
	Port *int64 `json:"port,omitempty"`
	// Username - Application server username.
	Username *string `json:"username,omitempty"`
	// Description - Text description of application server.
	Description *string `json:"description,omitempty"`
	// Password - Application server password.
	Password *string `json:"password,omitempty"`
	// ServerType - Application server type ({invalid|vss|vmware|cisco|stack_vision|container_node}).
	ServerType *NsAppServerType `json:"server_type,omitempty"`
	// Metadata - Key-value pairs that augment an application server's attributes.
	Metadata []*NsKeyValue `json:"metadata,omitempty"`
	// CreationTime - Time when this application server was created.
	CreationTime *int64 `json:"creation_time,omitempty"`
	// LastModified - Time when this application server was last modified.
	LastModified *int64 `json:"last_modified,omitempty"`
}

// EncodeApplicationServer - Transform ApplicationServer to applicationServer type
func EncodeApplicationServer(request interface{}) (*applicationServer, error) {
	reqApplicationServer := request.(*ApplicationServer)
	byte, err := json.Marshal(reqApplicationServer)
	objPtr := &applicationServer{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeApplicationServer Transform applicationServer to ApplicationServer type
func DecodeApplicationServer(request interface{}) (*ApplicationServer, error) {
	reqApplicationServer := request.(*applicationServer)
	byte, err := json.Marshal(reqApplicationServer)
	obj := &ApplicationServer{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
