package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ako1993/pokedexgo/internal/api"
	"github.com/ako1993/pokedexgo/internal/utils"
)

var c *api.Config
var location string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned_input := utils.CleanInput(text)
		user_command := cleaned_input[0]
		_, ok := commands[user_command]
		if ok {
			if user_command == "explore" {
				location = location + cleaned_input[1]
			}
			err := commands[user_command].callback(c, location)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command. Use help to see available commands.")
		}
	}
}
