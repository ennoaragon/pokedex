package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func commandMapB(cfg *config) error {

	if cfg.prev == "" {
		return fmt.Errorf("Reached the end of prev of the list")
	}

	res, err := http.Get(cfg.prev)

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
		cfg.prev = ""
	}

	if body.Next != nil {
		cfg.next = *body.Next
	} else {
		cfg.next = ""
	}

	return nil
}

