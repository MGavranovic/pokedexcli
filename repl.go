package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	var input string
	var cleanedInput []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input = scanner.Text()
			cleanedInput = cleanInput(input)
			if cmd, ok := commands[cleanedInput[0]]; ok {
				cmd.callback()
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(text string) []string {
	cleaned := strings.Split(strings.TrimSpace(strings.ToLower(text)), " ")

	return cleaned
}
