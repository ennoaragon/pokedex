package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type PokemonLocationResponse struct{
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []LocationArea `json:"results"`
}

type LocationArea struct{
	Name string `json:"name"`
	Url string `json:"url"`
}


func commandMap(cfg *config) error {

	if cfg.next == "" {
		return fmt.Errorf("reached the end of the list no new pokemon location areas")
	}
	res, err := http.Get(cfg.next)

	defer res.Body.Close()

	var body PokemonLocationResponse

	err = json.NewDecoder(res.Body).Decode(&body)

	if err != nil {
		return err
	}

	for _, val := range body.Results {
		fmt.Printf("%s\n", val.Name)
	}

	if body.Previous != nil {
		cfg.prev = *body.Previous
	} else {
		cfg.prev = cfg.next
	}

	if body.Next != nil {
		cfg.next = *body.Next
	} else {
		cfg.next = ""
	}

	return nil
}
