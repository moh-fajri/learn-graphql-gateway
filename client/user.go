package client

import (
	"context"
	"os"

	"github.com/moh-fajri/learn-garphql-gateway/util"

	"github.com/jinzhu/copier"
	"github.com/moh-fajri/learn-garphql-gateway/graph/model/user"

	pb "github.com/moh-fajri/learn-garphql-gateway/proto/user"
)

type UserClient struct{}

func NewUserClient() *UserClient {
	return &UserClient{}
}

func (u UserClient) List(ctx context.Context) ([]*user.User, error) {
	conn := util.Dial(os.Getenv("USER_GRPC"))
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	result, err := c.List(ctx, &pb.NoRequest{})
	if err != nil {
		return nil, err
	}
	var response []*user.User
	copier.Copy(&response, result.Users)
	return response, nil
}

func (u UserClient) Create(ctx context.Context, input user.NewUser) ([]*user.User, error) {
	conn := util.Dial(os.Getenv("USER_GRPC"))
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)
	result, err := c.Create(ctx, &pb.CreateRequest{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}
	var response []*user.User
	copier.Copy(&response, result.Users)
	return response, nil
}
