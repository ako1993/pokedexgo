package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
				err := command.callback()
				if err != nil {
					print(err.Error(), "\n")
				}
			}
		}
	}
}
