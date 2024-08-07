package commands

import (
	"ash/text-game/filesystem"
	"fmt"
)

func pwd(activeDirectory **filesystem.Directory) {
	fmt.Printf("\r\n%s\r\n", (*activeDirectory).Path())
}
