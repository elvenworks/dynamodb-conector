package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/elvenworks/dynamodb-conector/internal/domain"
)

type IDynamodb interface {
	GetConfig() *domain.DynamodbConfig
	GetItem(tableName string, key string, item string) (*dynamodb.GetItemOutput, error)
	GetCount(tableName string, limit int32) (int32, error)
}
