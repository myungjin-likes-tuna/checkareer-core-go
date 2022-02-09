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
	f := func(creater skills.Creater, reader skills.Reader) {
		skill, err := creater.Create(skills.WithName("golang"))
		assert.NoError(t, err)
		assert.NotZero(t, skill)

		_skill, err := reader.Read(skills.WithID(skill.ID))
		assert.NoError(t, err)
		assert.NotZero(t, _skill)
	}

	app := _test.NewForTest(t, TestModules, fx.Invoke(f))
	app.RequireStart()
}
