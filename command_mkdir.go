package main

import "fmt"

func command_mkdir(args []string, active_directory **Directory) {
	for _, new_dir_name := range args {
		if _, err := (*active_directory).AddChild(new_dir_name); err != nil {
			fmt.Println(err.Error())
		}
	}
}
