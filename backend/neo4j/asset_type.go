package neo4j

import (
	"strconv"

	"github.com/IzStriker/IT-Asset-Repository/backend/graph/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type assetType struct {
	driver neo4j.Driver
}

func (a assetType) List() ([]*model.AssetType, error) {
	session := a.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	types, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var assetTypes []*model.AssetType
		query := `MATCH (t:AssetType)
		OPTIONAL MATCH (t:AssetType)-[:EXTENDS]->(p:AssetType)
		RETURN id(t) as id, t.name as name, id(p) as extendsId;`

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
		query := `MATCH(t:AssetType)
		WHERE id(t) = $id
		OPTIONAL MATCH (t:AssetType)-[:EXTENDS]->(p:AssetType)
		RETURN t.name as name, id(p) as extendsId`

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, err
		}

		result, err := tx.Run(query, map[string]interface{}{"id": intId})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			record := result.Record()
			name, _ := record.Get("name")
			extendsId, _ := record.Get("extendsId")
			stringId := strconv.Itoa(int(extendsId.(int64)))

			assetType = &model.AssetType{ID: id, Name: name.(string), ExtendsID: &stringId}
		}
		return assetType, nil
	})
	if err != nil {
		return nil, err
	}

	return assetType.(*model.AssetType), nil
}
