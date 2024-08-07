package main

import "fmt"

// Todo figure out spaces in strings
// Remove quotes from strings
func command_append(command Command, active_directory *Directory) {
	name := command.args[0]
	content := command.args[1]

	file, err := (*active_directory).GetFile(name)

	if err != nil {
		fmt.Printf("\r\nappend: %s\r\n", err.Error())
	}

	file.AppendContent(content)
	new_line()
}
