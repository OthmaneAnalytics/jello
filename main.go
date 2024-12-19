package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)
func getDomainNameFromURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Hostname(), nil

}

func marshalAll[T any](items []T) ([][]byte, error) {
	var jbytes [][]byte
	for _, item := range items {
		data, err := json.Marshal(item) 
		if err != nil{
			return nil ,err
		} 
		jbytes = append(jbytes, data)
	}
	return jbytes, nil
}

const issueList = `{
	"ISSUE ONE":{
		"id": 0,
		"name": "Fix the thing",
		"estimate": 0.5,
		"completed": false
	},
	"ISSUE TWO":{
		"id":1,
		"name": "Unstick the widget",
		"estimate": 30,
		"completed": false
	}
}`



const userObject = `{
	"name": "Wayne Lagner",
	"role": "Developer",
	"remote": true
}`

const issueURL = "https://api.boot.dev/v1/courses_rest_api/learn-http/issues"

func main() {
	issues, err := getIssueData(issueURL)
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	// indent properly and print
	prettyData, err := prettify(string(issues))
	if err != nil {
		log.Fatalf("error prettifying data: %v", err)
	}
	fmt.Println(prettyData)
}

func prettify(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}

