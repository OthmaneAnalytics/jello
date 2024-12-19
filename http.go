package main

import (
	"fmt"
	"io"
	"net/http"
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

