package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
	"strings"
)

func mv(command Command, activeDirectory **filesystem.Directory) {
	source := command.Args[0]
	target := command.Args[1]

	sourceDirectory, err := (*activeDirectory).Traverse(source)
	if err != nil {
		fmt.Printf("\r\nmv: %s", err.Error())
		return
	}
	targetDirectory, err := (*activeDirectory).Traverse(target)
	if err != nil {
		fmt.Printf("\r\nmv: %s", err.Error())
		return
	}

	sourcePath := filesystem.ParseFilePath(source)
	toMove := sourcePath[len(sourcePath)-1]
	isFile := toMove != ".." && strings.Contains(toMove, ".")

	if isFile {
		file, err := sourceDirectory.GetFile(toMove)
		if err != nil {
			fmt.Printf("\r\nmv: %s", err.Error())
			return
		}

		err = targetDirectory.AddFile(file.Name, file.Extension, file.Contents)
		if err != nil {
			fmt.Printf("\r\nmv: %s", err.Error())
			return
		}

		err = sourceDirectory.RemoveFile(file.FullName())
		if err != nil {
			fmt.Printf("\r\nmv: %s", err.Error())
			return
		}
	} else {
		sourceParent := sourceDirectory.Parent
		sourceIsRoot := sourceParent == nil
		if sourceIsRoot {
			fmt.Printf("\r\nmv: cannot move: %s", sourceDirectory.Name)
			return
		}

		err := targetDirectory.AddExistingChild(sourceDirectory)
		if err != nil {
			fmt.Printf("\r\nmv: %s", err.Error())
			return
		}

		sourceDirectory.Parent = targetDirectory

		err = sourceParent.RemoveChild(sourceDirectory.Name)
		if err != nil {
			fmt.Printf("\r\nmv: %s", err.Error())
			return
		}
	}
	terminal.NewLine()
}
