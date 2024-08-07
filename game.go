package main

// Added an add file method to the directory struct
// Need to do:
// Implement rm recursive delete for files and folders
// touch command
// Text editor?

import (
	"ash/text-game/commands"
	"ash/text-game/filesystem"
	"ash/text-game/output"
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := nameInput(reader)

	fileSystem := filesystem.SetupFilesystem(name)

	activeDirectory := fileSystem

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)

	if err != nil {
		fmt.Fprint(os.Stderr, "Error setting raw mode: ", err.Error())
		os.Exit(1)
	}

	defer func() {
		fmt.Fprintln(os.Stdout, "Restoring terminal mode")
		if err := term.Restore(fd, oldState); err != nil {
			fmt.Fprintln(os.Stderr, "Error restoring terminal mode: ", err.Error())
		}
		os.Stdout.Sync()
	}()

	var inputBuilder strings.Builder
	var inputBuffer string
	var commandHistory []string
	historyIndex := 0

	fmt.Printf("%s(%s)%s $ ", output.Green, activeDirectory.Path(), output.Reset)

	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading rune", err.Error())
			continue
		}

		switch r {
		case '\r', '\n': // enter
			if inputBuilder.Len() == 0 {
				fmt.Printf("\r\n%s(%s)%s $ ", output.Green, activeDirectory.Path(), output.Reset)
				continue
			}

			commandHistory = append(commandHistory, inputBuilder.String())
			historyIndex = len(commandHistory)

			command, flags, args := commands.ParseCommand(inputBuilder.String())
			inputBuilder.Reset()

			commands.Execute(commands.Command{
				Command: command,
				Args:    args,
				Flags:   flags,
			}, &activeDirectory)
			fmt.Printf("%s(%s)%s $ ", output.Green, activeDirectory.Path(), output.Reset)
		case 127: // backspace
			if inputBuilder.Len() > 0 {
				input := inputBuilder.String()
				inputBuilder.Reset()

				if len(input) > 1 {
					inputBuilder.WriteString(input[:len(input)-1])
				}

				fmt.Print("\b \b")
			}
		case '\x03': // ctrl + c
			return
		case '\x1b':
			next, _, err := reader.ReadRune()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading rune", err.Error())
				continue
			}

			if next == '[' {
				next, _, err := reader.ReadRune()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error reading rune", err.Error())
					continue
				}

				switch next {
				case 'A': // up
					if historyIndex > 0 {
						historyIndex--
						inputBuffer = commandHistory[historyIndex]
						updatePrompt(inputBuilder.String(), inputBuffer, *activeDirectory)
						inputBuilder.Reset()
						inputBuilder.WriteString(inputBuffer)
					}
				case 'B': // down
					if historyIndex < len(commandHistory)-1 {
						historyIndex++
						inputBuffer = commandHistory[historyIndex]
						fmt.Printf("%s(%s)%s $ %s", output.Green, activeDirectory.Path(), output.Reset, inputBuffer)
						updatePrompt(inputBuilder.String(), inputBuffer, *activeDirectory)
						inputBuilder.Reset()
						inputBuilder.WriteString(inputBuffer)
					} else if historyIndex == len(commandHistory)-1 {
						historyIndex++
						inputBuffer = ""
						updatePrompt(inputBuilder.String(), inputBuffer, *activeDirectory)
						inputBuilder.Reset()
					}
				}
			}
		default: // normal characters
			inputBuilder.WriteRune(r)
			fmt.Print(string(r))
		}
	}
}

func updatePrompt(oldInput, newInput string, activeDirectory filesystem.Directory) {
	fmt.Print("\r" + strings.Repeat(" ", len(oldInput)+20) + "\r" + output.Green + "(" + activeDirectory.Path() + ")" + output.Reset + " $ " + newInput)
}

func nameInput(reader *bufio.Reader) string {
	fmt.Print("Please enter your name: ")

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	return text
}
