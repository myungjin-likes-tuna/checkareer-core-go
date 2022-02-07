package _neo4j

import (
	"checkareer-core/config"
	"checkareer-core/dbms"
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.uber.org/fx"
)

// registerHook invoke로 호출되므로 파라미터의 순서가 초기화의 순서
func registerHook(
	lifecycle fx.Lifecycle,
	settings *config.Settings,
	driver neo4j.Driver,
) {
	var neo4jContainer testcontainers.Container
	lifecycle.Append(fx.Hook{
		OnStart: func(context context.Context) error {
			if settings.CI {
				neo4jContainer = NewNeo4jTestContainer(settings)
			}
			return nil
		},
		OnStop: func(context context.Context) error {
			err := driver.Close()
			if err != nil {
				fmt.Println(err)
			}
			if settings.CI {
				return neo4jContainer.Terminate(context)
			}
			return err
		},
	})
}

// NewNeo4jTestContainer for neo4j test
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

// NewSettings constructor
func NewSettings() *config.Settings {
	var settings config.Settings
	settings.Neo4j.URI = "neo4j://localhost:7687"
	if settings.CI {
		settings.Neo4j.Username = "neo4j"
		settings.Neo4j.Password = "t2st2r"
	}
	return &settings
}

// TestModules of neo4j
var TestModules = fx.Options(
	fx.Provide(
		dbms.NewNeo4jDriver,
		dbms.NewNeo4jSessionGenerator,
		NewSettings,
		NewNeo4jTestContainer,
	),
	fx.Invoke(registerHook),
)
