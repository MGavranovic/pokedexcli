package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MGavranovic/pokedexcli/internal/pokecache"
)

func repl() {
	cfg := config{
		Next:  nil,
		Prev:  nil,
		cache: pokecache.NewCache(5 * time.Second),
	}

	var input string
	var cleanedInput []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input = scanner.Text()
			cleanedInput = cleanInput(input)
			if cmd, ok := commands[cleanedInput[0]]; ok {
				err := cmd.callback(&cfg)
				if err != nil {
					fmt.Printf("Error running %s: %s\n", cmd.name, err)
				}
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
