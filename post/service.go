package post

import (
	pb "gateway/post/pb"
	"gateway/util"
	"log"

	"google.golang.org/grpc"
)

type PostService struct {
	SVC pb.PostServiceClient
}

func NewPostServiceClient(c *util.Config) pb.PostServiceClient {
	cc, err := grpc.Dial(c.PostSVCUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Cannot connect to POSTSERVICE")
	}
	log.Println("Connected to POSTSERVICE")
	return pb.NewPostServiceClient(cc)
}
