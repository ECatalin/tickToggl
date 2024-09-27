package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
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
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	fmt.Println("Today's date:", today, "Tomorrow's date:", tomorrow)

	uri := "https://api.track.toggl.com/api/v9/me/time_entries"
	req, err := http.NewRequest(http.MethodGet, uri, nil)

	// Create a new request
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	q := req.URL.Query()
	q.Add("start_date", today)
	q.Add("end_date", tomorrow)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Set up basic authentication
	// TODO: Switch to session cookie
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
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, body, "", "\t")
	if error != nil {
		fmt.Println("JSON parse error: ", error)
		return
	}

	fmt.Println(prettyJSON.String())
}
