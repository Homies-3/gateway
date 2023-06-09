package auth

import (
	"context"
	pb "gateway/auth/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Client pb.AuthServiceClient
}

func NewAuthController(client pb.AuthServiceClient) controller {
	return controller{
		Client: client,
	}
}

func (c controller) Register(gc *gin.Context) {
	rR := RegisterRequest{}
	if err := gc.Bind(&rR); err != nil {
		log.Println(err)
		gc.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := c.Client.Register(context.Background(), &pb.RegisterRequest{
		Phone:    rR.Phone,
		Password: rR.Password,
		Company:  rR.Company,
	})
	if err != nil {
		log.Println(err)
		gc.AbortWithStatus(http.StatusBadGateway)
		return
	}
	gc.JSON(http.StatusOK, &res)
}

func (c controller) Validate(gc *gin.Context) {
	token, statusCode := ExtractJWTFromRequest(gc)
	if statusCode == http.StatusUnauthorized {
		gc.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if len(token) < 2 {
		gc.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil {
		log.Println(err)
		gc.AbortWithStatus(http.StatusBadGateway)
		return
	}
	gc.JSON(int(res.GetStatus()), &res)
}

func (c controller) Login(gc *gin.Context) {
	lR := LoginRequest{}
	if err := gc.Bind(&lR); err != nil {
		log.Println(err)
		gc.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := c.Client.Login(context.Background(), &pb.LoginRequest{
		Username: lR.Username,
		Password: lR.Password,
	})
	if err != nil {
		log.Println(err)
		gc.AbortWithStatus(http.StatusBadGateway)
		return
	}
	gc.JSON(http.StatusOK, &res)
}
