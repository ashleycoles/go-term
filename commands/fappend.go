package commands

import (
	"ash/go-term/filesystem"
	"ash/go-term/terminal"
	"fmt"
)

func fappend(command Command, activeDirectory *filesystem.Directory) {
	if command.ArgsCount() != 2 {
		terminal.NewLine()
		fmt.Print("append: must specify a filepath and content")
		terminal.NewLine()
		return
	}

	name := command.Args[0]
	content := command.Args[1]

	path := filesystem.ParsePath(name)

	if !path.HasFile() {
		terminal.NewLine()
		fmt.Printf("append: %s is not a file", name)
		terminal.NewLine()
	}

	targetDirectory, err := (*activeDirectory).Traverse(name)

	if err != nil {
		terminal.NewLine()
		fmt.Printf("append: %s", err.Error())
		terminal.NewLine()
		return
	}

	file, err := targetDirectory.GetFile(*path.File)

	if err != nil {
		terminal.NewLine()
		fmt.Printf("append: %s", err.Error())
		terminal.NewLine()
		return
	}

	file.AppendContent(content)
	terminal.NewLine()
}
