package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

// Todo:
// figure out spaces in strings
// Remove quotes from strings
func fappend(command Command, activeDirectory *filesystem.Directory) {
	name := command.Args[0]
	content := command.Args[1]

	file, err := (*activeDirectory).GetFile(name)

	if err != nil {
		fmt.Printf("\r\nappend: %s\r\n", err.Error())
	}

	file.AppendContent(content)
	terminal.NewLine()
}
