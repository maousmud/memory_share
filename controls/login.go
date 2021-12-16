package controls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginAction(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}
