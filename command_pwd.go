package main

import "fmt"

func command_pwd(active_directory **Directory) {
	fmt.Println((*active_directory).Path())
}
