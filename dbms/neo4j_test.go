package dbms_test

import (
	"checkareer-core/_test"
	"checkareer-core/config"
	"checkareer-core/dbms"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/fx"
)

// registerHook invoke로 호출되므로 파라미터의 순서가 초기화의 순서
func registerHook(
	lifecycle fx.Lifecycle,
	neo4jContainer testcontainers.Container,
	driver neo4j.Driver,
) {
	lifecycle.Append(fx.Hook{
		OnStop: func(context context.Context) error {
			err := driver.Close()
			if err != nil {
				fmt.Println(err)
			}
			return neo4jContainer.Terminate(context)
		},
	})
}

func NewNeo4jTestContainer(settings *config.Settings) testcontainers.Container {
	ctx := context.Background()
	container, err := startContainer(ctx, settings.Neo4j.Username, settings.Neo4j.Password)
	if err != nil {
		log.Fatal(err)
	}
	port, err := container.MappedPort(ctx, "7687")
	if err != nil {
		log.Fatal(err)
	}
	settings.Neo4j.URI = fmt.Sprintf("neo4j://localhost:%d", port.Int())
	return container
}

func startContainer(
	ctx context.Context,
	username, password string,
) (testcontainers.Container, error) {
	request := testcontainers.ContainerRequest{
		Image:        "neo4j",
		ExposedPorts: []string{"7687/tcp"},
		Env: map[string]string{
			"NEO4J_AUTH": fmt.Sprintf("%s/%s", username, password),
		},
		WaitingFor: wait.ForLog("hi"),
	}
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})
}

func NewSettings() *config.Settings {
	var settings config.Settings
	settings.Neo4j.URI = ""
	settings.Neo4j.Username = "neo4j"
	settings.Neo4j.Password = "t2st2r"
	return &settings
}

var TestModules = fx.Options(
	fx.Provide(
		dbms.NewNeo4jDriver,
		dbms.NewNeo4jSessionGenerator,
		NewSettings,
		NewNeo4jTestContainer,
	),
	fx.Invoke(registerHook),
)

func TestCreateCyper(t *testing.T) {
	f := func(settings *config.Settings, generator dbms.Neo4jSessionGenerator) {
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
