package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ako1993/pokedexgo/internal/pokecache"
)

type Config struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var base_url = "https://pokeapi.co/api/v2/location-area/"
var mapHasBeenCalled bool
var user_config *Config
var cache = pokecache.NewCache(7 * time.Second)
var url_to_cache string

func GetRequest(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response code error")
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetLocationPage(url string) (*Config, error) {
	data, ok := cache.Get(url)
	if !ok {
		var err error
		data, err = GetRequest(url)
		if err != nil {
			return nil, err
		}
		cache.Add(url, data)
	}
	var config Config
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}

func CommandMap(c *Config) error {
	if mapHasBeenCalled {
		url_to_cache = user_config.Next
		c, err := GetLocationPage(url_to_cache)
		if err != nil {
			return err
		}
		user_config = c
		for _, result := range user_config.Results {
			fmt.Println(result.Name)
		}
	}

	if !mapHasBeenCalled {
		url_to_cache = base_url
		c, err := GetLocationPage(base_url)
		if err != nil {
			return err
		}
		user_config = c
		for _, result := range user_config.Results {
			fmt.Println(result.Name)
		}
		mapHasBeenCalled = true

	}
	return nil
}

func CommandMapb(c *Config) error {
	if user_config == nil || user_config.Previous == "" {
		fmt.Println("You are on the first page. Use the map command to navigate forward")
	} else if user_config != nil && user_config.Previous == "" {
		fmt.Println("You are on the first page. Use the map command to navigate forward")
	} else if user_config != nil && user_config.Previous != "" {
		url_to_cache = user_config.Previous
		c, err := GetLocationPage(url_to_cache)
		if err != nil {
			return err
		}
		user_config = c
		for _, result := range user_config.Results {
			fmt.Println(result.Name)
		}

	}

	return nil
}
