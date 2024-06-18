package main

import (
	"fmt"
	"log"

	"github.com/nkosi-tauro/go-pokedex/pokeapi"
)

func callbackMap() error {
	pokeapiCLient := pokeapi.NewClient()

	resp, err := pokeapiCLient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return err
}