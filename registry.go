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
	callback    func(cfg *config, param string) error
}

type config struct {
	Next  *string
	Prev  *string
	cache *pokecache.Cache
}

var commands map[string]cliCommand

func commandExit(cfg *config, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, param string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func commandMap(cfg *config, param string) error {
	if cfg.Next != nil {
		val, ok := cfg.cache.Get(*cfg.Next)
		if ok {
			data, err := helpers.ByteToLocData(val)
			if err != nil {
				return err
			}
			pokeapi.PrintLocations(data)
			cfg.Next = data.Next
			cfg.Prev = data.Previous
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
				data, err := helpers.ByteToLocData(val)
				if err != nil {
					return err
				}
				pokeapi.PrintLocations(data)
				cfg.Next = data.Next
				cfg.Prev = data.Previous
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

func commandMapBack(cfg *config, param string) error {
	if cfg.Prev != nil {
		val, ok := cfg.cache.Get(*cfg.Prev)
		if ok {
			data, err := helpers.ByteToLocData(val)
			if err != nil {
				return err
			}
			pokeapi.PrintLocations(data)
			cfg.Next = data.Next
			cfg.Prev = data.Previous
		} else {
			err, next, prev, data := pokeapi.GetLocationAreas(*cfg.Prev)
			if err != nil {
				return err
			}
			byteData, err := helpers.LocDataToByte(data)
			if err != nil {
				return err
			}
			cfg.cache.Add(*cfg.Prev, byteData)
			cfg.Next = next
			cfg.Prev = prev
		}
	} else {
		fmt.Println("You're on the first page!")
	}
	return nil
}

func commandExplore(cfg *config, param string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", param)

	fmt.Printf("Exploring %s...\n", param)

	dataCache, ok := cfg.cache.Get(url)
	if ok {
		data, err := helpers.ByteToLocDetails(dataCache)
		if err != nil {
			return err
		}
		for _, pokemon := range data.PokemonEncounters {
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
	} else {
		data, err := pokeapi.GetLocationAreaDetails(url)
		if err != nil {
			fmt.Printf("Error exploring %s...\n", param)
			return err
		}
		fmt.Printf("Found pokemon:\n")
		for _, pokemon := range data.PokemonEncounters {
			fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
		}
		byteData, errCon := helpers.LocDetailsToByte(data)
		if errCon != nil {
			return errCon
		}
		cfg.cache.Add(url, byteData)
	}

	return nil
}
