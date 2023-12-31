package pokecache

// https://pokeapi.co/api/v2/pokemon/

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {

	urlToUse := base + "/pokemon/" + pokemon

	// Check Cache
	if value, ok := c.PokeCache.Get(pokemon); ok {
		resp := Pokemon{}
		err := json.Unmarshal(value, &resp)
		if err != nil {
			return Pokemon{}, err
		}
		return resp, nil
	}

	// Call API if not in cache
	r := Pokemon{}
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
		return Pokemon{}, err
	}

	c.PokeCache.Add(pokemon, body)

	return r, nil
}
