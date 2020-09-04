// Copyright 2020 Hewlett Packard Enterprise Development LP

package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/client/v1/nimbleos"
	"github.com/hpe-storage/nimble-golang-sdk/pkg/param"
)

const (
	groupURIFmt        = "https://%s:5392/%s"
	clientTimeout      = time.Second * 60 // 1 Minute
	maxLoginRetries    = 5
	retrySleepDuration = 2 // Seconds
)

// GroupMgmtClient :
type GroupMgmtClient struct {
	URL          string
	Client       *resty.Client
	SessionToken string
}

// DataWrapper is used to represent a generic JSON API payload
type DataWrapper struct {
	Data          interface{} `json:"data"`
	StartRow      *int        `json:"startRow,omitempty"`
	EndRow        *int        `json:"endRow,omitempty"`
	PageSize      *int        `json:"pageSize,omitempty"`
	TotalRows     *int        `json:"totalRows,omitempty"`
	OperationType *string     `json:"operationType,omitempty"`
}

// ErrorResponse is a serializer struct for representing a valid JSON API errors payload.
type ErrorResponse struct {
	Messages []*Message `json:"messages"`
}

// Message is an `Error` implementation as well as an implementation of the JSON API error object.
type Message struct {
	Code string `json:"code,omitempty"`
	Text string `json:"text,omitempty"`
}

// NewClient instantiates a new client to communicate with the Nimble group
func NewClient(ipAddress, username, password, apiVersion string) (*GroupMgmtClient, error) {
	// Create GroupMgmt Client
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{
		InsecureSkipVerify: true,
	})
	groupMgmtClient := &GroupMgmtClient{
		URL:    fmt.Sprintf("https://%s:5392/%s", ipAddress, apiVersion),
		Client: client,
	}

	// Get session token
	sessionToken, err := groupMgmtClient.login(username, password)
	if err != nil {
		return nil, err
	}
	groupMgmtClient.SessionToken = sessionToken
	return groupMgmtClient, nil
}

// EnableDebug : Enables debug logging of client request/response
func (client *GroupMgmtClient) EnableDebug() {
	client.Client.SetDebug(true)
}

func (client *GroupMgmtClient) login(username, password string) (string, error) {
	// Construct Payload
	appName := "Go sdkv1 client"
	token := &nimbleos.Token{
		Username: &username,
		Password: &password,
		AppName:  &appName,
	}
	token, err := client.GetTokenObjectSet().CreateObject(token)
	if err != nil {
		return "", err
	}
	return *token.SessionToken, err
}

// Post :
func (client *GroupMgmtClient) Post(path string, payload interface{}, respHolder interface{}) (interface{}, error) {
	// build the url
	url := fmt.Sprintf("%s/%s", client.URL, path)
	// Post it
	response, err := client.Client.R().
		SetHeader("X-Auth-Token", client.SessionToken).
		SetBody(&DataWrapper{
			Data: payload,
		}).
		Post(url)
	if err != nil {
		return nil, err
	}
	// Http error
	if !response.IsSuccess() {
		errResp, err := unwrapError(response.Body())
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("http response error: status (%d), messages: %v", response.StatusCode(), errResp)
	}
	return unwrap(response.Body(), respHolder)
}

// Put :
func (client *GroupMgmtClient) Put(path, id string, payload interface{}, respHolder interface{}) (interface{}, error) {
	// build the url
	url := fmt.Sprintf("%s/%s/%s", client.URL, path, id)

	// Put it
	response, err := client.Client.R().
		SetHeader("X-Auth-Token", client.SessionToken).
		SetBody(&DataWrapper{
			Data: payload,
		}).
		Put(url)
	if err != nil {
		return nil, err
	}
	// http Error
	if !response.IsSuccess() {
		errResp, err := unwrapError(response.Body())
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("http response error: status (%d), messages: %v", response.StatusCode(), errResp)
	}
	return unwrap(response.Body(), respHolder)
}

// Get : Only used to get a single object with the given ID
func (client *GroupMgmtClient) Get(path string, id string, intf interface{}) (interface{}, error) {
	// build the url
	url := fmt.Sprintf("%s/%s/%s", client.URL, path, id)

	// Get it
	response, err := client.Client.R().
		SetHeader("X-Auth-Token", client.SessionToken).
		Get(url)
	if err != nil {
		return nil, err
	}

	// convert a 404 (not found) to nil response
	if response.StatusCode() == 404 {
		return nil, nil
	}

	if response.IsSuccess() {
		// TODO: handle 202 accepted... this might translate into an async job that we need to wait for

		// unmarshal the response
		wrapper := &DataWrapper{
			Data: intf,
		}
		err = json.Unmarshal(response.Body(), wrapper)
		if err != nil {
			return nil, err
		}

		// return it
		return wrapper.Data, nil
	}
	// error condition unmarshalled
	wrapper := &ErrorResponse{}
	err = json.Unmarshal(response.Body(), wrapper)
	errResp, err := unwrapError(response.Body())
	if err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("http response error: status (%d), messages: %v", response.StatusCode(), errResp)
}

// Delete :
func (client *GroupMgmtClient) Delete(path string, id string) error {
	// build the url
	url := fmt.Sprintf("%s/%s/%s", client.URL, path, id)

	// delete it
	response, err := client.Client.R().
		SetHeader("X-Auth-Token", client.SessionToken).
		Delete(url)
	if err != nil {
		return err
	}

	// http Error
	if !response.IsSuccess() {
		errResp, err := unwrapError(response.Body())
		if err != nil {
			return err
		}
		return fmt.Errorf("http response error: status (%d), messages: %v", response.StatusCode(), errResp)
	}
	return nil
}

// List without any params
func (client *GroupMgmtClient) List(path string) (interface{}, error) {
	listResp, err := client.ListFromParams(path, nil)
	if err != nil {
		return nil, err
	}
	return listResp, nil
}

// ListFromParams :
func (client *GroupMgmtClient) ListFromParams(path string, params *param.GetParams) (interface{}, error) {

	wrapper, err := client.listGetOrPost(path, params)
	if err != nil {
		return nil, err
	}

	if params != nil && params.Page != nil {
		params.Page.TotalRows = wrapper.TotalRows
	}
	return wrapper.Data, nil
}

func (client *GroupMgmtClient) listGetOrPost(path string, params *param.GetParams) (*DataWrapper, error) {
	if params == nil {
		return client.listGet(path, nil)
	}

	// load the url query parameters
	queryParams, err := params.BuildQueryParts()
	if err != nil {
		return nil, err
	}

	// check if advanced criteria post is required
	if params.Filter != nil {
		simpleMap, _ := params.Filter.AsSimpleMap()
		if simpleMap == nil {
			fetch := "fetch"
			wrapper := &DataWrapper{
				Data:          params.Filter,
				StartRow:      params.Page.StartRow,
				EndRow:        params.Page.EndRow,
				OperationType: &fetch,
			}
			// complex filter, need to POST it
			postResp, err := client.listPost(path, wrapper, queryParams)
			if err != nil {
				return nil, err
			}
			return postResp, nil
		} else {
			// get request
			getResp, err := client.listGet(path, queryParams)
			if err != nil {
				return nil, err
			}
			return getResp, nil
		}
	} else {
		// get request
		getResp, err := client.listGet(path, queryParams)
		if err != nil {
			return nil, err
		}
		return getResp, nil
	}
}

// listPost uses a post request to get all objects on the path using an advanced criteria
func (client *GroupMgmtClient) listPost(
	path string,
	payload *DataWrapper,
	queryParams map[string]string,
) (*DataWrapper, error) {
	// build the url
	url := fmt.Sprintf("%s/%s/detail", client.URL, path)

	fmt.Printf("request = %v", payload)

	// Post it
	response, err := client.Client.R().
		SetQueryParams(queryParams).
		SetHeader("X-Auth-Token", client.SessionToken).
		SetBody(payload).
		Post(url)
	if err != nil {
		return nil, err
	}

	// http Error
	if !response.IsSuccess() {
		errResp, err := unwrapError(response.Body())
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("http response error: status (%d), messages: %v", response.StatusCode(), errResp)
	}

	// unmarshal the response
	wrapper := &DataWrapper{}
	err = json.Unmarshal(response.Body(), wrapper)
	if err != nil {
		return nil, err
	}
	// return it
	return wrapper, nil
}

// listGet uses a get request to get all objects on the path
func (client *GroupMgmtClient) listGet(
	path string,
	queryParams map[string]string,
) (*DataWrapper, error) {
	// build the url
	url := fmt.Sprintf("%s/%s/detail", client.URL, path)

	response, err := client.Client.R().
		SetQueryParams(queryParams).
		SetHeader("X-Auth-Token", client.SessionToken).
		Get(url)
	if err != nil {
		return nil, err
	}

	// http Error
	if !response.IsSuccess() {
		errResp, err := unwrapError(response.Body())
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("http response error: status (%d), messages: %v", response.StatusCode(), errResp)
	}

	// unmarshal the response
	wrapper := &DataWrapper{}
	err = json.Unmarshal(response.Body(), wrapper)
	if err != nil {
		return nil, err
	}

	// return it
	return wrapper, nil
}

// unwrap a response body
func unwrap(body []byte, payload interface{}) (interface{}, error) {
	// TODO: add some logging

	// unmarshal the response
	wrapper := &DataWrapper{
		Data: payload,
	}
	err := json.Unmarshal(body, wrapper)
	if err != nil {
		return nil, err
	}

	// return it
	return wrapper.Data, nil
}

//unwrap error response
func unwrapError(body []byte) (string, error) {
	errResp := ""
	wrapper := &ErrorResponse{}
	err := json.Unmarshal(body, wrapper)
	if err != nil {
		return errResp, err
	}

	for _, emsg := range wrapper.Messages {
		errResp += fmt.Sprintf("%+v", *emsg)
	}
	return errResp, nil
}
