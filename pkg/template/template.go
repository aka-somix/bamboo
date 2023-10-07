package template

import "github.com/aka-somix/bamboo/pkg/aws"


type Template struct {
	Author    string
	Name string
	Description string
}


func templateFromDynamoDB(item aws.DDBTemplate) Template {
	return Template{
		Name: item.Name,
		Author: item.Author,
		Description: item.Description,
	}
}

func templateFromDynamoDBList(items []aws.DDBTemplate) []Template {
	var templateList []Template

	for _, item := range items {
		templateList = append(templateList, templateFromDynamoDB(item))
	}

	return templateList
}
