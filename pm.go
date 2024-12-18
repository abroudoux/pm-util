package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed assets/ascii.txt
var asciiArt string

var referenceFilePath string = "/tmp/pm_reference_file"

func main() {
	err := checkIfReferenceFileExists()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		flagMode()
	}

	println("No command provided. Use 'pm --help' to see the available options.")
}

func checkIfReferenceFileExists() error {
	referenceFileContent, err := os.ReadFile(referenceFilePath)
	if err != nil {
		err := createReferenceFile()
		if err != nil {
			return fmt.Errorf("error creating reference file: %v", err)
		}
	}

	if string(referenceFileContent) == "" {
		return fmt.Errorf("reference file not found")
	}

	return nil
}

func createReferenceFile() error {
	file, err := os.Create(referenceFilePath)
	if err != nil {
		return fmt.Errorf("error creating reference file: %v", err)
	}

	defer file.Close()

	err = setReferenceFile("package.json")
	if err != nil {
		return fmt.Errorf("error setting reference file: %v", err)
	}

	println("Reference file successfully created!")
	return nil
}

func setReferenceFile(referenceFile string) error {
	err := os.WriteFile(referenceFilePath, []byte(referenceFile), 0644)
	if err != nil {
		return fmt.Errorf("error setting reference file: %v", err)
	}

	println("Target file set to", referenceFile)
	return nil
}

func flagMode() {
	flag := os.Args[1]

	if flag == "--file" || flag == "-f" {
		if len(os.Args) == 2 {
			err := printReferenceFile()
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
		} else {
			referenceFile := os.Args[2]
			err := setReferenceFile(referenceFile)
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
		}
	} else if flag == "--help" || flag == "-h" {
		printHelpMenu()
	} else if flag == "--root" || flag == "-r" {
		err := goToFileReference()
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	} else if flag == "--version" || flag == "-v" {
		fmt.Println(asciiArt)
		println("2.0.0")
	} else {
		err := runCommandInReferenceFileDirectory()
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func printReferenceFile() error {
	referenceFileContent, err := os.ReadFile(referenceFilePath)
	if err != nil {
		return fmt.Errorf("error reading reference file: %v", err)
	}

	println("Current reference file: " + strings.TrimSpace(string(referenceFileContent)))
	return nil
}

func getReferenceFile() (string, error) {
	referenceFileContent, err := os.ReadFile(referenceFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading reference file: %v", err)
	}

	if string(referenceFileContent) == "" {
		return "", fmt.Errorf("reference file not found")
	}

	return string(referenceFileContent), nil
}

func printHelpMenu() {
	println("Usage: pm [options] [command]")
	fmt.Printf("  %-20s %s\n", "pm [command]", "Go to the directory of the project and run the command")
	fmt.Printf("  %-20s %s\n", "pm [--file | -f] [file]", "Set the target file")
	fmt.Printf("  %-20s %s\n", "pm [--help | -h]", "Show this help menu")
}

func goToFileReference() error {
	for {
		if checkIfReferenceFileInCurrentDirectory() {
			currentDir, err := getCurrentWorkingDirectory()
			if err != nil {
				return fmt.Errorf("error getting current working directory: %v", err)
			}

			println("Current directory: " + currentDir)
			return nil
		}

		err := moveBack()
		if err != nil {
			return fmt.Errorf("error moving back: %v", err)
		}

		if isRootDirectory() {
			return fmt.Errorf("reference file not found")
		}
	}
}

func getCurrentWorkingDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: %v", err)
	}

	return dir, nil
}

func isRootDirectory() bool {
	currentDir, _ := os.Getwd()
	return currentDir == "/"
}

func moveBack() error {
	err := os.Chdir("..")
	if err != nil {
		return fmt.Errorf("error moving back: %v", err)
	}

	return nil
}

func checkIfReferenceFileInCurrentDirectory() bool {
	referenceFile, err := getReferenceFile()
	if err != nil {
		return false
	}

	_, err = os.Stat(referenceFile)
	return err == nil
}

func runCommandInReferenceFileDirectory() error {
	err := goToFileReference()
	if err != nil {
		return fmt.Errorf("error going to reference file directory: %v", err)
	}

	command := ""

	for _, arg := range os.Args[1:] {
		command += arg + " "
	}

	command = command[:len(command)-1]
	partsCommand := strings.Fields(command)

	if len(partsCommand) < 1 {
		return fmt.Errorf("command not found")
	}

	cmd := exec.Command(partsCommand[0], partsCommand[1:]...)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running command: %v", err)
	}

	println("Command executed successfully!")

	return nil
}