package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	commands map[string]cliCommand
	next string
	prev string
}

func getCommands() map[string]cliCommand{

	return map[string]cliCommand {
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: "map",
			description: "Get 20 locations",
			callback: commandMap,
		},
		"mapb": {
			name: "mapb",
			description: "Get previous 20 locations",
			callback: commandMapB,
		},
	}
}

func cleanInput(text string) [] string {
	words := strings.ToLower(text)
	word := strings.Fields(words)

	return word
}

func replStart(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0{
			continue
		}

		first_word := words[0]

		if cmd, ok := cfg.commands[first_word]; ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println()
			}

			continue
		} else {
			fmt.Println("Unkown command")

			continue
		}

	}

}
