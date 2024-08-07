package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

func ls(activeDirectory **filesystem.Directory) {
	for _, dir := range (*activeDirectory).Children {
		fmt.Printf("\r\n%s%s%s", terminal.Blue, dir.Name, terminal.Reset)
	}

	for _, file := range (*activeDirectory).Files {
		fmt.Printf("\r\n%s", file.FullName())
	}
	terminal.NewLine()
}
