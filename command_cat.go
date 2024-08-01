package main

import "fmt"

func command_cat(args []string, active_directory **Directory) {
	if len(args) != 1 {
		fmt.Println("cat: Must specify one file")
	}

	file, err := (*active_directory).GetFile(args[0])

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(file.contents)
}
