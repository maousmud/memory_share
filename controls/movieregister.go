package controls

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MovieregisterPostAction(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "movieregister.html", gin.H{})
}
