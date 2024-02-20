package api

import (
	"github.com/azizbek/middleware/api/handler"
	"github.com/azizbek/middleware/api/middleware"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	router := gin.New()
	handlerV1 := handler.New()

	router.Use(middleware.Auth)
	router.GET("/secure", handlerV1.Secure)
	router.GET("/jwt", handlerV1.GetJWTToken)

	v1 := router.Group("/v1", middleware.MiddleLogger)
	v1.GET("/ping", handlerV1.Ping)

	v2 := router.Group("/v2", middleware.MiddleLoggerV2)
	v2.GET("/ping", handlerV1.Ping)

	return router
}
