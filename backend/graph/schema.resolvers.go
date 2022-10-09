package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/IzStriker/IT-Asset-Repository/backend/graph/generated"
	"github.com/IzStriker/IT-Asset-Repository/backend/graph/model"
)

// AssetType is the resolver for the assetType field.
func (r *queryResolver) AssetType(ctx context.Context, id string) (*model.AssetType, error) {
	assetType, err := r.Resolver.Database.Repository.AssetType.Get(id)

	if err != nil {
		return nil, err
	}

	return assetType, nil
}

// AssetTypes is the resolver for the assetTypes field.
func (r *queryResolver) AssetTypes(ctx context.Context) ([]*model.AssetType, error) {
	types, err := r.Resolver.Database.Repository.AssetType.List()

	if err != nil {
		return nil, err
	}

	return types, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
