// Copyright 2020 Hewlett Packard Enterprise Development LP

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

var jar *cookiejar.Jar
var runId string
var postDataClient *http.Client

func loginDashboard() *http.Client {	
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}
	var jsonStr = []byte(`{"username":"test","password":"test"}`)
	req, _ := http.NewRequest("POST", "https://dcsdashboard.eng.nimblestorage.com:8080/apilogin", bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")	
	resp, err := client.Do(req)	
	if err != nil {
		panic(err)
	} else {		
		resp.Body.Close()
	}
	return client
}

func generateRunid() *http.Client {

	postDataClient := loginDashboard()
	if runId != "" {
		fmt.Println("Run Id already generated")
		fmt.Println(runId)
		return postDataClient
	}
	newBuild := "true"
	json_string := "{\"TestCaseResults\": {\"Status\": \"PASS\", \"Branch\": \"default\"," +
		"\"NewBuild\": \"" + newBuild + "\"," +
		"\"BuildId\": \"1\"," +
		"\"AppName\": \"Golang\"," +
		"\"ArrayBuild\": \"5.087\"," +
		"\"ProductName\": \"GolangSDK\"," +
		"\"TestCaseName\": \"san-1235\"," +
		"\"Description\": \"VolumeVolCollAssociate\"," +
		"\"PlatformName\": \"Linux\"," +
		"\"Duration\": \"2\"," +
		"\"StartTimeStamp\":  \"120102\"," +
		"\"RingId\": \"2\"," +
		"\"EndTimeStamp\": \"120103\"," +
		"\"TargetScope\": \"volume\"," +
		"\"Protocol\": \"iscsi\"," +
		"\"SuiteName\": \"Golang SDK\"," +
		"\"JenkinsJob\": \"\"" +
		"}}"
	fmt.Println(json_string)
	var testData = []byte(json_string)
	req, _ := http.NewRequest("POST", "https://dcsdashboard.eng.nimblestorage.com:8080/testcasevalues", bytes.NewBuffer(testData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := postDataClient.Do(req)
	if err != nil {
		fmt.Println(err)

	}else {
		body, _ := ioutil.ReadAll(resp.Body)
		var data map[string]interface{}
		error := json.Unmarshal([]byte(string(body)), &data)
		if error != nil {
			panic(error)
		}
		status, ok := data["status"].(map[string]interface{})
		if !ok {
			panic("Response status not found!")
		}
		runId = fmt.Sprintf("%v", status["RunId"])
		resp.Body.Close()
	}
	return postDataClient
}

func pushResultToDashboard(result string, testId string, testName string) {

	postDataClient := generateRunid()
	newBuild := "false"
	json_string := "{\"TestCaseResults\": {\"Status\": \"" + result + "\", \"Branch\": \"default\"," +
		"\"NewBuild\": \"" + newBuild + "\"," +
		"\"BuildId\": \"1\"," +
		"\"AppName\": \"Golang\"," +
		"\"ArrayBuild\": \"5.087\"," +
		"\"ProductName\": \"GolangSDK\"," +
		"\"TestCaseName\": \"" + testId + "\"," +
		"\"Description\": \"" + testName + "\"," +
		"\"PlatformName\": \"Linux\"," +
		"\"Duration\": \"2\"," +
		"\"StartTimeStamp\":  \"120102\"," +
		"\"RingId\": \"2\"," +
		"\"EndTimeStamp\": \"120103\"," +
		"\"TargetScope\": \"volume\"," +
		"\"Protocol\": \"iscsi\"," +
		"\"SuiteName\": \"Golang SDK\"," +
		"\"JenkinsJob\": \"\"," +
		"\"RunId\": \"" + runId + "\"" +
		"}}"
	fmt.Println(json_string)
	var testData = []byte(json_string)
	req, _ := http.NewRequest("POST", "https://dcsdashboard.eng.nimblestorage.com:8080/testcasevalues", bytes.NewBuffer(testData))
	req.Header.Add("Content-Type", "application/json")
	resp, err := postDataClient.Do(req)	
	if err != nil {
		fmt.Println(err)
	} else {		
		resp.Body.Close()
	}

}
