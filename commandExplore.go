package main

import "fmt"

func commandExplore(c *Config) error {
	fmt.Println(*c.Explore)
	pokemon, err := c.Client.GetEnounters(*c.Explore)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemon Encounters for area: %s\n", *c.Explore)
	for _, res := range pokemon.PokemonEncounters {
		fmt.Printf("- %s\n", res.Pokemon.Name)
	}

	*c.Explore = ""
	return nil
}
