package main

import (
	"fmt"
	"os"

	"github.com/MGavranovic/pokedexcli/helpers"
	"github.com/MGavranovic/pokedexcli/internal/pokeapi"
	"github.com/MGavranovic/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	Next  *string
	Prev  *string
	cache *pokecache.Cache
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
	if cfg.Next != nil {
		val, ok := cfg.cache.Get(*cfg.Next)
		if ok {
			fmt.Println(val)
		} else {
			err, next, prev, data := pokeapi.GetLocationAreas(*cfg.Next)
			if err != nil {
				return err
			}
			byteData, err := helpers.LocDataToByte(data)
			if err != nil {
				return err
			}
			cfg.cache.Add(*cfg.Next, byteData)
			cfg.Next = next
			cfg.Prev = prev

		}
	} else {
		if cfg.Next == nil && cfg.Prev != nil {
			fmt.Println("You're on the last page!")
		} else {
			val, ok := cfg.cache.Get("https://pokeapi.co/api/v2/location-area")
			if ok {
				fmt.Println(val)
			} else {
				err, next, prev, data := pokeapi.GetLocationAreas("https://pokeapi.co/api/v2/location-area")
				if err != nil {
					return err
				}
				byteData, err := helpers.LocDataToByte(data)
				if err != nil {
					return err
				}
				cfg.cache.Add("https://pokeapi.co/api/v2/location-area", byteData)
				cfg.Next = next
				cfg.Prev = prev
			}
		}
	}
	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.Prev != nil {
		err, next, prev, _ := pokeapi.GetLocationAreas(*cfg.Prev)
		if err != nil {
			return err
		} else {
			cfg.Next = next
			cfg.Prev = prev
		}
	} else {
		fmt.Println("You're on the first page!")
	}
	return nil
}
