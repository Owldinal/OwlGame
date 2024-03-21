package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"owl-backend/internal/server/api"
	"owl-backend/pkg/log"
)

func SetupEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = log.Logger.Writer()
	engine := gin.New()
	engine.Use(gin.LoggerWithWriter(gin.DefaultWriter), gin.Recovery())

	registerRouters(engine)

	return engine
}

func registerRouters(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/checkHealth", api.CheckHealth)

		apiGroup.GET("/getSignature", api.GetMintSignature)
	}
}
