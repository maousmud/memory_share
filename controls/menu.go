package controls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MenuPostAction(ctx *gin.Context) {
	username := ctx.PostForm("username")
	ctx.HTML(http.StatusOK, "menu.html", gin.H{"username": username})
}
