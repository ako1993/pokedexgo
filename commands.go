package main

import (
	"fmt"
	"os"

	"github.com/ako1993/pokedexgo/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*api.Config, string) error
}

func commandExit(c *api.Config, location string) error {
	print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *api.Config, location string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("map: Display the names of 20 location areas in the pokemon world")
	fmt.Println("mapb: Display the previous 20 locations")
	fmt.Println("explore: Display the Pokemon in a location")
	return nil
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Display a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Display the names of 20 location areas in the pokemon world",
		callback:    api.CommandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Allows a user to look at the previous page of locations",
		callback:    api.CommandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Shows the pokemon in a selected aread",
		callback:    api.CommandExplore,
	},
}
