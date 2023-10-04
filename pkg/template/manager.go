// Templates Manager
// ------------------
// TODO aka-somix: Add documentation

package template

import (
	"errors"
	"fmt"
	"os"
)


type templateManagerImpl struct {}


func NewTemplateManager() TemplateManager {
	return &templateManagerImpl{}
}

func (tm *templateManagerImpl) ListTemplatesInfo(author string) []Template {
	return []Template{
		{
			Name: "Test",
			Author: author,
			Description: "A Random Test Template",
			Path: "./ciao",
		},
		{
			Name: "Test2",
			Author: author,
			Description: "A Random Test Template",
			Path: "./ciao",
		},
	}
}

func (tm *templateManagerImpl) GetTemplateInfo(author string, name string) Template {
	return Template {
			Name: name,
			Author: author,
			Description: "A Random Test Template",
			Path: "./ciao",
		}
}

func (tm *templateManagerImpl) DownloadTemplate(author string, name string, folderPath string) error{

	// if the folder specified does not exist
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
        // ...try create the folder
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			// ... catch creation errors
			return errors.New("folder does not exist and could not be created")
		}
    }

	template := tm.GetTemplateInfo(author, name)

	fmt.Printf("Downloading template %s to folder: %s \n", template.Name, folderPath)

	return nil
}
