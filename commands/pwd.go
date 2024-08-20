package commands

import (
	"ash/go-term/filesystem"
	"ash/go-term/terminal"
	"fmt"
)

func pwd(activeDirectory **filesystem.Directory) {
	terminal.NewLine()
	fmt.Printf("%s", (*activeDirectory).Path())
	terminal.NewLine()
}
