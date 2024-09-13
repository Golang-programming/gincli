package cmd

import "fmt"

// Input Helper: Ask for missing input when the user didn't provide enough arguments
func PromptForInput(prompt string) string {
    fmt.Println(prompt)
    var input string
    fmt.Scanln(&input)
    return input
}
