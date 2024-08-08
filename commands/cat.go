package commands

import (
	"ash/text-game/filesystem"
	"fmt"
)

func cat(command Command, activeDirectory **filesystem.Directory) {
	if command.ArgsCount() != 1 {
		fmt.Print("\r\ncat: Must specify one file")
	}

	file, err := (*activeDirectory).GetFile(command.Args[0])

	if err != nil {
		fmt.Printf("\r\n%s", err.Error())
		return
	}

	fmt.Printf("\r\n%s\r\n", file.Contents)
}
