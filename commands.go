package main

import (
	"os"

	"github.com/ako1993/pokedexgo/internal/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*api.Config) error
}

func commandExit(c *api.Config) error {
	print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(c *api.Config) error {
	print("Welcome to the Pokedex!\n")
	print("Usage\n\n")
	print("help: Displays a help message\n")
	print("exit: Exit the Pokedex\n")
	print("map: Display the names of 20 location areas in the pokemon world\n")
	print("mapb: Display the previous 20 locations\n")
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
}
