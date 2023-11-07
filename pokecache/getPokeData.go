package pokecache

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokeData(pokemon *string) error {
	resp := Pokemon{}
	if value, ok := c.PokeCache.Get(*pokemon); ok {
		err := json.Unmarshal(value, &resp)
		if err != nil {
			return err
		}
		// Need to format this to show what I actually want it to show
		fmt.Println("Found Pokemon in Pokedex!")
		fmt.Printf("Name: %s\n", resp.Name)
		fmt.Printf("Name: %v\n", resp.Stats)
		fmt.Printf("Name: %v\n", resp.Abilities)
	} else {
		// This error never goes off
		return fmt.Errorf("Pokemon not in the pokedex")
	}
	return nil
}
