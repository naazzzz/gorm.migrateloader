package cmd

import (
	"github.com/spf13/cobra"
	"gorm.migrateloader/internal/generator"
)

var driver string

var generateCmd = &cobra.Command{
	Use:   "generate [path]",
	Short: "Generate models",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		generator.RunGenerate(path, driver)
	},
}

func init() {
	generateCmd.Flags().StringVar(&driver, "driver", "postgres", "db driver")
	rootCmd.AddCommand(generateCmd)
}
