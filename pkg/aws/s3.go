package aws

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type BambooBucket struct {
	Client *s3.Client
	BucketName string
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


// UploadFile reads from a file and puts the data into an object in a bucket.
func (b BambooBucket) UploadFile(objectKey string, localFileName string) error {
	file, err := os.Open(localFileName)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", localFileName, err)
	} else {
		defer file.Close()
		_, err = b.Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(b.BucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})
		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				localFileName, b.BucketName, objectKey, err)
		}
	}
	return err
}


// DownloadFile gets an object from a bucket and stores it in a local file.
func (b BambooBucket) DownloadFile(objectKey string, fileName string) error {
	result, err := b.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(b.BucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		log.Printf("Couldn't get object %v:%v. Here's why: %v\n", b.BucketName, objectKey, err)
		return err
	}
	defer result.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		log.Printf("Couldn't create file %v. Here's why: %v\n", fileName, err)
		return err
	}
	defer file.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Printf("Couldn't read object body from %v. Here's why: %v\n", objectKey, err)
	}
	_, err = file.Write(body)
	return err
}
