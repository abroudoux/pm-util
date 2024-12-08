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
	}
}

func getRootDir() {
	println("Root dir")
}