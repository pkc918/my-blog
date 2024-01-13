package controller

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/service"
)

func GetPhotos(c *gin.Context) {
	service.GetPhotos()
}
