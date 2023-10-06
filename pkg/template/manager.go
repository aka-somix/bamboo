// Templates Manager
// ------------------
// TODO aka-somix: Add documentation

package template

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aka-somix/bamboo/pkg/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type TemplateManager interface {
	ListTemplatesInfo(author string) []Template
	GetTemplateInfo(author string, name string) Template
	DownloadTemplate(author string, name string, folderPath string) error
	CreateTemplate(t Template, sourcePath string) error
}


type templateManagerImpl struct {}


func NewTemplateManager() TemplateManager {
	return &templateManagerImpl{}
}

func (tm *templateManagerImpl) ListTemplatesInfo(author string) []Template {

	cfg, _ := config.LoadDefaultConfig(context.TODO())

	table := aws.BambooTable{
		Client: dynamodb.NewFromConfig(cfg),
		TableName: "BambooTemplatesTable",
	}

	itemsFound, err := table.QueryTemplates(author)

	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}

	return templateFromDynamoDBList(itemsFound)
}

func (tm *templateManagerImpl) GetTemplateInfo(author string, name string) Template {
	return Template {
			Name: name,
			Author: author,
			Description: "A Random Test Template",
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

	cfg, _ := config.LoadDefaultConfig(context.TODO())

	table := aws.BambooTable{
		Client: dynamodb.NewFromConfig(cfg),
		TableName: "BambooTemplatesTable",
	}

	table.PutTemplate(&aws.DDBTemplate{
		Author: t.Author,
		Name: t.Name,
		Description: t.Description,
		S3Path: fmt.Sprintf("%s/%s", t.Author, strings.Replace(t.Name, " ", "-", -1)),
	})

	return nil
}
