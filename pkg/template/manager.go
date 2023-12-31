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

	return templateFromDynamoDBList(*itemsFound)
}

func (tm *templateManagerImpl) GetTemplateInfo(author string, name string) Template {

	cfg, _ := config.LoadDefaultConfig(context.TODO())

	table := aws.BambooTable{
		Client: dynamodb.NewFromConfig(cfg),
		TableName: "BambooTemplatesTable",
	}

	itemFound, err := table.GetTemplate(author, name)

	if err != nil {
		fmt.Printf("Error: %s \n", err)
	}


	return templateFromDynamoDB(*itemFound)
}

func (tm *templateManagerImpl) DownloadTemplate(author string, name string, folderPath string) error{

	// if the folder specified does not exist
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
        // ...try create the folder
		if err := os.MkdirAll(folderPath, 0777); err != nil {
			// ... catch creation errors
			return errors.New("folder does not exist and could not be created")
		}
    }

	t := tm.GetTemplateInfo(author, name)

	fmt.Printf("Downloading template %s to folder: %s \n", t.Name, folderPath)

	s3Path := fmt.Sprintf("%s/%s.zip", t.Author, strings.Replace(t.Name, " ", "-", -1))

	// Download from Origin
	err := FilesPacker{path: folderPath}.DownloadAndUnpack(s3Path)

	if err != nil {
		return err
	}

	return nil
}

func (tm *templateManagerImpl) CreateTemplate(t Template,  sourcePath string) error{

	fmt.Printf("Creating template %s from folder: %s \n", t.Name, sourcePath)

	cfg, _ := config.LoadDefaultConfig(context.TODO())

	table := aws.BambooTable{
		Client: dynamodb.NewFromConfig(cfg),
		TableName: "BambooTemplatesTable",
	}

	s3Path := fmt.Sprintf("%s/%s.zip", t.Author, strings.Replace(t.Name, " ", "-", -1))

	// Create template
	table.PutTemplate(&aws.DDBTemplate{
		Author: t.Author,
		Name: t.Name,
		Description: t.Description,
		S3Path: s3Path,
	})

	// Upload to Origin
	FilesPacker{path: sourcePath}.PackAndUpload(t.Name, s3Path)
	
	return nil
}
