package routers

import "github.com/gin-gonic/gin"

func MainRouter(group *gin.RouterGroup) {
	group.POST("/short")
	group.GET("/:id")
}
