package main

import (
	"github.com/goddamnnoob/notReddit/app"
	"github.com/goddamnnoob/notReddit/logger"
)

func main() {

	logger.Info("Starting the app")
	app.Start()
}
