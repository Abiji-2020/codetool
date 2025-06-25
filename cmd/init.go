package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project",
	Long:  `Initialize the project by setting up the necessary files and directories.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
