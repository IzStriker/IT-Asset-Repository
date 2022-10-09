package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/IzStriker/IT-Asset-Repository/backend/graph/generated"
	"github.com/IzStriker/IT-Asset-Repository/backend/graph/model"
	"github.com/google/uuid"
)

// UpsertAssetType is the resolver for the upsertAssetType field.
func (r *mutationResolver) UpsertAssetType(ctx context.Context, input model.AssetTypeInput) (*model.AssetType, error) {
	id := uuid.NewString()
	assetType := model.AssetType{
		ID:   id,
		Name: input.Name,
	}

	n := len(r.Resolver.AssetTypeStore)
	if n < 1 {
		r.Resolver.AssetTypeStore = make(map[string]model.AssetType)
	}

	r.Resolver.AssetTypeStore[id] = assetType

	return &assetType, nil
}

// UpsertAssetTypeAttribute is the resolver for the upsertAssetTypeAttribute field.
func (r *mutationResolver) UpsertAssetTypeAttribute(ctx context.Context, input model.AssetTypeAttributeInput) (*model.AssetTypeAttribute, error) {
	id := uuid.NewString()
	assetTypeAttribute := model.AssetTypeAttribute{
		ID:   id,
		Name: input.Name,
	}

	if len(r.Resolver.AssetTypeAttributeStore) < 1 {
		r.Resolver.AssetTypeAttributeStore = make(map[string]model.AssetTypeAttribute)
	}

	r.Resolver.AssetTypeAttributeStore[id] = assetTypeAttribute

	return &assetTypeAttribute, nil
}

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

// AssetTypeAttributes is the resolver for the assetTypeAttributes field.
func (r *queryResolver) AssetTypeAttributes(ctx context.Context) ([]*model.AssetTypeAttribute, error) {
	var attributesTypes []*model.AssetTypeAttribute

	for _, v := range r.Resolver.AssetTypeAttributeStore {
		attributesTypes = append(attributesTypes, &v)
	}

	return attributesTypes, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
