package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	"url-shortnner/routers"
)

func getPort() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	loadEnv()

	url := ginSwagger.URL("/docs/swagger.yaml") // The url pointing to API definition

	router.Use(gin.Logger())

	router.GET("/swagger/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.LoadHTMLGlob("views/*")

	v1Group := router.Group("/api/v1")

	routers.MainRouter(v1Group)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Nour and Tawfiek  Shortener ",
		})
	})

	port := getPort()
	router.Run(port)
}
