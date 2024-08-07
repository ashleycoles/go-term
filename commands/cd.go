package commands

import (
	"ash/text-game/filesystem"
	"fmt"
)

func cd(command Command, activeDirectory **filesystem.Directory) {
	fmt.Printf("\r\n")

	if len(command.Args) != 1 {
		fmt.Printf("Error, must specify a single directory\r\n")
		return
	}

	target, err := (*activeDirectory).Traverse(command.Args[0])

	if err != nil {
		fmt.Printf("%s\r\n", err.Error())
		return
	}

	*activeDirectory = target
}
