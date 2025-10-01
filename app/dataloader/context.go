package dataloader

import (
	"context"
)

// Context keys for DataLoader
type contextKey string

const (
	DataLoaderContextKey contextKey = "dataloader"
)

// ContextWithDataLoaders adds DataLoaders to context
func ContextWithDataLoaders(ctx context.Context, loaders *DataLoaders) context.Context {
	return context.WithValue(ctx, DataLoaderContextKey, loaders)
}

// DataLoadersFromContext retrieves DataLoaders from context
func DataLoadersFromContext(ctx context.Context) *DataLoaders {
	loaders, ok := ctx.Value(DataLoaderContextKey).(*DataLoaders)
	if !ok {
		return nil
	}
	return loaders
}

// GetDataLoadersFromContext retrieves DataLoaders from context or panics if not found
func GetDataLoadersFromContext(ctx context.Context) *DataLoaders {
	loaders := DataLoadersFromContext(ctx)
	if loaders == nil {
		panic("DataLoaders not found in context")
	}
	return loaders
}
