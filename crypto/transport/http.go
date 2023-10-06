package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/maxim12233/crypto-app-server/crypto/config"
	_ "github.com/maxim12233/crypto-app-server/crypto/docs"
	"github.com/maxim12233/crypto-app-server/crypto/endpoints"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(eps endpoints.AccountEndpoint) *gin.Engine {
	router := gin.Default()
	c := config.GetConfig()
	if c.GetString("env.mode") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.SetMode(gin.DebugMode)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("v1")
	{
		accountGroup := v1.Group("crypto")
		{
			accountGroup.GET("/:id", eps.GetAccount)
		}
	}

	return router
}
