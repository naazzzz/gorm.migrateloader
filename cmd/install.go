package cmd

import (
	"github.com/spf13/cobra"
	"gorm.migrateloader/internal/installer"
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
