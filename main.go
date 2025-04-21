// ./main.go
package main

import (
	"fmt"
	"os"
	"github.com/golang-programming/gincli/cmd"
)

var (
	version   = "dev"
	buildDate = "unknown"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("gincli version %s (built on %s)\n", version, buildDate)
		return
	}
	cmd.Execute()
}
