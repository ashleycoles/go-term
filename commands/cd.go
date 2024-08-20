package commands

import (
	"ash/go-term/filesystem"
	"ash/go-term/terminal"
	"fmt"
)

func cd(command Command, activeDirectory **filesystem.Directory) {
	terminal.NewLine()

	if command.ArgsCount() != 1 {
		fmt.Printf("cd: must specify a single directory")
		terminal.NewLine()
		return
	}

	target, err := (*activeDirectory).Traverse(command.Args[0])

	if err != nil {
		fmt.Printf("cd: %s", err.Error())
		terminal.NewLine()
		return
	}

	*activeDirectory = target
}
