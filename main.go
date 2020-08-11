package main

import (
	"go-rest-project/config"
	"go-rest-project/db"
	"go-rest-project/logger"
	"go-rest-project/model"
	"go-rest-project/web"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		panic("can't load config, error: " + err.Error())
	}

	log, err := logger.Init()
	if err != nil {
		panic("can't load logger, error: " + err.Error())
	}

	err = db.Load()
	if err != nil {
		panic("can't connect to db, error: " + err.Error())
	}

	model.Migrate(log)

	err = web.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
