package main

import (
	"fmt"
	"strings"
)

func command_touch(command Command, active_directory **Directory) {
	for _, new_file_name := range command.args {
		file_parts := strings.Split(new_file_name, ".")

		err := (*active_directory).AddFile(file_parts[0], file_parts[1], "")

		if err != nil {
			fmt.Printf("\r\n%s\r\n", err.Error())
		}
	}
	new_line()
}
