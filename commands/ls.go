package commands

import (
	"ash/go-term/filesystem"
	"ash/go-term/terminal"
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
				terminal.NewLine()
				return
			}

			if argsCount > 1 {
				terminal.NewLine()
				fmt.Printf("%s:", directory.Name)
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
		terminal.NewLine()
		fmt.Printf("%s%s%s", terminal.Blue, dir.Name, terminal.Reset)
	}

	for _, file := range (*directory).Files {
		terminal.NewLine()
		fmt.Printf("%s", file.FullName())
	}
}
