package commands

import (
	"ash/go-term/filesystem"
	"ash/go-term/terminal"
	"fmt"
)

func touch(command Command, activeDirectory **filesystem.Directory) {
	for _, newFilePath := range command.Args {
		target, err := (*activeDirectory).Traverse(newFilePath)

		if err != nil {
			fmt.Printf("cd: %s", err.Error())
			terminal.NewLine()
			return
		}

		path := filesystem.ParsePath(newFilePath)

		if !path.HasFile() {
			terminal.NewLine()
			fmt.Printf("touch: %s is not a file", newFilePath)
		}

		if name, extension, err := filesystem.GetFilenameParts(*path.File); err != nil {
			terminal.NewLine()
			fmt.Printf("touch: %s", err.Error())
		} else if err := target.AddFile(name, extension, ""); err != nil {
			terminal.NewLine()
			fmt.Printf("touch: %s", err.Error())
		}
	}
	terminal.NewLine()
}
