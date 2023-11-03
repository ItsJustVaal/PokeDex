package main

import (
	"fmt"
	"math/rand"
)

func catchPokemon(c *Config) error {
	pokemon, err := c.Client.GetPokemon(*c.Pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a pokeball at %s...\n", *c.Pokemon)
	chanceToCatch := pokemon.BaseExperience * randInt(1, 10)
	minToCatch := pokemon.BaseExperience * 5

	if chanceToCatch >= minToCatch {
		fmt.Println("Pokemon is caught! Adding to pokedex")
	} else {
		fmt.Println("Pokemon got away!")
	}

	*c.Pokemon = ""
	return nil
}

func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}
