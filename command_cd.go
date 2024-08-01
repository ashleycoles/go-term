package main

import "fmt"

func command_cd(args []string, active_directory **Directory) {
	if len(args) != 1 {
		fmt.Println("Error, must specify a single directory")
		return
	}

	target := args[0]

	if target == ".." {
		if (*active_directory).parent == nil {
			fmt.Println("Error, no parent directory")
			return
		}
		*active_directory = (*active_directory).parent
		return
	}

	for _, dir := range (*active_directory).children {
		if target == dir.name {
			*active_directory = dir
			return
		}
	}

	fmt.Println("Error, directory not found")
}
