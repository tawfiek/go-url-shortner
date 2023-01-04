package routers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os"
	controllers "url-shortnner/controllers"
)

func SetupApplicationRouters(server *controllers.IServer) *gin.Engine {
	router := gin.Default()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

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

	// Application Controllers Groups
	urlRouterGroup := v1Group.Group("/url")

	setupURLRouters(urlRouterGroup, server)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Nour and Tawfiek  Shortener ",
		})
	})

	return router
}
