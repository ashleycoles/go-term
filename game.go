package main

// Added an add file method to the directory struct
// Need to do:
// Update rm to include files
// Implement rm recursive delete for files and folders
// touch command
// Text editor?

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := name_input(reader)

	file_system := setup_file_system(name)

	active_directory := file_system

	for {
		command := command_prompt(reader)

		command_execute(command, &active_directory)
	}
}

func name_input(reader *bufio.Reader) string {
	fmt.Print("Please enter your name: ")

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	return text
}
