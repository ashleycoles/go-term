package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

func fappend(command Command, activeDirectory *filesystem.Directory) {
	name := command.Args[0]
	content := command.Args[1]

	file, err := (*activeDirectory).GetFile(name)

	if err != nil {
		terminal.NewLine()
		fmt.Printf("append: %s", err.Error())
		terminal.NewLine()
		return
	}

	file.AppendContent(content)
	terminal.NewLine()
}
