package client

import (
	"fmt"

	"github.com/RTS-1989/go-censor-svc/pkg/config"
	"github.com/RTS-1989/go-censor-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CensorServiceClient
}

func InitServiceClient(c *config.Config) pb.CensorServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CommentSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCensorServiceClient(cc)
}
