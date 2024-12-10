package main

import (
	"os"
)

var temp_file_path string = "/tmp/pm_last_dir"
var config_file_path string = "/tmp/pm_config_file"
var target_file string = "package.json"

func main() {
	checkIfTargetFileExists()

	if len(os.Args) > 1 {
		flagMode()
	}

	if hasReferenceFileInCurrentDirectory() {
		println("Reference file found")
	} else {
		println("Reference file not found")
	}
}

func checkIfTargetFileExists() {
	if (target_file == "") {
		println("Target file is not set, Use --file to set it")
		os.Exit(1)
	}
}

func flagMode() {
	arg := os.Args[1]

	if arg == "--root" || arg == "-r" {
		getRootDir()
	} else if arg == "--file" || arg == "-f" {
		setTargetFile()
	}
}

func getRootDir() {
	println("Root dir")
}

func hasReferenceFileInCurrentDirectory() bool {
	_, err := os.Stat(target_file)
	return !os.IsNotExist(err)
}

func setTargetFile() {
	if len(os.Args) < 3 {
		println("No file specified")
		os.Exit(1)
	}

	target_file = os.Args[2]
	println("Target file set to " + target_file)
}