package cmd

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"

    "github.com/spf13/cobra"
)

var (
    appName string
)

var newCmd = &cobra.Command{
    Use:   "new",
    Short: "Create a new Gin application",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 1 {
            fmt.Println("Please provide an application name.")
            return
        }
        appName = args[0]
        createNewApp(appName)
    },
}

func init() {
    rootCmd.AddCommand(newCmd)
}

func createNewApp(name string) {
    fmt.Printf("Creating new Gin application: %s\n", name)

    // Create the application directory
    if err := os.Mkdir(name, 0755); err != nil {
        fmt.Println("Error creating directory:", err)
        return
    }

    // Change to the application directory
    os.Chdir(name)

    // Initialize a new Go module
    cmd := exec.Command("go", "mod", "init", name)
    if output, err := cmd.CombinedOutput(); err != nil {
        fmt.Println("Error initializing go module:", string(output))
        return
    }

    // Get Gin package
    cmd = exec.Command("go", "get", "github.com/gin-gonic/gin")
    if output, err := cmd.CombinedOutput(); err != nil {
        fmt.Println("Error getting Gin package:", string(output))
        return
    }

    // Create a main.go file
    mainFilePath := filepath.Join("main.go")
    mainFileContent := `package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello, world!",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
`
    if err := os.WriteFile(mainFilePath, []byte(mainFileContent), 0644); err != nil {
        fmt.Println("Error creating main.go file:", err)
        return
    }

    fmt.Println("Application created successfully!")
}
