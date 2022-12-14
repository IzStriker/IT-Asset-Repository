package neo4j

import (
	"strconv"

	"github.com/IzStriker/IT-Asset-Repository/backend/graph/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
)

type assetType struct {
	driver neo4j.Driver
}

func (a assetType) List() ([]*model.AssetType, error) {
	session := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	types, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var assetTypes []*model.AssetType
		query := `
		// Get all types current and parent type extends, including it's self
		MATCH(type:AssetType)
		WITH [(type)-[:EXTENDS*0..]->(parent:AssetType) | parent] as types, type as baseType
		// For each type in types array, get attributes types that given type owns
		UNWIND types as type
		MATCH(type)-[:OWNS]->(a:AssetAttributeType)
		OPTIONAL MATCH (baseType)-[:EXTENDS]->(parent)
		return baseType.name as name, id(baseType) as id, id(parent) as extendsId, collect(a) as attributes`

		result, err := tx.Run(query, nil)
		if err != nil {
			return nil, err
		}

		for result.Next() {
			record := result.Record()

			id, _ := record.Get("id")
			stringId := strconv.Itoa(int(id.(int64)))
			if err != nil {
				panic(err)
			}
			name, _ := record.Get("name")
			assetType := model.AssetType{ID: stringId, Name: name.(string)}

			if extendsId, ok := record.Get("extendsId"); ok && extendsId != nil {
				stringExtendsId := strconv.Itoa(int(extendsId.(int64)))
				assetType.ExtendsID = &stringExtendsId
			}

			if attributes, ok := record.Get("attributes"); ok && attributes != nil {
				var attributeTypes []*model.AssetTypeAttribute
				attribute := attributes.([]interface{})
				for _, value := range attribute {
					node := value.(dbtype.Node)
					var attType model.Type

					for _, customType := range model.AllType {
						if string(customType) == node.Props["type"] {
							attType = customType
							break
						}
					}

					attributeTypes = append(attributeTypes, &model.AssetTypeAttribute{
						ID:   strconv.Itoa(int(node.Id)),
						Name: node.Props["name"].(string),
						Type: &attType,
					})
				}
				assetType.Attributes = attributeTypes
			}

			if err := result.Err(); err != nil {
				return nil, err
			}
			assetTypes = append(assetTypes, &assetType)
		}

		return assetTypes, nil
	})
	if err != nil {
		return nil, err
	}
	return types.([]*model.AssetType), nil
}

func (a assetType) Get(id string) (*model.AssetType, error) {
	session := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	assetType, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var assetType *model.AssetType
		query := `
		// Get all types current and parent type extends, including it's self
		MATCH(type:AssetType)
		WHERE id(type) = $id
		WITH [(type)-[:EXTENDS*0..]->(parent:AssetType) | parent] as types, type as baseType
		// For each type in types array, get attributes types that given type owns
		UNWIND types as type
		MATCH(type)-[:OWNS]->(a:AssetAttributeType)
		OPTIONAL MATCH (baseType)-[:EXTENDS]->(parent)
		return baseType.name as name, id(baseType) as id, id(parent) as extendsId, collect(a) as attributes`

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}

		result, err := tx.Run(query, map[string]interface{}{"id": intId})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			var stringId string
			record := result.Record()
			name, _ := record.Get("name")
			if extendsId, ok := record.Get("extendsId"); ok && extendsId != nil {
				stringId = strconv.Itoa(int(extendsId.(int64)))
			}
			assetType = &model.AssetType{ID: id, Name: name.(string), ExtendsID: &stringId}

			if attributes, ok := record.Get("attributes"); ok && attributes != nil {
				var attributeTypes []*model.AssetTypeAttribute
				attribute := attributes.([]interface{})
				for _, value := range attribute {
					node := value.(dbtype.Node)
					var attType model.Type

					for _, customType := range model.AllType {
						if string(customType) == node.Props["type"] {
							attType = customType
							break
						}
					}

					attributeTypes = append(attributeTypes, &model.AssetTypeAttribute{
						ID:   strconv.Itoa(int(node.Id)),
						Name: node.Props["name"].(string),
						Type: &attType,
					})
				}
				assetType.Attributes = attributeTypes
			}

		}
		return assetType, nil
	})
	if err != nil {
		return nil, err
	}

	return assetType.(*model.AssetType), nil
}
