package controllers

import (
	"gorm.io/gorm"
	"url-shortnner/config"
)

type IServer struct {
	Env config.Env
	DB  *gorm.DB
}
