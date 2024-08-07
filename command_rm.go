package main

import (
	"fmt"
	"strings"
)

func command_rm(command Command, active_directory **Directory) {
	if len(command.args) < 1 {
		fmt.Print("\r\nrm: No file or directory specified\r\n")
		return
	}

	for _, target := range command.args {
		if strings.Contains(target, ".") {
			if err := (*active_directory).RemoveFile(target); err != nil {
				fmt.Printf("\r\n%s", err.Error())

			}
		} else {
			if err := (*active_directory).RemoveChild(target); err != nil {
				fmt.Printf("\r\n%s", err.Error())
			}
		}
	}

	new_line()
}
