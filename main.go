package main

import "github.com/djd223/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaURL *string
	previousLocationAreaURL *string
}

func main() {
	cfg := config {
		pokeapiClient : pokeapi.NewClient(),
	}
	startRepl(&cfg)
}
