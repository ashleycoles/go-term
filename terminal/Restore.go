package terminal

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func Restore(state *term.State, fd int) {
	fmt.Print("Restoring terminal mode")
	if err := term.Restore(fd, state); err != nil {
		fmt.Fprintln(os.Stderr, "Error restoring terminal mode: ", err.Error())
	}
	os.Stdout.Sync()
}
