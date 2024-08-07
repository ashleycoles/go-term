package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/output"
	"fmt"
)

func mkdir(command Command, activeDirectory **filesystem.Directory) {
	for _, newDirName := range command.Args {
		if _, err := (*activeDirectory).AddChild(newDirName); err != nil {
			fmt.Printf("\r\n%s\r\n", err.Error())
			return
		}
	}

	output.NewLine()
}
