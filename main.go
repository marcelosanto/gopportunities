package main

import (
	"com.github/marcelosanto/gopportunities/config"
	"com.github/marcelosanto/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	//Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	//Initialize router
	router.Initialize()
}
