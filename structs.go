package main

import "pokedex/pokecache"

type Config struct {
	Client   pokecache.Client
	Next     *string
	Previous *string
	Explore  *string
}

type CliCommand struct {
	name        string
	description string
	callback    func(c *Config) error
}
