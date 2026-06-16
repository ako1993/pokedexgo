package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ako1993/pokedexgo/internal/api"
)

var c *api.Config
var base_url = "https://pokeapi.co/api/v2/location-area/"
var mapHasBeenCalled bool

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned_input := cleanInput(text)
		user_command := cleaned_input[0]
		for _, command := range commands {
			if strings.Contains(command.name, user_command) {
				if user_command == "map" && mapHasBeenCalled {
					c = api.GetRequest(c.Next, c)
					command.callback(c)
				}
				if user_command == "map" && !mapHasBeenCalled {
					c = api.GetRequest(base_url, c)
					command.callback(c)
					mapHasBeenCalled = true
				}
				err := command.callback(c)
				if err != nil {
					log.Fatal("Error executing command")
				}
			}
		}
	}
}
