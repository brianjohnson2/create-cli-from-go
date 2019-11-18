package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/gookit/color"
)

var filePath = "/usr/local/bin/"
var fileName = flag.String("file", "main.go", "The file to build.")

// Builds go file, defaults to main.go unless file specified with the -file flag
func buildFromGoFile() {

	flag.Parse()
	fmt.Println("Building file:", *fileName)
	buildCmd := exec.Command("go", "build", "-o", os.Args[1], *fileName)
	buildCmd.Stderr = os.Stderr
	err5 := buildCmd.Run()
	if err5 != nil {
		color.Red.Println("error: unable to build go file.")
		os.Exit(1)
	}
}

// Copys the built command into the file path specified in the FilePath var
func copyCommandtoPath() {

	buildFromGoFile()
	copyCommand := exec.Command("cp", os.Args[1], filePath)
	err4 := copyCommand.Run()
	if err4 != nil {
		color.Red.Println("error: unable to copy command to " + filePath)
		os.Exit(1)
	}

	color.Green.Println("success: created command " + os.Args[1])
}

func main() {

	// Checks to see if argument for command name was entered
	if len(os.Args) != 2 {
		err1 := errors.New("error: expecting name to create command")
		if err1 != nil {
			color.Red.Println(err1)
			os.Exit(1)
		}
	}

	filePath += os.Args[1]

	// Checks if command already exists and prompts to overwrite
	checkCommandAlreadyExists, _ := os.Stat(filePath)
	if checkCommandAlreadyExists != nil {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("The file already exists. Do you want to overwrite in? (y/n)")
		for scanner.Scan() {
			userInput := scanner.Text()
			fmt.Println(userInput)
			switch userInput {
			case "y":
				fmt.Println("Proceeding with command..")
				copyCommandtoPath()
				os.Exit(0)
			default:
				abort := errors.New("abort: cancelling command creation")
				if abort != nil {
					fmt.Println(abort)
					os.Exit(1)
				}
			}
		}
	} else {
		copyCommandtoPath()
		os.Exit(0)
	}
}
