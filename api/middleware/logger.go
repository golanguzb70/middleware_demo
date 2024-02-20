package middleware

import (
	"fmt"
	"net/http"

	"github.com/azizbek/middleware/api/jwt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func MiddleLogger(ctx *gin.Context) {
	fmt.Printf("V1 -> Method:%s, path: %s, IP: %s \n", ctx.Request.Method, ctx.Request.URL.Path, ctx.ClientIP())
	ctx.Next()
}

func MiddleLoggerV2(ctx *gin.Context) {
	fmt.Printf("V2 ->  Method:%s, path: %s, IP: %s \n", ctx.Request.Method, ctx.Request.URL.Path, ctx.ClientIP())
	ctx.Next()
}

func BlockPingV2(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/v2/ping" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Requesting to this url is blocked",
		})
		return
	}

	ctx.Next()
}

func Auth(ctx *gin.Context) {
	if ctx.Request.URL.Path == "/jwt" {
		ctx.Next()
		return
	}

	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized request",
		})
		return
	}

	claims, err := jwt.ExtractClaim(token, []byte("SecureSignINKey"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}

	if cast.ToString(claims["role"]) != "user" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "have no access to this path",
		})
		return
	}

	ctx.Next()
}
