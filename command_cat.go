package main

import "fmt"

func command_cat(command Command, active_directory **Directory) {
	if len(command.args) != 1 {
		fmt.Println("cat: Must specify one file")
	}

	file, err := (*active_directory).GetFile(command.args[0])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(file.contents)
}
