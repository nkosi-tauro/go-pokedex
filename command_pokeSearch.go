package main

import (
	"errors"
	"fmt"
)

func callbackPokeSearch(cfg *config, query string) error {
	if query == "" {
		return errors.New("please enter a pokemon name")
	}

	resp, err := cfg.pokeapiClient.GetPokemonByName(query)
	if err != nil {
		return err
	}
	// TODO: add more details
	fmt.Printf("Your pokemon: %s details:\n ", query)
	fmt.Printf("Species Name %s\n", resp.Species.Name)
	fmt.Printf("Moves: \n")
	for _, move := range resp.Moves{
		fmt.Printf(" - %s\n", move.Move.Name)
	}
	fmt.Printf("Abilities: \n")
	for _, ability := range resp.Abilities{
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}
	fmt.Printf("Order: \n")
	fmt.Printf(" - %v\n", resp.Order)

	return nil
}