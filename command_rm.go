package main

import "fmt"

func command_rm(args []string, active_directory **Directory) {
	if len(args) < 1 {
		fmt.Println("rm: No file or directory specified")
		return
	}

	for _, target := range args {
		if err := (*active_directory).RemoveChild(target); err != nil {
			fmt.Println(err.Error())
		}
	}
}
