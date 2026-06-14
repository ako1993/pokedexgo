package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned_input := cleanInput(text)
		user_command := cleaned_input[0]
		fmt.Printf("\rYour command was: %s\n", user_command)
	}
}
