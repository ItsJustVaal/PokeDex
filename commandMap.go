package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandMap(c *Config) error {
	urlToUse := ""
	if c.Page == 0 {
		urlToUse += c.Base
	} else {
		urlToUse += c.Next
	}
	r := JsonResponse{}
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
		log.Fatal(err)
	}
	c.Next = r.Next
	c.Previous = r.Previous
	c.Page++
	fmt.Println("Next set of 20 maps")
	for _, res := range r.Results {
		fmt.Printf("Map Name: %s\n", res.Name)
	}
	return nil
}
