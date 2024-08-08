package filesystem

func Setup(user string) *Directory {
	root := &Directory{Name: "root"}
	users, _ := root.AddChild("users")
	root.AddChild("etc")

	home, _ := users.AddChild(user)
	docs, _ := home.AddChild("Documents")
	docs.AddFile("test", "txt", "Hello, world!\r\nFoo\r\nBar")

	return root
}
