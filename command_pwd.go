package main

import (
	"fmt"
)

func command_pwd(active_directory **Directory) {
	fmt.Printf("\r\n%s\r\n", (*active_directory).Path())
}
