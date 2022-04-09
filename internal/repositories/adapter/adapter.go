package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Database struct {
	conn    *dynamodb.DynamoDB
	logMode bool
}

type Interface interface {
}

func NewInterface() *Interface {

}

func (db *Database) Health() bool {
	_, err := db.conn.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (db *Database) FindAll() {

}

func (db *Database) FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {
	parsedCondition, err := dynamodbattribute.MarshalMap(condition)

	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       parsedCondition,
	}
	return db.conn.GetItem(input)
}

func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {
	parsedEntity, err := dynamodbattribute.MarshalMap(entity)

	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      parsedEntity,
		TableName: aws.String(tableName),
	}

	return db.conn.PutItem(input)
}

func (db *Database) Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {
	parsedCondition, err := dynamodbattribute.MarshalMap(condition)

	if err != nil {
		return nil, err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       parsedCondition,
		TableName: aws.String(tableName),
	}

	return db.conn.DeleteItem(input)
}
