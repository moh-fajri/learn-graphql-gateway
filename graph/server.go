//go:generate go run github.com/99designs/gqlgen

package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/moh-fajri/learn-garphql-gateway/client"
	"github.com/moh-fajri/learn-garphql-gateway/graph/generated"
)

type Server struct {
	UserClient *client.UserClient
}

func NewGraphQLServer() (*Server, error) {
	userClient := client.NewUserClient()
	return &Server{
		UserClient: userClient,
	}, nil
}

func (s *Server) Mutation() generated.MutationResolver {
	return &MutationResolver{
		Server: s,
	}
}
func (s *Server) Query() generated.QueryResolver {
	return &QueryResolver{
		Server: s,
	}
}

func (s *Server) TOExecutableSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: s,
	})
}
