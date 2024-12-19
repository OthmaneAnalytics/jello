package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)


func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var dnss DNSResponse

	if err = json.Unmarshal(body,&dnss); err != nil {
		return "", err
	}
	return string(dnss.Answer[0].Data), nil
}


func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()
	data, err1 := io.ReadAll(res.Body)
	if err1 != nil {
		return nil, fmt.Errorf("error creading data: %w", err1)
	}
	return data, nil
}
/*
func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()
	data, err1 := io.ReadAll(res.Body)
	var issues []Issue
	if err1 != nil {
		return nil, fmt.Errorf("error creading data: %w", err1)
	}
	if err = json.Unmarshal(data, &issues); err != nil {
		return nil, err
	}
	return issues, nil
}
*/

func getIssues2(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	if err:= decoder.Decode(&issues); err != nil{
		return nil, fmt.Errorf("error decoding data: %w", err)
	}
	return issues, nil
}

type Issue struct {
	Title string
}

func getIssues(domain string) ([]Issue, error) {
	res, err := http.Get("https://" + domain + "/v1/courses_rest_api/learn-http/issues")
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func logIssues(issues []Issue) {
	for _, issue := range issues {
		fmt.Println(issue.Title)
	}
}

