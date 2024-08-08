package filesystem

import (
	"fmt"
	"strings"
)

type Directory struct {
	Name     string
	Parent   *Directory
	Children []*Directory
	Files    []*File
}

func (directory *Directory) AddChild(name string) (*Directory, error) {
	if strings.Contains(name, ".") {
		return nil, fmt.Errorf("directory names cannot include fullstops (.)")
	}

	for _, child := range directory.Children {
		if child.Name == name {
			return nil, fmt.Errorf("directory %q already exists", name)
		}
	}
	child := &Directory{
		Name:   name,
		Parent: directory,
	}

	directory.Children = append(directory.Children, child)
	return child, nil
}

func (directory *Directory) AddExistingChild(child *Directory) error {

	for _, directoryChild := range directory.Children {
		if directoryChild.Name == child.Name {
			return fmt.Errorf("directory %q already exists", child.Name)
		}
	}

	directory.Children = append(directory.Children, child)

	return nil
}

func (directory *Directory) RemoveChild(name string) error {
	index := -1

	for i, child := range directory.Children {
		if child.Name == name {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("directory %q not found", name)
	}

	directory.Children = append(directory.Children[:index], directory.Children[index+1:]...)
	return nil
}

func (directory *Directory) AddFile(name string, extension string, contents string) error {
	for _, file := range directory.Files {
		if file.Name == name && file.Extension == extension {
			return fmt.Errorf("file %s already exists", name)
		}
	}

	file := &File{
		Name:      name,
		Extension: extension,
		Contents:  contents,
	}

	directory.Files = append(directory.Files, file)
	return nil
}

func (directory *Directory) RemoveFile(name string) error {
	index := -1

	for i, file := range directory.Files {
		if file.FullName() == name {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("file %q not found", name)
	}

	directory.Files = append(directory.Files[:index], directory.Files[index+1:]...)
	return nil
}

func (directory *Directory) GetFile(name string) (*File, error) {
	for _, file := range directory.Files {
		if file.FullName() == name {
			return file, nil
		}
	}

	return nil, fmt.Errorf("file %s does not exist", name)
}

func (directory *Directory) Path() string {
	if directory.Parent == nil {
		return directory.Name
	}
	return directory.Parent.Path() + "/" + directory.Name
}

func (directory *Directory) FileExists(name string) bool {
	for _, file := range directory.Files {
		if name == file.FullName() {
			return true
		}
	}
	return false
}

func (directory *Directory) Traverse(path string) (*Directory, error) {
	parsedPath := ParseFilePath(path)

	// Remove the last item if it's a file
	last := parsedPath[len(parsedPath)-1]
	if last != ".." && strings.Contains(last, ".") {
		parsedPath = parsedPath[:len(parsedPath)-1]
	}

	var tempDirectory = directory

	for _, target := range parsedPath {
		if target == ".." {
			if tempDirectory.Parent == nil {
				return nil, fmt.Errorf("no Parent directory: %s", tempDirectory.Name)
			}
			tempDirectory = tempDirectory.Parent
			continue
		}

		found := false
		for _, dir := range tempDirectory.Children {
			if dir.Name == target {
				tempDirectory = dir
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("no such file directory: %s", path)
		}
	}

	return tempDirectory, nil
}
