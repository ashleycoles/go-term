package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
	"strings"
)

func rm(command Command, activeDirectory **filesystem.Directory) {
	if len(command.Args) < 1 {
		fmt.Print("\r\nrm: No file or directory specified\r\n")
		return
	}

	for _, target := range command.Args {
		if strings.Contains(target, ".") {
			if err := (*activeDirectory).RemoveFile(target); err != nil {
				fmt.Printf("\r\n%s", err.Error())

			}
		} else {
			if err := (*activeDirectory).RemoveChild(target); err != nil {
				fmt.Printf("\r\n%s", err.Error())
			}
		}
	}

	terminal.NewLine()
}
