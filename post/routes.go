package post

import (
	"gateway/auth"
	"gateway/util"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *util.Config, authSvc *auth.AuthService) *PostService {
	a := auth.InitAuthMiddleware(authSvc)
	svc := &PostService{
		SVC: NewPostServiceClient(c),
	}

	authC := NewPostController(svc.SVC)

	routes := r.Group("/post")
	routes.Use(a.AuthRequired)
	routes.POST("/create", authC.CreatePost)

	return svc
}
