package cmd

import (
	"github.com/Abiji-2020/codetool/pkg"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project",
	Long:  `Initialize the project by setting up the necessary files and directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := pkg.ConnectToDatabase(); err != nil {
			cmd.Println("Error initializing project:", err)
		} else {
			cmd.Println("Project initialized successfully.")
		}
	},
}
