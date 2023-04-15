package auth

import (
	"gateway/auth/pb"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	client pb.AuthServiceClient
}

func NewAuthController(client pb.AuthServiceClient) Controller {
	return Controller{
		client: client,
	}
}

func (c Controller) Register(gc *gin.Context) {

}
