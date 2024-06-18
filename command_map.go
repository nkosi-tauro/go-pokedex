package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, query string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	// // allow pagination
	cfg.nextLocationURL = resp.Next
	cfg.previousLocationURL = resp.Previous

	return err
}
func callbackMapBack(cfg *config, query string) error {
	if cfg.previousLocationURL == nil {
		return errors.New("no previous page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.previousLocationURL = resp.Previous

	return nil
}