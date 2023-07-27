package main

import (
	"github.com/dave-andrew/gqlgen-todos/database"
	"github.com/dave-andrew/gqlgen-todos/graph/model"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dave-andrew/gqlgen-todos/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := database.Instance()

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	resolver := &graph.Resolver{
		Db: db,
	}

	db.AutoMigrate(&model.User{}, &model.Post{})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
