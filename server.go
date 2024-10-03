package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andess86/gqlgen-todos/generator"
	"github.com/andess86/gqlgen-todos/graph"
)

const defaultPort = "8080"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    // Initialize the propeller generator and channel
    propellerGen := generator.NewPropellerGenerator()
    propellerChannel := propellerGen.StartGenerating()

    // Create a new resolver instance with the propeller channel
    resolver := &graph.Resolver{
        PropellerDataChannel: propellerChannel,
    }

    // Start the propeller generator in a separate goroutine
    go func() {
        for data := range propellerChannel {
            log.Printf("Main received data: %+v", data)
        }
    }()

    // Create a new server with the generated schema
    srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))
    srv.AddTransport(&transport.Websocket{}) // Open up a websocket connection

    // Setup routes
    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    // Start the HTTP server
    log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
