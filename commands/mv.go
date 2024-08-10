package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

func mv(command Command, activeDirectory **filesystem.Directory) {
	// TODO: Check for 2 args
	source := command.Args[0]
	target := command.Args[1]

	terminal.NewLine()

	sourceDirectory, err := (*activeDirectory).Traverse(source)
	if err != nil {
		fmt.Printf("mv: %s", err.Error())
		terminal.NewLine()
		return
	}
	targetDirectory, err := (*activeDirectory).Traverse(target)
	if err != nil {
		fmt.Printf("mv: %s", err.Error())
		terminal.NewLine()
		return
	}

	sourcePath := filesystem.ParsePath(source)
	toMove := sourcePath.GetLastFolder()

	if sourcePath.HasFile() {
		file, err := sourceDirectory.GetFile(toMove)
		if err != nil {
			fmt.Printf("mv: %s", err.Error())
			terminal.NewLine()
			return
		}

		err = targetDirectory.AddFile(file.Name, file.Extension, file.Contents)
		if err != nil {
			fmt.Printf("mv: %s", err.Error())
			terminal.NewLine()
			return
		}

		err = sourceDirectory.RemoveFile(file.FullName())
		if err != nil {
			fmt.Printf("mv: %s", err.Error())
			terminal.NewLine()
			return
		}
	} else {
		sourceParent := sourceDirectory.Parent
		sourceIsRoot := sourceParent == nil
		if sourceIsRoot {
			fmt.Printf("mv: cannot move: %s", sourceDirectory.Name)
			terminal.NewLine()
			return
		}

		err := targetDirectory.AddExistingChild(sourceDirectory)
		if err != nil {
			fmt.Printf("mv: %s", err.Error())
			terminal.NewLine()
			return
		}

		sourceDirectory.Parent = targetDirectory

		err = sourceParent.RemoveChild(sourceDirectory.Name)
		if err != nil {
			fmt.Printf("mv: %s", err.Error())
			terminal.NewLine()
			return
		}
	}
	terminal.NewLine()
}
