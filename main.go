package main

import (
	"time"
	"github.com/ennoaragon/pokedex/internal/pokeapi"
)

func main() {

	c := &config{
		commands: getCommands(),
		next: "https://pokeapi.co/api/v2/location-area/",
		prev: "",
	}

	replStart(c)
}

