package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type repository struct {
	AssetType assetType
}

type Database struct {
	Uri        string
	Username   string
	Password   string
	Repository repository
}

func (db *Database) Initialise() error {
	driver, err := neo4j.NewDriver(db.Uri, neo4j.BasicAuth(db.Username, db.Password, ""))
	if err != nil {
		return err
	}

	if err := driver.VerifyConnectivity(); err != nil {
		return err
	}

	db.Repository.AssetType.driver = driver
	return nil
}
