package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"pokedex/pokecache"
	"strings"
	"time"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UTC().UnixNano())
	c := Config{
		Client: pokecache.NewClient(5*time.Second, time.Minute*5),
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if commandName == "explore" {
			if len(words) > 1 && words[1] != "" {
				c.Explore = &words[1]
			}
		} else if commandName == "catch" {
			if len(words) > 1 && words[1] != "" {
				c.Pokemon = &words[1]
			}
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(&c)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon encounters for an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Displays the pokemon encounters for an area",
			callback:    catchPokemon,
		},
	}
}
