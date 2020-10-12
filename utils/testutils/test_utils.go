package testutils
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sirupsen/logrus"
	"testing"
)

type IntegrationTest struct{
	test *testing.T
	table string
	localClient *dynamodb.DynamoDB
}

func NewIntegrationTest(test *testing.T, localClient *dynamodb.DynamoDB, tableName string) * IntegrationTest{
	this := new(IntegrationTest)
	this.test = test
	this.table = tableName
	this.localClient = localClient
	return this
}

func (it IntegrationTest) createTable(table dynamodb.CreateTableInput) {
	resource, err := it.localClient.CreateTable(&table)
	if err != nil {
		it.test.Error(err)
	}
	logrus.Infof("Resource created %v", resource)
}

func (it IntegrationTest) deleteTable() {
	_, err := it.localClient.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String(it.table)})

	if err != nil {
		it.test.Error(err)
	}
}

