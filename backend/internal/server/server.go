package server

import (
	"github.com/gin-contrib/cors"
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

	// cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	engine.Use(cors.New(config))

	registerRouters(engine)

	return engine
}

func registerRouters(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/checkHealth", api.CheckHealth)

		apiGroup.POST("/generateSignature", api.GetMintSignature)

		apiGroup.GET("/user/info", api.GetUserInfo)
		apiGroup.GET("/user/owldinals", api.GetUserOwldinalList)
		apiGroup.GET("/user/boxes", api.GetUserBoxList)
		apiGroup.GET("/user/inviter", api.GetUserInviteList)
		apiGroup.GET("/user/mint_tx", api.GetRequestMintTx)
		apiGroup.POST("/user/boxes/claim", api.ClaimBoxes)

		apiGroup.GET("/job/mint", api.GetRequestJob)

		apiGroup.GET("/game/info", api.GetGameInfo)
		apiGroup.GET("/game/rewards_trend", api.GetRewardTrend)
		apiGroup.GET("/game/rewards_history", api.GetRewardHistory)

		apiGroup.POST("/admin/update_rewards", api.UpdateRewards)
		apiGroup.POST("/admin/retry_all_jobs", api.RetryAllJobs)
		apiGroup.POST("/admin/reload_log", api.ReloadLog)
	}
}
