package main

import (
	"ash/text-game/commands"
	"ash/text-game/filesystem"
	"ash/text-game/output"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/term"
)

func main() {
	commands.Clear()
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
		fmt.Print("Restoring terminal mode")
		if err := term.Restore(fd, oldState); err != nil {
			fmt.Fprintln(os.Stderr, "Error restoring terminal mode: ", err.Error())
		}
		os.Stdout.Sync()
	}()

	var inputBuilder strings.Builder
	var inputBuffer string
	var commandHistory []string
	historyIndex := 0
	cursorPos := 0

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
				updateHorizontal("", 0, *activeDirectory)
				continue
			}

			commandHistory = append(commandHistory, inputBuilder.String())
			historyIndex = len(commandHistory)

			command, flags, args := commands.ParseCommand(inputBuilder.String())
			inputBuilder.Reset()
			cursorPos = 0

			commands.Execute(commands.Command{
				Command: command,
				Args:    args,
				Flags:   flags,
			}, &activeDirectory)
			updateHorizontal("", 0, *activeDirectory)
		case 127: // backspace
			if inputBuilder.Len() > 0 && cursorPos > 0 {
				input := inputBuilder.String()
				inputBuilder.Reset()

				if cursorPos > 1 {
					inputBuilder.WriteString(input[:cursorPos-1] + input[cursorPos:])
				} else {
					inputBuilder.WriteString(input[cursorPos:])
				}
				cursorPos--
				updateHorizontal(inputBuilder.String(), cursorPos, *activeDirectory)
			}
		case '\x03': // ctrl + c
			return
		case '\x1b': // Arrow keys
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
						updateHorizontal(inputBuffer, len(inputBuffer), *activeDirectory)
						inputBuilder.Reset()
						inputBuilder.WriteString(inputBuffer)
						cursorPos = len(inputBuffer)
					}
				case 'B': // down
					if historyIndex < len(commandHistory)-1 {
						historyIndex++
						inputBuffer = commandHistory[historyIndex]
						updateHorizontal(inputBuffer, len(inputBuffer), *activeDirectory)
						inputBuilder.Reset()
						inputBuilder.WriteString(inputBuffer)
						cursorPos = len(inputBuffer)
					} else if historyIndex == len(commandHistory)-1 {
						historyIndex++
						inputBuffer = ""
						updateHorizontal(inputBuffer, 0, *activeDirectory)
						inputBuilder.Reset()
					}
				case 'D': // left
					if cursorPos > 0 {
						cursorPos--
						updateHorizontal(inputBuilder.String(), cursorPos, *activeDirectory)
					}
				case 'C': // right
					if cursorPos < inputBuilder.Len() {
						cursorPos++
						updateHorizontal(inputBuilder.String(), cursorPos, *activeDirectory)
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
			updateHorizontal(inputBuilder.String(), cursorPos, *activeDirectory)
		}
	}
}

func updateHorizontal(input string, cursorPos int, activeDirectory filesystem.Directory) {
	fmt.Print("\033[2K\r")
	prompt := fmt.Sprintf("%s(%s)%s $ ", output.Green, activeDirectory.Path(), output.Reset)
	fmt.Print(prompt)
	fmt.Print(input)
	moveCursor(cursorPos + len(stripANSI(prompt)))
}

func moveCursor(pos int) {
	fmt.Printf("\033[%dG", pos+1)
}

func nameInput(reader *bufio.Reader) string {
	fmt.Print("Please enter your name: ")

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	return text
}

func stripANSI(str string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(str, "")
}
