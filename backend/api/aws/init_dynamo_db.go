package aws

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	logrus "github.com/sirupsen/logrus"
)

func InitDynamoDBClient() *dynamodb.Client {
	// Load the AWS configuration from environment variables
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")

	// Create AWS Config with environment variables
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsAccessKeyID, awsSecretAccessKey, "")),
	)
	if err != nil {
		logrus.Fatalf("unable to load SDK config, %v", err)
	}

	// Create a DynamoDB client
	logrus.Info("DynamoDB client initialized successfully")
	return dynamodb.NewFromConfig(cfg)
}
