package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ako1993/pokedexgo/internal/api"
)

var c *api.Config

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned_input := cleanInput(text)
		user_command := cleaned_input[0]
		for _, command := range commands {
			if command.name == user_command {
				err := command.callback(c)
				if err != nil {
					log.Fatal("Error executing command")
				}
			}
		}
	}
}
