package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/moh-fajri/learn-garphql-gateway/graph"

	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s, err := graph.NewGraphQLServer()
	if err != nil {
		log.Fatal(err)
	}

	graphqlHandler := handler.NewDefaultServer(s.TOExecutableSchema())

	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err = e.Start(":" + port)
	if err != nil {
		log.Fatalln(err)
	}
}
