package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Configuration struct {
	ApiToken string
}

func readConfig() Configuration {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return configuration
}

func main() {
	conf := readConfig()

	// Create a new request
	req, err := http.NewRequest("GET", "https://api.track.toggl.com/api/v9/me", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set up basic authentication
	// Replace "your_username" and "your_password" with your actual credentials
	req.SetBasicAuth(conf.ApiToken, "api_token")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print response to console
	fmt.Println(string(body))
}

func auth() {

}
