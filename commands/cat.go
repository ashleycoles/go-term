package commands

import (
	"ash/go-term/filesystem"
	"ash/go-term/terminal"
	"fmt"
)

func cat(command Command, activeDirectory **filesystem.Directory) {
	terminal.NewLine()

	if command.ArgsCount() != 1 {
		fmt.Print("cat: must specify one file")
		terminal.NewLine()
		return
	}

	file, err := (*activeDirectory).GetFile(command.Args[0])

	if err != nil {
		fmt.Printf("cat: %s", err.Error())
		terminal.NewLine()
		return
	}

	fmt.Print(file.Contents)
	terminal.NewLine()
}
