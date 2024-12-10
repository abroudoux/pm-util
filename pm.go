package main

import (
	"fmt"
	"os"
	"os/exec"
)

var temp_file_path string = "/tmp/pm_last_dir"
var config_file_path string = "/tmp/pm_config_file"
var target_file string = "package.json"

func main() {
	checkIfTargetFileExists()

	if len(os.Args) > 1 {
		flagMode()
	}

	goToFileReference()
}

func checkIfTargetFileExists() {
	if (target_file == "") {
		println("Target file is not set, Use --file to set it")
		os.Exit(1)
	}
}

func flagMode() {
	arg := os.Args[1]

	if arg == "-" {
		goBackToPreviousDirectory()
	} else if arg == "--root" || arg == "-r" {
		goToFileReference()
	} else if arg == "--file" || arg == "-f" {
		setTargetFile()
	} else if arg == "--help" || arg == "-h" {
		printHelpMenu()
	}
}

func goToFileReference() {
	stop := false
	setLastDirectory()

	for !stop {
		moveBack()

		if hasReferenceFileInCurrentDirectory() {
			stop = true
			println("Reference file found in current directory")
		}

		if isInRootDirectory() {
			stop = true
			println("Root directory found, no reference file found")
			goBackToPreviousDirectory()
		}
	}
}

func hasReferenceFileInCurrentDirectory() bool {
	_, err := os.Stat(target_file)
	return !os.IsNotExist(err)
}

func moveBack() {
	cmd := exec.Command("cd", "..")
	err := cmd.Run()

	if err != nil {
		println("Error: " + err.Error())
	}
}

func isInRootDirectory() bool {
	_, err := os.Stat(temp_file_path)
	return os.IsNotExist(err)
}

func setTargetFile() {
	if len(os.Args) < 3 {
		println("No file specified")
		os.Exit(1)
	}

	target_file = os.Args[2]
	println("Target file set to " + target_file)
}

func setLastDirectory() {
	cmd := exec.Command("echo", "pwd > " + temp_file_path)
	err := cmd.Run()

	if err != nil {
		println("Error: " + err.Error())
	}
}

func goBackToPreviousDirectory() {
	cmd := exec.Command("cd", "$(cat " + temp_file_path + ")")
	err := cmd.Run()

	if err != nil {
		println("Error: " + err.Error())
	}
}

func printHelpMenu() {
	fmt.Println("Usage: pm [options] [command]")
	fmt.Println("Options:")
	fmt.Println("pm [--root | -r]    		Go to the root directory of the project")
	fmt.Println("pm [command]     	 		Go to the directory of the project and run the command")
	fmt.Println("pm [-]     		 		Go back to the previous directory")
	fmt.Println("pm [--file | -f] [file]    Set the target file")
	fmt.Println("pm [--help | -h]    		Show this help menu")

	os.Exit(0)
}