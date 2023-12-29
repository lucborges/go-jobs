package main

import (
	"github.com/lucborges/go-jobs/config"
	"github.com/lucborges/go-jobs/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}
	router.Initialize()
}
