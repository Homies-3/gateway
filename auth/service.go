package auth

import (
	pb "gateway/auth/pb"
	"gateway/util"
	"log"

	"google.golang.org/grpc"
)

type AuthService struct {
	SVC pb.AuthServiceClient
}

func NewAuthServiceClient(c *util.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSVCUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Cannot connect to auth service")
	}
	log.Println("Connected to AUTHSERVICE")
	return pb.NewAuthServiceClient(cc)
}
