package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetRequest(url string, c *Config) *Config {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request %v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error executing request %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Server returned non-200 status: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body %v", err)
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func CommandMap(c *Config) error {
	for _, result := range c.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func CommandMapb(c *Config) error {
	return nil
}
