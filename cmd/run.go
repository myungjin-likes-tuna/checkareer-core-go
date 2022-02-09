package cmd

import (
	"checkareer-core/api"
	"checkareer-core/config"
	"checkareer-core/dbms"
	"checkareer-core/modules/skills"
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Version: Version,
	Short:   fmt.Sprintf("%s Run", Name),
	Run: func(cmd *cobra.Command, args []string) {
		f := func() {}
		modules := fx.Options(
			config.Modules,
			dbms.Modules,
			skills.Modules,
			api.Modules,
			fx.Invoke(f),
		)
		fx.New(modules).Run()
	},
}
