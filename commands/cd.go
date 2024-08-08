package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

func cd(command Command, activeDirectory **filesystem.Directory) {
	terminal.NewLine()

	if len(command.Args) != 1 {
		fmt.Printf("Error, must specify a single directory\r\n")
		return
	}

	target, err := (*activeDirectory).Traverse(command.Args[0])

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		return
	}

	*activeDirectory = target
}
