package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	types "url-shortnner/Types"
)

func (*IServer) AddNewURL(context *gin.Context) {
	var body types.NewURL

	if err := context.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid URL",
		})

		log.Fatalln(err.Error())
	}

	// TODO: add the all the logic needed.
}

func (*IServer) GetUrlByID(context *gin.Context) {
	var body types.NewURL

	if err := context.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid URL",
		})

		log.Fatalln(err.Error())
	}

	// TODO: add the all the logic needed.
}
