package skills

import "go.uber.org/fx"

// Modules of skills
var Modules = fx.Options(
	fx.Provide(NewRepository),
	fx.Provide(NewCreater),
	fx.Provide(NewReader),
	fx.Provide(NewUpdater),
	fx.Provide(NewDeleter),
)
