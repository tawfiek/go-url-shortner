package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
	types "url-shortnner/Types"
	"url-shortnner/config"
	"url-shortnner/models"
	"url-shortnner/utils"
)

type IServer struct {
	Env config.Env
	DB  *gorm.DB
}

func (server *IServer) AddNewURL(context *gin.Context) {
	var body types.NewURL

	err := context.ShouldBindJSON(&body)

	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid URL",
		})

		return
	}

	fmt.Printf("Bounded Body: %v \n", body)

	encoding := utils.EncodeURL([]byte(body.LongUrl))
	uid := utils.RandomStringFromEncoding(10, []rune(encoding))

	durationToAdd := int64(server.Env.DefaultTTL) * int64(time.Hour) * 24
	ttl := time.Now().Add(time.Duration(durationToAdd))

	newURL := models.URL{
		OriginalURL: body.LongUrl,
		AddedAt:     time.Now(),
		UID:         uid,
		TTL:         ttl,
	}

	server.DB.Create(&newURL)

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"uid":     uid,
		"longURL": body.LongUrl,
	})
}

func (server *IServer) GetUrlByUniqueID(context *gin.Context) {
	var url models.URL
	uid := context.Param("uid")

	server.DB.First(&url, "uid=?", uid)
	context.Redirect(http.StatusPermanentRedirect, url.OriginalURL)
}
