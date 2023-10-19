package filesystem

import "os"

func PrepareFolderStructure() {
	CreateFolderIfNotExists("data")
	CreateFolderIfNotExists("data/projects")

	CreateFile("./data/test.txt", "Hello World\n")
}

func CreateFile(path string, content string) {

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		panic(err)
	}
}
func CreateFolder(path string) {
	err := os.Mkdir(path, 0755)
	if err != nil {
		panic(err)
	}
}

func CreateFolderIfNotExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		CreateFolder(path)
	}
}
