package template

import (
	"context"
	"fmt"
	"os"

	"github.com/aka-somix/bamboo/pkg/aws"
	"github.com/aka-somix/bamboo/pkg/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var baseFolder = "."

type FilesPacker struct {
	path string
}

func (fp FilesPacker) PackAndUpload(localName string, originPath string) error {

	zipFilePath := fmt.Sprintf("%s/%s.zip", baseFolder, localName)

	utils.ZipFolder(fp.path, zipFilePath)

	// Upload to S3
	cfg, _ := config.LoadDefaultConfig(context.TODO())

	bucket := aws.BambooBucket{
		Client: s3.NewFromConfig(cfg),
		BucketName: fmt.Sprintf("bamboo-templates-bucket-%s", cfg.Region), // TODO aka-somix: add account number or other specific info
	}

	err := bucket.UploadFile(originPath, zipFilePath)

	if err != nil {
		return err
	}

	// Remove the temp zip
	err = os.Remove(zipFilePath)

	if err != nil {
		return err
	}

	return nil
}

func (fp FilesPacker) DownloadAndUnpack(s3Path string) error {

	// Upload to S3
	cfg, _ := config.LoadDefaultConfig(context.TODO())

	bucket := aws.BambooBucket{
		Client: s3.NewFromConfig(cfg),
		BucketName: fmt.Sprintf("bamboo-templates-bucket-%s", cfg.Region), // TODO aka-somix: add account number or other specific info
	}

	zipFilePath := fmt.Sprintf("%s/temp.zip", baseFolder)

	err := bucket.DownloadFile(s3Path, zipFilePath)

	if err != nil {
		return err
	}

	err = utils.UnzipToFolder(zipFilePath, fp.path)

	if err != nil {
		return err
	}

	// Remove the temp zip
	err = os.Remove(zipFilePath)

	if err != nil {
		return err
	}

	return nil
}
