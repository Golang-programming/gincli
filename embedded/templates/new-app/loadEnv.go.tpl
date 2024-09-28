package main

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	fmt.Println("Loading environment variables from a configuration file")
	godotenv.Load()
}