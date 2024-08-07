package filesystem

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
