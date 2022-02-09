package api

import (
	"checkareer-core/api/v1/skills"

	"go.uber.org/fx"
)

// Modules of skills
var Modules = fx.Options(
	fx.Provide(NewServer),
	fx.Provide(skills.NewHandler),
	fx.Invoke(registerHook, skills.BindRoutes),
)
