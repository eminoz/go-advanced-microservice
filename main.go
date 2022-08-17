package main

import (
	"github.com/eminoz/go-advanced-microservice/config"
	"github.com/eminoz/go-advanced-microservice/database"
	"github.com/eminoz/go-advanced-microservice/router"
)

func main() {
	config.SetupConfig()
	database.SetDatabase()
	setup := router.Setup()
	setup.Listen(":" + config.GetConfig().Port)

}
