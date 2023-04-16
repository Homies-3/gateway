package post

import (
	"context"
	pb "gateway/post/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Client pb.PostServiceClient
}

func NewPostController(client pb.PostServiceClient) controller {
	return controller{
		Client: client,
	}
}

func (c controller) CreatePost(gc *gin.Context) {
	rR := PostRequest{}
	if err := gc.Bind(&rR); err != nil {
		log.Println(err)
		gc.AbortWithStatus(http.StatusBadRequest)
		return
	}
	res, err := c.Client.CreatePost(context.Background(), &pb.PostRequest{
		Title:   rR.Title,
		Content: rR.Content,
		User: &pb.User{
			Username: gc.GetString("userId"),
			Company:  "test",
		},
	})
	if err != nil {
		log.Println(err)
		gc.AbortWithStatus(http.StatusBadGateway)
		return
	}
	gc.JSON(http.StatusCreated, &res)
}
