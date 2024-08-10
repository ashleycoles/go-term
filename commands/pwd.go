package commands

import (
	"ash/text-game/filesystem"
	"ash/text-game/terminal"
	"fmt"
)

func pwd(activeDirectory **filesystem.Directory) {
	terminal.NewLine()
	fmt.Printf("%s", (*activeDirectory).Path())
	terminal.NewLine()
}
