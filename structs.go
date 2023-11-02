package main

type Config struct {
	Base     string
	Next     string
	Previous string
	Page     int
}

type JsonResponse struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type CliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}
