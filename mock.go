package dynamodb

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/elvenworks/dynamodb-conector/internal/domain"
	"github.com/stretchr/testify/mock"
)

type DynamodbMock struct {
	mock.Mock
}

func (m DynamodbMock) GetConfig() *domain.DynamodbConfig {
	args := m.Called()
	return args.Get(0).(*domain.DynamodbConfig)
}
func (m DynamodbMock) GetItem(tableName string, key string, item string) (*dynamodb.GetItemOutput, error) {
	args := m.Called(tableName, key, item)
	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

func (m DynamodbMock) GetCount(tableName string, limit int32) (int32, error) {
	args := m.Called(tableName, limit)
	return args.Get(0).(int32), args.Error(1)
}
