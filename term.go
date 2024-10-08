package main

import (
	"ash/go-term/commands"
	"ash/go-term/filesystem"
	"ash/go-term/terminal"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO: Better system for registering commands
// TODO: Update fappend to use traversal
// TODO: Update cat to use traversal
// TODO: HTTP request tool similar to curl but simpler

const (
	enter       = '\r'
	newline     = '\n'
	backspace   = 127
	ctrlC       = '\x03'
	escape      = '\x1b'
	arrowPrefix = '['
	upArrow     = 'A'
	downArrow   = 'B'
	leftArrow   = 'D'
	rightArrow  = 'C'
)

func main() {
	commands.Clear()
	reader := bufio.NewReader(os.Stdin)

	name := nameInput(reader)

	fileSystem := filesystem.Setup(name)

	activeDirectory := fileSystem

	oldState, fd := terminal.Setup()

	defer terminal.Restore(oldState, fd)

	var inputBuilder strings.Builder
	var inputBuffer string
	var commandHistory []string
	historyIndex := 0
	cursorPos := 0

	fmt.Printf("%s(%s)%s $ ", terminal.Green, activeDirectory.Path(), terminal.Reset)

	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading rune", err.Error())
			continue
		}

		switch r {
		case enter, newline:
			if inputBuilder.Len() == 0 {
				terminal.UpdatePrompt("", 0, *activeDirectory)
				continue
			}

			commandHistory = append(commandHistory, inputBuilder.String())
			historyIndex = len(commandHistory)

			command := commands.ParseCommand(inputBuilder.String())
			inputBuilder.Reset()
			cursorPos = 0

			commands.Execute(command, &activeDirectory)

			terminal.UpdatePrompt("", 0, *activeDirectory)
		case backspace:
			if inputBuilder.Len() > 0 && cursorPos > 0 {
				input := inputBuilder.String()
				inputBuilder.Reset()

				if cursorPos > 1 {
					inputBuilder.WriteString(input[:cursorPos-1] + input[cursorPos:])
				} else {
					inputBuilder.WriteString(input[cursorPos:])
				}
				cursorPos--
				terminal.UpdatePrompt(inputBuilder.String(), cursorPos, *activeDirectory)
			}
		case ctrlC:
			return
		case escape:
			next, _, err := reader.ReadRune()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading rune", err.Error())
				continue
			}

			if next == arrowPrefix {
				next, _, err := reader.ReadRune()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error reading rune", err.Error())
					continue
				}
				switch next {
				case upArrow:
					if historyIndex > 0 {
						historyIndex--
						inputBuffer = commandHistory[historyIndex]
						terminal.UpdatePrompt(inputBuffer, len(inputBuffer), *activeDirectory)
						inputBuilder.Reset()
						inputBuilder.WriteString(inputBuffer)
						cursorPos = len(inputBuffer)
					}
				case downArrow:
					if historyIndex < len(commandHistory)-1 {
						historyIndex++
						inputBuffer = commandHistory[historyIndex]
						terminal.UpdatePrompt(inputBuffer, len(inputBuffer), *activeDirectory)
						inputBuilder.Reset()
						inputBuilder.WriteString(inputBuffer)
						cursorPos = len(inputBuffer)
					} else if historyIndex == len(commandHistory)-1 {
						historyIndex++
						inputBuffer = ""
						terminal.UpdatePrompt(inputBuffer, 0, *activeDirectory)
						inputBuilder.Reset()
					}
				case leftArrow:
					if cursorPos > 0 {
						cursorPos--
						terminal.UpdatePrompt(inputBuilder.String(), cursorPos, *activeDirectory)
					}
				case rightArrow:
					if cursorPos < inputBuilder.Len() {
						cursorPos++
						terminal.UpdatePrompt(inputBuilder.String(), cursorPos, *activeDirectory)
					}
				}
			}
		default: // normal characters
			if cursorPos < inputBuilder.Len() {
				input := inputBuilder.String()
				inputBuilder.Reset()
				inputBuilder.WriteString(input[:cursorPos])
				inputBuilder.WriteRune(r)
				inputBuilder.WriteString(input[cursorPos:])
			} else {
				inputBuilder.WriteRune(r)
			}
			cursorPos++
			terminal.UpdatePrompt(inputBuilder.String(), cursorPos, *activeDirectory)
		}
	}
}

func nameInput(reader *bufio.Reader) string {
	fmt.Print("Please enter your name: ")

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	return text
}
