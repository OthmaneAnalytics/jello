package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

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

