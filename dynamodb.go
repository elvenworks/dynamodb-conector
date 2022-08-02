package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/elvenworks/dynamodb-conector/internal/domain"
	"github.com/elvenworks/dynamodb-conector/internal/driver"
	"github.com/sirupsen/logrus"
)

type Dynamodb struct {
	client *dynamodb.Client
	config domain.DynamodbConfig
}

type InitConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

func InitDynamodb(config InitConfig) *Dynamodb {

	clientDynamodb, err := driver.GetAWSDynamoDBClient(config.AccessKeyID, config.SecretAccessKey, config.Region)

	if err != nil {
		logrus.Error("unable to get dynamodb client")
	}

	return &Dynamodb{
		client: clientDynamodb,
		config: domain.DynamodbConfig{
			AccessKeyID:     config.AccessKeyID,
			SecretAccessKey: config.SecretAccessKey,
			Region:          config.Region,
		},
	}
}

func (d *Dynamodb) GetConfig() *domain.DynamodbConfig {
	return &d.config
}

func (d *Dynamodb) GetCount(tableName string, limit int32) (int32, error) {
	if limit == 0 {
		limit = 1
	}
	output, err := d.client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(tableName),
		Limit:     aws.Int32(limit),
	})

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return output.Count, err
}

func (d *Dynamodb) GetItem(tableName string, key string, item string) (*dynamodb.GetItemOutput, error) {

	output, err := d.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			key: &types.AttributeValueMemberS{Value: item},
		},
	})

	if err != nil {
		return nil, err
	}

	return output, err

}
