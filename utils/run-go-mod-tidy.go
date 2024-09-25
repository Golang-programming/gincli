// ./utils/run-go-mod-tidy.go
package utils

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
)

// RunGoModTidy executes 'go mod tidy' in the specified project directory.
func RunGoModTidy(projectDir string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Running `go mod tidy`..."
	s.Start()
	defer s.Stop()

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		LogError(fmt.Sprintf("Error running `go mod tidy`: %s", string(output)))
	}

	s.Stop()
	LogSuccess("`go mod tidy` completed successfully.")
}
