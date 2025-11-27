package main

import (
	"fmt"
	"os"

	"github.com/MGavranovic/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	Next string
	Prev string
}

var commands map[string]cliCommand

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	err := pokeapi.GetLocationAreas("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return err
	}
	return nil
}
