package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/dachichang/goath-stack-template/docs"
	"github.com/dachichang/goath-stack-template/internal/interface/api"
	"github.com/dachichang/goath-stack-template/internal/interface/middleware"
	"github.com/dachichang/goath-stack-template/internal/interface/web"
)

func NewRoutes(
	systemHandler *api.SystemHandler,
	webController *web.Controller,
	log *logrus.Logger,
) *gin.Engine {
	route := gin.Default()

	route.GET("/livez", systemHandler.Health)

	// NOTE: API
	apiV1 := route.Group("/api/v1", middleware.CorsMiddleware())
	apiV1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// NOTE: WEB
	authMiddleware := middleware.AuthorizationMiddleware(log)

	web := route.Group("", authMiddleware)
	web.Static("/assets", "./internal/interface/web/public/")
	web.GET("/", webController.Home)

	return route
}
