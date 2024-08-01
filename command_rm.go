package main

import "fmt"

func command_rm(command Command, active_directory **Directory) {
	if len(command.args) < 1 {
		fmt.Println("rm: No file or directory specified")
		return
	}

	for _, target := range command.args {
		if err := (*active_directory).RemoveChild(target); err != nil {
			fmt.Println(err.Error())
		}
	}
}
