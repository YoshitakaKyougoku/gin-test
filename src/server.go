package main

import (
	"gin-test/graph"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	r := gin.Default()

	r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	r.POST("/query", gin.WrapH(srv))

	log.Printf("GraphQL playground にアクセスするには http://localhost:%s/ に接続してください", port)
	log.Fatal(r.Run(":" + port))
}
