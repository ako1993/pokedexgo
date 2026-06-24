package main

import (
	"bufio"
	"fmt"
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
		_, ok := commands[user_command]
		if ok {
			err := commands[user_command].callback(c)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command. Use help to see available commands.")
		}
	}
}
