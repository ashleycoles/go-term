package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
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

		parsedTarget := filesystem.ParsePath(target)
		lastFolder := parsedTarget.GetLastFolder()

		if lastFolder == "." || lastFolder == ".." {
			fmt.Print("\r\nrm: \".\" and \"..\" may not be removed")
			return
		}

		if parsedTarget.HasFile() {
			if err := targetDirectory.RemoveFile(*parsedTarget.File); err != nil {
				fmt.Printf("\r\n%s", err.Error())
			}
		} else {
			if err := targetDirectory.Parent.RemoveChild(lastFolder); err != nil {
				fmt.Printf("\r\n%s", err.Error())
			}
		}
	}

	terminal.NewLine()
}
