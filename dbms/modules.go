package dbms

import "go.uber.org/fx"

// Modules of dbms
var Modules = fx.Options(
	fx.Provide(NewNeo4jDriver),
	fx.Provide(NewNeo4jSessionGenerator),
)
