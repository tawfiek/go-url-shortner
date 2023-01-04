package routers

import (
	"github.com/gin-gonic/gin"
	"url-shortnner/controllers"
)

func setupURLRouters(group *gin.RouterGroup, controllers *controllers.IServer) {
	group.POST("/short", controllers.AddNewURL)
	group.GET("/:id", controllers.GetUrlByID)
}
