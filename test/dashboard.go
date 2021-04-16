package test

import (
	"net/http"	
	"bytes"
	"net/http/cookiejar"
	"fmt"	
	"io/ioutil"
	"encoding/json"
)

var sessionCookie []*http.Cookie
var jar *cookiejar.Jar
var runId string
var postDataClient  *http.Client


func loginDashboard() (*http.Client) {
		
	sessionCookie  = nil			
	jar, _ := cookiejar.New(nil)
	
	client := &http.Client{
		Jar: jar,
	}
	
	var jsonStr = []byte(`{"username":"tester","password":"nimbledcstester"}`)	
	req, _ := http.NewRequest("POST", "http://10.18.223.131:4000/apilogin", bytes.NewBuffer(jsonStr))
	
	req.Header.Add("Content-Type", "application/json")
	fmt.Println(req)
	resp, err := client.Do(req)
	fmt.Println(resp)
	
	if err != nil {
		panic(nil)
	}
	if err == nil{
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		resp.Body.Close()
	}			
		
	return client
}


func generateRunid()(*http.Client) {
	
	postDataClient:= loginDashboard() 
	if runId != "" {
		fmt.Println("Run Id already generated")
		fmt.Println(runId)
		return postDataClient
	}	
	newBuild:= "true"	
	json_string := "{\"TestCaseResults\": {\"Status\": \"pass\", \"Branch\": \"default\","+
		"\"NewBuild\": \""+newBuild+"\","+
		"\"BuildId\": \"1\","+
		"\"AppName\": \"Golang\","+
		"\"ArrayBuild\": \"5.087\","+
		"\"ProductName\": \"GolangSDK\","+
		"\"TestCaseName\": \"san-1235\","+
		"\"Description\": \"VolumeVolCollAssociate\","+
		"\"PlatformName\": \"Linux\","+
		"\"Duration\": \"2\","+
		"\"StartTimeStamp\":  \"120102\","+
		"\"RingId\": \"2\","+
		"\"EndTimeStamp\": \"120103\","+
		"\"TargetScope\": \"volume\","+
		"\"Protocol\": \"iscsi\","+
		"\"SuiteName\": \"Golang SDK\","+
		"\"JenkinsJob\": \"\""+
		"}}"
	fmt.Println(json_string)
	var testData = []byte(json_string)	
	req, _:= http.NewRequest("POST", "http://10.18.223.131:4000/testcasevalues", bytes.NewBuffer(testData))
	
	req.Header.Add("Content-Type", "application/json")	
	resp, err:= postDataClient.Do(req)
	
	if err != nil {
		fmt.Println(err)
		
	}
	if err == nil{
		body, _ := ioutil.ReadAll(resp.Body)
		var data map[string]interface{}
		error := json.Unmarshal([]byte(string(body)), &data)
		if error != nil {
			panic(error)
		}
		
		status, ok := data["status"].(map[string]interface{})
		if !ok {
			panic("status not found!")
		}
		
		runId = fmt.Sprintf("%v", status["RunId"])				
		resp.Body.Close()
	}
	return postDataClient	
}

func pushResultToDashboard(result string,testId string,testName string) {
	
	postDataClient:= generateRunid()		
	newBuild:= "false"	
	json_string := "{\"TestCaseResults\": {\"Status\": \""+result+"\", \"Branch\": \"default\","+
		"\"NewBuild\": \""+newBuild+"\","+
		"\"BuildId\": \"1\","+
		"\"AppName\": \"Golang\","+
		"\"ArrayBuild\": \"5.087\","+
		"\"ProductName\": \"GolangSDK\","+
		"\"TestCaseName\": \""+testId+"\","+
		"\"Description\": \""+testName+"\","+
		"\"PlatformName\": \"Linux\","+
		"\"Duration\": \"2\","+
		"\"StartTimeStamp\":  \"120102\","+
		"\"RingId\": \"2\","+
		"\"EndTimeStamp\": \"120103\","+
		"\"TargetScope\": \"volume\","+
		"\"Protocol\": \"iscsi\","+
		"\"SuiteName\": \"Golang SDK\","+
		"\"JenkinsJob\": \"\","+
		"\"RunId\": \""+runId+"\""+
		"}}"
	fmt.Println(json_string)
	var testData = []byte(json_string)
	
	req, _:= http.NewRequest("POST", "http://10.18.223.131:4000/testcasevalues", bytes.NewBuffer(testData))
	
	req.Header.Add("Content-Type", "application/json")	
	resp, err:= postDataClient.Do(req)
	fmt.Println(resp)
	
	if err != nil {
		fmt.Println(err)
		
	}
	if err == nil{
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		resp.Body.Close()
	}
		
}