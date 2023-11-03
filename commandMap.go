package main

import (
	"fmt"
)

func commandMap(c *Config) error {
	locations, err := c.Client.GetLocations(c.Next)
	if err != nil {
		return err
	}

	fmt.Println("Next set of 20 maps")
	for _, res := range locations.Results {
		fmt.Printf("Map Name: %s\n", res.Name)
	}
	c.Next = locations.Next
	c.Previous = locations.Previous
	return nil
}

func commandMapb(c *Config) error {
	locations, err := c.Client.GetLocations(c.Previous)
	if err != nil {
		return err
	}

	fmt.Println("Next set of 20 maps")
	for _, res := range locations.Results {
		fmt.Printf("Map Name: %s\n", res.Name)
	}
	c.Next = locations.Next
	c.Previous = locations.Previous
	return nil
}
