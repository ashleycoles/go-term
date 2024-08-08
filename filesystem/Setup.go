package filesystem

func Setup(user string) *Directory {
	root := &Directory{Name: "root"}
	users, _ := root.AddChild("users")
	etc, _ := root.AddChild("etc")
	etc.AddFile("huh", "go", "package huh")
	etc.AddFile("whaaa", "go", "package whaa")

	home, _ := users.AddChild(user)
	home.AddFile("stuff", "txt", "things")
	docs, _ := home.AddChild("Documents")
	docs.AddFile("test", "txt", "Hello, world!\r\nFoo\r\nBar")
	docs.AddFile("foo", "txt", "Hello, world!\r\nBar\r\nBaz")

	return root
}
