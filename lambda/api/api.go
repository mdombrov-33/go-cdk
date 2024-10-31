package api

import (
	"lambda-func/database"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient // when we call api functions, we want to interact with the database
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}
