package main

import "os"

func callbackExit(cfg *config, query string) error {
	os.Exit(0)
	return nil
}