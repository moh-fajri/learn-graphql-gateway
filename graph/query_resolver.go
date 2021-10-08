package graph

import (
	"context"

	"github.com/moh-fajri/learn-garphql-gateway/graph/model/user"
)

type QueryResolver struct {
	Server *Server
}

func (r QueryResolver) Users(ctx context.Context) ([]*user.User, error) {
	return r.Server.UserClient.List(ctx)
}
