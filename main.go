package main

import (
	"github.com/eminoz/go-redis-project/config"
	"github.com/eminoz/go-redis-project/database"
	"github.com/eminoz/go-redis-project/router"
)

func main() {
	config.SetupConfig()
	database.SetDatabase()
	setup := router.Setup()
	setup.Listen(":" + config.GetConfig().Port)

}

