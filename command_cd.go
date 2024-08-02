package main

import (
	"fmt"
)

func command_cd(command Command, active_directory **Directory) {
	fmt.Printf("\r\n")

	if len(command.args) != 1 {
		fmt.Printf("Error, must specify a single directory\r\n")
		return
	}

	path := parse_file_path(command.args[0])

	for _, target := range path {

		if target == ".." {
			if (*active_directory).parent == nil {
				fmt.Printf("Error, no parent directory\r\n")
				return
			}
			*active_directory = (*active_directory).parent
			continue
		}

		found := false
		for _, dir := range (*active_directory).children {
			if dir.name == target {
				*active_directory = dir
				found = true
				break
			}
		}

		if !found {
			fmt.Printf("Error, directory %q not found in %s\r\n", target, (*active_directory).name)
			return
		}
	}
}
