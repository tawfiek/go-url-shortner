package routers

import (
	"github.com/gin-gonic/gin"
	shortner "url-shortnner/controllers"
)

func MainRouter(group *gin.RouterGroup) {
	group.POST("/short", shortner.AddNewURL)
	group.GET("/:id")
}
