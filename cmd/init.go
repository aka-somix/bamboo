/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/aka-somix/bamboo/pkg/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the underlying back-end infrastructure to manage your scaffolds",
	Long: `The init command is used to create the back-end infrastructure to host and manage your scaffolds`,


	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating Infrastructure to support Bamboo on your AWS Account")

		cfg, _ := config.LoadDefaultConfig(context.TODO())

		table := aws.BambooTable{
			Client: dynamodb.NewFromConfig(cfg),
			TableName: "BambooTemplatesTable",
		}

		bucket := aws.BambooBucket{
			Client: s3.NewFromConfig(cfg),
			BucketName: fmt.Sprintf("bamboo-templates-bucket-%s", cfg.Region), // TODO aka-somix: add account number or other specific info
		}

		if err := table.Create(); err != nil {
			fmt.Printf("Error while creating DynamoDB table: %v \n", err)
			os.Exit(1)
		}
		fmt.Println("DynamoDB table successfully created")
		
		if err := bucket.Create(cfg.Region); err != nil {
			fmt.Printf("Error while creating S3 Bucket: %v \n", err)
			os.Exit(1)
		}
		fmt.Println("S3 Bucket successfully created")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
