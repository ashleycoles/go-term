package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/output"
	"fmt"
)

func ls(activeDirectory **filesystem.Directory) {
	for _, dir := range (*activeDirectory).Children {
		fmt.Printf("\r\n%s%s%s", output.Blue, dir.Name, output.Reset)
	}

	for _, file := range (*activeDirectory).Files {
		fmt.Printf("\r\n%s", file.FullName())
	}
	output.NewLine()
}
