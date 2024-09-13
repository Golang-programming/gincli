package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load environment variables from a configuration file
	fmt.Println("Loading environment variables from a configuration file")
	godotenv.Load()
}