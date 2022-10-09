package graph

import "github.com/IzStriker/IT-Asset-Repository/backend/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AssetTypeStore          map[string]model.AssetType
	AssetTypeAttributeStore map[string]model.AssetTypeAttribute
}
