package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationArea struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(url string) (error, *string, *string, *LocationArea) {
	resp, err := http.Get(url)
	if err != nil {
		return err, nil, nil, &LocationArea{}
	}
	var data LocationArea
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return err, nil, nil, &LocationArea{}
	}
	PrintLocations(&data)
	return nil, data.Next, data.Previous, &data
}
