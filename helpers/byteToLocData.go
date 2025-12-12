package helpers

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io"

	"github.com/MGavranovic/pokedexcli/internal/pokeapi"
)

func ByteToLocData(la []byte) (*pokeapi.LocationArea, error) {
	var data pokeapi.LocationArea

	reader := bytes.NewReader(la)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&data)
	if !errors.Is(err, io.EOF) && err != nil {
		return &pokeapi.LocationArea{}, err
	}
	return &data, nil
}

func ByteToLocDetails(la []byte) (*pokeapi.LocationAreaDetails, error) {
	var data pokeapi.LocationAreaDetails

	reader := bytes.NewReader(la)
	decoder := gob.NewDecoder(reader)
	err := decoder.Decode(&data)
	if !errors.Is(err, io.EOF) && err != nil {
		return &pokeapi.LocationAreaDetails{}, err
	}
	return &data, nil
}
