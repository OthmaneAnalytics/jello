package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"bytes"
	"errors"
)
/*
func errIfNotHTTPS(URL string) error {
	url, err := url.Parse(URL)
	if err != nil {
		return err
	}
	if url.Scheme != "https" {
		return fmt.Errorf("URL scheme is not HTTPS: %s", URL)
	}
	return nil
}
*/
func getUsers(url string) ([]User, error) {
	fullURL := url + "?sort=experience"
	res, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var users []User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getResources(path string) []map[string]any {
	fullURL := "https://api.boot.dev" + path

	res, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	defer res.Body.Close()

	var resources []map[string]any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return resources
}

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil{
		return err
	}
	req.Header.Set("X-API-Key",apiKey)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil{
		return err
	}	
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("error")	
	}
	return nil
}

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	userdata, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(userdata))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key",apiKey)
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	req.Header.Set("X-API-Key",apiKey)
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}



func createUser(url, apiKey string, data User) (User, error) {
	userdata, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(userdata))
	if err != nil {
		return User{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key",apiKey)
	
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer res.Body.Close()

	var user User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
/*
func getUsers(url string) ([]User, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		return nil, err1
	}
	var users []User
	if err = json.Unmarshal(data, &users); err != nil{
		return nil, err
	}
	return users, nil
}
*/
type Project struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Assignees int    `json:"assignees"`
}
/*
func generateKey() string {
	const characters = "ABCDEF0123456789"
	result := ""
	rand.New(rand.NewSource(0))
	for i := 0; i < 16; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}
*/
func getProjectResponse(apiKey, url string) (Project, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Project{}, err
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return Project{}, err
	}
	defer resp.Body.Close()

	var project Project
	if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
		return Project{}, err
	}

	return project, nil
}

func putProject(apiKey, url string, project Project) error {
	client := &http.Client{}
	data, err := json.Marshal(project)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}

/*
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
*/

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
