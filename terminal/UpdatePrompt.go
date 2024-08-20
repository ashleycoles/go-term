package terminal

import (
	"ash/go-term/filesystem"
	"fmt"
	"regexp"
)

func UpdatePrompt(input string, cursorPos int, activeDirectory filesystem.Directory) {
	fmt.Print("\033[2K\r")
	prompt := fmt.Sprintf("%s(%s)%s $ ", Green, activeDirectory.Path(), Reset)
	fmt.Print(prompt)
	fmt.Print(input)
	moveCursor(cursorPos + len(stripANSI(prompt)))
}

func moveCursor(pos int) {
	fmt.Printf("\033[%dG", pos+1)
}

func stripANSI(str string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	return re.ReplaceAllString(str, "")
}
