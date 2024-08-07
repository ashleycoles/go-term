package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Command struct {
	command string
	args    []string
	flags   []string
}

func parse_command(command string) (string, []string, []string) {
	command_tokens := tokenise_command(command)

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

func tokenise_command(command string) []string {
	var tokens []string
	var token_builder strings.Builder
	in_quotes := false
	quote_char := rune(0)

	for _, char := range command {
		switch {
		case char == '"' || char == '\'':
			if in_quotes {
				if char == quote_char {
					in_quotes = false
					tokens = append(tokens, token_builder.String())
					token_builder.Reset()
				} else {
					token_builder.WriteRune(char)
				}
			} else {
				in_quotes = true
				quote_char = char
			}
		case unicode.IsSpace(char):
			if in_quotes {
				token_builder.WriteRune(char)
			} else if token_builder.Len() > 0 {
				tokens = append(tokens, token_builder.String())
				token_builder.Reset()
			}
		default:
			token_builder.WriteRune(char)
		}
	}

	if token_builder.Len() > 0 {
		tokens = append(tokens, token_builder.String())
	}

	return tokens
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
	case "append":
		command_append(command, *active_directory)
	case "clear":
		command_clear()
	default:
		fmt.Fprintln(os.Stdout, "\rCommand not found")
	}
}
