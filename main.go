package main

import (
	"fmt"
	"log"
)

func main() {
	issues, err := getIssueData()
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}
	s := string(issues)
	fmt.Println(s)
}

