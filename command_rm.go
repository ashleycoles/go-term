package main

import (
	"fmt"
	"strings"
)

func command_rm(command Command, active_directory **Directory) {
	if len(command.args) < 1 {
		fmt.Println("rm: No file or directory specified")
		return
	}

	for _, target := range command.args {
		if strings.Contains(target, ".") {
			if err := (*active_directory).RemoveFile(target); err != nil {
				fmt.Println(err.Error())

			}
		} else {
			if err := (*active_directory).RemoveChild(target); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
