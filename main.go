package main

import (
	log "github.com/sirupsen/logrus"
	"url-shortnner/config"
	"url-shortnner/controllers"
	"url-shortnner/routers"
)

func loadEnv() config.Env {
	env, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	return env
}

func main() {
	env := loadEnv()
	dbConnection := config.InitDB(env)
	controllers := new(controllers.IServer)

	controllers.DB = dbConnection
	controllers.Env = env

	engin := routers.SetupApplicationRouters(controllers)

	engin.Run(":" + env.Port)
}
