package cmd

import (
	"fmt"

	"github.com/Abiji-2020/codetool/pkg"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project",
	Long:  `Initialize the project by setting up the necessary files and directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := pkg.NewMindsDBClient("")
		if err := client.ConnectToDatabase(); err != nil {
			fmt.Println("Error connecting to database:", err)
			return
		}
		if err := client.CreateTable(); err != nil {
			fmt.Println("Error creating table:", err)
			return
		}
		fmt.Println("Project initialized successfully.")
	},
}
