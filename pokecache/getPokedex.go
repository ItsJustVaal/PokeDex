package pokecache

import "fmt"

func (c *Client) GetPokedex() error {
	if len(c.PokeCache.cache) == 0 {
		return fmt.Errorf("Pokedex Empty")
	}
	fmt.Println("Pokemon in Pokedex")
	for k, _ := range c.PokeCache.cache {
		fmt.Printf(" -%s\n", k)
	}
	return nil
}
