package router

import (
	"memory_share/controls"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", controls.LoginAction)
	r.POST("/menu", controls.MenuPostAction)
	r.POST("/movieregister", controls.MovieregisterPostAction)

	return r
}
