package skills_test

import (
	"checkareer-core/_test"
	"checkareer-core/_test/_neo4j"
	"checkareer-core/modules/skills"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

var TestModules = fx.Options(
	_neo4j.TestModules,
	fx.Provide(skills.NewRepository),
	fx.Provide(skills.NewCreater),
	fx.Provide(skills.NewReader),
)

func TestRepository(t *testing.T) {
	create := func(creater skills.Creater) {
		node, err := creater.Create(1, skills.WithName("golang"))
		assert.NoError(t, err)
		assert.NotZero(t, node)
	}
	read := func(reader skills.Reader) {
		node, err := reader.Read(skills.WithID(1))
		assert.NoError(t, err)
		assert.NotZero(t, node)
	}

	app := _test.NewForTest(t, TestModules, fx.Invoke(create, read))
	app.RequireStart()
}
