package main

import (
	"fmt"
	"os"
	"os/exec"
)

var temp_file_path string = "/tmp/pm_last_dir"
var reference_file_path string = "/tmp/pm_reference_file"
var reference_file string = "package.json"

func main() {
	checkIfReferenceFileExists()

	if len(os.Args) > 1 {
		flagMode()
	}

	goToFileReference()
}

func checkIfReferenceFileExists() {
	referenceFileContent, err := os.ReadFile(reference_file_path)

	if err != nil {
		println("Reference file not found: " + reference_file)
		createReferenceFile()
	}

	if string(referenceFileContent) == "" {
		println("Reference file not found: " + reference_file)
		println("Use 'pm --file [file]' to set the reference file")
		os.Exit(0)
	}
}

func setReFerenceFile() {
	reference_file = os.Args[2]
	cmd := exec.Command("echo", reference_file + " > " + reference_file_path)
	err := cmd.Run()

	if err != nil {
		println("Error setting reference file: " + err.Error())
		os.Exit(1)
	}

	println("Target file set to " + reference_file)
	os.Exit(0)
}

func createReferenceFile() {
	if _, err := os.Stat(reference_file); os.IsExist(err) {
		println("Reference file already exists")
		os.Exit(1)
	}

	cmd := exec.Command("touch", reference_file_path)
	err := cmd.Run()

	if err != nil {
		println("Error creating reference file: " + err.Error())
		os.Exit(1)
	}

	println("Reference file created: " + reference_file)
	os.Exit(0)
}

func printReferenceFile() {
	referenceFileContent, err := os.ReadFile(reference_file_path)

	if err != nil {
		println("Reference file not found: " + reference_file)
		os.Exit(1)
	}

	println("Current reference file: " + string(referenceFileContent))
	os.Exit(0)
}

func flagMode() {
	flag := os.Args[1]

	if flag == "--file" || flag == "-f" {
		if len(os.Args) < 3 {
			printReferenceFile()
		} else {
			setReFerenceFile()
		}
	}

	// if flag == "-" {
	// 	goBackToPreviousDirectory()
	// } else if flag == "--root" || flag == "-r" {
	// 	goToFileReference()
	// } else if flag == "--file" || flag == "-f" {
	// 	setReFerenceFile()
	// } else if flag == "--help" || flag == "-h" {
	// 	printHelpMenu()
	// }
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
	_, err := os.Stat(reference_file)
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