package main

import (
	"myapp/api/firebase"
	"myapp/auth"
	"myapp/dataloader"
	"myapp/directives"
	"myapp/graph/generated"
	"os"
	"time"

	"log"
	"net/http"

	"myapp/graph"
	"myapp/service"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {

	var Broker = new(service.Broker).New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	graph.Broker = Broker

	// firebase.SendToTokenLib()
	firebase.SendToToken()

	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	})

	graphConfig := generated.Config{Resolvers: &graph.Resolver{}}
	graphConfig.Directives.IsValidLogin = directives.IsValidLogin

	// Use New instead of NewDefaultServer in order to have full control over defining transports
	srv := handler.New(generated.NewExecutableSchema(graphConfig))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		InitFunc: auth.WebSocketMiddleware,
	})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	router.Use(auth.AuthMiddleware)
	router.Use(dataloader.DataloaderMiddleware)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
