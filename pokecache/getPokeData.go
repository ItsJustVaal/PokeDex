package pokecache

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetPokeData(pokemon *string) error {
	if value, ok := c.PokeCache.Get(*pokemon); ok {
		resp := Pokemon{}
		err := json.Unmarshal(value, &resp)
		if err != nil {
			return err
		}
		// Need to format this to show what I actually want it to show
		fmt.Println("Found Pokemon in Pokedex!")
		fmt.Printf("Name: %s\n", resp.Name)
		fmt.Printf("Name: %d\n", resp.Height)
		fmt.Printf("Weight: %d\n", resp.Weight)
		fmt.Println("Stats:")
		for x := range resp.Stats {
			fmt.Printf("  -%s: %d\n", resp.Stats[x].Stat.Name, resp.Stats[x].BaseStat)
		}
		fmt.Println("Types:")
		for x := range resp.Types {
			fmt.Printf("  -%s\n", resp.Types[x].Type.Name)
		}
	}
	return fmt.Errorf("Pokemon not in the pokedex")
}
