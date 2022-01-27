package cmd

import (
	"checkareer-core/config"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config",
	Version: Version,
	Short:   fmt.Sprintf("%s Config", Name),
	Run: func(cmd *cobra.Command, args []string) {
		log.Print(config.JSON())
	},
}
