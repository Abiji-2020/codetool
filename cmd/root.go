/*
Copyright © 2025 Abinand P
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "codetool",
	Short: "A tool for generating code templates",
	Long: `A simple code tool, which can generate 
	various code templates for different programming languages, 
	such as Python, Go, Java, Javascript and Ruby. 
	
	It generates the code with the help of knowledgebase from the dataset
	https://huggingface.co/datasets/claudios/code_search_net `,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
