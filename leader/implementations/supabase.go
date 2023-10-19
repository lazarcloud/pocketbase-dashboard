package implementations

import (
	"fmt"
	"os"
	"strings"

	"github.com/lazarcloud/lazarbase/leader/globals"
	cp "github.com/otiai10/copy"
)

func PrepareSupabase(projectFolderName string, uuid string, portOffset int) {
	dir := globals.ProjectPath + projectFolderName
	os.Mkdir(dir, 0777)

	content, err := os.ReadFile("./clean-data/editable-compose.yml")
	if err != nil {
		panic(err)
	}

	newContent := strings.ReplaceAll(string(content), "{special}", uuid)
	newContent = strings.ReplaceAll(newContent, "./volumes", globals.VMPath+"/projects/"+projectFolderName+"/volumes")

	err = os.WriteFile(dir+"/docker-compose.yml", []byte(newContent), 0777)
	if err != nil {
		panic(err)
	}

	content, err = os.ReadFile("./clean-data/editable.env.example")
	if err != nil {
		panic(err)
	}

	// Replace {special} with the string variable
	newContent = strings.ReplaceAll(string(content), "3000", fmt.Sprint(3000+portOffset))
	newContent = strings.ReplaceAll(newContent, "8000", fmt.Sprint(8000+portOffset))
	newContent = strings.ReplaceAll(newContent, "8443", fmt.Sprint(8443+portOffset))
	newContent = strings.ReplaceAll(newContent, "4000", fmt.Sprint(4000+portOffset))
	newContent = strings.ReplaceAll(newContent, "5432", fmt.Sprint(5432+portOffset))

	// Write the new content back to the file
	err = os.WriteFile(dir+"/.env", []byte(newContent), 0777)
	if err != nil {
		panic(err)
	}

	//copy folder volumes to test2

	cp.Copy("./clean-data/volumes", dir+"/volumes")

	content, err = os.ReadFile("./clean-data/volumes/logs/vector.yml")
	if err != nil {
		panic(err)
	}

	newContent = strings.ReplaceAll(string(content), "{special}", uuid)

	err = os.WriteFile(dir+"/volumes/logs/vector.yml", []byte(newContent), 0777)
	if err != nil {
		panic(err)
	}
}
