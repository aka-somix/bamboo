package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type BambooTable struct {
	Client *dynamodb.Client
	TableName      string
}


func (t BambooTable) QueryTemplates(author string) ([]map[string]types.AttributeValue, error) {

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


func (t BambooTable) Create() error {
	_, err := t.Client.CreateTable(
		context.TODO(),
		&dynamodb.CreateTableInput{
			TableName: aws.String(t.TableName),
			KeySchema: []types.KeySchemaElement{
				{
					AttributeName: aws.String("Author"),
					KeyType: types.KeyTypeHash,
				},
				{
					AttributeName: aws.String("Name"),
					KeyType: types.KeyTypeRange,
				},
			},
			AttributeDefinitions: []types.AttributeDefinition{
				{
					AttributeName: aws.String("Author"),
					AttributeType: types.ScalarAttributeTypeS,
				},
				{
					AttributeName: aws.String("Name"),
					AttributeType: types.ScalarAttributeTypeS,
				},
			},
			BillingMode: types.BillingModePayPerRequest,
			TableClass: types.TableClassStandardInfrequentAccess,
			// TODO aka-somix: Commented for testing, remove comment afterward
			// DeletionProtectionEnabled: aws.Bool(true), 
			Tags: []types.Tag{
				{
					Key: aws.String("project"),
					Value: aws.String("bamboo"),
				},
			},
		}, 
	)

	if err != nil {
		// TODO aka-somix: better error management
		return err
	}

	return nil
}
