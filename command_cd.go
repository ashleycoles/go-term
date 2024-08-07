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

	target, err := (*active_directory).Traverse(command.args[0])

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		return
	}

	*active_directory = target
}
