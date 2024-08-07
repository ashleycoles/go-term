package main

import (
	"fmt"
)

func command_ls(active_directory **Directory) {
	for _, dir := range (*active_directory).children {
		fmt.Printf("\r\n%s%s%s", Blue, dir.name, Reset)
	}

	for _, file := range (*active_directory).files {
		fmt.Printf("\r\n%s", file.FullName())
	}
	new_line()
}
