package filesystem

import (
	"strings"
)

func Setup(user string) *Directory {
	root := &Directory{Name: "root"}
	users, _ := root.AddChild("users")
	root.AddChild("etc")

	home, _ := users.AddChild(user)
	docs, _ := home.AddChild("Documents")
	docs.AddFile("test", "txt", "Hello, world!\r\nFoo\r\nBar")

	return root
}

// TODO: Move into own file and create a path struct
func ParseFilePath(path string) []string {
	return strings.Split(strings.TrimRight(path, "/"), "/")
}
