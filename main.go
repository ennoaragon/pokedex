package main

import (
	"time"
	"github.com/ennoaragon/pokedex/internal/pokeapi"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)

	c := &config{
		commands: getCommands(),
		pokeapiClient: pokeClient,
	}

	replStart(c)
}

