package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Loading environment variables from a configuration file")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found or error loading .env file")
	}
}
