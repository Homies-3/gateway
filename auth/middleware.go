package auth

import (
	"context"
	"net/http"
	"strings"

	pb "gateway/auth/pb"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *AuthService
}

func InitAuthMiddleware(svc *AuthService) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func ExtractJWTFromRequest(ctx *gin.Context) (token []string, status int) {
	authorization := ctx.Request.Header.Get("authorization")
	if authorization == "" {
		status = http.StatusUnauthorized
		return
	}
	token = strings.Split(authorization, "Bearer ")
	status = http.StatusOK
	return
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	token, statusCode := ExtractJWTFromRequest(ctx)

	if statusCode == http.StatusUnauthorized {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.svc.SVC.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}
