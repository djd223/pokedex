package main

import(
	"fmt"
	"log"

	"github.com/djd223/pokedex/internal/pokeapi"
)

func commandMapf() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	
	return nil
}