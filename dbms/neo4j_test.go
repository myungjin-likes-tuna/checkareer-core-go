package dbms_test

import (
	"checkareer-core/_test"
	"checkareer-core/config"
	"checkareer-core/dbms"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

var TestModules = fx.Options(
	config.Modules,
	dbms.Modules,
)

func TestCreateCyper(t *testing.T) {
	f := func(settings *config.Settings, generator dbms.Neo4jSessionGenerator) {
		if _, exist := settings.Extras["CI"]; exist {
			return
		}
		session := generator()
		defer session.Close()
		result, err := session.WriteTransaction(createItem)
		assert.NoError(t, err)
		assert.NotNil(t, result)

	}
	app := _test.NewForTest(t, TestModules, fx.Invoke(f))
	app.RequireStart()
}

func createItem(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run(
		"MERGE (n: Item {id: $id, name: $name}) RETURN n.id, n.name",
		map[string]interface{}{
			"id": 1, "name": "Item 1",
		},
	)
	if err != nil {
		return nil, err
	}
	record, err := records.Single()
	if err != nil {
		return nil, err
	}
	return &Item{
		ID:   record.Values[0].(int64),
		Name: record.Values[1].(string),
	}, nil
}

// Item
type Item struct {
	ID   int64
	Name string
}
