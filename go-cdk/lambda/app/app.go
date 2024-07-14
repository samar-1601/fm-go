package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct{
	ApiHandler api.ApiHandler
}

func NewApp() App{
	// we initialize aur DBStore which gets passed down to the api handler
	db:= database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}

