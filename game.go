package main

// Added an add file method to the directory struct
// Need to do:
// Implement rm recursive delete for files and folders
// touch command
// Text editor?

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := name_input(reader)

	file_system := setup_file_system(name)

	active_directory := file_system

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

	var inputBuider strings.Builder
	var inputBuffer string
	var commandHistory []string
	historyIndex := 0

	fmt.Printf("%s(%s)%s $ ", Green, active_directory.Path(), Reset)

	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading rune", err.Error())
			continue
		}

		switch r {
		case '\r', '\n': // enter
			if inputBuider.Len() == 0 {
				fmt.Printf("\r\n%s(%s)%s $ ", Green, active_directory.Path(), Reset)
				continue
			}

			commandHistory = append(commandHistory, inputBuider.String())
			historyIndex = len(commandHistory)

			command, flags, args := parse_command(inputBuider.String())
			inputBuider.Reset()

			command_execute(Command{
				command: command,
				args:    args,
				flags:   flags,
			}, &active_directory)
			fmt.Printf("%s(%s)%s $ ", Green, active_directory.Path(), Reset)
		case 127: // backspace
			if inputBuider.Len() > 0 {
				input := inputBuider.String()
				inputBuider.Reset()

				if len(input) > 1 {
					inputBuider.WriteString(input[:len(input)-1])
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
						updatePrompt(inputBuider.String(), inputBuffer, *active_directory)
						inputBuider.Reset()
						inputBuider.WriteString(inputBuffer)
					}
				case 'B': // down
					if historyIndex < len(commandHistory)-1 {
						historyIndex++
						inputBuffer = commandHistory[historyIndex]
						fmt.Printf("%s(%s)%s $ %s", Green, active_directory.Path(), Reset, inputBuffer)
						updatePrompt(inputBuider.String(), inputBuffer, *active_directory)
						inputBuider.Reset()
						inputBuider.WriteString(inputBuffer)
					} else if historyIndex == len(commandHistory)-1 {
						historyIndex++
						inputBuffer = ""
						updatePrompt(inputBuider.String(), inputBuffer, *active_directory)
						inputBuider.Reset()
					}
				}
			}
		default: // normal characters
			inputBuider.WriteRune(r)
			fmt.Print(string(r))
		}
	}
}

func updatePrompt(oldInput, newInput string, active_directory Directory) {
	fmt.Print("\r" + strings.Repeat(" ", len(oldInput)+20) + "\r" + Green + "(" + active_directory.Path() + ")" + Reset + " $ " + newInput)
}

func name_input(reader *bufio.Reader) string {
	fmt.Print("Please enter your name: ")

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	return text
}
