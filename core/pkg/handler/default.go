package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Say(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello wang")
}
