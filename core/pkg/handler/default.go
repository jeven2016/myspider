package handler

import (
	"core/pkg/client"
	"core/pkg/common/utils"
	"core/pkg/stream/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Say(ctx *gin.Context) {
	homeMessage := &message.HomeMessage{}
	homeMessage.SiteKey = "web1"
	msg, err := utils.ConvertToMap(homeMessage)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	err = client.GetRedisClient().PublishMessage(ctx, msg, message.HomeUrlStream)
	if err != nil {
		ctx.AbortWithError(500, err)
		return
	}
	ctx.String(http.StatusOK, "hello wang")
}
