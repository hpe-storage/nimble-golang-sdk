// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

var (
	client *GroupMgmtClient
)

func TestNewClient(t *testing.T) {

	// Create client
	var err error
	client, err = NewClient("10.234.48.245", "admin", "admin", "v1")
	if err != nil {
		t.Errorf("NewClient(): Unable to create client, err: %v", err.Error())
		return
	}
	fmt.Println("Session Token: ", client.SessionToken)

	tokenObjectSet := client.GetTokenObjectSet()

	payload := &nimbleos.Token{
		Username: "admin",
		Password: "admin",
	}

	// Create Token
	token, err := tokenObjectSet.CreateObject(payload)
	if err != nil {
		t.Errorf("Tokens:createObject(): Failed to get object, err: %v", err.Error())
		return
	}

	// Get Token
	token, err = tokenObjectSet.GetObject(token.ID)
	if err != nil {
		t.Errorf("Tokens:getObject(): Failed to get object, err: %v", err.Error())
		return
	}
	jsonutil.PrintPrettyJSONToConsole(token)

	// List Tokens
	tokens, err := tokenObjectSet.GetObjectList()
	if err != nil {
		t.Errorf("Tokens:getObject(): Failed to get object, err: %v", err.Error())
		return
	}
	jsonutil.PrintPrettyJSONToConsole(tokens)

	// Delete Token
	err = tokenObjectSet.DeleteObject(token.ID)
	if err != nil {
		t.Errorf("Tokens:deleteObject(): Failed to delete object, err: %v", err.Error())
		return
	}

	// Get Token again
	token, err = tokenObjectSet.GetObject(token.ID)
	if err != nil {
		t.Errorf("Tokens:getObject(): Token still exists")
		return
	}
}

func TestListGetOrPost(t *testing.T) {
	// Create GMD client
	var err error
	client, err := NewClient("gcostea-array2.rtpvlab.nimblestorage.com", "admin", "admin", "v1")
	if err != nil {
		t.Errorf("NewGmdClient(): Unable to create GMD client, err: %v", err.Error())
		return
	}
	//TODO: Create some volumes

	// List Volumes with Fields
	params := param.GetParams{
		Fields: []string{"id", "full_name", "creation_time"},
		SortBy: []param.SortOrder{{Field: "creation_time", Ascending: false}},
	}
	response, err := client.ListFromParams("volumes", &params)
	if err != nil {
		t.Errorf("Unable to list volumes, err: %v", err.Error())
	}
	switch value := response.(type) {
	case []interface{}:
		for _, val := range value {
			v := reflect.ValueOf(val)
			if v.Kind() != reflect.Map {
				t.Errorf("expected a map got %v", v.Kind())
			}
			if len(v.MapKeys()) != len(params.Fields) {
				t.Errorf("expected fields to be of len %d got %d", len(params.Fields), len(v.MapKeys()))
			}
		}
	default:
		t.Errorf("expected volumes, got kind :%v value :%v", reflect.TypeOf(value).Kind(), reflect.TypeOf(value))
	}

	// List Volumes with SearchFilter
}
