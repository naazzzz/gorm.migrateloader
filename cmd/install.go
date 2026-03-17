package cmd

import (
	"github.com/naazzzz/gorm-migrateloader/internal/installer"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dependencies",
	Run: func(cmd *cobra.Command, args []string) {
		installer.RunInstall()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
