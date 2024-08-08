package commands

import (
	"fmt"
	"strings"
	"unicode"

	"ash/text-game/filesystem"
)

type ValueFlag struct {
	Key   string
	Value string
}

type Command struct {
	Command    string
	Args       []string
	ValueFlags []ValueFlag
	Flags      []string
}

func (command *Command) ArgsCount() int {
	return len(command.Args)
}

func ParseCommand(command string) Command {
	commandTokens := tokeniseCommand(command)

	var flags []string
	var args []string
	var valueFlags []ValueFlag

	for _, token := range commandTokens[1:] {
		if strings.HasPrefix(token, "--") {
			trimmedFlag := strings.TrimLeft(token, "--")
			splitFlag := strings.Split(trimmedFlag, "=")

			valueFlags = append(valueFlags, ValueFlag{
				Key:   splitFlag[0],
				Value: splitFlag[1],
			})
		} else if strings.HasPrefix(token, "-") {
			trimmedFlag := strings.TrimLeft(token, "-")
			splitFlags := strings.Split(trimmedFlag, "")
			flags = append(flags, splitFlags...)
		} else {
			args = append(args, token)
		}
	}

	return Command{
		Command:    commandTokens[0],
		Args:       args,
		ValueFlags: valueFlags,
		Flags:      flags,
	}
}

func tokeniseCommand(command string) []string {
	var tokens []string
	var tokenBuilder strings.Builder
	inQuotes := false
	quoteChar := rune(0)

	for _, char := range command {
		switch {
		case char == '"' || char == '\'':
			if inQuotes {
				if char == quoteChar {
					inQuotes = false
					tokens = append(tokens, tokenBuilder.String())
					tokenBuilder.Reset()
				} else {
					tokenBuilder.WriteRune(char)
				}
			} else {
				inQuotes = true
				quoteChar = char
			}
		case unicode.IsSpace(char):
			if inQuotes {
				tokenBuilder.WriteRune(char)
			} else if tokenBuilder.Len() > 0 {
				tokens = append(tokens, tokenBuilder.String())
				tokenBuilder.Reset()
			}
		default:
			tokenBuilder.WriteRune(char)
		}
	}

	if tokenBuilder.Len() > 0 {
		tokens = append(tokens, tokenBuilder.String())
	}

	return tokens
}

func Execute(command Command, activeDirectory **filesystem.Directory) {
	switch command.Command {
	case "cd":
		cd(command, activeDirectory)
	case "mv":
		mv(command, activeDirectory)
	case "mkdir":
		mkdir(command, activeDirectory)
	case "ls":
		ls(command, activeDirectory)
	case "pwd":
		pwd(activeDirectory)
	case "rm":
		rm(command, activeDirectory)
	case "cat":
		cat(command, activeDirectory)
	case "touch":
		touch(command, activeDirectory)
	case "append":
		fappend(command, *activeDirectory)
	case "fetch":
		fetch(command)
	case "clear":
		Clear()
	default:
		fmt.Printf("\r\nCommand not found: %s\r\n", command.Command)
	}
}
