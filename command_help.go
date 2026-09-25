package main

import "fmt"

func commandHelp(c *config) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n\n\n")

	for _, val := range c.commands {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	fmt.Println()
	return nil
}

