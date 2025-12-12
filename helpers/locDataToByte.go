package helpers

import (
	"bytes"
	"encoding/gob"

	"github.com/MGavranovic/pokedexcli/internal/pokeapi"
)

func LocDataToByte(la *pokeapi.LocationArea) ([]byte, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(la)
	if err != nil {
		return nil, err
	}

	byteData := buffer.Bytes()
	return byteData, nil
}

func LocDetailsToByte(ld *pokeapi.LocationAreaDetails) ([]byte, error) {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(ld)
	if err != nil {
		return nil, err
	}
	byteData := buffer.Bytes()
	return byteData, nil
}
