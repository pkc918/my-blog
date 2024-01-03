package service

import (
	"github.com/gin-gonic/gin"
	"my-blog/app/controller"
)

func GetGithubProfile(c *gin.Context) {
	githubProfile := controller.GetGithubProfile()
	c.JSON(200, gin.H{"data": &githubProfile})
}
