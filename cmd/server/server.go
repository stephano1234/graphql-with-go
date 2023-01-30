package main

import (
	"log"
	"net/http"
	"os"
	"database/sql"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/stephano1234/graphql-go/graph"
	"github.com/stephano1234/graphql-go/internal/database"

	_ "github.com/mattn/go-sqlite3" 
)

const defaultPort = "8080"

func main() {
	db, err := sql.Open("sqlite3", "db/data.db")
	if err != nil {
		log.Fatalf("failed to open database %v", err)
	}
	defer db.Close()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: database.NewCategory(db),
		CourseDB: database.NewCourse(db),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
