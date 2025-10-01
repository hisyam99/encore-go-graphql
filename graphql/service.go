package graphql

import (
	"context"
	"net/http"

	"encore.app/app" // Import app package to access Service
	"encore.app/app/dataloader"
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

	// Initialize DataLoaders
	dataLoaders := dataloader.NewDataLoaders(db)

	// Create config with Resolver that uses db, services, and dataLoaders
	cfg := generated.Config{Resolvers: &Resolver{
		db:          db,
		services:    appServices,
		dataLoaders: dataLoaders,
	}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))

	pg := playground.Handler("GraphQL Playground", "/graphql")
	return &Service{srv: srv, playground: pg}, nil
}

var _ = initService

//encore:api public raw path=/graphql
func (s *Service) Query(w http.ResponseWriter, req *http.Request) {
	// Extract Authorization header and put it in context
	authHeader := req.Header.Get("Authorization")
	ctx := context.WithValue(req.Context(), utils.AuthHeaderKey, authHeader)

	// Initialize app service to get database connection for DataLoaders
	appService, err := app.New()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Add DataLoaders to context
	dataLoaders := dataloader.NewDataLoaders(appService.DB())
	ctx = dataloader.ContextWithDataLoaders(ctx, dataLoaders)
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
