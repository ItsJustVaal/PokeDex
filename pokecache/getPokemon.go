package pokecache

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemon(location string) (PokemonEncounters, error) {

	urlToUse := base + "/location-area/" + location
	fmt.Printf("URL In Use: %v\n", urlToUse)

	// Check Cache
	fmt.Println("Checking Cache")
	if value, ok := c.PokeCache.Get(location); ok {
		resp := PokemonEncounters{}
		err := json.Unmarshal(value, &resp)
		if err != nil {
			return PokemonEncounters{}, err
		}
		fmt.Println("Found Resposne")
		return resp, nil
	}

	// Call API if not in cache
	fmt.Println("Calling API")
	r := PokemonEncounters{}
	res, err := http.Get(urlToUse)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return PokemonEncounters{}, err
	}

	c.PokeCache.Add(location, body)
	return r, nil
}
