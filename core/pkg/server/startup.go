package server

import (
	"core/pkg/handler"
	"core/pkg/middleware"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"time"
)

// Start corresponding services
func Start() *gin.Engine {
	registerGobType()

	gin.SetMode(gin.ReleaseMode)
	var engine = gin.Default()
	engine.Use(
		middleware.GinLogger(),
		middleware.GinRecovery(false))

	engine.GET("/say", handler.Say)

	return engine
}

func registerGobType() {
	gob.Register(time.Time{})
}
