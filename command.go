package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Command struct {
	command string
	args    []string
}

func command_prompt(reader *bufio.Reader) Command {
	print("$ ")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	command_tokens := strings.Fields(input)

	command := Command{
		command: command_tokens[0],
		args:    command_tokens[1:],
	}

	return command
}

func command_execute(command Command, active_directory **Directory) {
	switch command.command {
	case "cd":
		command_cd(command.args, active_directory)
	case "mkdir":
		command_mkdir(command.args, active_directory)
	case "ls":
		command_ls(active_directory)
	case "pwd":
		command_pwd(active_directory)
	case "rm":
		command_rm(command.args, active_directory)
	case "cat":
		command_cat(command.args, active_directory)
	default:
		fmt.Println("Command not found")
	}
}
