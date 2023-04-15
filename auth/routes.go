package auth

import (
	"gateway/util"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *util.Config) *AuthService {
	svc := &AuthService{
		SVC: NewAuthServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", nil)
	routes.POST("/login", nil)

	return svc
}
