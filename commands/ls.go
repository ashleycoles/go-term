package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

func ls(command Command, activeDirectory **filesystem.Directory) {
	argsCount := command.ArgsCount()

	if argsCount == 0 {
		displayContents(activeDirectory)
	} else {
		for i, arg := range command.Args {
			directory, err := (*activeDirectory).Traverse(arg)
			if err != nil {
				fmt.Printf("ls: %s", err.Error())
			}

			if argsCount > 1 {
				fmt.Printf("\r\n%s:", directory.Name)
			}

			displayContents(&directory)

			if argsCount > 1 && i < argsCount-1 {
				terminal.NewLine()
			}
		}
	}

	terminal.NewLine()
}

func displayContents(directory **filesystem.Directory) {
	for _, dir := range (*directory).Children {
		fmt.Printf("\r\n%s%s%s", terminal.Blue, dir.Name, terminal.Reset)
	}

	for _, file := range (*directory).Files {
		fmt.Printf("\r\n%s", file.FullName())
	}
}
