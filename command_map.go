package main

import (
	"fmt"
)


func commandMap(cfg *config) error {

	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	for _, location := range locationRes.Results {
		fmt.Printf("%s\n", location.Name)
	}

	cfg.prev = locationRes.Previous
	cfg.next = locationRes.Next

	return nil
}

func commandMapB(cfg *config) error {

	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.prev)
	if err != nil {
		return err
	}

	for _, val := range locationRes.Results {
		fmt.Printf("%s\n", val.Name)
	}

	cfg.prev = locationRes.Previous
	cfg.next = locationRes.Next

	return nil
}

