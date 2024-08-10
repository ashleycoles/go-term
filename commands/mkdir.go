package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

func mkdir(command Command, activeDirectory **filesystem.Directory) {
	for _, newDirName := range command.Args {
		if _, err := (*activeDirectory).AddChild(newDirName); err != nil {
			terminal.NewLine()
			fmt.Printf("%s", err.Error())
			terminal.NewLine()
			return
		}
	}

	terminal.NewLine()
}
