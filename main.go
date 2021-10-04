package main

import (
	"github.com/goddamnnoob/notReddit/app"
	"github.com/goddamnnoob/notReddit/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Error(".env File not found " + err.Error())
		panic(err)
	}
}

func main() {

	logger.Info("Starting the app")
	app.Start()
}
