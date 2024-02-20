package handler

import (
	"fmt"
	"net/http"

	"github.com/azizbek/middleware/api/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Ping(ctx *gin.Context) {
	fmt.Println("I am in the ping")
	ctx.JSON(200, gin.H{
		"message": "Pong",
	})
}

func (h *Handler) Secure(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Congratulations, you have access to secure",
	})
}

func (h *Handler) GetJWTToken(ctx *gin.Context) {
	jwtToken := jwt.JWTHandler{
		Sub:     uuid.NewString(),
		Role:    "user",
		Timeout: 300,
	}

	accessToken, err := jwtToken.GenerateAuthJWT()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Oops something went wrong!!!",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"access_token": accessToken,
	})
}
