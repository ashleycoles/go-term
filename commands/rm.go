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
		targetDirectory, err := (*activeDirectory).Traverse(target)
		if err != nil {
			fmt.Printf("\r\n%s", err.Error())
			return
		}

		parsedTarget := filesystem.ParseFilePath(target)
		last := parsedTarget[len(parsedTarget)-1]

		if last == "." || last == ".." {
			fmt.Print("\r\nrm: \".\" and \"..\" may not be removed")
			return
		}

		if strings.Contains(last, ".") {
			if err := targetDirectory.RemoveFile(last); err != nil {
				fmt.Printf("\r\n%s", err.Error())
			}
		} else {
			if err := targetDirectory.Parent.RemoveChild(last); err != nil {
				fmt.Printf("\r\n%s", err.Error())
			}
		}
	}

	terminal.NewLine()
}
