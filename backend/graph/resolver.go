package graph

import (
	"github.com/IzStriker/IT-Asset-Repository/backend/graph/model"
	"github.com/IzStriker/IT-Asset-Repository/backend/neo4j"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AssetTypeStore          map[string]model.AssetType
	AssetTypeAttributeStore map[string]model.AssetTypeAttribute
	Database                neo4j.Database
}
