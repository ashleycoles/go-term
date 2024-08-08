package filesystem

import "strings"

type Path struct {
	Folders []string
	File    *string
}

func (path *Path) GetLastFolder() string {
	return path.Folders[len(path.Folders)-1]
}

func (path *Path) HasFile() bool {
	return path.File != nil
}

func ParsePath(path string) Path {
	splitPath := strings.Split(strings.TrimRight(path, "/"), "/")
	pathLen := len(splitPath)
	last := splitPath[pathLen-1]
	lastIsFile := last != "." && last != ".." && strings.Contains(last, ".")
	if lastIsFile {
		return Path{
			Folders: splitPath[:pathLen-1],
			File:    &last,
		}
	}

	return Path{
		Folders: splitPath,
		File:    nil,
	}
}
