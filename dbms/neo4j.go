package dbms

import (
	"checkareer-core/config"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// NewNeo4jDriver 새로운 neo4j driver 생성
func NewNeo4jDriver(settings *config.Settings) neo4j.Driver {
	driver, err := neo4j.NewDriver(
		settings.Neo4j.URI,
		neo4j.BasicAuth(settings.Neo4j.Username, settings.Neo4j.Password, ""),
	)
	if err != nil {
		log.Fatal(err)
	}
	return driver
}

// Neo4jSessionGenerator generator
type Neo4jSessionGenerator func() neo4j.Session

// NewNeo4jSessionGenerator 새로운 neo4j session 생성
func NewNeo4jSessionGenerator(driver neo4j.Driver) Neo4jSessionGenerator {
	return func() neo4j.Session {
		return driver.NewSession(neo4j.SessionConfig{})
	}
}
