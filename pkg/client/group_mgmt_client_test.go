// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hpe-storage/common-host-libs/jsonutil"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
	"github.com/stretchr/testify/assert"
)

var (
	client *GroupMgmtClient
)

func TestNewClient(t *testing.T) {

	// Create client
	var err error
	client, err = NewClient("1.1.1.1", "xxx", "xxx", "v1", true, false)
	if err != nil {
		t.Errorf("NewClient(): Unable to create client, err: %v", err.Error())
		return
	}
	fmt.Println("Session Token: ", client.SessionToken)

	tokenObjectSet := client.GetTokenObjectSet()
	user := "admin"
	pass := "admin"
	payload := &nimbleos.Token{
		Username: &user,
		Password: &pass,
	}

	// Create Token
	token, err := tokenObjectSet.CreateObject(payload)
	if err != nil {
		t.Errorf("Tokens:createObject(): Failed to get object, err: %v", err.Error())
		return
	}

	// Get Token
	token, err = tokenObjectSet.GetObject(*token.ID)
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
	err = tokenObjectSet.DeleteObject(*token.ID)
	if err != nil {
		t.Errorf("Tokens:deleteObject(): Failed to delete object, err: %v", err.Error())
		return
	}

	// Get Token again
	token, err = tokenObjectSet.GetObject(*token.ID)
	if err != nil {
		t.Errorf("Tokens:getObject(): Token still exists")
		return
	}
}

func TestListGetOrPost(t *testing.T) {
	// Create GMD client
	var err error
	client, err := NewClient("10.157.82.90", "admin", "admin", "v1", true, false)
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
func TestVolumeAvgStatLast5mins(t *testing.T) {
	respHolder := &nimbleos.Volume{}
	body := []byte(`
         {
                "data": {
                   "id": "0644aeb604d6b5e7c6000000000000000000000002",
                   "name": "hiqa-win25-ESX-Host-sjc-array2094-datastore-for-VMs",
                   "pool_name": "default",
                   "size": 7340032,
                   "usage_valid": true,
                   "vol_usage_mapped_bytes": 301340909568,
                   "reserve": 0,
                   "snap_usage_uncompressed_bytes": 0,
                   "avg_stats_last_5mins": {
                          "combined_iops": 192,
                          "combined_latency": 69,
                          "combined_throughput": 2974470,
                          "read_iops": 10,
                          "read_latency": 155,
                          "read_throughput": 1529765,
                          "write_iops": 180,
                          "write_latency": 70,
                          "write_throughput": 1444702
                   }
                }
         }`)

	dataIntf, totalRows, err := unwrapData(body, respHolder)

	assert.Nil(t, err)
	assert.Nil(t, totalRows)
	assert.NotNil(t, dataIntf)
	volume := dataIntf.(*nimbleos.Volume)
	assert.NotNil(t, volume)
	assert.NotNil(t, volume.AvgStatsLast5mins)
	assert.Equal(t, int64(10), *volume.AvgStatsLast5mins.ReadIops)
}
