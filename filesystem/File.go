package filesystem

import (
	"fmt"
	"strings"
)

type File struct {
	Name      string
	Extension string
	Contents  string
}

func (file *File) FullName() string {
	return file.Name + "." + file.Extension
}

func (file *File) AppendContent(content string) {
	file.Contents += "\r\n" + content
}

func IsValidFilename(name string) bool {
	return name != "." && name != ".." && strings.Contains(name, ".")
}

func GetFilenameParts(name string) (string, string, error) {
	parts := strings.SplitN(name, ".", 2)

	if len(parts) == 1 {
		return "", "", fmt.Errorf("%s is not a valid filename", name)
	}
	return parts[0], parts[1], nil
}
