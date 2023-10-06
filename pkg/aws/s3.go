package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type BambooBucket struct {
	Client *s3.Client
	BucketName string
	Prefix string
}

func (b BambooBucket) Create(region string) error {
	_, err := b.Client.CreateBucket(
		context.TODO(), 
		&s3.CreateBucketInput{
			Bucket: aws.String(b.BucketName),
			CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraint(region),
			},
		},
	)

	b.Client.PutBucketTagging(
		context.TODO(),
		&s3.PutBucketTaggingInput{
			Bucket: aws.String(b.BucketName),
			Tagging: &types.Tagging{
				TagSet: []types.Tag{
					{
						Key: aws.String("project"),
						Value: aws.String("bamboo"),
					},
				},
			},
		},
	)

	// TODO aka-somix: add error management

	return err
}