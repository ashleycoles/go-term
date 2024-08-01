package main

import "fmt"

func command_cd(command Command, active_directory **Directory) {
	if len(command.args) != 1 {
		fmt.Println("Error, must specify a single directory")
		return
	}

	path := parse_file_path(command.args[0])

	for _, target := range path {

		if target == ".." {
			if (*active_directory).parent == nil {
				fmt.Println("Error, no parent directory")
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
			fmt.Printf("Error, directory %q not found in %s\n", target, (*active_directory).name)
			return
		}
	}
}
