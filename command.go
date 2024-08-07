package main

import (
	"fmt"
	"os"
	"strings"
)

type Command struct {
	command string
	args    []string
	flags   []string
}

func parse_command(command string) (string, []string, []string) {
	command_tokens := strings.Fields(command)

	var flags []string
	var args []string

	for _, token := range command_tokens[1:] {
		if strings.HasPrefix(token, "-") {
			flags = append(flags, token)
		} else {
			args = append(args, token)
		}
	}

	return command_tokens[0], flags, args
}

func command_execute(command Command, active_directory **Directory) {
	switch command.command {
	case "cd":
		command_cd(command, active_directory)
	case "mkdir":
		command_mkdir(command, active_directory)
	case "ls":
		command_ls(active_directory)
	case "pwd":
		command_pwd(active_directory)
	case "rm":
		command_rm(command, active_directory)
	case "cat":
		command_cat(command, active_directory)
	case "touch":
		command_touch(command, active_directory)
	case "clear":
		command_clear()
	default:
		fmt.Fprintln(os.Stdout, "\rCommand not found")

	}
}
