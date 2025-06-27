package cmd

import (
	"fmt"

	"github.com/Abiji-2020/codetool/config"
	"github.com/Abiji-2020/codetool/pkg"
	"github.com/spf13/cobra"
)

var Client *pkg.MindsDBClient
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the project",
	Long:  `Initialize the project by setting up the necessary files and directories.`,
	Run: func(cmd *cobra.Command, args []string) {
		Client := pkg.NewMindsDBClient("")
		if err := Client.ConnectToDatabase(); err != nil {
			fmt.Println("Error connecting to database:", err)
			return
		}

		/*	if err := Client.CreateTable(); err != nil {
				fmt.Println("Error creating table:", err)
				return
			}
		*/
		if err := Client.CreateKnowledgeBase(config.AgentKnowledgeBase); err != nil {
			fmt.Println("Error creating knowledge base:", err)
			return
		}

		if err := Client.CreateEngine(); err != nil {
			fmt.Println("Error creating engine:", err)
			return
		}

		if err := Client.CreateModel(); err != nil {
			fmt.Println("Error creating model:", err)
			return
		}

		if err := Client.CreateAgent(config.AgentName); err != nil {
			fmt.Println("Error creating agent:", err)
			return
		}

		fmt.Println("Project initialized successfully.")
	},
}
