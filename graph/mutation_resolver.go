package graph

import (
	"context"

	"github.com/moh-fajri/learn-garphql-gateway/graph/model/user"
)

type MutationResolver struct {
	Server *Server
}

func (r MutationResolver) CreateUser(ctx context.Context, input user.NewUser) ([]*user.User, error) {
	return r.Server.UserClient.Create(ctx, input)
}
