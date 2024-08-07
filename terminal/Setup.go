package terminal

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func Setup() (*term.State, int) {
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)

	if err != nil {
		fmt.Fprint(os.Stderr, "Error setting raw mode: ", err.Error())
		os.Exit(1)
	}

	return oldState, fd
}
