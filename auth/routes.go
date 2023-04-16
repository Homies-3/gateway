package auth

import (
	"gateway/util"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *util.Config) *AuthService {
	svc := &AuthService{
		SVC: NewAuthServiceClient(c),
	}

	authC := NewAuthController(svc.SVC)

	routes := r.Group("/auth")
	routes.POST("/register", authC.Register)
	routes.POST("/login", authC.Login)

	return svc
}
