package pokeapi

import "fmt"

func PrintLocations(la *LocationArea) {
	for _, loc := range la.Results {
		fmt.Println(loc.Name)
	}
}
