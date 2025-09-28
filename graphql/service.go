package graphql

import (
	"context"
	"net/http"

	"encore.app/app" // Import app package to access Service
	"encore.app/app/services"
	"encore.app/app/utils"
	"encore.app/graphql/generated"
	"encore.dev"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

//go:generate go run github.com/99designs/gqlgen generate

//encore:service
type Service struct {
	srv        *handler.Server
	playground http.Handler
}

func initService() (*Service, error) {
	// Initialize app.Service to get the db
	appService, err := app.New() // Call constructor from app/app.go
	if err != nil {
		return nil, err
	}

	// Use appService's db
	db := appService.DB()

	// Initialize services
	appServices := services.NewServices(db)

	// Create config with Resolver that uses db and services
	cfg := generated.Config{Resolvers: &Resolver{db: db, services: appServices}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	pg := playground.Handler("GraphQL Playground", "/graphql")
	return &Service{srv: srv, playground: pg}, nil
}

//encore:api public raw path=/graphql
func (s *Service) Query(w http.ResponseWriter, req *http.Request) {
	// Extract Authorization header and put it in context
	authHeader := req.Header.Get("Authorization")
	ctx := context.WithValue(req.Context(), utils.AuthHeaderKey, authHeader)
	req = req.WithContext(ctx)

	// Serve GraphQL with the updated context
	s.srv.ServeHTTP(w, req)
}

//encore:api public raw path=/graphql/playground
func (s *Service) Playground(w http.ResponseWriter, req *http.Request) {
	if encore.Meta().Environment.Type == encore.EnvProduction {
		http.Error(w, "Playground disabled", http.StatusNotFound)
		return
	}
	s.playground.ServeHTTP(w, req)
}
