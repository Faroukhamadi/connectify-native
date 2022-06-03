package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Faroukhamadi/connectify-native/ent"
	"github.com/Faroukhamadi/connectify-native/graph"
	"github.com/Faroukhamadi/connectify-native/graph/generated"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

const defaultPort = "9500"

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// dbUrl := os.Getenv("DATABASE_URL") + "?sslmode=disable"
	dbUrl := "postgresql://postgres:16%2F04%2F2002farouk@localhost:5432/Connectify?sslmode=disable"

	client, err := ent.Open("postgres", dbUrl)
	fmt.Println("database url: ", dbUrl)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	example(ctx, client)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func example(ctx context.Context, client *ent.Client) {
	farouk := client.User.
		Create().
		SetFirstName("Farouk").
		SetLastName("Hamadi").
		SetUsername("farakrakk").
		SetPassword("1234").SaveX(ctx)
	fmt.Println("User created: ", farouk)
}
