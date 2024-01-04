package controller

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/service"
)

func GetGithubProfile(c *gin.Context) {
	githubProfile := service.GetGithubProfile()
	c.JSON(200, gin.H{"data": &githubProfile})
}
