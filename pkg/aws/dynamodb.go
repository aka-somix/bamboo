package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type TemplatesTable struct {
	Client *dynamodb.Client
	TableName      string
}


func (t TemplatesTable) QueryTemplates(author string) ([]map[string]types.AttributeValue, error) {

	response, err := t.Client.Query(
		context.TODO(), 
		&dynamodb.QueryInput{
			TableName:                 aws.String(t.TableName),
			ExpressionAttributeNames:  map[string]string{"#hashkey": "Author",},
			ExpressionAttributeValues: map[string]types.AttributeValue{":hashkey": &types.AttributeValueMemberS{Value:author}},
			KeyConditionExpression:    aws.String("#hashKey = :hashKeyValue"),
		},
	)

	if err != nil {
		// TODO aka-somix: better error management
		return nil, err
	}

	return response.Items, nil

}