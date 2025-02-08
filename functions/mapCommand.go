package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// CommandMap fetches and displays the next 20 location areas
func MapCommand(cfg *Config) error {
	// Use the NextURL if available, otherwise use the default endpoint
	url := cfg.NextURL
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	// Fetch data from the PokeAPI
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch location areas: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the JSON response
	var result struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
		} `json:"results"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Update the config with the Next and Previous URLs
	cfg.NextURL = result.Next
	cfg.PreviousURL = result.Previous

	// Display the location areas
	fmt.Println("Location areas:")
	for _, area := range result.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func MapbCommand(cfg *Config) error {
	// Use the PreviousURL if available
	if cfg.PreviousURL == "" {
		fmt.Println("You're already at the beginning of the list.")
		return nil
	}

	// Fetch data from the PokeAPI
	resp, err := http.Get(cfg.PreviousURL)
	if err != nil {
		return fmt.Errorf("failed to fetch location areas: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the JSON response
	var result struct {
		Next     string `json:"next"`
		Previous string `json:"previous"`
		Results  []struct {
			Name string `json:"name"`
		} `json:"results"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Update the config with the Next and Previous URLs
	cfg.NextURL = result.Next
	cfg.PreviousURL = result.Previous

	// Display the location areas
	fmt.Println("Location areas:")
	for _, area := range result.Results {
		fmt.Println(area.Name)
	}

	return nil
}
