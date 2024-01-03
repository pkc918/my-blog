package initial

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-blog/router"
)

func Server() {
	r := gin.Default()
	fmt.Println("initial")
	router.AddRoutes(r)
	r.Run()
}
