package driver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetAWSDynamoDBClient(accessKeyID string, secretAccessKey string, region string) (*dynamodb.Client, error) {
	var cfg aws.Config
	var err error

	if len(accessKeyID) == 0 || len(secretAccessKey) == 0 {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			return nil, err
		}
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")))
		if err != nil {
			return nil, err
		}
		cfg.Region = region
	}

	return dynamodb.NewFromConfig(cfg), nil
}
