package utils

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func RunGoModTidy(projectDir string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Running `go mod tidy`..."
	s.Start()
	defer s.Stop()

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectDir
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running `go mod tidy`: %v\n", err)
		os.Exit(1)
	}

	s.Stop()
	fmt.Println(color.New(color.FgGreen).Sprint("`go mod tidy` completed successfully."))
}
