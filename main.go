package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"log"
)

func main() {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}
	defer res.Body.Close()

	var projects []Project
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&projects)
	if err != nil {
		log.Fatalf("error decoding response: %v", err)
	}

	logProjects(projects)
}

type Project struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Assigness int    `json:"assignees"`
}

func logProjects(projects []Project) {
	for _, p := range projects {
		fmt.Printf("Project: %s, Complete: %v\n", p.Title, p.Completed)
	}
}

func getMailtoLinkForEmail(email string) string {
	return "mailto:" + email
}

func getDomainNameFromURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return parsedURL.Hostname(), nil

}

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}
	pass, ok := parsedUrl.User.Password()
	if !ok {
		pass = ""
	}
	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: pass,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
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
/*
func main() {
	issues, err := getIssues(domain)
	if err != nil {
		log.Fatalf("error getting issues data: %v", err)
	}
	logIssues(issues)
}
*/
func prettify(data string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(data), "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting JSON: %w", err)
	}
	return prettyJSON.String(), nil
}

