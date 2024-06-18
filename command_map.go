package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		log.Fatal(err)
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
func callbackMapBack(cfg *config) error {
	if cfg.previousLocationURL == nil {
		return errors.New("no previous page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.previousLocationURL = resp.Previous

	return nil
}