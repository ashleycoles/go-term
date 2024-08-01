package main

import "fmt"

func command_ls(active_directory **Directory) {
	for _, dir := range (*active_directory).children {
		fmt.Println(dir.name)
	}

	for _, file := range (*active_directory).files {
		fmt.Println(file.FullName())
	}
}
