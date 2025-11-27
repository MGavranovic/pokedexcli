package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationArea struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous any       `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	var data LocationArea
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&data); err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}
