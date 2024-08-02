package main

import "fmt"

func command_cat(command Command, active_directory **Directory) {
	if len(command.args) != 1 {
		fmt.Print("\r\ncat: Must specify one file")
	}

	file, err := (*active_directory).GetFile(command.args[0])

	if err != nil {
		fmt.Printf("\r\n%s", err.Error())
		return
	}

	fmt.Printf("\r\n%s\r\n", file.contents)
}
