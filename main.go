package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/gookit/color"
)

var (
	filePath = "/usr/local/bin/"
	fileName = flag.String("file", "main.go", "The file to build.")
	err      error
)

// Builds go file, defaults to main.go unless file specified with the -file flag
func buildFromGoFile() {

	flag.Parse()
	buildCmd := exec.Command("go", "build", "-o", os.Args[1], *fileName)
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		color.Red.Println("error: unable to build go file.")
	} else {
		fmt.Println("Building file:", *fileName)
	}
}

// Copys the built command into the file path specified in the FilePath var
func copyCommandtoPath() {

	buildFromGoFile()
	copyCommand := exec.Command("cp", os.Args[1], filePath)
	if err := copyCommand.Run(); err != nil {
		color.Red.Println("error: unable to copy command to " + filePath)
	} else {
		color.Green.Println("success: created command " + os.Args[1])
	}
}

func main() {

	// Checks to see if argument for command name was entered
	if len(os.Args) != 2 {
		color.Red.Println("Usage: [command-name]")
		return
	}

	filePath += os.Args[1]

	// Checks if command already exists and prompts to overwrite
	checkCommandAlreadyExists, _ := os.Stat(filePath)
	if checkCommandAlreadyExists != nil {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("The file already exists. Do you want to overwrite in? (y)")
		for scanner.Scan() {
			userInput := scanner.Text()
			fmt.Println(userInput)
			switch userInput {
			case "y":
				fmt.Println("Proceeding with command..")
				copyCommandtoPath()
				return
			default:
				color.Red.Println("abort: cancelling command creation")
			}
		}
	} else {
		copyCommandtoPath()
	}
}
