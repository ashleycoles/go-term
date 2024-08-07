package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/output"
	"fmt"
	"strings"
)

func touch(command Command, activeDirectory **filesystem.Directory) {
	for _, newFileName := range command.Args {
		fileParts := strings.Split(newFileName, ".")

		err := (*activeDirectory).AddFile(fileParts[0], fileParts[1], "")

		if err != nil {
			fmt.Printf("\r\n%s\r\n", err.Error())
		}
	}
	output.NewLine()
}
