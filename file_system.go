package main

import (
	"fmt"
	"strings"
)

type File struct {
	name      string
	extension string
	contents  string
}

func (file *File) FullName() string {
	return file.name + "." + file.extension
}

type Directory struct {
	name     string
	parent   *Directory
	children []*Directory
	files    []*File
}

func (directory *Directory) AddChild(name string) (*Directory, error) {
	if strings.Contains(name, ".") {
		return nil, fmt.Errorf("Directory names cannot include fullstops (.)")
	}

	for _, child := range directory.children {
		if child.name == name {
			return nil, fmt.Errorf("Directory %q already exists", name)
		}
	}
	child := &Directory{
		name:   name,
		parent: directory,
	}

	directory.children = append(directory.children, child)
	return child, nil
}

func (directory *Directory) RemoveChild(name string) error {
	index := -1

	for i, child := range directory.children {
		if child.name == name {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("Directory %q not found", name)
	}

	directory.children = append(directory.children[:index], directory.children[index+1:]...)
	return nil
}

func (directory *Directory) AddFile(name string, extension string, contents string) error {
	for _, file := range directory.files {
		if file.name == name && file.extension == extension {
			return fmt.Errorf("File %s already exists", name)
		}
	}

	file := &File{
		name:      name,
		extension: extension,
		contents:  contents,
	}

	directory.files = append(directory.files, file)
	return nil
}

func (directory *Directory) RemoveFile(name string) error {
	index := -1

	for i, file := range directory.files {
		if file.FullName() == name {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("File %q not found", name)
	}

	directory.files = append(directory.files[:index], directory.files[index+1:]...)
	return nil
}

func (directory *Directory) GetFile(name string) (*File, error) {
	for _, file := range directory.files {
		if file.FullName() == name {
			return file, nil
		}
	}

	return nil, fmt.Errorf("File %q not found", name)
}

func (directory *Directory) Path() string {
	if directory.parent == nil {
		return directory.name
	}
	return directory.parent.Path() + "/" + directory.name
}

func (directory *Directory) Traverse(path string) (*Directory, error) {
	parsed_path := parse_file_path(path)

	// Remove the last item if it's a file
	last := parsed_path[len(parsed_path)-1]
	if last != ".." && strings.Contains(last, ".") {
		return nil, fmt.Errorf("not a directory: %s", last)
	}

	var temp_directory = directory

	for _, target := range parsed_path {
		if target == ".." {
			if temp_directory.parent == nil {
				return nil, fmt.Errorf("no parent directory: %s", temp_directory.name)
			}
			temp_directory = directory.parent
			continue
		}

		found := false
		for _, dir := range temp_directory.children {
			if dir.name == target {
				temp_directory = dir
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("no such file directory: %s", path)
		}
	}

	return temp_directory, nil
}

func setup_file_system(user string) *Directory {
	root := &Directory{name: "root"}
	users, _ := root.AddChild("users")
	root.AddChild("etc")

	home, _ := users.AddChild(user)
	docs, _ := home.AddChild("Documents")
	docs.AddFile("test", "txt", "Hello, world!\r\nFoo\r\nBar")

	return root
}

func parse_file_path(path string) []string {
	return strings.Split(strings.TrimRight(path, "/"), "/")
}
