package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

func touch(command Command, activeDirectory **filesystem.Directory) {
	for _, newFileName := range command.Args {
		if name, extension, err := filesystem.GetFilenameParts(newFileName); err != nil {
			fmt.Printf("\r\ntouch: %s", err.Error())
		} else if err := (*activeDirectory).AddFile(name, extension, ""); err != nil {
			fmt.Printf("\r\ntouch: %s", err.Error())
		}
	}
	terminal.NewLine()
}
