// Templates Manager
// ------------------
// TODO aka-somix: Add documentation

package template

import (
	"errors"
	"fmt"
	"os"

	"github.com/aka-somix/bamboo/pkg/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)


type templateManagerImpl struct {}


func NewTemplateManager() TemplateManager {
	return &templateManagerImpl{}
}

func (tm *templateManagerImpl) ListTemplatesInfo(author string) []Template {
	templateTable := aws.BambooTable{
		Client: dynamodb.New(dynamodb.Options{Region: "eu-west-1"}),
		TableName: "BambooTemplatesTable",
	}

	itemsFound, _ := templateTable.QueryTemplates(author)

	fmt.Println("Items Found on dynamodb:")
	for _, item := range itemsFound {
		fmt.Printf("%v \n", item)
	}

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
	};
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

func (tm *templateManagerImpl) CreateTemplate(t Template,  sourcePath string) error{

	fmt.Printf("Creating template %s from folder: %s \n", t.Name, sourcePath)

	return nil
}
