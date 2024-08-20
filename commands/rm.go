package commands

import (
	"ash/go-term/filesystem"
	"ash/go-term/terminal"
	"fmt"
)

func rm(command Command, activeDirectory **filesystem.Directory) {
	terminal.NewLine()

	if command.ArgsCount() < 1 {
		fmt.Print("rm: No file or directory specified")
		terminal.NewLine()
		return
	}

	for _, target := range command.Args {
		targetDirectory, err := (*activeDirectory).Traverse(target)
		if err != nil {
			fmt.Printf("%s", err.Error())
			terminal.NewLine()
			return
		}

		parsedTarget := filesystem.ParsePath(target)
		lastFolder := parsedTarget.GetLastFolder()

		if lastFolder == "." || lastFolder == ".." {
			fmt.Print("rm: \".\" and \"..\" may not be removed")
			terminal.NewLine()
			return
		}

		if parsedTarget.HasFile() {
			if err := targetDirectory.RemoveFile(*parsedTarget.File); err != nil {
				fmt.Printf("rm: %s", err.Error())
				terminal.NewLine()
				return
			}
		} else {
			if err := targetDirectory.Parent.RemoveChild(lastFolder); err != nil {
				fmt.Printf("rm: %s", err.Error())
				terminal.NewLine()
				return
			}
		}
	}

	terminal.NewLine()
}
