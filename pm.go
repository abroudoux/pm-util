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

var previousWorkdingDirectoryPath string = "/tmp/pm_last_working_directory"
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

	err = goToFileReference()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
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

func setReferenceFile(referenceFile string) error {
	err := os.WriteFile(referenceFilePath, []byte(referenceFile), 0644)
	if err != nil {
		return fmt.Errorf("error setting reference file: %v", err)
	}

	println("Target file set to", referenceFile)
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

func flagMode() {
	flag := os.Args[1]

	if flag == "-" {
		err := goToPreviousWorkingDirectory()
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
	} else if flag == "--file" || flag == "-f" {
		if len(os.Args) < 3 {
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

func printHelpMenu() {
	println("Usage: pm [options] [command]")
	println("Options:")
	println("pm [--root | -r]    		Go to the root directory of the project")
	println("pm [command]     	 		Go to the directory of the project and run the command")
	println("pm [-]     		 		Go back to the previous directory")
	println("pm [--file | -f] [file]    Set the target file")
	println("pm [--help | -h]    		Show this help menu")

	os.Exit(0)
}

func goToFileReference() error {
	currentWorkingDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	err = setCurrentWorkingDirectory(currentWorkingDirectory)
	if err != nil {
		return fmt.Errorf("error setting current working directory: %v", err)
	}

	referenceFile, err := os.ReadFile(referenceFilePath)
	if err != nil {
		return fmt.Errorf("error reading reference file: %v", err)
	}

	println("Going to " + string(referenceFile))
	stop := false

	for !stop {
		err := moveBack()
		if err != nil {
			return fmt.Errorf("error moving back: %v", err)
		}

		if checkIfReferenceFileInCurrentDirectory() {
			stop = true
			referenceFile, err := getReferenceFile()
			if err != nil {
				return fmt.Errorf("error getting reference file: %v", err)
			}
			println("Reference file '%s' found", referenceFile)
		}
	}

	return nil
}

func moveBack() error {
	cmd := exec.Command("cd", "..")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error moving back: %v", err)
	}

	return nil
}

func goToPreviousWorkingDirectory() error {
	previousWorkingDirectory, err := os.ReadFile(previousWorkdingDirectoryPath)
	if err != nil {
		return fmt.Errorf("error reading last directory: %v", err)
	}

	if string(previousWorkingDirectory) == "" {
		return fmt.Errorf("last directory not found")
	}

	println("Going back to " + string(previousWorkingDirectory))
	cmd := exec.Command("cd", string(previousWorkingDirectory))
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error going back to last directory: %v", err)
	}

	return nil
}

func checkIfReferenceFileInCurrentDirectory() bool {
	info, err := os.Stat(referenceFilePath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func setCurrentWorkingDirectory(currentWorkingDirectory string) error {
	err := os.WriteFile(previousWorkdingDirectoryPath, []byte(currentWorkingDirectory), 0644)
	if err != nil {
		return fmt.Errorf("error setting current working directory: %v", err)
	}

	return nil
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

	println("Command '%s' executed", command)

	err = goToPreviousWorkingDirectory()
	if err != nil {
		return fmt.Errorf("error going back to previous directory: %v", err)
	}

	return nil
}